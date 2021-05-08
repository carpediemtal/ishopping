[toc]

# 说明

- 对于接口有问题的、某些功能问题的、有一些good idea的，在大群里说或者私聊组长。
- release2开发周期：**第6周** 至 **第8周**，请尽快。
- release1的add接口合并到了seller 的 commodity_edit里
- 最新更新时间：**2021-04-24 10:45**
- 2021-04-17 23:30修改：Visitor index&category channel
- 2021-04-18 20:30新增：Seller category_list
- 2021-04-23 21:00删除：Seller order_delivery的物流单号删除，接口负责人请跟进
- 2021-04-23 21:40修改：Seller commodity_edit的商品介绍单词拼错了，已订正
- 2021-04-23 21:50修改：Seller category_list接口定义错了，已订正
- 2021-04-24 10:45修改：Seller commodity_edit、Visitor index&category channel新增thumbnail
- 2021-5-8 11:11:53修改：修改 /api/commodity_detail 接口返回的内容

## 待完成

- 主页、分类页（分类页有选页功能）
- 买家的个人信息预览、个人信息修改
- 买家购买商品时的确认页面（类比手机淘宝流程）
- 卖家的店铺信息预览、个人信息修改
- 卖家的商品添加、修改页面（release1的添加接口有变动）
- 卖家的待处理订单列表
- 卖家的发货（可以在待处理订单列表里加个按钮，点击时请求接口进行发货；也可以自行商议具体页面……）

# Visitor（不需要登录就可以查看）

## 1. index&category channel

### 请求

| 请求类型 | path                        |
| -------- | --------------------------- |
| get      | /api/index_category_channel |

| 字段名      | 类型 | 说明                            |
| ----------- | ---- | ------------------------------- |
| category_id | int  | 分类id，主页为0，否则为分类的id |
| page_num    | int  | 页数                            |

### 返回

| 字段名         | 类型        | 说明     |
| -------------- | ----------- | -------- |
| commodity_list | []list_item | 商品列表 |

list_item:

| 字段名       | 类型   | 说明       |
| ------------ | ------ | ---------- |
| commodity_id | int    | 商品id     |
| name         | string | 商品名     |
| price        | float  | 商品价格   |
| thumbnail    | string | 商品缩略图 |

## 2. Detail

### 请求

| 请求类型 | path                  |
| -------- | --------------------- |
| get      | /api/commodity_detail |

| 字段名 | 类型 | 说明   |
| ------ | ---- | ------ |
| cid    | int  | 商品id |

### 返回

| 字段名       | 类型                   | 说明       |
| ------------ | ---------------------- | ---------- |
| name         | string                 | 商品名称   |
| inventory    | int                    | 商品库存   |
| sales        | int                    | 商品销量   |
| price        | float                  | 商品价格   |
| introduction | string                 | 商品介绍   |
| thumbnail    | string                 | 商品缩略图 |
| images       | []string（字符串数组） | 商品图片   |



# Buyer

## 1.information

### 请求

| 请求类型 | path          |
| -------- | ------------- |
| get | /api/buyer/information |

### 返回

| 字段名      | 类型   | 说明     |
| ----------- | ------ | -------- |
| name        | string | 买家姓名 |
| address     | string | 地址     |
| phone_num | string | 电话     |

## 2.information_modify

### 请求

| 请求类型 | path                          |
| -------- | ----------------------------- |
| post     | /api/buyer/information_modify |

| 字段名      | 类型   | 说明      |
| ----------- | ------ | -------- |
| uid         | int    | 买家id   |
| name        | string | 买家姓名 |
| address     | string | 地址     |
| phone_num   | string | 电话     |

### 返回

无data，成功code=0，错误code=-1，错误信息返回在msg里

## 3.commodity_buy

### 请求

| 请求类型 | path       |
| -------- | ---------- |
| post | /api/buyer/commodity_buy |

| 字段名       | 类型 | 说明   |
| :----------- | ---- | ------ |
| commodity_id | int  | 商品id |


### 返回

| 字段名   | 类型 | 说明   |
| :------- | ---- | ------ |
| order_id | int  | 订单id |

# Seller

## 1.shop_information

### 请求

| 请求类型 | path                         |
| -------- | ---------------------------- |
| get      | /api/seller/shop_information |

### 返回

| 字段名             | 类型   | 说明     |
| ------------------ | ------ | -------- |
| shopname           | string | 店铺名   |
| address            | string | 发货地址 |
| seller_phonenumber | string | 卖家电话 |

## 2.shop_information_modify

### 请求

| 请求类型 | path                                |
| -------- | ----------------------------------- |
| post     | /api/seller/shop_information_modify |

| 字段名             | 类型   | 说明     |
| ------------------ | ------ | -------- |
| shopname           | string | 店铺名   |
| address            | string | 发货地址 |
| seller_phonenumber | string | 卖家电话 |

### 返回

无data，成功code=0，错误code=-1，错误信息返回在msg里

## 3.category_list

### 请求

| 请求类型 | path                      |
| :------- | :------------------------ |
| get      | /api/seller/category_list |

### 返回

| 字段名        | 类型            | 说明       |
| ------------- | --------------- | ---------- |
| category_list | []category_item | 所有的分类 |

category_item：

| 字段名        | 类型   | 说明   |
| ------------- | ------ | ------ |
| category_id   | int    | 分类id |
| category_name | string | 分类名 |

## 4.commodity_edit

### 请求

| 请求类型 | path              |
| -------- | ----------------- |
| post     | /api/seller/commodity_edit |

|字段名| 类型  | 说明 |
| -------------- | ------ | ---- |
| category_id | int | 分类id，可通过请求category_list获取所有分类 |
| commodity_id | int | 商品id（增加时传0，否则传商品id） |
| commodity_name | string | 商品名 |
| inventory | int | 商品库存 |
| introduction | string | 商品介绍 |
| price | float | 商品价格 |
| thumbnail | string | 商品缩略图 |
| image | []string | 商品图片 |
| edit_type | int | 0表示修改，1表示增加 |

### 返回
无data，成功code=0，失败code=-1，错误信息返回在msg里

### 说明

**add功能已被合并到edit**


## 5.order_list

### 请求

| 请求类型 | path              |
| -------- | ----------------- |
| get   | /api/seller/order_list |

| 字段名       | 类型 | 说明                                       |
| ------------ | ---- | ------------------------------------------ |
| order_status | int  | 订单状态（1表示未发货，2表示已发货未签收） |

### 返回

| 字段名 | 类型   | 说明                       |
| ------ | ------ | -------------------------- |
| order_list | []order |订单列表 |

order：

| 字段名      | 类型  | 说明     |
| ----------- | ----- | -------- |
| order_id    | int   | 订单号   |
| total_price | float | 总价格   |
| ……          |       | 想到再加 |

## 6.order_delivery

### 请求

| 请求类型 | path                       |
| -------- | -------------------------- |
| post     | /api/seller/order_delivery |

| 字段名   | 类型 | 说明     |
| -------- | ---- | -------- |
| order_id | int  | 订单单号 |

### 返回

无data，成功code=0，失败code=-1，错误信息返回在msg里



# User 

## 1.userType

### 请求

| 请求类型 | path          |
| -------- | ------------- |
| post     | /api/userType |

### 返回

| 字段名 | 类型 | 说明                    |
| ------ | ---- | ----------------------- |
| type   | int8 | 用户类型（1卖家 0买家） |

返回type，成功code=0，失败code=-1，错误信息返回在msg里

