# api

## mysql 立ち上げと db 作成

```
docker run -d  -p 3306:3306 -e MYSQL_ALLOW_EMPTY_PASSWORD=yes mysql:5.7

# dockerのmysqlの中に入って，以下のsqlでplaylistというdbを作成する
CREATE DATABASE playlist CHARACTER SET utf8mb4;
```

## mysql migration

```
cd schema
bash hack.sh sqldef_apply
```

## api server 立ち上げ

```
cp .env.sample .env
# .env に値当てはめ

go run main.go
```
