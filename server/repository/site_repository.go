package repository

import (
	"context"

	"github.com/renoinn/bookmark-go/datasource/ent"
	"github.com/renoinn/bookmark-go/datasource/ent/site"
)

type SiteRepository interface {
	CreateSite(ctx context.Context, title string, url string) (entity *ent.Site, err error)
	FindById(ctx context.Context, id int) (entity *ent.Site, err error)
}

type siteRepository struct {
	client *ent.Client
}

// CreateSite implements SiteRepository
func (sr *siteRepository) CreateSite(ctx context.Context, title string, url string) (entity *ent.Site, err error) {
	b := sr.client.Site.Create()
	b.SetTitle(title)
	b.SetURL(url)
	e, err := b.Save(ctx)
	if err != nil {
		return nil, err
	}

	q := sr.client.Site.Query().Where(site.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}

	return e, nil
}

// FindById implements SiteRepository
func (sr *siteRepository) FindById(ctx context.Context, id int) (entity *ent.Site, err error) {
	e, err := sr.client.Site.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func NewSiteRepository(client *ent.Client) SiteRepository {
	return &siteRepository{client}
}
