package patient

import (
	models "github.com/riskycase/my-go-app/models"
)

type Usecase interface {
	Fetch() ([]models.Patient, error)
	GetById(id int64) (models.Patient, error)
	New(patient models.Patient) error
}