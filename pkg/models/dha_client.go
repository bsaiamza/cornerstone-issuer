package models

type DHAResponse struct {
	Root Root `xml:"ROOT"`
}

type Root struct {
	Error  string `xml:"ERROR"`
	Person Person `xml:"PERSON"`
}

type Person struct {
	IDNumber       string `xml:"IDNUMBER"`
	NewIDNo        string `xml:"NEWIDNO"`
	Surname        string `xml:"SURNAME"`
	Names          string `xml:"NAMES"`
	Gender         string `xml:"GENDER"`
	DeathStatus    string `xml:"DEATH_STATUS"`
	BirthPlace     string `xml:"BIRTH_PLACE"`
	Nationality    string `xml:"NATIONALITY"`
	MarriageStatus string `xml:"MARRIAGE_STATUS"`
	EmailAddress   string `xml:"EMAIL_ADDRESS"`
	AddrLine1      string `xml:"ADDR_LINE1"`
	AddrLine2      string `xml:"ADDR_LINE2"`
	AddrPostalCode string `xml:"ADDR_POSTAL_CODE"`
	PostAddr1      string `xml:"POST_ADDR1"`
	PostAddr2      string `xml:"POST_ADDR2"`
	PostPostalCode string `xml:"POST_POSTAL_CODE"`
	NewResAddr1    string `xml:"NEW_RES_ADDR1"`
	NewResAddr2    string `xml:"NEW_RES_ADDR2"`
	NewResAddr3    string `xml:"NEW_RES_ADDR3"`
	NewResAddr4    string `xml:"NEW_RES_ADDR4"`
	NewResPcode    string `xml:"NEW_RES_PCODE"`
	NewPostAddr1   string `xml:"NEW_POST_ADDR1"`
	NewPostAddr2   string `xml:"NEW_POST_ADDR2"`
	NewPostAddr3   string `xml:"NEW_POST_ADDR3"`
	NewPostAddr4   string `xml:"NEW_POST_ADDR4"`
	NewPostPcode   string `xml:"NEW_POST_PCODE"`
}

type DHASimulatorResponse struct {
	IDNumber                   string `json:"Identity_Number"`
	Names                      string `json:"Names"`
	Surname                    string `json:"Surname"`
	Sex                        string `json:"Sex"`
	IssueDate                  string `json:"Issue_Date"`
	DateOfBirth                string `json:"Date_of_Birth"`
	BiometricsPhoto            string `json:"Biometrics-photo"`
	BiometricsFingerprint      string `json:"Biometrics-fingerprint"`
	BiometricsFingerprintMatch int    `json:"Biometrics-fingerprint_match"`
	Nationality                string `json:"Nationality"`
	CountryOfBirth             string `json:"Country_of_Birth"`
	Status                     string `json:"Status"`
	BarCode                    string `json:"Bar_Code"`
}

type DHAUser struct {
	IDNumber       string `json:"id_number"`
	FirstNames     string `json:"first_names"`
	Surname        string `json:"surname"`
	Gender         string `json:"gender"`
	DateOfBirth    string `json:"date_of_birth"`
	CountryOfBirth string `json:"Country_of_Birth"`
}

// type DHAErrorResponse []DHAError

// type DHAError struct {
// 	OriginatingDate string `json:"originatingDate"`
// 	ResponseCode    int64  `json:"responseCode"`
// 	ResponseDesc    string `json:"responseDesc"`
// 	ZaID            string `json:"ZA ID"`
// 	EndToEndID      string `json:"endToEndId"`
// }
