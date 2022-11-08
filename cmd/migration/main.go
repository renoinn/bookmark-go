package main

import (
	"context"
	"fmt"
	"log"
	"os"

    atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"
	"github.com/renoinn/bookmark-go/datasource/ent/migrate"
)

func main() {
    dataSourceName := fmt.Sprintf("mysql://%s:%s@%s/%s?charset=utf8&parseTime=True", "bkmk_user", "bkmk_password", "localhost:3306", "bookmark_db")
    ctx := context.Background()

    // Create a local migration directory able to understand Atlas migration file format for replay.
    dir, err := atlas.NewLocalDir("datasource/ent/migrate/migrations")
    if err != nil {
        log.Fatalf("failed creating atlas migration directory: %v", err)
    }

    // Migrate diff options.
    opts := []schema.MigrateOption{
        schema.WithDir(dir),                         // provide migration directory
        schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
        schema.WithDialect(dialect.MySQL),           // Ent dialect to use
        schema.WithFormatter(atlas.DefaultFormatter),
    }

    // Generate migrations using Atlas support for MySQL (note the Ent dialect option passed above).
    err = migrate.NamedDiff(ctx, dataSourceName, os.Args[1], opts...)
    if err != nil {
        log.Fatalf("failed generating migration file: %v", err)
    }

    /*
    client, err := ent.Open("mysql", dataSourceName)
    if err != nil {
        log.Fatalf("failed opening connection to mysql: %v", err)
    }
    defer client.Close()
    // Run the auto migration tool.
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
    */
}
