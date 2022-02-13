package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/oaraujocesar/buildbox-webchallenge-api/db/sqlc"
)

type createPostRequest struct {
	ImageURL string `json:"image_url" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Message  string `json:"message" binding:"required"`
}

func (server *Server) createPost(ctx *gin.Context) {
	var req createPostRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))

		return
	}

	arg := db.CreatePostParams{
		ImageUrl: req.ImageURL,
		Name:     req.Name,
		Message:  req.Message,
	}

	post, err := server.store.CreatePost(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))

		return
	}

	ctx.JSON(http.StatusOK, post)
}

type listPostsRequest struct {
	Page    int32 `form:"page"  binding:"required,min=1"`
	PerPage int32 `form:"perPage" binding:"required,min=5,max=10"`
}

func (server *Server) listPosts(ctx *gin.Context) {
	var req listPostsRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))

		return
	}

	arg := db.ListPostsParams{
		Limit:  req.PerPage,
		Offset: (req.Page - 1) * req.PerPage,
	}

	accounts, err := server.store.ListPosts(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))

		return
	}

	ctx.JSON(http.StatusOK, accounts)
}
