package src

import (
	"github.com/gin-gonic/gin"
)

// PopulateRoutes populates the routes of the Moujin instance.
func (m *Moujin) PopulateRoutes() {
	m.R.GET("/", Home)
	m.R.GET("/home", GetHomePage)
	m.R.GET("/about", GetAboutPage)
	m.R.Static("/css", "./web/stylesheets")
	m.R.Static("/js", "./web/scripts")
	m.R.Static("/ass", "./web/assets")
}

func Home(c *gin.Context) {
	err := HomeTemplate().Render(c, c.Writer)
	if err != nil {
		return
	}
}

func GetHomePage(c *gin.Context) {
	err := HomePage().Render(c, c.Writer)
	if err != nil {
		return
	}
}

func GetAboutPage(c *gin.Context) {
	err := AboutPage().Render(c, c.Writer)
	if err != nil {
		return
	}
}
