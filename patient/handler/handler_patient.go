package handler

import(
	patient "github.com/riskycase/my-go-app/patient"
	models "github.com/riskycase/my-go-app/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PatientHandler struct {
	Usecase patient.Usecase
}

func NewPatientHandler(router *gin.RouterGroup, usecase patient.Usecase) {
	handler := &PatientHandler{
		Usecase: usecase,
	}
	router.GET("", handler.GetAllPatients)
	router.GET("/:id", handler.GetPatientById)
	router.POST("", handler.AddPatient)
}

func (p *PatientHandler) GetAllPatients (c *gin.Context) {
	allPatients, err := p.Usecase.Fetch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ "message" : err })
	}
	c.JSON(http.StatusOK, allPatients)
}

func (p *PatientHandler) GetPatientById (c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ "message" : err })
	}
	patient, err := p.Usecase.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ "message" : err })
	}
	c.JSON(http.StatusOK, patient)
}

func (p *PatientHandler) AddPatient (c *gin.Context) {
	patient := models.Patient{}
	c.BindJSON(&patient)
	err := p.Usecase.New(patient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ "message" : err })
	}
	c.JSON(http.StatusOK, gin.H{ "message" : "Success" })
}
