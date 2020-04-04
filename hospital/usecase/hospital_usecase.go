package repository

import(
	models "github.com/riskycase/my-go-app/models"
	hospital "github.com/riskycase/my-go-app/hospital"
)

type HospitalUsecase struct {
	Repo hospital.Repository
}

func NewHospitalUsecase(repository hospital.Repository) hospital.Usecase {
	return &HospitalUsecase{
		Repo: repository,
	}
}

func (p *HospitalUsecase) Fetch() ([]models.Hospital, error) {
	return p.Repo.Fetch()
}

func (p *HospitalUsecase) GetById(id int64) (models.Hospital, error) {
	return p.Repo.GetById(id)
}

func (p *HospitalUsecase) New(hospital models.Hospital) error {
	return p.Repo.New(hospital)
}
