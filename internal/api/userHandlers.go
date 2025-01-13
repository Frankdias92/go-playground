package api

import (
	"errors"
	"net/http"

	"github.com/frankdias92/go-playground/internal/jsonutils"
	"github.com/frankdias92/go-playground/internal/services"
	"github.com/frankdias92/go-playground/internal/usecase/user"
)

func (api *Api) handleSignUpUser(w http.ResponseWriter, r *http.Request) {
	data, problems, err := jsonutils.DecodeValidJson[user.CreateUser](r)
	if err != nil {
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, problems)
		return
	}

	id, err := api.UserService.CreateUser(r.Context(),
		data.UserName,
		data.Email,
		string(data.Password),
		data.Bio,
	)
	if err != nil {
		if errors.Is(err, services.ErrDuplicatedEmailOrUsername) {
			_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, map[string]any{
				"error": "email or password already exists",
			})
			return
		}
	}

	_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, map[string]any{
		"user_id": id,
	})
}

func (api *Api) handlesLoginUser(w http.ResponseWriter, r *http.Request) {
	panic("todo")
	data, problems, err := jsonutils.DecodeValidJson[user.LoginUserReq](r)
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, problems)
	}

	id, err := api.UserService.AuthenticateUser(r.Context(), data.Email, data.Password)
	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) {
			jsonutils.EncodeJson(w, r, http.StatusBadRequest, map[string]any{
				"error": "invalid email or password",
			})
			return
		}
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": "unexpected internal server error",
		})
		return
	}
	err = api.Sessions.RenewToken(r.Context())
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": "unexpected internal server error",
		})
		return
	}

	api.Sessions.Put(r.Context(), "AuthenticatedUserId", id)

	jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{
		"message": "logged in, succesfully",
	})
}

func (api *Api) handleLogOut(w http.ResponseWriter, r *http.Request) {
	panic("todo")
	err := api.Sessions.RenewToken(r.Context())
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": "unexpected internal server error",
		})
		return
	}

	api.Sessions.Remove(r.Context(), "AuthenticatedUserId")
	jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{
		"message": "logged out succesfully",
	})
}
