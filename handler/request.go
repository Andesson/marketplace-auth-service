package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
}

func (r *CreateUserRequest) Validate() error {
	if r.Email == "" && r.FullName == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Email == "" {
		return errParamIsRequired("email", "string")
	}
	if r.FullName == "" {
		return errParamIsRequired("full_name", "string")
	}
	if r.Password == "" {
		return errParamIsRequired("password", "string")
	}
	return nil
}
