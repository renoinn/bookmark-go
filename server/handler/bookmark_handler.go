package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renoinn/bookmark-go/server/response"
	"github.com/renoinn/bookmark-go/server/repository"
)

type BookmarkHandler interface {
	GetBookmarks(ctx *gin.Context)
	PostBookmark(ctx *gin.Context)
}

type bookmarkHandler struct {
	userRepository     repository.UserRepository
	siteRepository     repository.SiteRepository
	bookmarkRepository repository.BookmarkRepository
}

func NewBookmarkHandler(
	userRepository repository.UserRepository,
	siteRepository repository.SiteRepository,
	bookmarkRepository repository.BookmarkRepository,
) BookmarkHandler {
	return &bookmarkHandler{
		userRepository,
		siteRepository,
		bookmarkRepository,
	}
}

func (bh *bookmarkHandler) GetBookmarks(ctx *gin.Context) {
	bookmarks, err := bh.bookmarkRepository.GetBookmarks(ctx)
	if err != nil {
		print(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	res := []response.Bookmark{}
	for _, value := range bookmarks {
		data := response.Bookmark{
			BookmarkID: uint64(value.ID),
			SiteTitle:  value.Edges.HaveSite.Title,
			SiteURL:    value.Edges.HaveSite.URL,
			Note:       value.Note,
		}
		res = append(res, data)
	}

	ctx.JSON(http.StatusOK, res)
}

func (bh *bookmarkHandler) PostBookmark(ctx *gin.Context) {
	type PostParam struct {
		Title string `form:"title" json:"title" valid:"Required; MaxSize(100)"`
		URL   string `form:"url" json:"url" valid:"Required; MaxSize(2048)"`
		Note  string `form:"note" json:"note" valid:"MaxSize(1000)"`
	}
	var form PostParam

	err := ctx.BindJSON(form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "{msg: bad request}")
	}

	u, err := bh.userRepository.FindById(ctx, 1)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "{msg: user not found}")
	}

	s, err := bh.siteRepository.CreateSite(ctx, form.Title, form.URL)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "{msg: faild create site}")
	}

	b, err := bh.bookmarkRepository.CreateBookmark(ctx, u, s, form.Note)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "{msg: faild create bookmark}")
	}

	res := response.Bookmark{
		BookmarkID: uint64(b.ID),
		SiteTitle:  b.Edges.HaveSite.Title,
		SiteURL:    b.Edges.HaveSite.URL,
		Note:       b.Note,
	}

	ctx.JSON(http.StatusOK, res)
}
