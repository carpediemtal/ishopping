package service

import (
	"ishopping/src/db"
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

func GetBuyerProfileByUsername(username string) (buyer Buyer, err error) {
	uid, err := getIdByUsername(username)
	if err != nil {
		return
	}
	return GetBuyerProfileById(uid)
}
