package hospital

import (
	models "github.com/riskycase/my-go-app/models"
)

type Repository interface {
	Fetch() ([]models.Hospital, error)
	GetById(id int64) (models.Hospital, error)
	New(hospital models.Hospital) error
}