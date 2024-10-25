package handler

import (
	"github.com/Zkeai/MuPay-Go/common/conf"
	"github.com/Zkeai/MuPay-Go/common/logger"
	"github.com/Zkeai/MuPay-Go/internal/dto"
	"github.com/Zkeai/MuPay-Go/internal/repo/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

// createCommodity 增加商品
// @Tags  commodity
// @Summary 增加商品
// @Param req body db.YuCommodity true "增加商品提交参数"
// @Router /commodity/projected/add [post]
// @Success 200 {object} conf.Response
// @Failure 400 {object} string "参数错误"
// @Failure 500 {object} string "内部错误"
// @Produce json
// @Accept json
func createCommodity(c *gin.Context) {
	r := new(db.YuCommodity)
	if err := c.Bind(r); err != nil {
		logger.Error(err)
		c.JSON(http.StatusInternalServerError, conf.Response{Code: 500, Msg: "err", Data: err.Error()})
		return
	}
	query, err := svc.CreateCommodity(c.Request.Context(), r)
	if err != nil {
		logger.Error(err)
		c.JSON(http.StatusInternalServerError, conf.Response{Code: 500, Msg: "err", Data: err.Error()})
		return
	}

	c.JSON(http.StatusOK, conf.Response{Code: 200, Msg: "success", Data: query})
}

// commodityQuery 商品查询
// @Tags  commodity
// @Summary 商品查询
// @Router /commodity/public/query [get]
// @Success 200 {object} conf.Response
// @Failure 400 {object} conf.ResponseError
// @Failure 500 {object} string "内部错误"
func commodityQuery(c *gin.Context) {
	r := new(dto.CommodityQueryReq)
	if err := c.Bind(r); err != nil {
		logger.Error(err)
		c.JSON(http.StatusInternalServerError, conf.Response{Code: 500, Msg: "err", Data: err.Error()})
		return
	}
	query, err := svc.GetCommodity(c.Request.Context(), r.CategoryID)
	if err != nil {
		logger.Error(err)
		c.JSON(http.StatusInternalServerError, conf.Response{Code: 500, Msg: "err", Data: err.Error()})
		return
	}

	c.JSON(http.StatusOK, conf.Response{Code: 200, Msg: "success", Data: query})
}
