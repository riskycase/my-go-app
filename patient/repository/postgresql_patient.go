package repository

import(
	models "github.com/riskycase/my-go-app/models"
	patient "github.com/riskycase/my-go-app/patient"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type PatientRepository struct {
	Db *sqlx.DB
}

func NewPatientRepository(db *sqlx.DB) patient.Repository {
	return &PatientRepository{
		Db: db,
	}
}

func (p *PatientRepository) Fetch() ([]models.Patient, error) {
	patients := []models.Patient{}
	err := p.Db.Select(&patients, "SELECT * FROM patient")
	return patients, err
}

func (p *PatientRepository) GetById(id int64) (models.Patient, error) {
	patient := models.Patient{}
	err := p.Db.Get(&patient, "SELECT * FROM patient where id = $1", strconv.FormatInt(id, 10))
	return patient, err
}

func (p *PatientRepository) New(patient models.Patient) error {
	_, err := p.Db.NamedExec("INSERT INTO patient (name, illness, birth_date, location_id) VALUES (:name, :illness, :birth_date, :location_id)", &patient)
	return err
}
