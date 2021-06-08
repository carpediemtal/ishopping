package service

import (
	"errors"
	"ishopping/src/db"
)

type Order struct {
	Oid     int     `json:"oid" db:"oid"`
	Cid     int     `json:"commodity_id" db:"cid"`
	Price   float64 `json:"total_price" db:"price"`
	Name    string  `json:"buyer_name" db:"name"`
	Phone   string  `json:"buyer_phone" db:"phone_num"`
	Address string  `json:"buyer_address" db:"address"`
	Count   int     `json:"count" db:"count"`
}

func GetOrderListByStatus(status, userID int) (orderList []Order, err error) {
	sid, err := getSidByUid(userID)
	if err != nil {
		return
	}
	err = db.DB.Select(&orderList, `select oid, commodity.cid, price, buyer.name, buyer.phone_num, buyer.address, count from purchase_order, commodity, buyer where purchase_order.cid = commodity.cid and purchase_order.uid = buyer.uid and status = ? and sid = ?`, status, sid)
	if err != nil {
		return
	}

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
