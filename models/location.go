package models

type Location struct {
	Id int `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	HospitalId int `db:"hospital_id" json:"hospitalId"`
}
