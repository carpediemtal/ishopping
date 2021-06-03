package service

import (
	"database/sql"
	"errors"
	"fmt"
	"ishopping/src/db"
)

type SearchedComment struct {
	Coid    int    `json:"evaluation_id" map:"coid"`
	Uid     int    `json:"buyer_id" map:"uid"`
	Bname   string `json:"buyer_name" map:"name"`
	Cid     int    `json:"commodity_id" map:"cid"`
	Cname   string `json:"commodity_name" map:"name"`
	Content string `json:"content" map:"conetnt"`
}

type SearchedBannedUser struct {
	Uid  int    `json:"user_id" map:"uid"`
	Name string `json:"name"`
}

// 通过uid获取姓名
func GetBuyerNameByUid(uid int) (name string, err error) {
	row := db.DB.QueryRow("select name from buyer where uid = ?", uid)
	err = row.Scan(&name)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			err = errors.New(fmt.Sprintf("no buyer found uid:'%d'", uid))
		}
	}
	return
}

func GetCommodityNameByCid(cid int) (name string, err error) {
	row := db.DB.QueryRow("select name from commodity where cid = ?", cid)
	err = row.Scan(&name)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			err = errors.New(fmt.Sprintf("no commodity found cid:'%d'", cid))
		}
	}
	return
}

//根据页面大小倒序获取对应页数的评论列表
func GetCommentListByPageAndPageSize(page_sizde, page int) (comments []SearchedComment, err error) {
	var rows *sql.Rows
	if page <= 0 || page_sizde <= 0 {
		err = errors.New("invalid page or page_size")
	} else {
		rows, err = db.DB.Query("select coid, cid, uid, content from comment order by coid desc limit ?, ?", (page-1)*page_sizde, page_sizde)
	}
	if err != nil {
		return
	}

	var comment SearchedComment
	for rows.Next() {
		err = rows.Scan(&comment.Coid, &comment.Cid, &comment.Uid, &comment.Content)
		if err != nil {
			return
		}

		comment.Bname, err = GetBuyerNameByUid(comment.Uid)
		if err != nil {
			return
		}

		comment.Cname, err = GetCommodityNameByCid(comment.Cid)
		if err != nil {
			return
		}

		comments = append(comments, comment)
	}

	// if len(comments) == 0 {
	// 	err = errors.New("no comment found")
	// }
	return
}

//根据评论id删除评论
func DeleteCommentByCoid(coid int) (err error) {
	_, err = db.DB.Exec("delete from comment where coid = ?", coid)
	return
}

//根据用户id ban用户
func BanUserByUid(uid int) (err error) {
	row := db.DB.QueryRow(`select uid from user_ban where uid = ?`, uid)
	var x int
	if err = row.Scan(&x); err == nil {
		return errors.New("the user has been disabled already")
	}

	_, err = db.DB.Exec("insert into user_ban (uid) values (?)", uid)
	return
}

//根据用户id unban用户
func UnBanUserByUid(uid int) (err error) {
	row := db.DB.QueryRow(`select uid from user_ban where uid = ?`, uid)
	var x int
	if err = row.Scan(&x); err != nil {
		return errors.New("the user is not disabled")
	}

	_, err = db.DB.Exec("delete from user_ban where uid = ?", uid)
	return
}

//根据用户id获取用户类别
func GetUserTypeByUid(uid int) (userType int, err error) {
	row := db.DB.QueryRow("select type from user where uid = ?", uid)
	err = row.Scan(&userType)
	return
}

//根据用户id获取买家名或卖家的店铺名
func GetNameByUid(uid, usertype int) (name string, err error) {
	var row *sql.Row
	if usertype == 0 {
		row = db.DB.QueryRow("select name from buyer where uid = ?", uid)
	} else {
		row = db.DB.QueryRow("select shop_name from shop where uid = ?", uid)
	}
	err = row.Scan(&name)
	return
}

//根据页面大小倒序获取对应页数的banned用户列表
func GetBannedUserListByPageAndPageSize(page_size, page int) (bannedusers []SearchedBannedUser, err error) {
	var rows *sql.Rows
	if page <= 0 || page_size <= 0 {
		err = errors.New("invalid page or page_size")
	} else {
		rows, err = db.DB.Query("select uid from user_ban order by uid desc limit ?, ?", (page-1)*page_size, page_size)
	}
	if err != nil {
		fmt.Print("here2")
		return
	}

	var banneduser SearchedBannedUser
	var usertype int
	for rows.Next() {
		err = rows.Scan(&banneduser.Uid)
		if err != nil {
			fmt.Print("here")
			return
		}
		usertype, err = GetUserTypeByUid(banneduser.Uid)
		if err != nil {
			return
		}
		banneduser.Name, err = GetNameByUid(banneduser.Uid, usertype)
		if err != nil {
			fmt.Print("here3")
			return
		}
		bannedusers = append(bannedusers, banneduser)
	}

	if len(bannedusers) == 0 {
		err = errors.New("no banned user found")
	}

	return
}
