# lets-go-chat

## Start using Docker

build

```
docker build . -t go-chat
```

run

```
docker run -p 8080:8080 -it go-chat
```

run DB

```
docker run -it --rm --name go-postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=secret postgres:14.0
```

```
CREATE TABLE public.users (
	id varchar NULL,
	username varchar NULL,
	"password" varchar NULL
);

```

## Rest API

create user

```
curl --location --request POST 'http://localhost:8080/v1/user' \
--header 'Content-Type: application/json' \
--data-raw '{
    "userName":"Alex2",
    "password":"password"
}'
```

login user

```
curl --location --request POST 'http://localhost:8080/v1/user/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "userName":"Alex",
    "password":"password"
}'
```

Test covearge
```
go test -cover
go test -cover -coverprofile=c.out
go tool cover -html=c.out -o coverage.html 
go test -bench=.
go test ./...
```