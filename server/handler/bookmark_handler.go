package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renoinn/bookmark-go/server/repository"
	"github.com/renoinn/bookmark-go/server/response"
)

type BookmarkHandler interface {
	GetBookmarks(ctx *gin.Context)
	PostBookmark(ctx *gin.Context)
}

type bookmarkHandler struct {
	userRepository     repository.UserRepository
	bookmarkRepository repository.BookmarkRepository
}

func NewBookmarkHandler(
	userRepository repository.UserRepository,
	bookmarkRepository repository.BookmarkRepository,
) BookmarkHandler {
	return &bookmarkHandler{
		userRepository,
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
			SiteTitle:  value.Title,
			SiteURL:    value.URL,
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

	b, err := bh.bookmarkRepository.CreateBookmark(ctx, u, form.Title, form.URL, form.Note)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "{msg: faild create bookmark}")
	}

	res := response.Bookmark{
		BookmarkID: uint64(b.ID),
		SiteTitle:  b.Title,
		SiteURL:    b.URL,
		Note:       b.Note,
	}

	ctx.JSON(http.StatusOK, res)
}
