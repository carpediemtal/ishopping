package service

import (
	"errors"
	"ishopping/src/db"
	"time"
)

type Cart struct {
	Uid   int `json:"uid" map:"uid"`
	Cid   int `json:"cid" map:"cid"`
	Count int `json:"count" map:"count"`
}

type CartCommodity struct {
	Cid       int     `json:"commodity_id" map:"commodity_id"`
	Count     int     `json:"count" map:"count"`
	Name      string  `json:"commodity_name" map:"commodity_name"`
	Price     float64 `json:"price" map:"price"`
	Thumbnail string  `json:"thumbnail" map:"thumbnail"`
}

type OrderHistory struct {
	Oid        int    `json:"order_id" map:"oid" db:"oid"`
	Status     int    `json:"order_status" map:"status" db:"status"`
	Cid        string `json:"commodity_id" map:"cid" db:"cid"`
	CreateTime int    `json:"create_time" map:"create_time" db:"create_time"`
	ModifyTime int    `json:"modify_time" map:"modify_time" db:"modify_time"`
	Count      int    `json:"item_num" map:"count" db:"count"`
	ShopName   string `json:"shop_name" map:"shop_name" db:"shop_name"`
	Price      string `json:"pay_price" map:"pay_price" db:"pay_price"`
	Img        string `json:"item_img" map:"item_img" db:"item_img"`
}

func CartAddCommodity(uid, cid, count int) (err error) {
	inventory, _ := getInventoryByCid(cid)
	var nowCount int
	nowCount, err = GetCountById(uid, cid) //获取当前购物车中 该cid 商品的数量
	if err == nil {
		var caCount = nowCount + count
		if caCount < 0 {
			return errors.New("the count of commodity can't be a negative number")
		} else if caCount == 0 {
			err = CartDeleteCommodity(uid, cid)
		} else if caCount > inventory {
			return errors.New("the count of commodity is larger than the inventory of the commodity")
		} else {
			_, err = db.DB.Exec(`update cart set count = ? where cid = ? and uid = ?`, caCount, cid, uid)
		}
	} else { //如果不存在当前cid的记录
		if count <= 0 { //如果添加数量为负数，则不准添加进购物车
			return errors.New("the count can't be a negative or a zero")
		} else if inventory < count {
			return errors.New("the count is larger than the inventory of the commodity")
		} else {
			_, err = db.DB.Exec(`insert into cart (cid, uid, count) values (?, ?, ?)`, cid, uid, count)
		}
	}
	return err
}

func GetCountById(uid, cid int) (count int, err error) {
	row := db.DB.QueryRow("select count from cart where uid = ? and cid = ?", uid, cid)
	err = row.Scan(&count)
	return
}

func ConfirmCartAndCommodity(uid, cid int) (err error) {
	count, err := GetCountById(uid, cid)
	if err != nil {
		return errors.New("no this record")
	}
	inventory, err := getInventoryByCid(cid)
	if err != nil {
		return errors.New("don't have this commodity")
	}
	if count > inventory {
		return errors.New("in a short inventory")
	}
	return
}

func CartToOrder(uid, cid int) (err error) {
	count, err := GetCountById(uid, cid)
	if err != nil {
		return errors.New("no this record")
	} else {
		timestamp := time.Now().Unix()
		_, err = db.DB.Exec(`insert into purchase_order (status,uid, cid, create_time,modify_time,count) values (?, ?, ?, ?, ?, ?)`, 1, uid, cid, timestamp, timestamp, count)
		if err != nil {
			return errors.New("can't insert into order")
		}
		err = DecreaseInventoryByCid(cid, count)
		if err != nil {
			return err
		}
		err = CartDeleteCommodity(uid, cid)
		if err != nil {
			return err
		}
	}
	return
}

func CartDeleteCommodity(uid, cid int) (err error) {
	_, err = GetCountById(uid, cid)
	if err != nil {
		return errors.New("no this record")
	} else {
		_, err = db.DB.Exec(`delete from cart where cid = ? and uid = ?`, cid, uid)
	}
	return
}

func CartGetCommodityByUid(uid int) (results []CartCommodity, err error) {
	if err = db.DB.Select(&results, `select cid, count from ub_cart where uid = ? order by cartid desc`, uid); err != nil {
		return
	}
	for i, commodity := range results {
		row1 := db.DB.QueryRow(`select name, price from ub_commodity where cid = ?`, commodity.Cid)
		//err = row1.Scan(&commodity.Name, &commodity.Price)
		if err = row1.Scan(&results[i].Name, &results[i].Price); err != nil {
			return results, errors.New("no this commodity")
		}
		row2 := db.DB.QueryRow(`select meta_val from commodity_meta where cid = ? and meta_key = ?`, commodity.Cid, "thumbnail")
		if err = row2.Scan(&results[i].Thumbnail); err != nil {
			results[i].Thumbnail = Image404
		}
	}
	return results, nil
}

func GetOrderHistoryByUid(uid int) (results []OrderHistory, err error) {
	if err = db.DB.Select(&results, `select oid, status, cid, create_time, modify_time, count from purchase_order where uid = ? order by oid desc`, uid); err != nil {
		return
	}
	for i, commodity := range results {
		row1 := db.DB.QueryRow(`select price, sid from ub_commodity where cid = ?`, commodity.Cid)
		var sid int
		if err = row1.Scan(&results[i].Price, &sid); err != nil {
			return results, errors.New("no this commodity")
		}
		row2 := db.DB.QueryRow(`select shop_name from shop where sid = ?`, sid)
		if err = row2.Scan(&results[i].ShopName); err != nil {
			return results, errors.New("no this shop")
		}
		row3 := db.DB.QueryRow(`select meta_val from commodity_meta where cid = ? and meta_key = ?`, commodity.Cid, "thumbnail")
		if err = row3.Scan(&results[i].Img); err != nil {
			results[i].Img = Image404
		}
	}
	return results, nil
}
