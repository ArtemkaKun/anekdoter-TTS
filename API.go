package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
	"strconv"
)

var Router = mux.NewRouter()

func init() {
	Router.HandleFunc("/audioJoke", GetAudioJoke).Methods("GET")
}

func GetAudioJoke(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "audio/ogg")

	CreateVoiceFile()
	WriteFileToRequest(writer)
}

func WriteFileToRequest(writer http.ResponseWriter) {
	audioFile, err := os.Open("output.ogg")
	if err != nil {
		// return 404 HTTP response code for File not found
		http.Error(writer, "Update file not found.", 404)
		return
	}

	defer audioFile.Close()

	fileHeader := make([]byte, 512)                // 512 bytes is sufficient for http.DetectContentType() to work
	audioFile.Read(fileHeader)                     // read the first 512 bytes from the updateFile
	fileType := http.DetectContentType(fileHeader) // set the type

	fileInfo, _ := audioFile.Stat()
	fileSize := fileInfo.Size()

	//Transmit the headers
	writer.Header().Set("Expires", "0")
	writer.Header().Set("Content-Transfer-Encoding", "binary")
	writer.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")
	writer.Header().Set("Content-Disposition", "attachment; filename="+"output.ogg")
	writer.Header().Set("Content-Type", fileType)
	writer.Header().Set("Content-Length", strconv.FormatInt(fileSize, 10))

	//Send the file
	audioFile.Seek(0, 0)       // reset back to position since we've read first 512 bytes of data previously
	io.Copy(writer, audioFile) // transmit the updatefile bytes to the client
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
