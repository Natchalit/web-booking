package routes

import (
	"web-booking/services/rtx"
	"web-booking/src/controllers"

	"github.com/gin-gonic/gin"
)

var MODULE = rtx.ROUTEX{
	MODULE: ``,
	PORT:   8000,
}

func Routes() {
	MODULE.Connect(func(e *gin.Engine, rg *gin.RouterGroup) {
		e.GET(`/`, controllers.HomePageController)
	})
}
