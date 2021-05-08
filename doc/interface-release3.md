# 说明

- 对于接口有问题的，在大群里说或私聊群里的大佬们。
- release3开发周期：**第9周** 至 **第12周**，请尽快。
-  接口有修改请在次写下修改的地方。
- 最新更新时间：**2021-05-08 23:53**

# Visitor

## search enhance

### 说明

把release1的搜索合并到改接口，该接口可以实现搜索特定价格区间的商品。

### 请求

| 请求类型 | path        |
| -------- | ----------- |
| get      | /api/search |

| 字段名    | 类型   | 说明         |
| --------- | ------ | ------------ |
| content   | string | 搜索框内内容 |
| page      | id     | 页数         |
| price_min | float  | 最低价格     |
| price_max | float  | 最高价格     |

### 返回

| 字段名 | 类型        | 说明     |
| ------ | ----------- | -------- |
| list   | []list_item | 商品列表 |

list_item

| 字段名          | 类型   | 说明       |
| --------------- | ------ | ---------- |
| commodity_id    | id     | 商品id     |
| name            | string | 商品名     |
| commodity_price | float  | 商品价格   |
| thumbnail       | string | 商品缩略图 |
| sales           | int    | 销量       |

# Buyer

## order history

### 说明


订单状态：1-3：未发货、已发货、已确认收货

### 请求

| 请求类型 | path                     |
| -------- | ------------------------ |
| get      | /api/buyer/order_history |

| 字段名    | 类型 | 说明   |
| --------- | ---- | ------ |
| page      | int  | 当前页 |
| page_size | int  | 页大小 |

### 返回

| 字段名       | 类型   | 说明                   |
| ------------ | ------ | ---------------------- |
| create_time  | string | 订单创建的时间         |
| paid_time    | string | 订单付款的时间         |
| pay_price    | string | 买家确认收货的付款金额 |
| order_id     | int    | 订单id                 |
| item_img     | string | 商品图片               |
| item_num     | int    | 商品数量               |
| commodity_id | int    | 商品id                 |
| shop_name    | string | 店铺名称               |
| order_status | int    | 订单状态               |

## evaluate

### 说明

用户评价

### 请求

请求类型| path
-------- | -----
post| /api/buyer/evaluate 

字段名| 类型 | 说明
-------- | ----- | -----
order_id| int | 订单id 
rate| int | 评分1-5 
content| string | 评价内容 

### 返回

无data，添加成功code=0，错误code=-1，错误信息返回在msg里

## commodity evaluation

### 说明

在详情页用这个接口

### 请求

| 请求类型 | path                            |
| -------- | ------------------------------- |
| get      | /api/buyer/commodity_evaluation |

| 字段名       | 类型 | 说明   |
| ------------ | ---- | ------ |
| commodity_id | int  | 商品id |

### 返回

| 字段名 | 类型        | 说明     |
| ------ | ----------- | -------- |
| list   | []list_item | 评价列表 |

list_item

| 字段名    | 类型   | 说明           |
| --------- | ------ | -------------- |
| rate      | int    | 评分1-5        |
| content   | string | 商品评价       |
| timestamp | int    | 评价当时时间戳 |

## cart add

### 请求

请求类型| path
-------- | -----
post| /api/buyer/cart_add 

| 字段名       | 类型 | 描述       |
| ------------ | ---- | ---------- |
| commodity_id | int  | 商品id     |
| count        | int  | 购买的数量 |

### 返回

无data，添加成功code=0，错误code=-1，错误信息返回在msg里

## cart_get

### 请求

| 请求类型 | path                |
| -------- | ------------------- |
| get      | /api/buyer/Cart_get |

| 字段名    | 类型 | 说明   |
| --------- | ---- | ------ |
| page      | int  | 当前页 |
| page_size | int  | 页大小 |

### 返回

| 字段名 | 类型 | 说明       |
| ------ | ---- | ---------- |
| list   | []   | 购物车数据 |

list结构

| Parameter      | 类型   | 描述       |
| -------------- | ------ | ---------- |
| commodity_id   | int    | 商品id     |
| commodity_name | string | 商品名称   |
| thumbnail      | string | 缩略图     |
| price          | double | 价格       |
| count          | int    | 购买的数量 |

# Manager

## admin login

单独的管理员登录界面

## evaluation list

### 说明

请求该接口后配合使用delete evaluation、ban buyer

### 请求

请求类型| path
-------- | -----
 get      | /api/manager/evaluation_list 

| 字段名    | 类型 | 说明     |
| --------- | ---- | -------- |
| page      | int  | 页数     |
| page_size | int  | 每页个数 |

### 返回

字段名| 类型 | 说明
-------- | ----- | -----
evaluation_list| []list_evaluation| 评价列表

list_evalution

字段名| 类型 | 说明
-------- | ----- | -----
 evaluation_id  | int | 评价的id 
 buyer_id       | int | 评论人 
 buyer_name     | string | 评论人 
commodity_id| int| 商品id
commodity_name| string| 商品名
 content        | string| 商品评价

## delete evaluation

### 请求

请求类型| path
-------- | -----
post| /api/manager/delete_evaluation

字段名| 类型 | 说明
-------- | ----- | -----
evaluation_id|int | 评论的id 

### 返回

无data，成功code=1，错误code=0,错误信息返回在msg里

## ban buyer

### 请求

请求类型| path
-------- | -----
post| /api/manager/ban_buyer 

字段名| 类型 | 说明
-------- | ----- | -----
buyer_id|int| 买家id 

### 返回

无data，成功code=1，错误code=0，错误信息返回在msg里

## search seller id

### 说明

配合ban seller使用

### 请求

| 请求类型 | path                          |
| -------- | ----------------------------- |
| post     | /api/manager/search_seller_id |

| 字段名    | 类型   | 说明   |
| --------- | ------ | ------ |
| shop_name | string | 店铺名 |

### 返回

| 字段名    | 类型 | 说明   |
| --------- | ---- | ------ |
| seller_id | int  | 卖家id |

## ban seller

### 请求

请求类型| path
-------- | -----
post| /api/manager/ban_seller 

字段名| 类型 | 说明
-------- | ----- | -----
seller_id|int| 卖家id 

### 返回

无data，成功code=1，错误code=0，错误信息返回在msg里