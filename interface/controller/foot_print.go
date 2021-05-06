package controller

import (
	"finder/usecase/interactor"
	"net/http"
)

type footPrintController struct {
	footPrintInteractor interactor.FootPrintInteractor
}

func NewFootPrintController(fpi interactor.FootPrintInteractor) *footPrintController {
	return &footPrintController{
		footPrintInteractor: fpi,
	}
}

func (c *footPrintController) Index(ctx Context) {
	currentUserUid := ctx.Value("currentUserUid").(string)
	footPrints, err := c.footPrintInteractor.GetFootPrintsByUid(currentUserUid)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, footPrints)
}

// 未読の足跡数を返すエンドポイント
