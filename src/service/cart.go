package service

import (
	"errors"
	"ishopping/src/db"
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
	Oid       int    `json:"order_id" map:"oid"`
	Status    int    `json:"order_status" map:"status"`
	Cid       string `json:"commodity_id" map:"cid"`
	Cre_time  int    `json:"create_time" map:"create_time"`
	Mod_time  int    `json:"modify_time" map:"modify_time"`
	Count     int    `json:"item_num" map:"item_num"`
	Shop_name string `json:"shop_name" map:"shop_name"`
	Price     string `json:"pay_price" map:"price"`
	Img       string `json:"item_img" map:"item_img"`
}

func CartAddCommodity(uid, cid, count int) (err error) {
	var nowCount int
	nowCount, err = GetCountById(uid, cid)
	if err == nil {
		var caCount = nowCount + count
		if caCount < 0 {
			return errors.New("the count of commodity can't be a negative number")
		} else if caCount == 0 {
			err = CartDeleteCommodity(uid, cid)
		} else {
			_, err = db.DB.Exec(`update cart set count = ? where cid = ? and uid = ?`, caCount, cid, uid)
		}
	} else {
		_, err = db.DB.Exec(`insert into cart (cid, uid, count) values (?, ?, ?)`, cid, uid, count)
	}
	return err
}

func GetCountById(uid, cid int) (count int, err error) {
	row := db.DB.QueryRow("select count from cart where uid = ? and cid = ?", uid, cid)
	err = row.Scan(&count)
	return
}

func CartDeleteCommodity(uid, cid int) (err error) {
	_, err = db.DB.Exec(`delete from cart where cid = ? and uid = ?`, cid, uid)
	return
}

func CartGetCommodityByUid(uid int) (results []CartCommodity, err error) {
	if err = db.DB.Select(&results, `select cid, count from cart where uid = ? order by cartid desc`, uid); err != nil {
		return
	}
	for i, commodity := range results {
		row1 := db.DB.QueryRow(`select name, price from commodity where cid = ?`, commodity.Cid)
		//err = row1.Scan(&commodity.Name, &commodity.Price)
		if err = row1.Scan(&results[i].Name, &results[i].Price); err != nil {
			return results, errors.New("no this commodity")
		}
		row2 := db.DB.QueryRow(`select meta_val from commodity_meta where cid = ? and meta_key = ?`, commodity.Cid, "thumbnail")
		if err = row2.Scan(&commodity.Thumbnail); err != nil {
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
		row1 := db.DB.QueryRow(`select price, sid from commodity where cid = ?`, commodity.Cid)
		var sid int
		if err = row1.Scan(&results[i].Price, &sid); err != nil {
			return results, errors.New("no this commodity")
		}
		row2 := db.DB.QueryRow(`select shop_name from shop where sid = ?`, sid)
		if err = row2.Scan(&results[i].Shop_name); err != nil {
			return results, errors.New("no this shop")
		}
		row3 := db.DB.QueryRow(`select meta_val from commodity_meta where cid = ? and meta_key = ?`, commodity.Cid, "thumbnail")
		if err = row3.Scan(&commodity.Img); err != nil {
			results[i].Img = Image404
		}
	}
	return results, nil
}

func CartConfimCommodity(uid int, cids []int) (err error) {
	for _, cid := range cids {
		err = CartDeleteCommodity(uid, cid)
		if err != nil {
			return errors.New("confim error")
		}
	}
	return
}
