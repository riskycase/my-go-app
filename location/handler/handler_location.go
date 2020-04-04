package handler

import(
	location "github.com/riskycase/my-go-app/location"
	models "github.com/riskycase/my-go-app/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type LocationHandler struct {
	Usecase location.Usecase
}

func NewLocationHandler(router *gin.RouterGroup, usecase location.Usecase) {
	handler := &LocationHandler{
		Usecase: usecase,
	}
	router.GET("", handler.GetAllLocations)
	router.GET("/:id", handler.GetLocationById)
	router.POST("", handler.AddLocation)
}

func (p *LocationHandler) GetAllLocations (c *gin.Context) {
	allLocations, err := p.Usecase.Fetch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ "message" : err })
	}
	c.JSON(http.StatusOK, allLocations)
}

func (p *LocationHandler) GetLocationById (c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ "message" : err })
	}
	location, err := p.Usecase.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ "message" : err })
	}
	c.JSON(http.StatusOK, location)
}

func (p *LocationHandler) AddLocation (c *gin.Context) {
	location := models.Location{}
	c.BindJSON(&location)
	err := p.Usecase.New(location)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ "message" : err })
	}
	c.JSON(http.StatusOK, gin.H{ "message" : "Success" })
}
