# ISHOPPING

## 文档

- Release1：[ishopping/interface-release1.md at main · carpediemtal/ishopping (github.com)](https://github.com/carpediemtal/ishopping/blob/main/doc/interface-release1.md)
- Release2：[ishopping/interface-release2.md at main · carpediemtal/ishopping (github.com)](https://github.com/carpediemtal/ishopping/blob/main/doc/interface-release2.md)
- Release3：[ishopping/interface-release3.md at main · carpediemtal/ishopping (github.com)](https://github.com/carpediemtal/ishopping/blob/main/doc/interface-release3.md)

## 目录结构说明

所有后端代码都放在`ishopping/src`下，前端代码放在`ishopping/view`下

## 后端运行方法

```bash
git clone git@github.com:carpediemtal/ishopping.git
cd ishopping
go run src/main.go
```

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
    sid       int            not null,
    name      varchar(32)    not null,
    price     decimal(10, 2) not null,
    sales     int default 0  not null,
    inventory int default 0  not null,
    caid      int            not null
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

create table cart
(
    cartid int auto_increment
        primary key,
    cid    int           not null,
    uid    int           not null,
    count  int default 0 not null,
    constraint cid
        foreign key (cid) references commodity (cid),
    constraint userid
        foreign key (uid) references user (uid)
);

create table user_ban
(
    uid int null,
    constraint uid
        foreign key (uid) references user (uid)
);

create table user_meta
(
    mid      int auto_increment
        primary key,
    uid      int         not null,
    meta_key varchar(64) not null,
    meta_val text        not null
);

create definer = root@localhost view ub_cart as
select `ishopping`.`cart`.`cartid` AS `cartid`,
       `ishopping`.`cart`.`cid`    AS `cid`,
       `ishopping`.`cart`.`uid`    AS `uid`,
       `ishopping`.`cart`.`count`  AS `count`
from `ishopping`.`cart`
where `ishopping`.`cart`.`cid` in (select `ub_commodity`.`cid` from `ishopping`.`ub_commodity`);

create definer = root@localhost view ub_commodity as
select `ishopping`.`commodity`.`cid`       AS `cid`,
       `ishopping`.`commodity`.`sid`       AS `sid`,
       `ishopping`.`commodity`.`name`      AS `name`,
       `ishopping`.`commodity`.`price`     AS `price`,
       `ishopping`.`commodity`.`sales`     AS `sales`,
       `ishopping`.`commodity`.`inventory` AS `inventory`,
       `ishopping`.`commodity`.`caid`      AS `caid`
from `ishopping`.`commodity`
where `ishopping`.`commodity`.`sid` in (select `ub_shop`.`sid` from `ishopping`.`ub_shop`);

create definer = root@localhost view ub_shop as
select `ishopping`.`shop`.`sid`       AS `sid`,
       `ishopping`.`shop`.`uid`       AS `uid`,
       `ishopping`.`shop`.`shop_name` AS `shop_name`
from `ishopping`.`shop`
where `ishopping`.`shop`.`uid` in (select `ishopping`.`user_ban`.`uid` from `ishopping`.`user_ban`) is false;

create definer = root@localhost view ub_user as
select `ishopping`.`user`.`uid`      AS `uid`,
       `ishopping`.`user`.`username` AS `username`,
       `ishopping`.`user`.`password` AS `password`,
       `ishopping`.`user`.`type`     AS `type`
from `ishopping`.`user`
where `ishopping`.`user`.`uid` in (select `ishopping`.`user_ban`.`uid` from `ishopping`.`user_ban`) is false;



```

## 用到的框架

### 数据库

[sqlx](https://github.com/jmoiron/sqlx)

[sqlx user documentation](http://jmoiron.github.io/sqlx/)

### web

[GIN](https://github.com/gin-gonic/gin#quick-start)

## 文档示例

### 请求

使用 json 数据封装请求，看前端代码

### 返回

返回的是 json 格式，后端的看后端代码怎么处理，前端的看前端代码怎么处理。

### 通用定义

| 字段名 | 类型   | 说明                          |
| ------ | ------ | ----------------------------- |
| code   | int    | 0 表示成功，非 0 表示发生错误 |
| msg    | string | 错误信息                      |
| data   | {}     | 成功时返回的数据              |

### data 字段

该字段为可变字段，视业务需要定义，比如示例代码中返回为：{name: "XXX"}，写文档时写成如下：

| 字段名 | 类型   | 说明     |
| ------ | ------ | -------- |
| name   | string | 返回名字 |