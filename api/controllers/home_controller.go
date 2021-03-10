package controllers

import (
	"net/http"

	"backend/api/responses"
)

//Home function
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}

// func (server *Server) GetData(w http.ResponseWriter, r *http.Request) {
// 	datas := Data{
// 		Pesan{Nama: "Budi", Pesan: "Oke"},
// 		Pesan{Nama: "Ujang", Pesan: "Iya"},
// 	}

// 	if err := json.NewEncoder(w).Encode(datas); err != nil {
// 		panic(err)
// 	}
// }
