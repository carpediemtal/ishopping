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
