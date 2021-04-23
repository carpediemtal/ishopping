package service

import (
	"database/sql"
	"errors"
	"ishopping/src/db"
)

var page_max_limits int = 15 //页面最大展示商品数量

type SearchedCommodity struct{
	Cid       int     `json:"cid" map:"cid"`
	Name      string  `json:"name" map:"name"`
	Price     float64 `json:"price" map:"price"`
	Picture_path   string  `json:"picture_path" map:"picture_path"`
}

//通过cid与meta_key查找meta_vul
func GetCommodityExtraInfoByCid(cid int, meta_key string )(meta_val string, err error){
	row := db.DB.QueryRow("select meta_val from commodity_meta where cid = ? and meta_key = ?", cid, meta_key)
	err = row.Scan(&meta_val)
	return
}

//获取某个分类下对应页数的商品列表
func GetCommodityListByCategoryAndPage(caid int, page int)(commoditylist []SearchedCommodity, err error){
	var rows *sql.Rows
	if caid ==0{
		rows, err = db.DB.Query("select cid,name,price from commodity limit ?, ?",(page-1)*page_max_limits, page_max_limits)
	}else{
		rows, err = db.DB.Query("select cid,name,price from commodity where caid = ? limit ?, ?", caid,(page-1)*page_max_limits, page_max_limits)
	}
	if err != nil{
		return
	}
	var commodity SearchedCommodity
	for rows.Next(){
		err = rows.Scan(&commodity.Cid, &commodity.Name, &commodity.Price)
		if err != nil{
			return
		}
		commodity.Picture_path, err = GetCommodityExtraInfoByCid(commodity.Cid, "picture_path")
		if err != nil{
			if err.Error() == "sql: no rows in result set"{
				commodity.Picture_path = "defult"
				err = nil
			}else{
				return
			}
		}
		commoditylist = append(commoditylist, commodity)
	}
	if len(commoditylist) == 0{
		err = errors.New("no commodity found")
	}
	return
}

