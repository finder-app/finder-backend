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
	uid := ctx.Value("currentUserUid").(string)
	users, err := c.footPrintInteractor.GetFootPrintUsersByUid(uid)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// 未読の足跡数を返すエンドポイント
