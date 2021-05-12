package controller

import (
	"finder/usecase"
	"net/http"
)

type footPrintController struct {
	footPrintUsecase usecase.FootPrintUsecase
}

func NewFootPrintController(fpi usecase.FootPrintUsecase) *footPrintController {
	return &footPrintController{
		footPrintUsecase: fpi,
	}
}

func (c *footPrintController) Index(ctx Context) {
	currentUserUid := ctx.Value("currentUserUid").(string)
	footPrints, err := c.footPrintUsecase.GetFootPrintsByUid(currentUserUid)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, footPrints)
}

// 未読の足跡数を返すエンドポイント
