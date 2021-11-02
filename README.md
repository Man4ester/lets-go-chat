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
curl --location --request POST 'http://localhost:8080/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"Sasha",
    "userName":"userName",
    "password":"passsdasdasdasdword"
}'
```

get user
```
curl --location --request GET 'http://localhost:8080/users/{user_id}'
```