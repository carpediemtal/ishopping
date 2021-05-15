# 说明

-   对于接口有问题的，在大群里说或私聊群里的大佬们。
-   release3 开发周期：**第 9 周** 至 **第 12 周**，请尽快。
-   Important：`请不要随意更改接口文档内容`。
-   Important：请严格按照接口文档进行数据交换，不要充耳不闻，比如：
    -   商品 id 时 commodity_id，而不是 cid；
    -   价格是 float 类型，而不是 int。
-   05-13 21:10：`ban buyer`和`ban seller` 合并为`ban user`
-   05-12 22:20：修改`visitor commodity detail`入参、返回
-   05-14 10:48：修改用户类型`manager`为`admin`，及修改相关接口的path
-   05-14 10:48：增加`unban`接口
-  05-15 14:35：修改`order history `接口的返回字段 `paid_time` 为 `modify_time `
-  05-15 15:14：修改`admin search seller id`请求方法为`get`
-  05-15 15:39：修改`buyer cart get`入参、`buyer cart delete`返回值
-  05-15 15:52：新增`buyer cart confirm`接口


# Visitor

## search

### 说明

把 release1 的搜索合并到该接口，该接口可以实现搜索特定价格区间的商品。

### 请求

| 请求类型 | path        |
| -------- | ----------- |
| get      | /api/search |

| 字段名     | 类型   | 说明               |
| ---------- | ------ | ------------------ |
| content    | string | 搜索框内内容       |
| page_index | id     | 第几页（从零开始） |
| page_size  | int    | 每页的记录数       |
| price_min  | float  | 最低价格           |
| price_max  | float  | 最高价格           |

### 返回

| 字段名 | 类型        | 说明     |
| ------ | ----------- | -------- |
| list   | []list_item | 商品列表 |

list_item

| 字段名    | 类型   | 说明       |
| --------- | ------ | ---------- |
| cid       | int    | 商品 id    |
| name      | string | 商品名     |
| price     | float  | 商品价格   |
| thumbnail | string | 商品缩略图 |
| sales     | int    | 销量       |

查询越界时返回错误信息：“page out of range”

## commodity detail

### 说明

该接口在release2已有，但要规范了请求字段

### 请求

| 请求类型 | path                  |
| -------- | --------------------- |
| get      | /api/commodity_detail |

| 字段名 | 类型 | 说明   |
| ------ | ---- | ------ |
| commodity_id    | int  | 商品id |

### 返回

| 字段名       | 类型                   | 说明       |
| ------------ | ---------------------- | ---------- |
| name         | string                 | 商品名称   |
| inventory    | int                    | 商品库存   |
| sales        | int                    | 商品销量   |
| price        | float                  | 商品价格   |
| introduction | string                 | 商品介绍   |
| thumbnail    | string                 | 商品缩略图 |
| images       | []string | 商品图片   |

## commodity evaluation list

### 说明

在详情页用这个接口

### 请求

| 请求类型 | path                      |
| -------- | ------------------------- |
| get      | /api/commodity_evaluation |

| 字段名       | 类型 | 说明    |
| ------------ | ---- | ------- |
| commodity_id | int  | 商品 id |

### 返回

| 字段名 | 类型        | 说明     |
| ------ | ----------- | -------- |
| list   | []list_item | 评价列表 |

list_item

| 字段名    | 类型   | 说明           |
| --------- | ------ | -------------- |
| rate      | int    | 评分 1-5       |
| content   | string | 商品评价       |
| timestamp | int    | 评价当时时间戳 |

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
| modify_time  | string | 订单最近状态变更的时间 |
| pay_price    | string | 买家确认收货的付款金额 |
| order_id     | int    | 订单 id                |
| item_img     | string | 商品图片               |
| item_num     | int    | 商品数量               |
| commodity_id | int    | 商品 id                |
| shop_name    | string | 店铺名称               |
| order_status | int    | 订单状态               |

## evaluate

### 说明

用户评价

### 请求

| 请求类型 | path                |
| -------- | ------------------- |
| post     | /api/buyer/evaluate |

| 字段名   | 类型   | 说明     |
| -------- | ------ | -------- |
| order_id | int    | 订单 id  |
| rate     | int    | 评分 1-5 |
| content  | string | 评价内容 |

### 返回

无 data，添加成功 code=0，错误 code=-1，错误信息返回在 msg 里

## cart add

### 说明

在商品详情页点击加入购物车时调用本接口

### 请求

| 请求类型 | path                |
| -------- | ------------------- |
| post     | /api/buyer/cart_add |

| 字段名       | 类型 | 描述       |
| ------------ | ---- | ---------- |
| commodity_id | int  | 商品 id    |
| count        | int  | 购买的数量 |

### 返回

无 data，添加成功 code=0，错误 code=-1，错误信息返回在 msg 里

## cart delete

### 说明

购物车里删除商品

### 请求

| 请求类型 | path                   |
| -------- | ---------------------- |
| post     | /api/buyer/cart_delete |

| 字段名       | 类型 | 描述          |
| ------------ | ---- | ------------- |
| commodity_id | int  | 要删除商品 id |

### 返回

无 data，添加成功 code=0，错误 code=-1，错误信息返回在 msg 里

## cart get

### 说明

在我的购物车页面调用

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

list 结构

| Parameter      | 类型   | 描述       |
| -------------- | ------ | ---------- |
| commodity_id   | int    | 商品 id    |
| commodity_name | string | 商品名称   |
| thumbnail      | string | 缩略图     |
| price          | double | 价格       |
| count          | int    | 购买的数量 |

## cart confirm

### 说明

购物车商品下单

### 请求

| 请求类型 | path                    |
| -------- | ----------------------- |
| post     | /api/buyer/cart_confirm |

| 字段名       | 类型  | 描述              |
| ------------ | ----- | ----------------- |
| commodity_id | []int | 要删除商品 id列表 |

### 返回

无 data，添加成功 code=0，错误 code=-1，错误信息返回在 msg 里

# Seller

## shop's commodity list

### 说明

卖家的全部商品列表

### 请求

| 请求类型 | path                       |
| -------- | -------------------------- |
| get      | /api/seller/commodity_list |

| 字段名    | 类型 | 说明     |
| --------- | ---- | -------- |
| page      | int  | 页数     |
| page_size | int  | 每页个数 |

### 返回

| 字段名         | 类型   | 说明     |
| -------------- | ------ | -------- |
| commodity_list | []item | 评价列表 |

item

| 字段名          | 类型   | 说明       |
| --------------- | ------ | ---------- |
| commodity_id    | id     | 商品 id    |
| name            | string | 商品名     |
| commodity_price | float  | 商品价格   |
| thumbnail       | string | 商品缩略图 |
| sales           | int    | 销量       |

## commodity detail

### 说明

在编辑商品前，需先获取到之前添加的商品详情，并渲染到前端里。

### 请求

| 请求类型 | path                         |
| -------- | ---------------------------- |
| get      | /api/seller/commodity_detail |

| 字段名       | 类型 | 说明    |
| ------------ | ---- | ------- |
| commodity_id | int  | 商品 id |

### 返回

| 字段名         | 类型     | 说明       |
| -------------- | -------- | ---------- |
| category_id    | int      | 分类 id    |
| commodity_name | string   | 商品名     |
| inventory      | int      | 商品库存   |
| introduction   | string   | 商品介绍   |
| price          | float    | 商品价格   |
| thumbnail      | string   | 商品缩略图 |
| image          | []string | 商品图片   |

# Admin

## admin login

### 请求

| 请求类型 | path             |
| -------- | ---------------- |
| POST     | /api/admin/login |

| 字段名   | 类型   |
| -------- | ------ |
| username | string |
| password | string |

### 返回

成功 code=0，错误 code=-1，错误信息在 msg 里。成功时 data 嵌套数据如下：

| 字段名 | 类型   | 说明                                     |
| ------ | ------ | ---------------------------------------- |
| token  | string | 身份 token，前端把它设置 cookie 名为 jwt |

## all evaluation list

### 说明

请求该接口后配合使用 delete evaluation、ban user

### 请求

| 请求类型 | path                       |
| -------- | -------------------------- |
| get      | /api/admin/evaluation_list |

| 字段名    | 类型 | 说明     |
| --------- | ---- | -------- |
| page      | int  | 页数     |
| page_size | int  | 每页个数 |

### 返回

| 字段名          | 类型              | 说明     |
| --------------- | ----------------- | -------- |
| evaluation_list | []list_evaluation | 评价列表 |

list_evalution

| 字段名         | 类型   | 说明      |
| -------------- | ------ | --------- |
| evaluation_id  | int    | 评价的 id |
| buyer_id       | int    | 评论人    |
| buyer_name     | string | 评论人    |
| commodity_id   | int    | 商品 id   |
| commodity_name | string | 商品名    |
| content        | string | 商品评价  |

## delete evaluation

### 请求

| 请求类型 | path                         |
| -------- | ---------------------------- |
| post     | /api/admin/delete_evaluation |

| 字段名        | 类型 | 说明      |
| ------------- | ---- | --------- |
| evaluation_id | int  | 评论的 id |

### 返回

无 data，成功 code返回0值，失败时返回非0值，失败原因在 msg 里

## search seller id

### 说明

配合 ban user 使用

### 请求

| 请求类型 | path                        |
| -------- | --------------------------- |
| get      | /api/admin/search_seller_id |

| 字段名    | 类型   | 说明   |
| --------- | ------ | ------ |
| shop_name | string | 店铺名 |

### 返回

| 字段名    | 类型 | 说明      |
| --------- | ---- | --------- |
| seller_id | int  | 卖家的 id |

## ban

### 请求

| 请求类型 | path           |
| -------- | -------------- |
| post     | /api/admin/ban |

| 字段名  | 类型 | 说明    |
| ------- | ---- | ------- |
| user_id | int  | 用户 id |

### 返回

无 data，成功 code返回0值，失败时返回非0值，失败原因在 msg 里

## unban

### 请求

| 请求类型 | path             |
| -------- | ---------------- |
| post     | /api/admin/unban |

| 字段名  | 类型 | 说明    |
| ------- | ---- | ------- |
| user_id | int  | 用户 id |

### 返回

无 data，成功 code返回0值，失败时返回非0值，失败原因在 msg 里