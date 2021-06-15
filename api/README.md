# api

## mysql 立ち上げ

```
docker run -d  -p 3306:3306 -e MYSQL_ALLOW_EMPTY_PASSWORD=yes mysql:5.7
```

## api server 立ち上げ

```
go run main.go
```
