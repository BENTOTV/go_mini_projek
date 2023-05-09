package contoller

import (
	database "myapp/lib/database"
	models "myapp/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// create new Transaksi
func GetTransaksisController(c echo.Context) error {
	Transaksis, err := database.GetTransaksis()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success getting Transaksis",
		"data":    Transaksis,
	})
}

func GetTransaksiController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrInvalidID.Error())
	}

	Transaksi, err := database.GetTransaksiByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success getting Transaksi",
		"data":    Transaksi,
	})
}

func CreateTransaksiController(c echo.Context) error {
	Transaksi := models.Transaksi{}
	c.Bind(&Transaksi)

	newTransaksi, err := database.CreateTransaksi(Transaksi)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success creating Transaksi",
		"data":    newTransaksi,
	})
}

func DeleteTransaksiController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrInvalidID.Error())
	}

	err = database.DeleteTransaksi(id)
	if err != nil {
		if err == database.ErrIDNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success deleting Transaksi",
	})
}

func UpdateTransaksiController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrInvalidID.Error())
	}

	Transaksi := models.Transaksi{}
	c.Bind(&Transaksi)

	err = database.UpdateTransaksi(id, Transaksi)
	if err != nil {
		if err == database.ErrIDNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success updating Transaksi",
	})
}
