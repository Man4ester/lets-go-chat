package handlers

import (
	"testing"
	"net/http"
	"strings"
	"net/http/httptest"
	"github.com/stretchr/testify/mock"
	"github.com/gorilla/websocket"
	"lets-go-chat/mocks"
	"lets-go-chat/internal/models"
)

func TestGetActiveUsers(t *testing.T) {
	url := "/v1/user/active"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetActiveUsers)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "{\"count\":0}\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}

func TestUserCreation_CreateUserOK(t *testing.T)  {
	url := "/v1/user"
	r := strings.NewReader("{\n    \"userName\":\"Alex2@gmail.com\",\n    \"password\":\"password\"\n}")
	req, err := http.NewRequest(http.MethodPost, url, r)
	if err != nil {
		t.Fatal(err)
	}

	repo := new(mocks.UserRepository)
	repo.On("SaveUser", mock.Anything).Return(nil)

	rr := httptest.NewRecorder()

	hUserCreation := UserCreation{
		Repo: repo,
	}
	handler := http.HandlerFunc(hUserCreation.CreateUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestUserCreation_CreateUserKO_BadPayload(t *testing.T)  {
	url := "/v1/user"
	r := strings.NewReader("{\n    \"userName\":\"Alex2@gmail.com\"}")
	req, err := http.NewRequest(http.MethodPost, url, r)
	if err != nil {
		t.Fatal(err)
	}

	repo := new(mocks.UserRepository)
	repo.On("SaveUser", mock.Anything).Return(nil)

	rr := httptest.NewRecorder()

	hUserCreation := UserCreation{
		Repo: repo,
	}
	handler := http.HandlerFunc(hUserCreation.CreateUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestUserLogin_LoginUser(t *testing.T) {
	url := "/v1/user/login"
	r := strings.NewReader("{\n    \"userName\":\"Alex3@gmail.com\",\n    \"password\":\"password\"\n}")
	req, err := http.NewRequest(http.MethodPost, url, r)
	if err != nil {
		t.Fatal(err)
	}

	repo := new(mocks.UserRepository)
	repo.On("GetUserByUserName", mock.Anything).Return(models.User {
		Id : "id",
		UserName: "Alex3@gmail.com",
		Password: "5f4dcc3b5aa765d61d8327deb882cf99",
	}, nil)

	rr := httptest.NewRecorder()

	hUserCreation := UserLogin{
		Repo: repo,
	}
	handler := http.HandlerFunc(hUserCreation.LoginUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestWsRTMStartKO(t *testing.T) {
	url := "/v1/chat/ws.rtm.start"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	hWS := WsRTM{
		Upgrader: websocket.Upgrader{},
	}

	handler := http.HandlerFunc(hWS.WsRTMStart)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}