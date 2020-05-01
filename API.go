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
	writer.Header().Set("Content-Type", "audio/ogg")

	WriteFileToRequest(writer)
}

func WriteFileToRequest(writer http.ResponseWriter) {
	audioContent := CreateVoiceFile()

	//fileHeader := make([]byte, 512) // 512 bytes is sufficient for http.DetectContentType() to work
	//audioFile.Read(fileHeader)                     // read the first 512 bytes from the updateFile
	//fileType := http.DetectContentType(fileHeader) // set the type
	//
	//fileInfo, _ := audioFile.Stat()
	//fileSize := fileInfo.Size()
	//
	////Transmit the headers
	writer.Header().Set("Expires", "0")
	writer.Header().Set("Content-Transfer-Encoding", "binary")
	writer.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")
	writer.Header().Set("Content-Disposition", "attachment; filename="+"output.ogg")
	//writer.Header().Set("Content-Length", strconv.FormatInt(fileSize, 10))
	//
	////Send the file
	//audioFile.Seek(0, 0)       // reset back to position since we've read first 512 bytes of data previously

	writer.Write(audioContent)
	//io.Copy(writer, audioContent) // transmit the updatefile bytes to the client
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
