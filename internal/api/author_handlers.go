package api

import (
	"errors"
	"net/http"

	"github.com/Luiz-Hen-Reis/go-book-library/internal/jsonutils"
	"github.com/Luiz-Hen-Reis/go-book-library/internal/services"
	"github.com/Luiz-Hen-Reis/go-book-library/internal/usecases/authors"
	"github.com/go-chi/chi/v5"
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
		if errors.Is(err, services.ErrDuplicatedName) {
			_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, map[string]any{
				"error": err.Error(),
			})
			return
		}
		
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, map[string]any{
			"error": services.ErrUnexpectedError.Error(),
		})
		return
	}
	
	_ = jsonutils.EncodeJson(w, r, http.StatusCreated, map[string]any{
		"author": author,
	})
}

func (api *Api) handleListAuthors(w http.ResponseWriter, r *http.Request) {
	data, err := api.AuthorService.ListAuthors(r.Context())

	if err != nil {
		if errors.Is(err, services.ErrUnexpectedError) {
			_ = jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
				"error": err.Error(),
			})
			return
		}

		_ = jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": services.ErrUnexpectedError.Error(),
		})
		return
	}

	_ = jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{
		"authors": data,
	})
}

func (api *Api) handleGetAuthorByID (w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	data, err := api.AuthorService.GetAuthorByID(r.Context(), idParam)

	if err != nil {
		if errors.Is(err, services.ErrUnexpectedError) {
			_ = jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
				"error": err.Error(),
			})
			return
		}

		_ = jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
		return
	}

	_ = jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{
		"author": data,
	})
}

func (api *Api) handleDeleteAuthorByID (w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	err := api.AuthorService.DeleteAuthorByID(r.Context(), idParam)

	if err != nil {
		if errors.Is(err, services.ErrAuthorNotFound) {
			_ = jsonutils.EncodeJson(w, r, http.StatusNotFound, map[string]any{
				"error": err.Error(),
			})
			return
		}
		
		_ = jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
		return
	}

	_ = jsonutils.EncodeJson(w, r, http.StatusNoContent, map[string]any{})
}