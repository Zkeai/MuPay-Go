package handler

import (
	"github.com/Zkeai/MuPay-Go/common/conf"
	"github.com/Zkeai/MuPay-Go/common/logger"
	"github.com/Zkeai/MuPay-Go/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

// categoryAdd 增加商品分类
// @Tags  category
// @Summary 增加商品分类
// @Param req body dto.CategoryAddReq true "管理员注册提交参数"
// @Router /category/public/add [post]
// @Success 200 {object} conf.Response
// @Failure 400 {object} string "参数错误"
// @Failure 500 {object} string "内部错误"
// @Produce json
// @Accept json
func categoryAdd(c *gin.Context) {
	r := new(dto.CategoryAddReq)
	if err := c.Bind(r); err != nil {
		logger.Error(err)
		c.JSON(http.StatusInternalServerError, conf.Response{Code: 500, Msg: "err", Data: err.Error()})
		return
	}
	query, err := svc.CategoryAdd(c.Request.Context(), r.Name, r.UserID, r.Sort, r.Icon)
	if err != nil {
		logger.Error(err)
		c.JSON(http.StatusInternalServerError, conf.Response{Code: 500, Msg: "err", Data: err.Error()})
		return
	}

	c.JSON(http.StatusOK, conf.Response{Code: 200, Msg: "success", Data: query})
}

// categoryQuery 商品分类查询
// @Tags  category
// @Summary 商品分类查询
// @Router /category/public/query [get]
// @Success 200 {object} conf.Response
// @Failure 400 {object} conf.ResponseError
// @Failure 500 {object} string "内部错误"
func categoryQuery(c *gin.Context) {
	r := new(dto.CategoryQueryReq)
	if err := c.Bind(r); err != nil {
		logger.Error(err)
		c.JSON(http.StatusInternalServerError, conf.Response{Code: 500, Msg: "err", Data: err.Error()})
		return
	}
	query, err := svc.CategoryQuery(c.Request.Context(), r.Userid)
	if err != nil {
		logger.Error(err)
		c.JSON(http.StatusInternalServerError, conf.Response{Code: 500, Msg: "err", Data: err.Error()})
		return
	}

	c.JSON(http.StatusOK, conf.Response{Code: 200, Msg: "success", Data: query})
}
