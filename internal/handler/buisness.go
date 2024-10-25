package handler

import (
	"github.com/Zkeai/MuPay-Go/common/conf"
	"github.com/Zkeai/MuPay-Go/common/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// businessQuery 商铺信息查询
// @Tags  business
// @Summary 商铺信息查询
// @Router /user/public/query [get]
// @Success 200 {object} conf.Response
// @Failure 400 {object} conf.ResponseError
// @Failure 500 {object} string "内部错误"
func businessQuery(c *gin.Context) {
	query, err := svc.BusinessQuery(c.Request.Context(), c.Request.Host)
	if err != nil {
		logger.Error(err)
		c.JSON(http.StatusInternalServerError, conf.Response{Code: 500, Msg: "err", Data: err.Error()})
		return
	}

	c.JSON(http.StatusOK, conf.Response{Code: 200, Msg: "success", Data: query})
}
