package api

import (
	"errors"
	"net/http"

	"github.com/Luiz-Hen-Reis/go-book-library/internal/jsonutils"
	"github.com/Luiz-Hen-Reis/go-book-library/internal/services"
	"github.com/Luiz-Hen-Reis/go-book-library/internal/usecases/authors"
)

func (api *Api) handleCreateAuthor(w http.ResponseWriter, r *http.Request) {
	data, problems, err := jsonutils.DecodeValidJson[authors.CreateAuthorReq](r)

	if err != nil {
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, problems)
		return
	}

	author, err := api.AuthorService.CreateAuthor(
		r.Context(),
		data.Name,
		data.Bio,
	)

	if err != nil {
		if errors.Is(err, services.ErrUnexpectedError) {
			_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, map[string]any{
				"error": err.Error(),
			})
			return
		}
		
	}
	
	_ = jsonutils.EncodeJson(w, r, http.StatusCreated, map[string]any{
		"author": author,
	})
}