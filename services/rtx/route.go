package rtx

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"web-booking/services/logx"

	"github.com/gin-gonic/gin"
)

func (r *ROUTEX) Connect(rx func(*gin.Engine, *gin.RouterGroup)) {

	router := gin.New()

	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/"},
	}))

	// Serve static files from the specified directory
	logx.Infof(`Serve static files from the specified directory`)
	router.Static("/css", "./src/templates/assets/css")
	router.Static("/js", "./src/templates/assets/js")

	// Load HTML templates
	logx.Infof(`Load HTML templates`)
	path := filepath.Join("./src/templates", "**", "*.html")
	logx.Infof(`%s`, path)
	router.LoadHTMLGlob(path)
	logx.Infof(`Load HTML SUCCESS`)

	rx(router, router.Group(r.MODULE))

	// router.Use(UseHeader)
	router.Use(gin.Recovery())

	addr := ``
	appEnv := os.Getenv(`APP_ENV`)
	if appEnv == `develop` {
		addr = fmt.Sprintf("127.0.0.1:%d", r.PORT)
	} else {
		addr = fmt.Sprintf("0.0.0.0:%d", 8080)
	}

	s := &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
