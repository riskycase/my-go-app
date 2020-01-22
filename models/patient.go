package models

type Patient struct {
	Id int `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Illness string `db:"illness" json:"illness"`
	BirthDate string `db:"birth_date" json:"birthDate"`
	LocationId int `db:"location_id" json:"locationId"`
}
