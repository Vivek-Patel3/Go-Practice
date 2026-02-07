package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Server struct {
	users map[string]userInfo
}

type user struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type userInfo struct {
	email string
	age   int
}

func New() *Server {
	return &Server{make(map[string]userInfo)}
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

func (s *Server) HandleUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := make([]user, len(s.users))
	i := 0
	for k, v := range s.users {
		res[i] = user{k, v.email, v.age}
		i++
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) HandleCreateUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost, http.MethodPut:
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
		s.users[u.Name] = userInfo{
			email: u.Email,
			age:   u.Age,
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// read user on the basis of name
func (s *Server) HandleReadUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// get the query parameters
		name := r.URL.Query().Get("name") // now we have the map of query parameters and their values

		// now find the name in the stored map
		u, ok := s.users[name]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// converting to DTO
		result := user {
			Name: name,
			Age: u.age,
			Email: u.email,
		}

		// now convert to JSON
		msg, err := json.Marshal(result)
		if err != nil {
			log.Printf("Could not marshal request: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(msg)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}