package main

//
//import (
//	"errors"
//	"fmt"
//	"github.com/gin-gonic/gin"
//	_ "github.com/go-sql-driver/mysql"
//	"github.com/jmoiron/sqlx"
//	"log"
//)
//
//type Commodity struct {
//	Uid           int     `db:"uid" json:"user_id"`
//	Sid           int     `db:"sid" json:"sid"`
//	Name          string  `db:"name" json:"commodity_name"`
//	Price         float32 `db:"price" json:"price"`
//	Inventory     int     `db:"inventory" json:"inventory"`
//	Caid          int64   `db:"caid" json:"caid"`
//	Introdunction string  `db:"meta_val" json:"introduction"`
//	Cid           int64   `db:"cid" json:"cid"`
//}
//
//func AddCommodityToDb(commodity *Commodity, db *sqlx.DB) (int64, error) {
//	tx := db.MustBegin()
//	sql := "insert into commodity (sid,name,price,inventory,caid) values (:sid, :name, :price, :inventory, :caid);"
//	result, err := tx.NamedExec(sql, commodity)
//	cid, _ := result.LastInsertId()
//	return cid, err
//}
//func AddCommodityExtraInfomation(commodity *Commodity, db *sqlx.DB) error {
//	tx := db.MustBegin()
//	sql := "insert into commodity_meta (cid, meta_key, meta_val) values (:cid, 'introduction', :meta_val);"
//	_, err := tx.NamedExec(sql, commodity)
//	return err
//}
//
//func GetSidByUid(uid int, db *sqlx.DB) int {
//	sql := "select sid from shop where uid = ?"
//	rows, _ := db.Query(sql, uid)
//	var sid int
//	rows.Next()
//	rows.Scan(&sid)
//	return sid
//}
//
//func GetCaidByCategoryName(name string, db *sqlx.DB) int {
//	sql := "select caid from category where name = ?"
//	rows, _ := db.Query(sql, name)
//	var caid int
//	rows.Next()
//	rows.Scan(&caid)
//	return caid
//}
//
//func AddCommodity(commodity_name, introduction string, price float32, user_id, inventory int) (int64, error) {
//	dbuser := "root"
//	dbpwd := "root"
//	dbhost := "127.0.0.1"
//	dbname := "shopping"
//	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", dbuser, dbpwd, dbhost, dbname)
//	db, err := sqlx.Open("mysql", dns)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	defer db.Close()
//	sid := GetSidByUid(user_id, db)
//	if sid == 0 {
//		return 0, errors.New("No seller found")
//	}
//	commodity := &Commodity{
//		Sid:           sid,
//		Name:          commodity_name,
//		Price:         price,
//		Inventory:     inventory,
//		Introdunction: introduction,
//	}
//	cid, err := AddCommodityToDb(commodity, db)
//	if err != nil {
//		return 0, err
//	}
//	commodity.Cid = cid
//	return cid, AddCommodityExtraInfomation(commodity, db)
//}
//
//func AddCommodityHandle() {
//	r := gin.Default()
//	r.Use(func(c *gin.Context) {
//		c.Header("Access-Control-Allow-Origin", "*")
//	})
//
//	r.POST("/api/commodity_add", func(c *gin.Context) {
//		req := Commodity{}
//		err := c.BindJSON(&req)
//		if err != nil {
//			JsonErr(c, "发生XXX错误")
//			return
//		}
//		cid, err := AddCommodity(req.Name, req.Introdunction, req.Price, req.Uid, req.Inventory)
//		if err != nil {
//			JsonErr(c, err.Error())
//			return
//		}
//		JsonOK(c, gin.H{"cid": cid})
//	})
//
//	r.Run(":7000")
//}
