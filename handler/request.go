package handler

import "fmt"

func errParamIsRequired(name, typ string) error {

	return fmt.Errorf("Parameter: %s, of type: %s, is required", name, typ)
}

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (req *CreateOpeningRequest) validate() error {

	if req.Role == "" && req.Company == "" && req.Location == "" && req.Remote == nil && req.Link == "" && req.Salary <= 0 {
		return fmt.Errorf("Request body is empty or malformed")
	}

	if req.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if req.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if req.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if req.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}
	if req.Link == "" {
		return errParamIsRequired("link", "string")
	}
	if req.Salary <= 0 {
		return errParamIsRequired("salary", "int64")
	}
	return nil
}

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (req *UpdateOpeningRequest) validate() error {

	// se nenhum campo for vazio ou salario for maios que 0 então retorna nulo, e valida os dados
	if req.Role != "" || req.Company != "" || req.Location != "" || req.Remote != nil || req.Link != "" || req.Salary > 0 {

		return nil
	}

	// se algum dado dos vazio ou salario for menor ou igual a zero retorna o erro
	return fmt.Errorf("at least one valid field muts be provided")
}
