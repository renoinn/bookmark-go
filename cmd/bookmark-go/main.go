package main

import (
	"context"
	"fmt"
	"log"

	"github.com/renoinn/bookmark-go/datasource/ent"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
    dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True", "user", "p@ssword", "localhost:3306", "bookmark_db")
    client, err := ent.Open("mysql", dataSourceName)
    if err != nil {
        log.Fatalf("failed opening connection to mysql: %v", err)
    }
    defer client.Close()
    // Run the auto migration tool.
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
}
