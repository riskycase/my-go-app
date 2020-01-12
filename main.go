package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"encoding/json"
	
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type Patient struct {
	Id int `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Illness string `db:"illness" json:"illness"`
	BirthDate string `db:"birth_date" json:"birth_date"`
	LocationId int `db:"location_id" json:"location_id"`
}

type Hospital struct {
	Id int `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	MaxPatientAmount int `db:"max_patient_amount" json:"max_patient_amount"`
}

type Location struct {
	Id int `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	HospitalId int `db:"hospital_id" json:"hospital_id"`
}
	
/*
 * Initialise the database as a global variable
 */
func initDatabase() {
	url := "postgres://ghuvahbe:8Y7zUWnG1VZSF8vRsZPvo_W-AdD1EVkP@arjuna.db.elephantsql.com:5432/ghuvahbe"
	var err error //Seperate initialisation allows use of = instead of := in next line
	//Thus we modify the global variable db instead of declaring a local one with :=
	db, err = sqlx.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

/*
 * Return a json of all patients
 */
func GetAllPatients(c *gin.Context){
	patients := []Patient{}
	err := db.Select(&patients, "SELECT * FROM patient")
	if err != nil {
		panic(err)
	}
	output, e := json.MarshalIndent(patients, "", "	")
	if e != nil {
		panic(e)
	}
	c.String(http.StatusOK, string(output))
}

/*
 * Return a json of all hospitals
 */
func GetAllHospitals(c *gin.Context){
	hospitals := []Hospital{}
	err := db.Select(&hospitals, "SELECT * FROM hospital")
	if err != nil {
		panic(err)
	}
	output, e := json.MarshalIndent(hospitals, "", "	")
	if e != nil {
		panic(e)
	}
	c.String(http.StatusOK, string(output))
}

/*
 * Return a json of all locations
 */
func GetAllLocations(c *gin.Context){
	locations := []Location{}
	err := db.Select(&locations, "SELECT * FROM location")
	if err != nil {
		panic(err)
	}
	output, e := json.MarshalIndent(locations, "", "	")
	if e != nil {
		panic(e)
	}
	c.String(http.StatusOK, string(output))
}

/*
 * Initialise the routes
 */
func initAPI() {

	//Initialise a router
	var router = gin.Default()

	// Return a preset JSON on receiving a get request at /hello
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{ "message" : "Hello world!" })
	})

	//Construct and send a JSON greeting the person in the parameter
	router.GET("/greet/:name", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{ "message" : "Hello " + c.Param("name") + "!"})
	})

	//Return list of all patients
	router.GET("/patient/all", GetAllPatients)

	//Return a list of all hospitals
	router.GET("/hospital/all", GetAllHospitals)

	//Return a list of all locations
	router.GET("/location/all", GetAllLocations)
	
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

func main() {

	//call the above two functions
	initDatabase()
	initAPI()
	
}
