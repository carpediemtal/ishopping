package service

import (
	"database/sql"
	"errors"
	"ishopping/src/db"
	"log"
	"time"
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

type CommodityDetail struct {
	Name         string   `json:"name" map:"name"`
	Inventory    int      `json:"inventory" map:"inventory"`
	Sales        int      `json:"sales" map:"sales"`
	Price        float64  `json:"price" map:"price"`
	Introduction string   `json:"introduction" map:"introduction"`
	Thumbnail    string   `json:"thumbnail" map:"thumbnail"`
	Images       []string `json:"images" map:"images"`
}

type CommodityEdit struct {
	Caid         int      `json:"category_id"`
	Cid          int      `json:"commodity_id"`
	Name         string   `json:"commodity_name"`
	Inventory    int      `json:"inventory"`
	Introduction string   `json:"introduction"`
	Price        float64  `json:"price"`
	Thumbnail    string   `json:"thumbnail"`
	Image        []string `json:"image"`
	EditType     int      `json:"edit_type"`
}

func GetCommodityByCid(cid int) (commodity Commodity, err error) {
	row1 := db.DB.QueryRow("select * from commodity where cid = ?", cid)
	err = row1.Scan(&commodity.Cid, &commodity.Sid, &commodity.Name, &commodity.Price, &commodity.Sales, &commodity.Inventory, &commodity.Caid)
	return
}

func GetCommodityDetailByCid(cid int) (detail CommodityDetail, err error) {
	row := db.DB.QueryRow(`select name, inventory, sales, price from commodity where cid = ?`, cid)
	err = row.Scan(&detail.Name, &detail.Inventory, &detail.Sales, &detail.Price)
	if err != nil {
		return detail, errors.New("commodity_detail searching error")
	}

	row = db.DB.QueryRow(`select meta_val from commodity_meta where cid = ? and meta_key = ?`, cid, "introduction")
	err = row.Scan(&detail.Introduction)
	if err != nil {
		log.Println("no introduction found", err)
		detail.Introduction = "There Is No Introduction"
	}

	row = db.DB.QueryRow(`select meta_val from commodity_meta where cid = ? and meta_key = ?`, cid, "thumbnail")
	err = row.Scan(&detail.Thumbnail)
	if err != nil {
		log.Println("no thumbnail found", err)
		detail.Thumbnail = "http://ishopping.gq/static/img/404.jpg"
	}

	err = db.DB.Select(&detail.Images, `select meta_val from commodity_meta where cid = ? and meta_key = ?`, cid, "image")
	if err != nil || len(detail.Images) == 0 {
		log.Println("no image found", err)
		detail.Images = []string{"http://ishopping.gq/static/img/404.jpg"}
	}

	return detail, nil
}

func GetCommodities() (commodities []Commodity, err error) {
	err = db.DB.Select(&commodities, `select * from commodity`)
	return
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

func getSidByCid(cid int) (sid int, err error) {
	row := db.DB.QueryRow(`select sid from commodity where cid = ?`, cid)
	err = row.Scan(&sid)
	return
}

type Category struct {
	Caid int    `json:"category_id"`
	Name string `json:"category_name"`
}

func GetAllCategories() (categories []Category, err error) {
	err = db.DB.Select(&categories, `select * from category`)
	return
}

func UpdateCommodityInfo(cm CommodityEdit) error {
	if _, err := db.DB.Exec(`update commodity set name = ?, price = ?, inventory = ?, caid = ?`, cm.Name, cm.Price, cm.Inventory, cm.Caid); err != nil {
		return err
	}

	if _, err := db.DB.Exec(`delete from commodity_meta where cid = ?`, cm.Cid); err != nil {
		return err
	}

	if _, err := db.DB.Exec(`insert into commodity_meta (cid, meta_key, meta_val) values (?, ?, ?)`, cm.Cid, "introduction", cm.Introduction); err != nil {
		return err
	}

	for _, image := range cm.Image {
		if _, err := db.DB.Exec(`insert into commodity_meta (cid, meta_key, meta_val) values (?, ?, ?)`, cm.Cid, "image", image); err != nil {
			return err
		}
	}

	return nil
}

func AddCommodity(cm CommodityEdit, uid int) error {
	sid, err := getSidByUid(uid)
	if err != nil {
		return errors.New("no matched shop id")
	}

	var res sql.Result
	if res, err = db.DB.Exec(`insert into commodity (sid, name, price, inventory, caid) VALUES (?, ?, ?, ?, ?)`, sid, cm.Name, cm.Price, cm.Inventory, cm.Caid); err != nil {
		return err
	}
	// 获取插入操作自动生成的主键
	cid, _ := res.LastInsertId()

	if _, err := db.DB.Exec(`insert into commodity_meta (cid, meta_key, meta_val) values (?, ?, ?)`, cid, "introduction", cm.Introduction); err != nil {
		return err
	}

	if _, err := db.DB.Exec(`insert into commodity_meta (cid, meta_key, meta_val) values (?, ?, ?)`, cid, "thumbnail", cm.Thumbnail); err != nil {
		return err
	}

	for _, image := range cm.Image {
		if _, err = db.DB.Exec(`insert into commodity_meta (cid, meta_key, meta_val) values (?, ?, ?)`, cid, "image", image); err != nil {
			return err
		}
	}

	return nil
}

type Pur_order struct {
	Oid       int `json:"oid" map:"oid"`
	Status    int `json:"status" map:"status"`
	Uid       int `json:"uid" map:"uid"`
	Cid       int `json:"cid" map:"cid"`
	TimeStamp int `json:"time" map:"time"`
}

func CreatPurchaseOrderByCid(uid int, cid int) (pur_order Pur_order, err error) {
	timestamp := time.Now().Unix()
	_, err = db.DB.Exec(`insert into purchase_order (status, uid, cid, timestamp) values (?, ?, ?, ?)`, 1, uid, cid, timestamp)
	row := db.DB.QueryRow("select * from purchase_order  where uid = ? and cid= ? order by oid desc", uid, cid)
	err = row.Scan(&pur_order.Oid, &pur_order.Status, &pur_order.Uid, &pur_order.Cid, &pur_order.TimeStamp)
	return
}

type list_item struct {
	Rate      int    `json:"rate" map:"rate"`
	Content   string `json:"content" map:"content"`
	Timestamp int    `json:"timestamp" map:"timestamp"`
}

func GetCommodityEvaluationListByCommodityId(cid int) (list []list_item, err error) {
	err = db.DB.Select(&list, `select rate, content, timestamp from comment where cid = ?`, cid)
	if len(list) == 0 {
		return list, errors.New("Commodity evaluation-No records selected")
	}
	return
}
