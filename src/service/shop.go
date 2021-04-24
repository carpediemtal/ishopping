package service

import "ishopping/src/db"

type Shop struct {
	Uid            int    `json:"uid" map:"uid"`           //buyer 的 id
	Sid            int    `json:"Sid" map:"sid"`           //店的id
	Name           string `json:"name" map:"name"`         //buyer 的name
	ShopName       string `json:"shopName" map:"shopName"` //shop 的name 店名
	Address        string `json:"address" map:"address"`
	SellerPhoneNum string `json:"phone_num" map:"phone_num"`
}

func GetShopProfileById(uid int) (shop Shop, err error) {
	row := db.DB.QueryRow("select * from buyer natural join shop where buyer.uid = ?", uid)
	err = row.Scan(&shop.Uid, &shop.Name, &shop.Address, &shop.SellerPhoneNum, &shop.Sid, &shop.ShopName)
	return
}

func UpdateShopInfo(uid int, shopName, address, phoneNum string) (err error) {
	shop, err := GetShopProfileById(uid)
	// 如果买家的信息已经存在
	if err == nil {
		_, err = db.DB.Exec(`update shop set shop_name = ? where uid = ?`, shopName, uid)
		_, err = db.DB.Exec(`update buyer set address = ?, phone_num = ? where uid = ?`, address, phoneNum, uid)
	} else {
		_, err = db.DB.Exec(`insert into shop (sid,uid, shop_name) values (?, ?, ?, ?)`, shop.Sid, uid, shopName)
		_, err = db.DB.Exec(`insert into buyer (uid, name, address, phone_num) values (?, ?, ?, ?)`, uid, shop.Name, address, phoneNum)
	}
	return
}
