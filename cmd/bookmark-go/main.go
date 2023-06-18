package main

import (
	"log"

	"entgo.io/ent/dialect"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/renoinn/bookmark-go/datasource/ent"
	"github.com/renoinn/bookmark-go/server/handler"
	"github.com/renoinn/bookmark-go/server/repository"
)

func main() {
	client, err := ent.Open(dialect.MySQL, "bkmk_user:bkmk_password@tcp(bookmark_mysql:3306)/bookmark_db")
	if err != nil {
		log.Fatal(err)
	}

	u := repository.NewUserRepository(client)
	b := repository.NewBookmarkRepository(client)

	h := handler.NewBookmarkHandler(u, b)

	g := gin.Default()
	g.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	g.GET("/bookmarks", h.GetBookmarks)
	g.POST("/bookmark", h.PostBookmark)
	g.PUT("/bookmark", h.PutBookmark)
	g.DELETE("/bookmark", h.DeleteBookmark)

	g.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
