package service

import (
	"database/sql"
	"errors"
	"ishopping/src/db"
)

const (
	PageMaxLimits = 15 //页面最大展示商品数量
)

type SearchedCommodity struct {
	Cid         int     `json:"commodity_id" map:"cid"`
	Name        string  `json:"name" map:"name"`
	Price       float64 `json:"price" map:"price"`
	Thumbnail string  `json:"thumbnail" map:"thumbnail"`
}
//
// 通过cid与meta_key查找meta_val
func GetCommodityExtraInfoByCid(cid int, metaKey string) (metaVal string, err error) {
	row := db.DB.QueryRow("select meta_val from commodity_meta where cid = ? and meta_key = ?", cid, metaKey)
	err = row.Scan(&metaVal)
	return
}

//获取某个分类下对应页数的商品列表
func GetCommodityListByCategoryAndPage(caid int, page int) (commodities []SearchedCommodity, err error) {
	var rows *sql.Rows
	if caid == 0 { // AllCategories
		rows, err = db.DB.Query("select cid, name, price from commodity limit ?, ?", (page-1)*PageMaxLimits, PageMaxLimits)
	} else { // Specific Category
		rows, err = db.DB.Query("select cid, name, price from commodity where caid = ? limit ?, ?", caid, (page-1)*PageMaxLimits, PageMaxLimits)
	}
	if err != nil {
		return
	}

	var commodity SearchedCommodity
	for rows.Next() {
		err = rows.Scan(&commodity.Cid, &commodity.Name, &commodity.Price)
		if err != nil {
			return
		}

		commodity.Thumbnail, err = GetCommodityExtraInfoByCid(commodity.Cid, "thumbnail")
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				commodity.Thumbnail = "http://ishopping.gq/static/img/404.jpg"
				err = nil
			} else {
				return
			}
		}
		commodities = append(commodities, commodity)
	}

	if len(commodities) == 0 {
		err = errors.New("no commodity found")
	}
	return
}
