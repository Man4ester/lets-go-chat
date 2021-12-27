package main

import (
	"net/http"
	"strconv"
	"log"
	"flag"
	"database/sql"
	"fmt"
	"time"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/urfave/negroni"
	"github.com/justinas/alice"
	"github.com/gorilla/websocket"
	"lets-go-chat/configs"
	"lets-go-chat/internal/handlers"
	rep "lets-go-chat/internal/repositories"
	"lets-go-chat/internal/services"
	"lets-go-chat/pkg/jwt"
)

func main() {

	configFile := flag.String("config", "config.json", "Configuration file in JSON-format")
	flag.Parse()

	config, err := configs.LoadConfig(*configFile)
	if err != nil {
		log.Fatal("can't load configuration")
	}

	jwt.ApplySecret(config.JWTSecret)

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DBConfig.DBHost, config.DBConfig.DBPort, config.DBConfig.DBUser, config.DBConfig.DBPassword, config.DBConfig.DBName)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal("can't connect to DB")
	}
	defer db.Close()

	rep.RegisterUserRepository(rep.NewUsersDataRepository(db))
	rep.NewUsersDataRepository(db)

	commonHandlers := alice.New(errorMiddleware, recoverHandler)
	authHandlers := alice.New(errorMiddleware, recoverHandler, authMiddleware)

	r := mux.NewRouter()

	hUserCreation := handlers.UserCreation{
		Repo: rep.GetUserRepository(),
	}

	hUserLogin := handlers.UserLogin{
		Repo: rep.GetUserRepository(),
	}

	hWS := handlers.WsRTM{
		Upgrader: websocket.Upgrader{},
	}
	services.StartRedis(config.RedisUrl)

	r.Handle("/v1/user", commonHandlers.Then(http.HandlerFunc(hUserCreation.CreateUser))).Methods(http.MethodPost)
	r.Handle("/v1/user/login", commonHandlers.Then(http.HandlerFunc(hUserLogin.LoginUser))).Methods(http.MethodPost)
	r.Handle("/v1/user/active", commonHandlers.Then(http.HandlerFunc(handlers.GetActiveUsers))).Methods(http.MethodGet)
	r.Handle("/v1/chat/ws.rtm.start", authHandlers.Then(http.HandlerFunc(hWS.WsRTMStart))).Methods(http.MethodGet)
	err =  http.ListenAndServe(":"+strconv.Itoa(config.ServerPort), requestLogger(r))
	if err != nil {
		log.Fatal("server failed to add listener")
	}

}


func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		keys, ok := r.URL.Query()["token"]
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Can't read token parameter")
			return
		}
		token :=keys[0]
		err := services.ApplyTokenFromRegistry(token)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			log.Println(err)
			return
		}
		userName, err := jwt.DecodeJWT(token)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		log.Println(userName)
		next.ServeHTTP(w, r)

	})
}

func errorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lrw := negroni.NewResponseWriter(w)
		next.ServeHTTP(lrw, r)
		statusCode := lrw.Status()

		if lrw.Status() < http.StatusOK || lrw.Status() > http.StatusMultipleChoices {
			log.Printf("<-- %d %s", statusCode, http.StatusText(statusCode))
		}
	})
}

// The recoverHandler will catch any panic and stops the panicking sequence.
// Instead of panic it will give a possibility to manage response in appropriate way.
// It is possible using recover() with defer only, otherwise it will not stop panicking sequence.
// In our case, we will return to user 500 response instead of crashed call at all.
func recoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func requestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		targetMux.ServeHTTP(w, r)
		requesterIP := r.RemoteAddr
		log.Printf(
			"%s\t\t%s\t\t%s\t\t%v",
			r.Method,
			r.RequestURI,
			requesterIP,
			time.Since(start),
		)
	})
}