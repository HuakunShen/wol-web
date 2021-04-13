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
