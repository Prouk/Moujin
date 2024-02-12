package src

import "github.com/gin-gonic/gin"

// PopulateRoutes populates the routes of the Moujin instance.
func (m *Moujin) PopulateRoutes() {
	m.R.GET("/", Home)
	m.R.GET("/about", About)
	m.R.Static("/css", "./web/stylesheets")
	m.R.Static("/js", "./web/scripts")
	m.R.Static("/ass", "./web/assets")
}

func Home(c *gin.Context) {
	err := HomePage("Home").Render(c, c.Writer)
	if err != nil {
		return
	}
}

func About(c *gin.Context) {
	err := AboutPage("About").Render(c, c.Writer)
	if err != nil {
		return
	}
}
