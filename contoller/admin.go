package contoller

import (
	database "myapp/lib/database"
	utils "myapp/lib/middleware"
	models "myapp/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// create new Admin
func GetAdminsController(c echo.Context) error {
	Admins, err := database.GetAdmins()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success getting Admins",
		"data":    Admins,
	})
}

func GetAdminController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrInvalidID.Error())
	}

	Admin, err := database.GetAdminByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success getting Admin",
		"data":    Admin,
	})
}

func CreateAdminController(c echo.Context) error {
	Admin := models.Admin{}
	c.Bind(&Admin)

	newAdmin, err := database.CreateAdmin(Admin)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success creating Admin",
		"data":    newAdmin,
	})
}

func DeleteAdminController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrInvalidID.Error())
	}

	err = database.DeleteAdmin(id)
	if err != nil {
		if err == database.ErrIDNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success deleting Admin",
	})
}

func UpdateAdminController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrInvalidID.Error())
	}

	Admin := models.Admin{}
	c.Bind(&Admin)

	if Admin.Password != "" {
		hash, err := utils.HashPassword(Admin.Password)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		Admin.Password = hash
	}

	err = database.UpdateAdmin(id, Admin)
	if err != nil {
		if err == database.ErrIDNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success updating Admin",
	})
}

func LoginAdminController(c echo.Context) error {
	AdminRequest := models.Admin{}
	c.Bind(&AdminRequest)

	Admin, err := database.LoginAdmin(AdminRequest)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusUnauthorized, ErrInvalidCredentials.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if !utils.ComparePassword(AdminRequest.Password, Admin.Password) {
		return echo.NewHTTPError(http.StatusUnauthorized, ErrInvalidCredentials.Error())
	}

	token, err := utils.GenerateTokenAdmin(Admin)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"token":   token,
	})
}
