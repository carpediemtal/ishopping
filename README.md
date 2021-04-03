# ishopping

## 数据库建表

```sql
mysql -u root -p

create database ishopping;

use ishopping;

create table buyer
(
    uid       int         not null
        primary key,
    name      varchar(16) not null,
    address   varchar(64) not null,
    phone_num char(11)    not null,
    constraint buyer_phone_num_uindex
        unique (phone_num)
);

create table cart
(
    cid int not null,
    uid int not null
);

create table category
(
    caid int auto_increment
        primary key,
    name varchar(32) not null
);

create table comment
(
    coid      int auto_increment
        primary key,
    cid       int  not null,
    uid       int  not null,
    content   text not null,
    timestamp int  not null,
    rate      int  not null
);

create table commodity
(
    cid       int auto_increment
        primary key,
    sid       int           not null,
    name      varchar(32)   not null,
    price     int           not null,
    sales     int default 0 not null,
    inventory int default 0 not null,
    caid      int           not null
);

create table commodity_meta
(
    mid      int auto_increment
        primary key,
    cid      int         not null,
    meta_key varchar(32) not null,
    meta_val text        not null
);

create table photo
(
    pid  int auto_increment
        primary key,
    path text not null
);

create table purchase_order
(
    oid       int auto_increment
        primary key,
    status    int not null,
    uid       int not null,
    cid       int not null,
    timestamp int not null
);

create table shop
(
    sid       int auto_increment
        primary key,
    uid       int         not null,
    shop_name varchar(32) not null,
    constraint shop_shop_name_uindex
        unique (shop_name)
);

create table user
(
    uid      int auto_increment
        primary key,
    username varchar(32) not null,
    password char(32)    not null,
    type     tinyint(1)  not null,
    constraint user_username_uindex
        unique (username)
);

create table user_meta
(
    mid      int auto_increment
        primary key,
    uid      int         not null,
    meta_key varchar(64) not null,
    meta_val text        not null
);


```

## 框架

### 数据库

[sqlx](https://github.com/jmoiron/sqlx)

[sqlx user documentation](http://jmoiron.github.io/sqlx/)

### web

[GIN](https://github.com/gin-gonic/gin#quick-start)