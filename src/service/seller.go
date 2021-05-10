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
	row := db.DB.QueryRow(`select status, cid from purchase_order where oid = ?`, oid)
	var status, cid int
	err := row.Scan(&status, &cid)
	if err != nil {
		return err
	}

	if status != 1 {
		return errors.New("this order had been delivered")
	}

	order_sid, err := getSidByCid(cid)
	if err != nil {
		return err
	}

	user_sid, err := getSidByUid(uid)
	if err != nil {
		return err
	}
	if order_sid != user_sid {
		return errors.New("can't deliver order not belogng to you")
	}

	_, err = db.DB.Exec(`update purchase_order set status = 2 where oid = ?`, oid)
	return err
}
