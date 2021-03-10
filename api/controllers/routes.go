package controllers

import (
	"backend/api/middlewares"
	"backend/code"
)

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST", "OPTIONS")

	// s.Router.HandleFunc("/login", s.Login).Methods("POST")

	s.Router.HandleFunc("/test", s.EndpointTest)
	// s.Router.HandleFunc("/home/getjson", s.GetData)

	// s.Router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Test") })
	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Role routes
	s.Router.HandleFunc("/roles", middlewares.SetMiddlewareJSON(s.GetRoles)).Methods("GET")
	s.Router.HandleFunc("/roles/{email}", middlewares.SetMiddlewareJSON(s.GetRolesEmail)).Methods("GET")
	s.Router.HandleFunc("/roles/create", middlewares.SetMiddlewareJSON(s.CreateRole)).Methods("POST")
	s.Router.HandleFunc("/roles/edit/{id}", middlewares.SetMiddlewareJSON(s.UpdateRole)).Methods("PUT")
	s.Router.HandleFunc("/roles/delete/{id}", middlewares.SetMiddlewareJSON(s.DeleteRole)).Methods("DELETE")

	//Other routes
	s.Router.HandleFunc("/file", middlewares.SetMiddlewareJSON(code.Client))

}
