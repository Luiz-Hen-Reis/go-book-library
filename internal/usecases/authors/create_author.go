package authors

import (
	"context"

	"github.com/Luiz-Hen-Reis/go-book-library/internal/validator"
)

type CreateAuthorReq struct {
		Name 	string     	`json:"name"`
		Bio  	string 			`json:"bio"`
}

type DefaultAuthorRes struct {
	ID 	string     	`json:"id"`
	Name 	string     	`json:"name"`
	Bio  	string 			`json:"bio"`
}

func (req CreateAuthorReq) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator

	eval.CheckField(validator.NotBlank(req.Name), "name", "this field cannot be empty")
	eval.CheckField(validator.MinChars(req.Name, 3) &&
	 validator.MaxChars(req.Name, 20), "name", "this field must have a length between 3 and 20")

	eval.CheckField(validator.NotBlank(req.Bio), "bio", "this field cannot be empty")
	eval.CheckField(validator.MinChars(req.Bio, 10) &&
	 validator.MaxChars(req.Bio, 255), "bio", "this field must have a length between 3 and 20")

	return eval
}