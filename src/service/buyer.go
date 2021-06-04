package service

import (
	"ishopping/src/db"
	"time"
)

type Buyer struct {
	Uid      int    `json:"uid" map:"uid"`
	Name     string `json:"name" map:"name"`
	Address  string `json:"address" map:"address"`
	PhoneNum string `json:"phone_num" map:"phone_num"`
}

func UpdateBuyerInfo(uid int, name, address, phoneNum string) (err error) {
	_, err = GetBuyerProfileById(uid)
	// 如果买家的信息已经存在
	if err == nil {
		_, err = db.DB.Exec(`update buyer set name = ?, address = ?, phone_num = ? where uid = ?`, name, address, phoneNum, uid)
	} else {
		_, err = db.DB.Exec(`insert into buyer (uid, name, address, phone_num) values (?, ?, ?, ?)`, uid, name, address, phoneNum)
	}
	return
}

func GetBuyerProfileById(uid int) (buyer Buyer, err error) {
	row := db.DB.QueryRow("select * from buyer where uid = ?", uid)
	err = row.Scan(&buyer.Uid, &buyer.Name, &buyer.Address, &buyer.PhoneNum)
	return
}

func BuyerEvaluateCommodity(uid, order_id, rate int, content string) (err error) {
	var cid int
	row := db.DB.QueryRow("select cid from purchase_order where oid = ?", order_id)
	err = row.Scan(&cid)
	if err != nil {
		return err
	}
	timestamp := time.Now().Unix()
	_, err = db.DB.Exec(`insert into comment (cid, uid, content, timestamp, rate) values (?, ?, ?, ?, ?)`, cid, uid, content, timestamp, rate)
	return
}

func CommoditySignFor(oid int) (err error) {
	_, err = db.DB.Exec(`update purchase_order set status = ?, modify_time = ? where oid = ?`, 3, time.Now().Unix(), oid)
	return
}

func CommodityUpdateSales(oid int) (err error) {
	row := db.DB.QueryRow(`select cid, count from purchase_order where oid = ?`, oid)
	var cid, count int
	if err = row.Scan(&cid, &count); err != nil {
		return
	}

	row = db.DB.QueryRow(`select sales from commodity where cid = ?`, cid)
	var sales int
	if err = row.Scan(&sales); err != nil {
		return
	}

	_, err = db.DB.Exec(`update commodity set sales = ? where cid = ?`, sales+count, cid)
	return
}
