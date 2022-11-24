package repository

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/renoinn/bookmark-go/datasource/ent"
	"github.com/renoinn/bookmark-go/datasource/ent/bookmark"
)

type BookmarkRepository interface {
    GetBookmarks(ctx *gin.Context) (bookmarks []*ent.Bookmark, err error)
    CreateBookmark(ctx *gin.Context, user *ent.User, site *ent.Site, note string) (id uint64, err error)
}

type bookmarkRepository struct {
    client *ent.Client
}

func (br *bookmarkRepository) GetBookmarks(ctx *gin.Context) (bookmarks []*ent.Bookmark, err error) {
    results, err := br.client.Bookmark.Query().All(ctx)
    if err != nil {
		log.Print(err)
		return nil, err
    }
    return results, nil
}

func (br *bookmarkRepository) CreateBookmark(ctx *gin.Context, user *ent.User, site *ent.Site, note string) (id uint64, err error) {
    b := br.client.Bookmark.Create()
    b.SetHaveSite(site)
    b.SetOwner(user)
    e, err := b.Save(ctx)
	if err != nil {
		return 0, err
	}

	q := br.client.Bookmark.Query().Where(bookmark.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return 0, err
	}

	return uint64(e.ID), nil
}

func NewBookmarkRepository(client *ent.Client) BookmarkRepository {
    return &bookmarkRepository{client}
}

