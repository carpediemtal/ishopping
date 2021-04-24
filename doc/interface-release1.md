[toc]

# 基本知识

json基本数据类型为：

| 类型             | 表示       |
| ---------------- | ---------- |
| 字符串           | string     |
| 数字             | int、float |
| 对象（json对象） | {}         |
| 数组             | []         |
| 布尔             | bool       |
| 无               | Null       |

基本上看一遍就知道

# 接口请求

所有接口请求使用json封装

# 接口返回外层

任何返回都要按照这个格式，数据放在data段。

| 字段名 | 类型   | 说明                             |
| ------ | ------ | -------------------------------- |
| code   | bool   | 0表示成功，非0表示发生错误       |
| msg    | string | 错误时信息                       |
| data   | {}     | 成功时返回数据集，类型为嵌套json |

# 接口

## 1.register

### 请求

| 请求类型 | path          |
| -------- | ------------- |
| POST     | /api/register |

| 字段名    | 类型   | 说明                 |
| --------- | ------ | -------------------- |
| username  | string | 用户名               |
| password  | string | 密码                 |
| user_type | int    | 0表示卖家，1表示商家 |

### 返回

无data，成功code=0，错误code=-1，错误信息返回在msg里

## 2.login

### 请求

| 请求类型 | path       |
| -------- | ---------- |
| POST     | /api/login |

| 字段名   | 类型   |
| -------- | ------ |
| username | string |
| password | string |

### 返回

成功code=0，错误code=-1，错误信息在msg里。成功时data嵌套数据如下：

| 字段名 | 类型   | 说明                                 |
| ------ | ------ | ------------------------------------ |
| token  | string | 身份token，前端把它设置cookie名为jwt |

## 3.search

### 请求

| 请求类型 | path                  |
| -------- | --------------------- |
| GET      | /api/commodity_search |

| 字段名  | 类型   | 说明         |
| ------- | ------ | ------------ |
| content | string | 搜索框内内容 |

### 返回

返回的是json格式，后端的看后端代码怎么处理，前端的看前端代码怎么处理。

| 字段名         | 类型 | 说明              |
| -------------- | ---- | ----------------- |
| commodity_list | []   | 粗略commodity列表 |

list_item:

| 字段名       | 类型   | 说明     |
| ------------ | ------ | -------- |
| commodity_id | int    | 商品id   |
| name         | string | 商品名   |
| price        | float  | 商品价格 |

## 4.detail

### 请求

| 请求类型 | path                  |
| -------- | --------------------- |
| GET      | /api/commodity_detail |

| 字段名       | 类型 | 说明   |
| ------------ | ---- | ------ |
| commodity_id | int  | 商品id |

### 返回

| 字段名         | 类型   | 说明                                         |
| -------------- | ------ | -------------------------------------------- |
| commodity_name | string | 商品名                                       |
| price          | float  | 商品价格                                     |
| sales          | int    | 商品已出售量                                 |
| introduction   | string | 商品文本介绍，该字段以后会改，第一版先这样吧 |
| inventory      | int    | 商品库存                                     |

## 5.commodity add

### 请求

| 请求类型 | path               |
| -------- | ------------------ |
| POST     | /api/commodity_add |

| 字段名         | 类型   | 说明     |
| -------------- | ------ | -------- |
| user_id        | int    | 用户名   |
| commodity_name | string | 商品名   |
| inventory      | int    | 商品库存 |
| introduction   | string | 商品介绍 |
| price          | float  | 商品价格 |

### 返回

成功code=0，错误code=-1，错误信息在msg里。成功时data里字段如下：

| 字段名       | 类型 | 说明   |
| ------------ | ---- | ------ |
| commodity_id | int  | 商品id |

# 示例

/api/commodity_add ：

## 请求

```json
{
    "content": "世界好组长"
}
```

## 返回

```json
{
    "code": 0,
    "msg": "",
    "data": {
        "commodity_list": [
            {
                "commodity_id": 1,
                "name": "commodity1",
                "price": 1.01
            },
            {
                "commodity_id": 2,
                "name": "commodity2",
                "price": 2.09
            }
        ]
    }
}
```