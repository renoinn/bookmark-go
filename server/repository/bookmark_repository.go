package repository

import (
	"context"
	"log"

	"github.com/renoinn/bookmark-go/datasource/ent"
	"github.com/renoinn/bookmark-go/datasource/ent/bookmark"
)

type BookmarkRepository interface {
    GetBookmarks(ctx context.Context) (bookmarks []*ent.Bookmark, err error)
    CreateBookmark(ctx context.Context, user *ent.User, site *ent.Site, note string) (entity *ent.Bookmark, err error)
}

type bookmarkRepository struct {
    client *ent.Client
}

func (br *bookmarkRepository) GetBookmarks(ctx context.Context) (bookmarks []*ent.Bookmark, err error) {
    results, err := br.client.Bookmark.Query().All(ctx)
    if err != nil {
		log.Print(err)
		return nil, err
    }
    return results, nil
}

func (br *bookmarkRepository) CreateBookmark(ctx context.Context, user *ent.User, site *ent.Site, note string) (entity *ent.Bookmark, err error) {
    b := br.client.Bookmark.Create()
    b.SetHaveSite(site)
    b.SetOwner(user)
    b.SetNote(note)
    e, err := b.Save(ctx)
	if err != nil {
		return nil, err
	}

	q := br.client.Bookmark.Query().Where(bookmark.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}

	return e, nil
}

func NewBookmarkRepository(client *ent.Client) BookmarkRepository {
    return &bookmarkRepository{client}
}

