package handler

import (
	"ishopping/src/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CommentsGetHandler(c *gin.Context) {
	Page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		JsonErr(c, "invalid page")
		return
	}

	Pagesize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		JsonErr(c, "invalid page_size")
		return
	}

	ans, err := service.GetCommentListByPageAndPageSize(Pagesize, Page)
	if err != nil {
		JsonErr(c, err.Error())
		return
	}else if len(ans) == 0 {
		JsonOKWithMsg(c, gin.H{"evaluation_list": ans}, "no comment found")
		return
	}
	JsonOK(c, gin.H{"evaluation_list": ans})
}

func CommentDeleteHandler(c *gin.Context) {
	type Params struct {
		Coid int `json:"evaluation_id"`
	}
	var p Params
	err := c.ShouldBindJSON(&p)
	if err != nil {
		JsonErr(c, "BindJsonError when delete comment: "+err.Error())
		return
	}
	err = service.DeleteCommentByCoid(p.Coid)
	if err != nil {
		JsonErr(c, "delete comment error: "+err.Error())
		return
	}
	JsonOK(c, gin.H{})
}

func BanUserHandler(c *gin.Context) {
	type Params struct {
		Uid int `json:"user_id"`
	}
	var p Params
	err := c.ShouldBindJSON(&p)
	if err != nil {
		JsonErr(c, "BindJsonError when ban user: "+err.Error())
		return
	}

	err = service.BanUserByUid(p.Uid)
	if err != nil {
		JsonErr(c, "ban user error: "+err.Error())
		return
	}
	JsonOK(c, gin.H{})
}

func UnBanUserHandler(c *gin.Context) {
	type Params struct {
		Uid int `json:"user_id"`
	}
	var p Params
	err := c.ShouldBindJSON(&p)
	if err != nil {
		JsonErr(c, "BindJsonError when ban user: "+err.Error())
		return
	}

	err = service.UnBanUserByUid(p.Uid)
	if err != nil {
		JsonErr(c, "unban user error: "+err.Error())
		return
	}
	JsonOK(c, gin.H{})
}

func BannedUsersGetHandler(c *gin.Context) {
	Page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		JsonErr(c, "invalid page")
		return
	}

	Pagesize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		JsonErr(c, "invalid page_size")
		return
	}

	ans, err := service.GetBannedUserListByPageAndPageSize(Pagesize, Page)
	if err != nil {
		JsonErr(c, err.Error())
		return
	}
	JsonOK(c, gin.H{"list": ans})
}
