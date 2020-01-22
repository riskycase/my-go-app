package main

import (
	"github.com/gin-gonic/gin"
	
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
	
	patientR "github.com/riskycase/my-go-app/patient/repository"
	patientU "github.com/riskycase/my-go-app/patient/usecase"
	patientH "github.com/riskycase/my-go-app/patient/handler"
	
	locationR "github.com/riskycase/my-go-app/location/repository"
	locationU "github.com/riskycase/my-go-app/location/usecase"
	locationH "github.com/riskycase/my-go-app/location/handler"
	
	hospitalR "github.com/riskycase/my-go-app/hospital/repository"
	hospitalU "github.com/riskycase/my-go-app/hospital/usecase"
	hospitalH "github.com/riskycase/my-go-app/hospital/handler"
)

func main() {
	r := gin.Default()
	url := "postgres://ghuvahbe:8Y7zUWnG1VZSF8vRsZPvo_W-AdD1EVkP@arjuna.db.elephantsql.com:5432/ghuvahbe"
	db := sqlx.MustConnect("postgres", url)

	patientRepository := patientR.NewPatientRepository(db)
	patientUsecase := patientU.NewPatientUsecase(patientRepository)
	patientHandler := r.Group("/patient/")
	patientH.NewPatientHandler(patientHandler, patientUsecase)

	hospitalRepository := hospitalR.NewHospitalRepository(db)
	hospitalUsecase := hospitalU.NewHospitalUsecase(hospitalRepository)
	hospitalHandler := r.Group("/hospital/")
	hospitalH.NewHospitalHandler(hospitalHandler, hospitalUsecase)

	locationRepository := locationR.NewLocationRepository(db)
	locationUsecase := locationU.NewLocationUsecase(locationRepository)
	locationHandler := r.Group("/location/")
	locationH.NewLocationHandler(locationHandler, locationUsecase)
	
	err := r.Run()
	if err != nil {
		panic(err)
	}
	
}
