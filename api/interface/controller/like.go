package controller

import (
	"net/http"

	"github.com/finder-app/finder-backend/api/finder-protocol-buffers/pb"

	"github.com/gin-gonic/gin"
)

type LikeController struct {
	likeClinet pb.LikeServiceClient
}

func NewLikeController(likeClinet pb.LikeServiceClient) *LikeController {
	return &LikeController{
		likeClinet: likeClinet,
	}
}

func (c *LikeController) Create(ctx *gin.Context) {
	req := &pb.CreateLikeReq{
		SentUserUid:     ctx.Value("currentUserUid").(string),
		RecievedUserUid: ctx.Param("uid"),
	}
	like, err := c.likeClinet.CreateLike(ctx, req)
	if err != nil {
		ErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}
	ctx.JSON(http.StatusCreated, like)
}

func (c *LikeController) Index(ctx *gin.Context) {
	req := &pb.GetOldestLikeReq{
		CurrentUserUid: ctx.Value("currentUserUid").(string),
	}
	like, err := c.likeClinet.GetOldestLike(ctx, req)
	if err != nil {
		// NOTE: レコードが1件もなかった場合は200を返す 理由はfrontでエラーを発生させたくないため
		if IsRecordNotFoundError(err) {
			ctx.JSON(http.StatusOK, nil)
			return
		}
		ErrorResponse(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, like)
}

func (c *LikeController) Skip(ctx *gin.Context) {
	req := &pb.SkipLikeReq{
		SentUserUid:     ctx.Param("sent_uesr_uid"),
		RecievedUserUid: ctx.Value("currentUserUid").(string),
	}
	empty, err := c.likeClinet.SkipLike(ctx, req)
	if err != nil {
		ErrorResponse(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, empty)
}

func (c *LikeController) Consent(ctx *gin.Context) {
	req := &pb.ConsentLikeReq{
		RecievedUserUid: ctx.Value("currentUserUid").(string),
		SentUserUid:     ctx.Param("sent_uesr_uid"),
	}
	res, err := c.likeClinet.ConsentLike(ctx, req)
	if err != nil {
		ErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"like": res.Like,
		"room": res.Room,
	})
}
