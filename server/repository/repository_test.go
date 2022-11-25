package repository

import (
	"context"
	"testing"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
	"github.com/renoinn/bookmark-go/datasource/ent"
)

func TestCreateUser(t *testing.T) {
    client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
    if err != nil {
        t.Fatalf("failed opening connection to sqlite: %v", err)
    }
    defer client.Close()
    ctx := context.Background()
    if err := client.Schema.Create(ctx); err != nil {
        t.Fatalf("failed creating schema resources: %v", err)
    }

    ur := userRepository{client}
    got, err := ur.CreateUser(ctx, "hoge1", "hoge@hoge.com")
    if err != nil {
        t.Fatalf("failed creating user")
    }
    t.Logf("create user id: %d\n", got.ID)
}

func TestBookmarkRepository(t *testing.T) {
    client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
    if err != nil {
        t.Fatalf("failed opening connection to sqlite: %v", err)
    }
    defer client.Close()
    ctx := context.Background()
    if err := client.Schema.Create(ctx); err != nil {
        t.Fatalf("failed creating schema resources: %v", err)
    }

    ur := userRepository{client}
    u, err := ur.CreateUser(ctx, "hoge2", "hoge@hoge.com")
    if err != nil {
        t.Fatalf("failed creating user: %v", err)
    }

    sr := siteRepository{client}
    s, err := sr.CreateSite(ctx, "google", "https://www.google.com")
    if err != nil {
        t.Fatalf("failed creating site: %v", err)
    }

    br := bookmarkRepository{client}
    b, err := br.CreateBookmark(ctx, u, s, "testing")
    if err != nil {
        t.Fatalf("failed creating bookmark: %v", err)
    }

    t.Logf("create bookmark id: %d\n", b.ID)
}
