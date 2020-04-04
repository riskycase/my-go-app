package location

import (
	models "github.com/riskycase/my-go-app/models"
)

type Usecase interface {
	Fetch() ([]models.Location, error)
	GetById(id int64) (models.Location, error)
	New(location models.Location) error
}