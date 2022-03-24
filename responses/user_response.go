package responses

import "github.com/labstack/echo/v4"

type UserResponse struct {
	Status  int       `json:"Status"`
	Message string    `json:"message"`
	Data    *echo.Map `json:"data"`
}
