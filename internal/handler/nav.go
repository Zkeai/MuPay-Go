package handler

import (
	"github.com/Zkeai/MuPay-Go/common/conf"
	"github.com/Zkeai/MuPay-Go/common/logger"
	"github.com/Zkeai/MuPay-Go/internal/repo/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

// createNav 增加侧边栏
// @Tags  nav
// @Summary 增加侧边栏
// @Param req body db.NavItem true "增加侧边栏提交参数"
// @Router /nav/projected/add [post]
// @Success 200 {object} conf.Response
// @Failure 400 {object} string "参数错误"
// @Failure 500 {object} string "内部错误"
// @Produce json
// @Accept json
func createNav(c *gin.Context) {
	r := new(db.NavItem)
	if err := c.Bind(r); err != nil {
		logger.Error(err)
		c.JSON(http.StatusInternalServerError, conf.Response{Code: 500, Msg: "err", Data: err.Error()})
		return
	}
	query, err := svc.NavAdd(c.Request.Context(), r)
	if err != nil {
		logger.Error(err)
		c.JSON(http.StatusInternalServerError, conf.Response{Code: 500, Msg: "err", Data: err.Error()})
		return
	}

	c.JSON(http.StatusOK, conf.Response{Code: 200, Msg: "success", Data: query})
}

// navQuery 侧边栏查询
// @Tags  nav
// @Summary 侧边栏查询
// @Router /nav/protected/query [get]
// @Success 200 {object} conf.Response
// @Failure 400 {object} conf.ResponseError
// @Failure 500 {object} string "内部错误"
func navQuery(c *gin.Context) {

	//获取address
	wallet, _ := c.Get("wallet")
	walletAddress := wallet.(string)

	query, err := svc.QueryNav(c.Request.Context(), walletAddress)
	if err != nil {
		logger.Error(err)
		c.JSON(http.StatusInternalServerError, conf.Response{Code: 500, Msg: "err", Data: err.Error()})
		return
	}

	c.JSON(http.StatusOK, conf.Response{Code: 200, Msg: "success", Data: query})
}
