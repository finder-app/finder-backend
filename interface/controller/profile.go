package controller

import (
	"finder/domain"
	"finder/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	profileUsecase usecase.ProfileUsecase
}

func NewProfileController(pu usecase.ProfileUsecase) *ProfileController {
	return &ProfileController{
		profileUsecase: pu,
	}
}

func (c *ProfileController) Index(ctx *gin.Context) {
	currentUserUid := ctx.Value("currentUserUid").(string)
	user, err := c.profileUsecase.GetProfileByUid(currentUserUid)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *ProfileController) Update(ctx *gin.Context) {
	updateUser := &domain.UpdateUser{}
	ctx.BindJSON(&updateUser)
	currentUserUid := ctx.Value("currentUserUid").(string)
	user, err := c.profileUsecase.UpdateUser(currentUserUid, updateUser)
	if err != nil {
		ErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}
	ctx.JSON(http.StatusOK, user)
}