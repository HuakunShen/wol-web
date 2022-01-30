# wol-web backend

## Environment Variables `.env`

Sample `.env`

```
PORT=9090
JWT_SECRET=secret
JWT_VALID_TIME=14400

NUM_USER_ALLOWED=1
```

Use `NUM_USER_ALLOWED=x` to allow x users to signup, default is `1`. If you would like to allow more users to sign up modify this environment variable.

Use `JWT_VALID_TIME` to set valid login time. The unit is **minute**.

## Database

The database used has switched from postgresql to sqlite since the app is simple enough. to be handled by sqlite. The following schemas are a reference to its database design.

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
    owner to wol;

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
    port    integer      default 9                                    not null,
    constraint computers_pk_2
        unique (user_id, name)
);

alter table computers
    owner to wol;


```
