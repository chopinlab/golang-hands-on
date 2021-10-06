package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"golang-hands-on/docs"
	"golang-hands-on/models"
	"golang.org/x/sync/errgroup"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {

	fmt.Println("hello world")

	g := errgroup.Group{}
	g.Go(func() error { return serveEcho() })
	g.Wait()
}

func serveEcho() (err error) {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	sErr := LoadSwaggerConfig()
	if sErr != nil {
		log.Printf("%v", sErr)
		return
	}

	e.GET("/*", echoSwagger.WrapHandler)

	g := e.Group("/api/v1" )
	g.GET("/health", GetHealthCheck())
	g.GET("/cloudAccounts/:cloudAccountId", GetCloudAccount())

	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}
	e.Listener = listener
	e.Logger.Fatal(e.StartServer(e.Server))
	return
}

func LoadSwaggerConfig() error {

	docs.SwaggerInfo.Title = "golang-hands-on"
	docs.SwaggerInfo.Version = "v1"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Host = "127.0.0.1:3000"
	docs.SwaggerInfo.Schemes = []string{"http"}
	docs.SwaggerInfo.Description = "golang-hands-on"
	return nil
}

// GetHealthCheck
// @Tags Health
// @Summary Health Check
// @Description Health Check
// @Accept json
// @Produce json
// @Success 200 {object} models.SuccessHealthResponse
// @Router /health [get]
func GetHealthCheck() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		return c.JSON(http.StatusOK, models.SuccessHealthResponse{Message: "OK"})
	}
}

// GetCloudAccount
// @Tags Cloud Account
// @Summary Get Cloud Account
// @Description Get Cloud Account
// @Accept json
// @Produce json
// @Param cloudAccountId path string true "Cloud Account Id"
// @Success 200 {object} models.CloudAccount
// @Failure 400 {object} models.ErrorBadRequestResponse
// @Failure 404 {object} models.ErrorNotFoundResponse
// @Router /cloudAccounts/{cloudAccountId} [get]
func GetCloudAccount() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		cloudAccountId := c.Param("cloudAccountId")
		if len(cloudAccountId) == 0 {
			return c.JSON(http.StatusBadRequest,
				models.ErrorBadRequestResponse{Message: "Bad Request", Reason: "Bad Request"})
		}

		// Service나 Dao가 들어갈 부분
		mockUpData := models.CloudAccount{
			CloudAccountId:   "xxxxxx-xxxxxx-xxxxxx",
			CloudAccountName: "aws",
			Provider:         "aws",
			DisplayName:      "ec2",
			UseFlag:          1,
			Created:          time.Now().UTC(),
			Modified:         time.Now().UTC(),
		}
		return c.JSON(http.StatusOK, mockUpData)
	}
}