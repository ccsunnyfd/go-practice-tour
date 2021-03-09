package di

import (
	"net/http"

	"github.com/ccsunnyfd/practice/blog/ent"
	"github.com/ccsunnyfd/practice/blog/internal/service"
	"github.com/gin-gonic/gin"
)

func InitApp() (*App, func(), error) {
	dao, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer dao.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	server := gin.Default()
	service := service.NewService()
	NewApp(, srv)
}
