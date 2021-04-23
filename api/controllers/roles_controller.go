package controllers

import (
	"backend/api/auth"
	"backend/api/models"
	"backend/api/responses"
	"backend/api/utils/formaterror"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (server *Server) validToken(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if tokenID != uint32(uid) {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
}

//Get All Roles function
func (server *Server) GetRoles(w http.ResponseWriter, r *http.Request) {

	role := models.Roles{}

	roles, err := role.FindAllRole(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, roles)
}

//Get Roles by Email function
func (server *Server) GetRolesEmail(w http.ResponseWriter, r *http.Request) {
	// contents, err := ioutil.ReadAll(r.Body)
	vars := mux.Vars(r)
	email := vars["email"]
	role := models.Roles{}
	rolesGet, err := role.FindAllRoleBymail(server.DB, email)
	// err = json.Unmarshal(contents, &rolesGet)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, rolesGet)
}

// Create Role function
func (server *Server) CreateRole(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	role := models.Roles{}
	err = json.Unmarshal(body, &role)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	if tokenID != uint32(uid) {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	role.Prepare()
	err = role.Validate("")

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	roleCreated, err := role.SaveRole(server.DB)

	if err != nil {

		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, roleCreated.ID))
	responses.JSON(w, http.StatusCreated, roleCreated)
}

// Update Role function
func (server *Server) UpdateRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	role := models.Roles{}
	err = json.Unmarshal(body, &role)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	role.Prepare()
	err = role.Validate("update")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	updatedRole, err := role.UpdateRole(server.DB, uint32(uid))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedRole)
}

// Delete Role function
func (server *Server) DeleteRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	user := models.Roles{}

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = user.DeleteRole(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	responses.JSON(w, http.StatusNoContent, "")
}
