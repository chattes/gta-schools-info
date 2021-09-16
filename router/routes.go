package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/chattes/gta-schools-info/common"
	"github.com/chattes/gta-schools-info/controller"
	"github.com/chattes/gta-schools-info/database"
)

type SchoolHandler struct {
}

type searchParam struct {
	params url.Values
}

func newSchoolHandler() *SchoolHandler {
	return &SchoolHandler{}
}

func (h *SchoolHandler) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func (h *SchoolHandler) searchHandler(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "get" || r.Method == "GET":
		err := r.ParseForm()
		if err != nil {
			fmt.Printf("An error occured while parsing form data %e", err)
			w.WriteHeader(http.StatusInternalServerError)
			resp, _ := json.Marshal(common.ErrorResponse{
				Message: "Unexpected error",
			})
			w.Write(resp)
			return
		}

		queryParams := searchParam{
			params: r.Form,
		}

		if len(queryParams.params.Get("name")) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			msg, _ := json.Marshal(common.ErrorResponse{
				Message: "school name cannot be empty",
			})
			w.Write(msg)
			return
		}
		schoolController := controller.SchoolController{
			Db: database.NewMySql(),
		}
		schoolRes, err := schoolController.Search(queryParams.params.Get("name"))

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			msg, _ := json.Marshal(common.ErrorResponse{
				Message: "An error has occured while searching",
			})
			w.Write(msg)
			return
		}
		resp, err := json.Marshal(schoolRes)
		if err != nil {
			panic(err)
		}

		w.Write(resp)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func SetupRoutes() {
	schoolHandler := newSchoolHandler()

	http.HandleFunc("/health", schoolHandler.healthCheck)
	http.HandleFunc("/search", schoolHandler.searchHandler)

}
