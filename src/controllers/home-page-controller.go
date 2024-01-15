package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomePageController(c *gin.Context) {
	c.HTML(http.StatusOK, "home/index.html", gin.H{
		`css_path`:  "/css/index.css",
		`title`:     `HOME PAGE`,
		`page_name`: `HOME PAGE`,
	})
}
