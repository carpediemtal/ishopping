package service

import (
	"errors"
	"ishopping/src/db"
	"log"
)

type Commodity struct {
	Cid       int     `json:"cid" map:"cid"`
	Sid       int     `json:"sid" map:"sid"`
	Name      string  `json:"name" map:"name"`
	Price     float64 `json:"price" map:"price"`
	Sales     int     `json:"sales" map:"sales"`
	Inventory int     `json:"inventory" map:"inventory"`
	Caid      int     `json:"caid" map:"caid"`
}

func GetCommodityProfileByCid(cid int) (commodity Commodity, err error) {
	row1 := db.DB.QueryRow("select * from commodity where cid = ?", cid)
	err = row1.Scan(&commodity.Cid, &commodity.Sid, &commodity.Name, &commodity.Price, &commodity.Sales, &commodity.Inventory, &commodity.Caid)
	return
}

func GetCommodities() (commodities []Commodity, err error) {
	err = db.DB.Select(&commodities, `select * from commodity`)
	return
}

// 添加商品成功时返回商品ID
func AddCommodity(uid int, price float64, inventory int, name, introduction string) (cid int, err error) {
	sid, err := getSidByUid(uid)
	if err != nil {
		return 0, errors.New("no matched shop id")
	}

	// TODO: sales 和 caid 先插0，以后的版本再改
	_, err = db.DB.Exec(`insert into commodity (sid, name, price, sales, inventory, caid) VALUES (?, ?, ?, ?, ?, ?)`, sid, name, price, 0, inventory, 0)
	if err != nil {
		return 0, err
	}

	cid, err = getCidBySid(sid)
	if err != nil {
		return 0, err
	}
	_, err = db.DB.Exec(`insert into commodity_meta (cid, meta_key, meta_val) values (?, ?, ?)`, cid, "introduction", introduction)
	return cid, err
}

func getSidByUid(uid int) (sid int, err error) {
	log.Println("uid:", uid)
	row := db.DB.QueryRow(`select sid from shop where uid = ?`, uid)
	err = row.Scan(&sid)
	return
}

func getCidBySid(sid int) (cid int, err error) {
	row := db.DB.QueryRow(`select cid from commodity where sid = ?`, sid)
	err = row.Scan(&cid)
	return
}

type Order struct {
	Oid   int     `json:"oid" map:"oid"`
	Price float64 `json:"total_price" map:"total_price"`
}

func GetOrderListByStatus(status, userID int) (orderList []Order, err error) {
	err = db.DB.Select(&orderList, `select oid, price from commodity, purchase_order where commodity.cid = purchase_order.cid && status = ? && uid = ?`, status, userID)
	return
}
