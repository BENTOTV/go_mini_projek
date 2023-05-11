package routes

import (
	// admins "myapp/contoller/admin"
	// layanans "myapp/contoller/layanan"
	// notifikasis "myapp/contoller/notifikasi"
	// transaksis "myapp/contoller/transaksi"
	// users "myapp/contoller/user"
	"myapp/config"
	controllers "myapp/contoller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// func InitRoutes() *echo.Echo {
// 	e := echo.New()

// 	// use middleware log
// 	middlewares.LogMiddleware(e)

// 	// use middleware jwt
// 	m := e.Group("")
// 	m.Use(middleware.JWT([]byte("legal")))

// 	//users routes
// 	m.POST("/login", users.LoginUsersController)
// 	m.POST("/users", users.CreateUserController)
// 	m.GET("/users", users.GetUsersController)
// 	m.GET("/users/:id", users.GetUserController)
// 	m.DELETE("/users/:id", users.DeleteUserController)
// 	m.PUT("/users/:id", users.UpdateUserController)
// 	//admin routes
// 	m.GET("/admin", admins.GetAdminsController)
// 	m.GET("/admin/:id", admins.GetAdminController)
// 	m.POST("/admin", admins.CreateAdminController)
// 	m.DELETE("/admin/:id", admins.DeleteAdminController)
// 	m.PUT("/admin/:id", admins.UpdateAdminController)
// 	//layanan routes
// 	m.GET("/layanan", layanans.GetLayanansController)
// 	m.GET("/layanan/:id", layanans.GetLayananController)
// 	m.POST("/layanan", layanans.CreateLayananController)
// 	m.DELETE("/layanan/:id", layanans.DeleteLayananController)
// 	m.PUT("/layanan/:id", layanans.UpdateLayananController)
// 	//transaksi routes
// 	m.GET("/transaksi", transaksis.GetTransaksisController)
// 	m.GET("/transaksi/:id", transaksis.GetTransaksiController)
// 	m.POST("/transaksi", transaksis.CreateTransaksiController)
// 	m.DELETE("/transaksi/:id", transaksis.DeleteTransaksiController)
// 	m.PUT("/transaksi/:id", transaksis.UpdateTransaksiController)
// 	//notifikasi routes
// 	m.GET("/notifikasi", notifikasis.GetNotifikasisController)
// 	m.GET("/notifikasi/:id", notifikasis.GetNotifikasiController)
// 	m.POST("/notifikasi", notifikasis.CreateNotifikasiController)
// 	m.DELETE("/notifikasi/:id", notifikasis.DeleteNotifikasiController)
// 	m.PUT("/notifikasi/:id", notifikasis.UpdateNotifikasiController)
// 	return e
// }

func New() *echo.Echo {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	NewUserControllers(e)
	NewAdminControllers(e)
	NewLayananControllers(e)
	NewNotifikasiControllers(e)
	NewTransaksiControllers(e)

	return e
}

func NewUserControllers(e *echo.Echo) {
	secureGroup := e.Group("")
	secureGroup.Use(middleware.JWT([]byte(config.JWT_SECRET)))

	// No auth
	e.POST("/cusers", controllers.CreateUserController)
	e.POST("/login", controllers.LoginController)

	// Auth
	secureGroup.GET("/users", controllers.GetUsersController)
	secureGroup.GET("/users/:id", controllers.GetUserController)
	secureGroup.DELETE("/users/:id", controllers.DeleteUserController)
	secureGroup.PUT("/users/:id", controllers.UpdateUserController)
}

func NewAdminControllers(e *echo.Echo) {
	secureGroup := e.Group("")
	secureGroup.Use(middleware.JWT([]byte(config.JWT_SECRET)))

	// No auth
	e.POST("/cadmins", controllers.CreateAdminController)
	e.POST("/loginadmin", controllers.LoginAdminController)

	// Auth
	secureGroup.GET("/admins/:id", controllers.GetAdminController)
	secureGroup.GET("/admins", controllers.GetAdminsController)
	secureGroup.DELETE("/admins/:id", controllers.DeleteAdminController)
	secureGroup.PUT("/admins/:id", controllers.UpdateAdminController)
}

func NewLayananControllers(e *echo.Echo) {
	secureGroup := e.Group("")
	secureGroup.Use(middleware.JWT([]byte(config.JWT_SECRET)))

	secureGroup.GET("/layanans", controllers.GetLayanansController)
	secureGroup.GET("/layanans/:id", controllers.GetLayananController)

	secureGroup.POST("/listLayanans", controllers.CreateLayananController)
	secureGroup.DELETE("/layanans/:id", controllers.DeleteLayananController)
	secureGroup.PUT("/layanans/:id", controllers.UpdateLayananController)
}

func NewTransaksiControllers(e *echo.Echo) {
	secureGroup := e.Group("")
	secureGroup.Use(middleware.JWT([]byte(config.JWT_SECRET)))

	secureGroup.GET("/transaksis", controllers.GetTransaksisController)
	secureGroup.GET("/transaksis/:id", controllers.GetTransaksiController)

	// Auth
	secureGroup.POST("/transaksis", controllers.CreateTransaksiController)
	secureGroup.DELETE("/transaksis/:id", controllers.DeleteTransaksiController)
	secureGroup.PUT("/transaksis/:id", controllers.UpdateTransaksiController)
}
func NewNotifikasiControllers(e *echo.Echo) {
	secureGroup := e.Group("")
	secureGroup.Use(middleware.JWT([]byte(config.JWT_SECRET)))

	// No auth
	e.GET("/notifikasis", controllers.GetNotifikasisController)
	e.GET("/notifikasis/:id", controllers.GetNotifikasiController)

	// Auth
	secureGroup.POST("/notifikasis", controllers.CreateNotifikasiController)
	secureGroup.DELETE("/notifikasis/:id", controllers.DeleteNotifikasiController)
	secureGroup.PUT("/notifikasis/:id", controllers.UpdateNotifikasiController)
}
