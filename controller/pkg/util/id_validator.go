package util

import (
	"fmt"
	"strconv"
)

func IDValidator(id, gender, countryOfBirth string) (string, error) {
	// check if numeric
	_, err := strconv.Atoi(id)
	if err != nil {
		return "", fmt.Errorf("ID must be numeric")
	}

	// check if id is 13 digits
	if len(id) != 13 {
		return "", fmt.Errorf("ID must be 13 characters long")
	}

	// validate date of birth month
	idMonth := string(id[2:4])
	if idMonth < "01" || idMonth > "12" {
		return "", fmt.Errorf("ID contains invalid DOB month")
	}

	// validate date of birth month
	idDay := string(id[4:6])
	if idDay < "01" || idDay > "31" {
		return "", fmt.Errorf("ID contains invalid DOB day")
	}

	// validate gender
	idGender := string(id[6])
	if idGender >= "5" {
		idGender = "Male"
	} else {
		idGender = "Female"
	}

	if idGender != gender {
		return "", fmt.Errorf("genders don't match")
	}

	// validate ZA citizenship
	citizen := string(id[10])
	if citizen > "1" {
		return "", fmt.Errorf("invalid citizenship status")
	}

	if !(citizen == "0" && countryOfBirth == "South Africa") {
		return "", fmt.Errorf("invalid citizenship status")
	}

	// validate checksum
	valid := validate(id)

	if !valid {
		return "", fmt.Errorf("invalid ID number")
	}

	return id, nil
}

func validate(pan string) bool {
	/* Validate string with Luhn (mod-10) */
	var alter bool
	var checksum int

	for position := len(pan) - 1; position > -1; position-- {
		digit := int(pan[position] - 48)
		if alter {
			digit = digit * 2
			if digit > 9 {
				digit = (digit % 10) + 1
			}
		}
		alter = !alter
		checksum += digit
	}
	return checksum%10 == 0
}
