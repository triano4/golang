package controllers

import (
	"backend/api/models"
	"backend/api/responses"
	"backend/api/utils/formaterror"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Get All Roles function
func (server *Server) GetMenu(w http.ResponseWriter, r *http.Request) {

	nav := models.Nav{}

	navs, err := nav.FindAllMenu(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, navs)
}

// Create Menu function
func (server *Server) CreateMenu(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	nav := models.Nav{}
	err = json.Unmarshal(body, &nav)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = nav.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	navCreated, err := nav.SaveNav(server.DB)

	if err != nil {

		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, navCreated.Id))
	responses.JSON(w, http.StatusCreated, navCreated)
}
