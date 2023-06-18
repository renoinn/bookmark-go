package repository

import (
	"context"
	"log"

	"github.com/renoinn/bookmark-go/datasource/ent"
	"github.com/renoinn/bookmark-go/datasource/ent/bookmark"
)

type BookmarkRepository interface {
	GetBookmarks(ctx context.Context) (bookmarks []*ent.Bookmark, err error)
	CreateBookmark(ctx context.Context, user *ent.User, title string, url string, note string) (entity *ent.Bookmark, err error)
	UpdateBookmark(ctx context.Context, user *ent.User, bookmark ent.Bookmark, title string, url string, note string) (entity *ent.Bookmark, err error)
	DeleteBookmark(ctx context.Context, user *ent.User, bookmark ent.Bookmark) (id int, err error)
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

func (br *bookmarkRepository) CreateBookmark(ctx context.Context, user *ent.User, title string, url string, note string) (entity *ent.Bookmark, err error) {
	b := br.client.Bookmark.Create()
	b.SetOwner(user)
	b.SetTitle(title)
	b.SetURL(url)
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

// UpdateBookmark implements BookmarkRepository.
func (br *bookmarkRepository) UpdateBookmark(ctx context.Context, user *ent.User, bookmark ent.Bookmark, title string, url string, note string) (entity *ent.Bookmark, err error) {
	builder := br.client.Bookmark.UpdateOneID(bookmark.ID)
	builder.SetURL(url)
	builder.SetTitle(title)
	builder.SetNote(note)
	e, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return e, nil
}

// DeleteBookmark implements BookmarkRepository.
func (br *bookmarkRepository) DeleteBookmark(ctx context.Context, user *ent.User, bookmark ent.Bookmark) (id int, err error) {
	builder := br.client.Bookmark.DeleteOneID(bookmark.ID)
	err = builder.Exec(ctx)
	if err != nil {
		return 0, err
	}
	return bookmark.ID, nil
}

func NewBookmarkRepository(client *ent.Client) BookmarkRepository {
	return &bookmarkRepository{client}
}
