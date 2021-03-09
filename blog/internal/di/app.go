package di

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/ccsunnyfd/practice/blog/internal/service"
)

// App App
type App struct {
	svc  *service.Service
	http *http.Server
}

// NewApp NewApp
func NewApp(svc *service.Service, h *http.Server) (app *App, closeFunc func(), err error) {
	app = &App{
		svc:  svc,
		http: h,
	}
	closeFunc = func() {
		ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
		defer cancel()

		if err := h.Shutdown(ctx); err != nil {
			log.Panicf("httpSrv.Shutdown error(%v)", err)
		}

		log.Println("Server exiting")
	}
}
