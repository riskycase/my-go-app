package models

type Hospital struct {
	Id int `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	MaxPatientAmount int `db:"max_patient_amount" json:"maxPatientAmount"`
}
