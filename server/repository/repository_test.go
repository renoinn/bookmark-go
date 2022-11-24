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
    got, err := ur.CreateUser(ctx, "hoge", "hoge@hoge.com")
    if err != nil {
        t.Fatalf("failed creating user")
    }
    t.Logf("create user id: %d\n", got)
}
