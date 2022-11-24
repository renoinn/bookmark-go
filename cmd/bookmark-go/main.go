package main

import (
	"log"

	"entgo.io/ent/dialect"
	"github.com/gin-gonic/gin"
	"github.com/renoinn/bookmark-go/datasource/ent"
	"github.com/renoinn/bookmark-go/server/handler"
	"github.com/renoinn/bookmark-go/server/repository"
)

func main() {
    client, err := ent.Open(dialect.MySQL, "sample_user:sample_password@tcp(localhost:3306)/sample_db")
    if err != nil {
        log.Fatal(err)
    }

    r := repository.NewBookmarkRepository(client)

    h := handler.NewBookmarkHandler(r)

	g := gin.Default()
	g.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	g.GET("/bookmarks", h.GetBookmarks)
	g.POST("/bookmark", h.PostBookmark)

	g.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
