package main

import (
	"fmt"
    "log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"encoding/json"
)

var sessions []string

func genSessionHandler(w http.ResponseWriter, r *http.Request) {
	secret := r.FormValue("secret")

	/// verify secret key
	if secret == "abcdef" {
		sessions = append(sessions, "") /// create new session
		fmt.Fprint(w, len(sessions) - 1) /// return session index as ID
	}
}

func updateSessionHandler(w http.ResponseWriter, r *http.Request) {
	secret := r.FormValue("secret")

	if secret == "abcdef" {
		playback := r.FormValue("playback")
		vars := mux.Vars(r)

		i, err := strconv.Atoi(vars["sessionKey"])

		var jsonMap map[string]string
		if (err == nil) {
			sessions[i] = playback
			jsonMap = make(map[string]string)

			jsonMap["Success"] = "true"
		} else {
			jsonMap = make(map[string]string)

			jsonMap["Success"] = "false"
		}

		response, _ := json.Marshal(jsonMap)
		fmt.Fprint(w, string(response))
	}
}

func joinSessionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	i, err := strconv.Atoi(vars["sessionKey"])

	var jsonMap map[string]string

	jsonMap = make(map[string]string)
	fmt.Printf("%d %d %t\n", i, len(sessions), err == nil)

	if err == nil && i < len(sessions) {
		jsonMap["Success"] = "true"
		jsonMap["Playback"] = sessions[i]

		fmt.Println("No ERR")
	} else {
		jsonMap["Success"] = "false"
	}

	response, _ := json.Marshal(jsonMap)
	fmt.Fprint(w, string(response))
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", welcomeHandler)
	r.HandleFunc("/create/", genSessionHandler)
	r.HandleFunc("/update/{sessionKey}", updateSessionHandler)
	r.HandleFunc("/join/{sessionKey}", joinSessionHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}