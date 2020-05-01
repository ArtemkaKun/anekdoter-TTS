package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var Router = mux.NewRouter()

func init() {
	Router.HandleFunc("/audioJoke", GetAudioJoke).Methods("GET")
}

func GetAudioJoke(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "application/ogg")

	WriteFileToRequest(writer)
}

func WriteFileToRequest(writer http.ResponseWriter) {
	audioContent := CreateVoiceFile()

	writer.Header().Set("Expires", "0")
	writer.Header().Set("Content-Transfer-Encoding", "binary")
	writer.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")
	writer.Header().Set("Content-Disposition", "attachment; filename="+"output.ogg")

	writer.Write(audioContent)
}

func GetJoke() (joke string) {
	res, err := http.Get("http://anekdoter.oa.r.appspot.com/joke")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&joke)

	return
}
