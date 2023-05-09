package contoller

import (
	database "myapp/lib/database"
	models "myapp/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// create new Layanan
func GetLayanansController(c echo.Context) error {
	Layanans, err := database.GetLayanans()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success getting Layanans",
		"data":    Layanans,
	})
}

func GetLayananController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrInvalidID.Error())
	}

	Layanan, err := database.GetLayananByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success getting Layanan",
		"data":    Layanan,
	})
}

func CreateLayananController(c echo.Context) error {
	Layanan := models.Layanan{}
	c.Bind(&Layanan)

	newLayanan, err := database.CreateLayanan(Layanan)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success creating Layanan",
		"data":    newLayanan,
	})
}

func DeleteLayananController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrInvalidID.Error())
	}

	err = database.DeleteLayanan(id)
	if err != nil {
		if err == database.ErrIDNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success deleting Layanan",
	})
}

func UpdateLayananController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrInvalidID.Error())
	}

	Layanan := models.Layanan{}
	c.Bind(&Layanan)

	err = database.UpdateLayanan(id, Layanan)
	if err != nil {
		if err == database.ErrIDNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success updating Layanan",
	})
}
