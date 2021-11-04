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