package service

import (
	"errors"
	"ishopping/src/db"
)

type Order struct {
	Oid   int     `json:"oid" map:"oid"`
	Price float64 `json:"total_price" map:"total_price"`
}

func GetOrderListByStatus(status, userID int) (orderList []Order, err error) {
	sid, err := getSidByUid(userID)
	if err != nil {
		return
	}
	err = db.DB.Select(&orderList, `select oid, price from commodity, purchase_order where commodity.cid = purchase_order.cid and status = ? and sid = ?`, status, sid)
	return
}

func OrderDelivery(oid, uid int) error {
	row := db.DB.QueryRow(`select status from purchase_order where oid = ?`, oid)
	var status int
	err := row.Scan(&status)
	if err != nil {
		return err
	}

	if status == 2 {
		return errors.New("this order had been delivered")
	}

	_, err = db.DB.Exec(`update purchase_order set status = 2 where oid = ? and uid = ?`, oid, uid)
	return err
}
