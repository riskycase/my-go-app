package repository

import(
	models "github.com/riskycase/my-go-app/models"
	patient "github.com/riskycase/my-go-app/patient"
)

type PatientUsecase struct {
	Repo patient.Repository
}

func NewPatientUsecase(repository patient.Repository) patient.Usecase {
	return &PatientUsecase{
		Repo: repository,
	}
}

func (p *PatientUsecase) Fetch() ([]models.Patient, error) {
	return p.Repo.Fetch()
}

func (p *PatientUsecase) GetById(id int64) (models.Patient, error) {
	return p.Repo.GetById(id)
}

func (p *PatientUsecase) New(patient models.Patient) error {
	return p.Repo.New(patient)
}
