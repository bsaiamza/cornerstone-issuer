package models

type DhaData struct {
	Users []User `json:"users"`
}

type User struct {
	ID                     string `json:"id"`
	GivenName              string `json:"givenName"`
	FamilyName             string `json:"familyName"`
	Gender                 string `json:"gender"`
	LprNumber              string `json:"lprNumber"`
	LprCategory            string `json:"lprCategory"`
	ResidentSince          string `json:"residentSince"`
	CommuterClassification string `json:"commuterClassification"`
	BirthDate              string `json:"birthDate"`
	BirthCountry           string `json:"birthCountry"`
}
