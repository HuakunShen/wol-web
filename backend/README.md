# golang-auth

## Set Up MySQL Database with docker

```bash
# start database
docker run --name dev-mysql -e MYSQL_ROOT_PASSWORD=<password> -d -p 3306:3306 mysql:latest
# connect to database
docker exec -it dev-mysql bash
mysql -u root -p
```

## Set Up PostgreSQL Database with docker

```bash
docker run --name wol-db \
-e POSTGRES_PASSWORD=password \
-e PGDATA=/var/lib/postgresql/data/pgdata \
-v $PWD/pgdata:/var/lib/postgresql/data \
-d -p 5432:5432 \
postgres:latest
```

## Environment Variables `.env`

```
PORT=9090
JWT_SECRET=secret
JWT_VALID_TIME=14400

POSTGRES_HOST=localhost

POSTGRES_PORT=5432
POSTGRES_USER=wol
POSTGRES_PASSWORD=wakeonlan
POSTGRES_DB=wol
POSTGRES_TIMESZONE=America/Toronto

NUM_USER_ALLOWED=1
```

The **PORT** can be modified to any port but **POSTGRES_PORT** cannot be changed for now.

Use `NUM_USER_ALLOWED=x` to allow x users to signup, default is `1`. If you would like to allow more users to sign up modify this environment variable.

Use `JWT_VALID_TIME` to set valid login time. The unit is **minute**.

### MySQL Schema

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

create table if not exists computers
(
    id      bigint unsigned auto_increment
        primary key,
    user_id bigint unsigned                        not null,
    name    varchar(255)                           not null,
    mac     varchar(18)                            not null,
    ip      varchar(26)  default '255.255.255.255' not null,
    port    int unsigned default '9'               not null,
    constraint macs_pk
        unique (user_id, name),
    constraint macs_users_id_fk
        foreign key (user_id) references users (id)
            on update cascade on delete cascade
);
```

### Postgresql Schema

```sql
create table if not exists users
(
    id       bigserial not null
        constraint users_pkey
            primary key,
    username text
        constraint users_username_key
            unique,
    password bytea
);

alter table users
    owner to postgres;


create table computers
(
    id      serial                                                    not null
        constraint computers_pk
            primary key,
    user_id integer                                                   not null
        constraint computers_users_id_fk
            references users
            on update cascade on delete cascade,
    name    varchar(255)                                              not null,
    mac     varchar(17)                                               not null,
    ip      varchar(100) default '255.255.255.255'::character varying not null,
    port    integer      default 9                                    not null
);

alter table computers
    owner to postgres;
```
