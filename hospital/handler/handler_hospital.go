package handler

import(
	hospital "github.com/riskycase/my-go-app/hospital"
	models "github.com/riskycase/my-go-app/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type HospitalHandler struct {
	Usecase hospital.Usecase
}

func NewHospitalHandler(router *gin.RouterGroup, usecase hospital.Usecase) {
	handler := &HospitalHandler{
		Usecase: usecase,
	}
	router.GET("", handler.GetAllHospitals)
	router.GET("/:id", handler.GetHospitalById)
	router.POST("", handler.AddHospital)
}

func (p *HospitalHandler) GetAllHospitals (c *gin.Context) {
	allHosptials, err := p.Usecase.Fetch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ "message" : err })
	}
	c.JSON(http.StatusOK, allHosptials)
}

func (p *HospitalHandler) GetHospitalById (c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ "message" : err })
	}
	hospital, err := p.Usecase.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ "message" : err })
	}
	c.JSON(http.StatusOK, hospital)
}

func (p *HospitalHandler) AddHospital (c *gin.Context) {
	hospital := models.Hospital{}
	c.BindJSON(&hospital)
	err := p.Usecase.New(hospital)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ "message" : err })
	}
	c.JSON(http.StatusOK, gin.H{ "message" : "Success" })
}
