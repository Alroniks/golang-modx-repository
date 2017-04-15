package main

import (
	"./controllers"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", controllers.Index())
	e.GET("/verify", controllers.Verify())
	//e.GET("/home", handlers.Home)
	//e.GET("/repository", handlers.Repositories)
	//e.GET("/repository/any", handlers.Repository)
	//e.GET("/package", handlers.Package)
	//e.GET("/package/versions", handlers.PackageVersions)
	//e.GET("/package/update", handlers.PackageUpdate)
	//e.GET("/package/download/any", handlers.PackageDownload)

	e.Logger.Fatal(e.Start(":8555"))
}
