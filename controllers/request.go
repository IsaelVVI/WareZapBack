package controllers

import "fmt"

// Function for return err in request body
func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// Create Opening
type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

// Function for verify request body
func (request *CreateOpeningRequest) Validate() error {
	if request.Role == "" &&
		request.Company == "" &&
		request.Link == "" &&
		request.Location == "" &&
		request.Remote == nil &&
		request.Salary <= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}

	if request.Role == "" {
		return errParamIsRequired("role", "string")
	}

	if request.Company == "" {
		return errParamIsRequired("company", "string")
	}

	if request.Location == "" {
		return errParamIsRequired("location", "string")
	}

	if request.Link == "" {
		return errParamIsRequired("link", "string")
	}

	if request.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}

	if request.Salary <= 0 {
		return errParamIsRequired("salary", "int64")
	}

	return nil
}

// Update Opening
type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (request *UpdateOpeningRequest) Validate() error {
	// If any field is provieded, validation is truthy

	if request.Role != "" ||
		request.Company != "" ||
		request.Location != "" ||
		request.Link != "" ||
		request.Salary > 0 ||
		request.Remote != nil {
		return nil
	}

	// If none of the fields were provided, return falsy
	return fmt.Errorf("at least one valid field must be provided")

}
