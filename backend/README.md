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

```sql
create table if not exists users
(
    id       bigint unsigned auto_increment
    primary key,
    username varchar(255) null,
    password longblob     null,
    constraint username
    unique (username)
);

create table macs
(
    id      bigint unsigned not null
        primary key,
    user_id bigint unsigned not null,
    name    varchar(255)    not null,
    mac     varchar(18)     not null,
    constraint macs_users_id_fk
        foreign key (user_id) references users (id)
            on update cascade on delete cascade
);
```