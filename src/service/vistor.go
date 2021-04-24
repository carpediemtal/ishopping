package service

import (
	"database/sql"
	"errors"
	"ishopping/src/db"
)

const PageMaxLimits = 15 //页面最大展示商品数量

type SearchedCommodity struct {
	Cid         int     `json:"cid" map:"cid"`
	Name        string  `json:"name" map:"name"`
	Price       float64 `json:"price" map:"price"`
	PicturePath string  `json:"picture_path" map:"picture_path"`
}

// 通过cid与meta_key查找meta_val
func GetCommodityExtraInfoByCid(cid int, metaKey string) (metaVal string, err error) {
	row := db.DB.QueryRow("select meta_val from commodity_meta where cid = ? and meta_key = ?", cid, metaKey)
	err = row.Scan(&metaVal)
	return
}

//获取某个分类下对应页数的商品列表
func GetCommodityListByCategoryAndPage(caid int, page int) (commodities []SearchedCommodity, err error) {
	var rows *sql.Rows
	if caid == 0 {
		rows, err = db.DB.Query("select cid, name, price from commodity limit ?, ?", (page-1)*PageMaxLimits, PageMaxLimits)
	} else {
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
		commodity.PicturePath, err = GetCommodityExtraInfoByCid(commodity.Cid, "picture_path")
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				commodity.PicturePath = "defult"
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
