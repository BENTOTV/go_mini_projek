package contoller

import (
	database "myapp/lib/database"
	models "myapp/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// create new Notifikasi
func GetNotifikasisController(c echo.Context) error {
	Notifikasis, err := database.GetNotifikasis()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success getting Notifikasis",
		"data":    Notifikasis,
	})
}

func GetNotifikasiController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrInvalidID.Error())
	}

	Notifikasi, err := database.GetNotifikasiByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success getting Notifikasi",
		"data":    Notifikasi,
	})
}

func CreateNotifikasiController(c echo.Context) error {
	Notifikasi := models.Notifikasi{}
	c.Bind(&Notifikasi)

	newNotifikasi, err := database.CreateNotifikasi(Notifikasi)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success creating Notifikasi",
		"data":    newNotifikasi,
	})
}

func DeleteNotifikasiController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrInvalidID.Error())
	}

	err = database.DeleteNotifikasi(id)
	if err != nil {
		if err == database.ErrIDNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success deleting Notifikasi",
	})
}

func UpdateNotifikasiController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrInvalidID.Error())
	}

	Notifikasi := models.Notifikasi{}
	c.Bind(&Notifikasi)

	err = database.UpdateNotifikasi(id, Notifikasi)
	if err != nil {
		if err == database.ErrIDNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success updating Notifikasi",
	})
}
