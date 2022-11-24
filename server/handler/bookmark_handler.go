package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renoinn/bookmark-go/server/repository"
)

type BookmarkHandler interface {
    GetBookmarks(ctx *gin.Context)
    PostBookmark(ctx *gin.Context)
}

type bookmarkHandler struct {
    repository repository.BookmarkRepository
}

func NewBookmarkHandler(repository repository.BookmarkRepository) BookmarkHandler {
    return &bookmarkHandler{repository}
}

func (bh *bookmarkHandler) GetBookmarks(ctx *gin.Context) {
    bookmarks, err := bh.repository.GetBookmarks(ctx)
    if err != nil {
        print(err)
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }

    type ResponseData struct {
		BookmarkID   uint64 `json:"bookmark_id"`
		SiteTitle string `json:"site_title"`
		SiteURL string `json:"site_title"`
		Note string `json:"note"`
	}

    response := []ResponseData{}
    for _, value := range bookmarks {
        data := ResponseData{
            BookmarkID: uint64(value.ID),
            SiteTitle: value.Edges.HaveSite.Title,
            SiteURL: value.Edges.HaveSite.URL,
            Note: value.Note,
        }
        response = append(response, data)
    }

    ctx.JSON(http.StatusOK, response)
}

func (bh *bookmarkHandler) PostBookmark(ctx *gin.Context) {
    type PostParam struct {
		Title string `form:"title" json:"title" valid:"Required; MaxSize(100)"`
		URL string `form:"url" json:"url" valid:"Required; MaxSize(2048)"`
		Note string `form:"note" json:"note" valid:"MaxSize(1000)"`
	}
	var form PostParam

    err := ctx.BindJSON(form)
	if err != nil {
        ctx.JSON(http.StatusBadRequest, "{msg: bad request}")
	}
}
