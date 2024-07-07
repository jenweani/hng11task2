package services

import (
	"errors"
	"hng11task2/internal/models"
)

func UserFieldValidations(u models.User) []error {
	var errs []error

	// Test for uniqueness
	// check for other objects with the same value as these fields
	// fields to test: email
	users := GetUsersByEmail(u.Email)
	
	if len(users) > 0 {
		errs = append(errs, errors.New("user account with the same email found"))
	}

	// Test for nullness: check fields that are null
	switch "" {
		case u.Email :
			errs = append(errs, errors.New("email cannot be empty"))
		case u.FirstName:
			errs = append(errs, errors.New("firstName cannot be empty"))
		case u.LastName:
			errs = append(errs, errors.New("lastName cannot be empty"))
		case u.Password: 
		errs = append(errs, errors.New("email cannot be empty"))
	}

	if len(errs) != 0 {
		return errs
	}
	return nil
}

func  OrganisationFieldValidations(o models.Organisation) []error {
	var errs []error

	// Test for nullness
	// check fields that are null
	// fields: name
	if o.Name == "" {
		errs = append(errs, errors.New("name cannot be empty"))
	}

	if len(errs) != 0 {
		return errs
	}
	return nil
}