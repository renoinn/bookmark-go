package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/renoinn/bookmark-go/datasource/ent"
	"github.com/renoinn/bookmark-go/datasource/ent/site"
)

type SiteRepository interface {
	CreateSite(ctx *gin.Context, title string, url string) (id uint64, err error)
	FindById(ctx *gin.Context, id int) (site *ent.Site, err error)
}

type siteRepository struct {
	client *ent.Client
}

// CreateSite implements SiteRepository
func (sr *siteRepository) CreateSite(ctx *gin.Context, title string, url string) (id uint64, err error) {
	b := sr.client.Site.Create()
	b.SetTitle(title)
	b.SetURL(url)
	e, err := b.Save(ctx)
	if err != nil {
		return 0, err
	}

	q := sr.client.Site.Query().Where(site.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return 0, err
	}

	return uint64(e.ID), nil
}

// FindById implements SiteRepository
func (sr *siteRepository) FindById(ctx *gin.Context, id int) (site *ent.Site, err error) {
	e, err := sr.client.Site.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func NewSiteRepository(client *ent.Client) SiteRepository {
	return &siteRepository{client}
}
