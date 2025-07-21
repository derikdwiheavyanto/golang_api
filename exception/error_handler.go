package exception

import (
	"golang_restful_api/helper"
	"golang_restful_api/model/web"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err any) {

	if notFoundError(w, r, err) {
		return
	}

	if validationError(w, r, err) {
		return
	}

	if unauthorized(w, r, err) {
		return
	}

	InternalServerError(w, r, err)
}

func notFoundError(w http.ResponseWriter, r *http.Request, err any) bool {
	exception, ok := err.(NotFoundError)

	if ok {
		w.Header().Set("Content-Type", "aplication/json")
		w.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func validationError(w http.ResponseWriter, r *http.Request, err any) bool {
	exception, ok := err.(validator.ValidationErrors)

	if ok {
		w.Header().Set("Content-Type", "aplication/json")
		w.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func unauthorized(w http.ResponseWriter, r *http.Request, err any) bool {
	_, ok := err.(Unauthorized)

	if ok {
		w.Header().Set("Content-Type", "aplication/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func InternalServerError(w http.ResponseWriter, r *http.Request, err any) {
	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToResponseBody(w, webResponse)
}
