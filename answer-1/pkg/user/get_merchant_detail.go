package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"m4j00/pkg/helper"
	"m4j00/pkg/models"
)

type GetMerchantDetailRequest struct {
	Size int `form:"size"`
	Page int `form:"page"`
}

type GetMerchantDetailResponse struct {
	MerchantName string `json:"merchant_name"`
	OutletName   string `json:"outlet_name"`
	Omzet        int    `json:"omzet"`
}

func GetDetailMerchant(c *gin.Context) {
	user, _ := c.Get("user")
	req := GetMerchantDetailRequest{}
	err := c.BindQuery(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
	}

	var ts []*models.Transaction
	//var merchant []*models.Merchant

	db := c.MustGet("db").(*gorm.DB)

	limit, offset := helper.GetLimitOffset(req.Page, req.Size)

	queryDB := db.
		Preload("Merchant").
		Preload("Merchant.User").
		Preload("Outlet").
		Joins("LEFT JOIN merchants m on transactions.merchant_id = m.id").
		Joins("LEFT JOIN outlets o ON m.id = o.merchant_id").
		Joins("LEFT JOIN users u ON m.user_id = u.id").
		Where("m.user_id = ? ", user.(models.User).ID).
		Limit(limit).
		Offset(offset).
		Find(&ts).Debug()

	err = queryDB.Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
	}

	// Handle response
	merchantDetailResponse := []GetMerchantDetailResponse{}
	for _, m := range ts {

		newMerchantResponse := &GetMerchantDetailResponse{
			MerchantName: m.Merchant.MerchantName,
			OutletName:   m.Outlet.OutletName,
			Omzet:        int(m.BillTotal),
		}

		merchantDetailResponse = append(merchantDetailResponse, *newMerchantResponse)
	}

	//Count all datas
	var total int64
	if res := queryDB.Limit(-1).Offset(-1).Count(&total); res.Error != nil {
		total = 0
	}
	var previous, next bool
	if req.Page != 1 && req.Page != 0 && req.Size != 0 {
		previous = true
	}

	if offset+limit < int(total) && req.Page != 0 && req.Size != 0 {
		next = true
	}

	response := helper.PagedResponse{
		Data:     merchantDetailResponse,
		Page:     req.Page,
		Size:     req.Size,
		Previous: previous,
		Next:     next,
		Total:    int(total),
	}

	c.JSON(http.StatusOK, response)
}
