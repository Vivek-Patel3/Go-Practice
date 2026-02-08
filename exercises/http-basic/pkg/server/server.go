package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"github.com/google/uuid"
)

type Server struct {
	users map[string]user
}

type user struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func New() *Server {
	return &Server{make(map[string]user)}
}

var index = `
<!DOCTYPE html>
<html>
	<body>
		<h1> My First Heading </h1>
		<p> My First Paragraph </p>
	</body>
</html>
`

func (s *Server) HandleIndex(w http.ResponseWriter, r *http.Request) {
	// since / is fallback path, we need to reject those endpoints which do not exist
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// nothing to read from the request
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(index))
}

func (s *Server) HandleReadAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := make([]user, len(s.users))
	i := 0
	for _, v := range s.users {
		res[i] = user{v.Name, v.Email, v.Age}
		i++
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) HandleCreateUsers(w http.ResponseWriter, r *http.Request) {
	// check the header of the request: validating content-type
	if contentType := r.Header.Get("Content-Type"); contentType != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	// now extract the body of the request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Could not read request body: %v", err)
		w.WriteHeader(http.StatusInternalServerError) // HTTP 500
		return
	}

	defer r.Body.Close()

	// unmarshal the body
	var u user
	err = json.Unmarshal(body, &u)

	if err != nil {
		log.Printf("Could not unmarshal the request body: %v", err)
		w.WriteHeader(http.StatusBadRequest) // HTTP 400
		return
	}

	log.Printf("Create user: %v", u.Name)
	id := uuid.NewString()
	
	s.users[id] = user{
		Name: u.Name,
		Email: u.Email,
		Age:   u.Age,
	}

	w.WriteHeader(http.StatusCreated) // 201
	w.Header().Set("location", "/users/" + id)
	w.Write([]byte(id))
	
}

// read user on the basis of primary key (id)
func (s *Server) HandleReadUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	u, ok := s.users[id]
	if !ok {
		http.NotFound(w, r)
		return
	}

	if err := json.NewEncoder(w).Encode(u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}