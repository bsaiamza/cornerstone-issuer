package utils

import (
	"fmt"
	"strconv"
)

func IDValidator(id, gender, countryOfBirth string) (string, error) {
	err := numericCheck(id)
	if err != nil {
		return "", err
	}

	err = digitCount(id)
	if err != nil {
		return "", err
	}

	err = validBirthMonth(id)
	if err != nil {
		return "", err
	}

	err = validBirthDay(id)
	if err != nil {
		return "", err
	}

	err = validGender(id, gender)
	if err != nil {
		return "", err
	}

	err = citizenshipCheck(id, countryOfBirth)
	if err != nil {
		return "", err
	}

	err = validChecksum(id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func numericCheck(id string) error {
	_, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("ID must be numeric")
	}

	return nil
}

func digitCount(id string) error {
	if len(id) != 13 {
		return fmt.Errorf("ID must be 13 characters long")
	}

	return nil
}

func validBirthMonth(id string) error {
	birthMonth := string(id[2:4])
	if birthMonth < "01" || birthMonth > "12" {
		return fmt.Errorf("ID contains invalid DOB month")
	}

	return nil
}

func validBirthDay(id string) error {
	birthDay := string(id[4:6])
	if birthDay < "01" || birthDay > "31" {
		return fmt.Errorf("ID contains invalid DOB day")
	}

	return nil
}

func validGender(id, gender string) error {
	idGender := string(id[6])
	if idGender >= "5" {
		idGender = "Male"
	} else {
		idGender = "Female"
	}

	if idGender != gender {
		return fmt.Errorf("genders don't match")
	}

	return nil
}

func citizenshipCheck(id, countryOfBirth string) error {
	citizen := string(id[10])
	if citizen > "1" {
		return fmt.Errorf("invalid citizenship status")
	}

	if !(citizen == "0" && countryOfBirth == "South Africa") {
		return fmt.Errorf("invalid citizenship status")
	}

	return nil
}

func validChecksum(id string) error {
	valid := validateChecksum(id)

	if !valid {
		return fmt.Errorf("invalid ID number")
	}

	return nil
}

func validateChecksum(pan string) bool {
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
