package grades

import (
	"net/http"
	"strconv"
	"strings"
)

func RegisterHandlers() {
	handler := new(studentHandler)
	http.Handle("/students", handler)
	http.Handle("/students/", handler)
}

type studentHandler struct{}

// 2 Students
// 3 Students/{id}
// 4 Students/{id}/grades
func (sh studentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathSegments := strings.Split(r.URL.Path, "/")
	switch len(pathSegments) {
	case 2:
		sh.getAll(w, r)
	case 3:
		id, err := strconv.Atoi(pathSegments[2])
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		sh.getOne(w, r, id)
	case 4:
		id, err := strconv.Atoi(pathSegments[2])
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
		}
		sh.addGrade(w, r, id)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (sh studentHandler) getAll(w http.ResponseWriter, r *http.Request) {

}

func (sh studentHandler) getOne(w http.ResponseWriter, r *http.Request, id int) {

}

func (sh studentHandler) addGrade(w http.ResponseWriter, r *http.Request, id int) {

}
