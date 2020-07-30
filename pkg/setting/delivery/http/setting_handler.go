package http

import (
	"net/http"
	"strconv"

	"github.com/abmid/icanvas-analytics/pkg/auth"
	"github.com/abmid/icanvas-analytics/pkg/setting/entity"
	"github.com/abmid/icanvas-analytics/pkg/setting/usecase"
	echo "github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type ResponseError struct {
	Message string `json:"message"`
}

type ResponseSuccess struct {
	Message string `json:"message"`
}

type SettingHandler struct {
	SettingUseCase usecase.SettingUseCase
}

type FormData struct {
	Name     string `form:"name" validate:"required"`
	Category string `form:"category" validate:"required"`
	Value    string `form:"value" validate:"required"`
}

func (SH *SettingHandler) All() echo.HandlerFunc {
	return func(c echo.Context) error {
		filter := entity.Setting{}
		settings, err := SH.SettingUseCase.FindByFilter(filter)
		if err != nil {
			logrus.Error(err)
			return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		}
		if len(settings) == 0 {
			return c.JSON(http.StatusOK, echo.Map{"message": "Data not found"})
		}
		return c.JSON(http.StatusOK, settings)
	}
}

func (SH *SettingHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		formData := new(FormData)
		// Bind
		err := c.Bind(formData)
		if err != nil {
			logrus.Error(err)
			return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		}
		// validate
		if err = c.Validate(formData); err != nil {
			return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		}
		// Store data
		setting := entity.Setting{
			Name:     formData.Name,
			Category: formData.Category,
			Value:    formData.Value,
		}
		err = SH.SettingUseCase.Create(&setting)
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()})
		}
		return c.JSON(http.StatusCreated, echo.Map{"id": setting.ID})
	}
}

func (SH *SettingHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		settingID := c.Param("id")
		formData := new(FormData)
		// Bind
		if err := c.Bind(formData); err != nil {
			logrus.Error(err)
			return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		}
		// validate
		if err := c.Validate(formData); err != nil {
			return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		}
		// Update data
		setting := entity.Setting{
			Name:     formData.Name,
			Category: formData.Category,
			Value:    formData.Value,
		}
		castSettingID, err := strconv.Atoi(settingID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		}

		err = SH.SettingUseCase.Update(uint32(castSettingID), setting)
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()})
		}
		return c.JSON(http.StatusOK, echo.Map{"status": "success"})
	}
}

func (SH *SettingHandler) FindByFilter() echo.HandlerFunc {
	return func(c echo.Context) error {
		filter := new(entity.Setting)
		// Bind
		err := c.Bind(filter)
		if err != nil {
			logrus.Error(err)
			return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		}
		// validate
		if err = c.Validate(filter); err != nil {
			return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		}
		// Get Data
		res, err := SH.SettingUseCase.FindByFilter(*filter)
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()})
		}
		return c.JSON(http.StatusOK, res)
	}
}

func NewHandler(path string, g *echo.Group, JWTKey string, settingUC usecase.SettingUseCase) {
	handler := SettingHandler{
		SettingUseCase: settingUC,
	}
	r := g.Group(path)
	r.Use(auth.MiddlewareAuthJWT(JWTKey))
	// Get All Configuration
	r.GET("", handler.All())
	// Store one configuration
	r.POST("", handler.Create())
	// Update One configuration
	r.PUT("/:id", handler.Update())
	// Get by Filter
	r.GET("/filter", handler.FindByFilter())
	// Store spesific for canvas setting
	r.POST("/canvas", handler.CreateOrUpdateCanvas())
}
