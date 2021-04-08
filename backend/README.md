# golang-auth

## Set Up MySQL Database with docker

```bash
# start database
docker run --name dev-mysql -e MYSQL_ROOT_PASSWORD=<password> -d -p 3306:3306 mysql:latest
# connect to database
docker exec -it dev-mysql bash
mysql -u root -p
```

```
PORT=...
DB_USERNAME=...
DB_PASSWORD=...
DB_DATABASE=...
JWT_VALID_TIME=...
```
