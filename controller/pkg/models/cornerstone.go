package models

type PrepareCornerstoneData struct {
	Email          string `json:"email"`
	IDNumber       string `json:"id_number"`
	Surname        string `json:"surname"`
	Forenames      string `json:"forenames"`
	Gender         string `json:"gender"`
	DateOfBirth    string `json:"date_of_birth"`
	CountryOfBirth string `json:"country_of_birth"`
}

type CornerstoneData struct {
	IDNumber       string `json:"id_number"`
	Surname        string `json:"surname"`
	Forenames      string `json:"forenames"`
	Gender         string `json:"gender"`
	DateOfBirth    string `json:"date_of_birth"`
	CountryOfBirth string `json:"country_of_birth"`
}
