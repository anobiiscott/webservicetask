package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	. "github.com/scottbeaman/webservice-exercise/developers"
	"strconv"
)

type Response struct {
	Code   int    `json:"code"`
	Result string `json:"result"`
}

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", HelloHandler)
	r.HandleFunc("/developers", DevelopersGET).Methods("GET")
	r.HandleFunc("/developers", DevelopersPOST).Methods("POST")
	r.HandleFunc("/developers/{key}", DeveloperGET)
	return r
}

func SendResponse(res []byte, w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	w.Write(res)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// Build the response struct
	name := r.FormValue("name")
	code := 200
	if len(name) == 0 {
		res, _ := json.Marshal(Response{http.StatusUnprocessableEntity, "Unprocessable Entity"})
		SendResponse(res, w, http.StatusUnprocessableEntity)
		return
	} else {
		res, _ := json.Marshal(fmt.Sprintf("%s %s", "hello", name))
		SendResponse(res, w, code)
		return
	}
}

func DeveloperGET(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["key"])

	if key > len(Developers)  || key <= 0 {
		res, _ := json.Marshal(Response{http.StatusNotFound, "Developer not found"})
		SendResponse(res, w, http.StatusNotFound)
		return
	}
	res, _ := json.Marshal(Developers[key])
	SendResponse(res, w, 200)
}

func DevelopersGET(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal(Developers)
	SendResponse(res, w, 200)
}

func DevelopersPOST(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	age, _ := strconv.Atoi(r.FormValue("age"))
	language := r.FormValue("language")
	floor, _ := strconv.Atoi(r.FormValue("floor"))

	if len(name) == 0 || age == 0 || len(language) == 0 || floor ==0 {
		res, _ := json.Marshal(Response{http.StatusUnprocessableEntity, "Unprocessable Entity"})
		SendResponse(res, w, http.StatusUnprocessableEntity)
	}

	index := len(Developers) + 1
	Developers[index] = Developer{Name: name, Age: age, Language: language, Floor: floor}
	DevelopersGET(w, r)
}
