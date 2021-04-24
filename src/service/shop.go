package service

import "ishopping/src/db"

type Shop struct {
	Uid            int    `json:"uid" map:"uid"` //buyer 的 id
	ShopName       string `json:"shopName" map:"shopName"`
	Address        string `json:"address" map:"address"`
	SellerPhoneNum string `json:"phone_num" map:"phone_num"`
}

func GetShopProfileById(uid int) (shop Shop, err error) {
	shop.Uid = uid
	shop.Address, err = GetShopAddressById(uid)
	if err != nil {
		return shop, err
	}
	shop.SellerPhoneNum, err = GetShopPhoneNumById(uid)
	if err != nil {
		return shop, err
	}
	shop.ShopName, err = GetShopNameById(uid)
	if err != nil {
		return shop, err
	}

	return
}

func GetShopAddressById(uid int) (address string, err error) {
	row := db.DB.QueryRow("select meta_val from user_meta  where uid = ? and meta_key= ?", uid, "shop_address")
	err = row.Scan(&address)
	return
}

func GetShopPhoneNumById(uid int) (phoneNum string, err error) {
	row := db.DB.QueryRow("select meta_val from user_meta  where uid = ? and meta_key= ?", uid, "shop_phonenum")
	err = row.Scan(&phoneNum)
	return
}

func GetShopNameById(uid int) (shopName string, err error) {
	//row := db.DB.QueryRow("select meta_val from user_meta  where uid = ? and meta_key= ?", uid, "shop_name")
	row := db.DB.QueryRow("select shop_name from shop where uid = ?", uid)
	err = row.Scan(&shopName)
	return
}

func UpdateShopInfo(uid int, shopName, address, phoneNum string) (err error) {
	_, err = GetShopProfileById(uid)
	// 如果买家的信息已经存在
	if err == nil {
		_, err = db.DB.Exec(`update user_meta set meta_val = ? where uid = ? and meta_key= ?`, address, uid, "shop_address")
		_, err = db.DB.Exec(`update user_meta set meta_val = ? where uid = ? and meta_key= ?`, phoneNum, uid, "shop_phonenum")
		//shop_name 放在user_meta里
		//_, err = db.DB.Exec(`update user_meta set meta_val = ? where uid = ? and meta_key= ?`, shopName, uid, "shop_name")
		//shop_name 在shop里修改
		_, err = db.DB.Exec(`update shop set shop_name = ? where uid = ?`, shopName, uid)
	} else {
		_, err = db.DB.Exec(`insert into user_meta (uid, meta_key,meta_val) values (?, "shop_address", ?)`, uid, address)
		_, err = db.DB.Exec(`insert into user_meta (uid, meta_key,meta_val) values (?, "shop_phonenum", ?)`, uid, phoneNum)
		//同上
		//_, err = db.DB.Exec(`insert into user_meta (uid, meta_key,meta_val) values (?, "shop_name", ?)`, uid, shopName)

		_, err = db.DB.Exec(`update shop set shop_name = ? where uid = ?`, shopName, uid)

	}

	return
}
