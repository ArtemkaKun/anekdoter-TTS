package main

import (
	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"context"
	"google.golang.org/api/option"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
	"io/ioutil"
	"log"
	"math/rand"
)

const (
	POLINA string = "ru-RU-Wavenet-A"
	STASIK string = "ru-RU-Wavenet-B"
	MARINA string = "ru-RU-Wavenet-C"
	GEORGY string = "ru-RU-Wavenet-D"
)

var Client *texttospeech.Client
var CTX context.Context

func init() {
	CTX = context.Background()
	var err error

	Client, err = texttospeech.NewClient(CTX, option.WithCredentialsFile("Anekdoter-5359d800761b.json"))
	if err != nil {
		log.Fatal(err)
	}
}

func PrepareSynthesizer(joke string) (req texttospeechpb.SynthesizeSpeechRequest) {
	var voice string

	switch rand.Intn(4) {
	case 0:
		voice = POLINA
	case 1:
		voice = STASIK
	case 2:
		voice = MARINA
	case 3:
		voice = GEORGY
	}

	req = texttospeechpb.SynthesizeSpeechRequest{
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: joke},
		},
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "ru-RU",
			Name:         voice,
		},
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_OGG_OPUS,
		},
	}

	return
}

func CreateVoiceFile() {
	synthesizer := PrepareSynthesizer(GetJoke())
	resp, err := Client.SynthesizeSpeech(CTX, &synthesizer)
	if err != nil {
		log.Fatal(err)
	}

	filename := "output.ogg"
	err = ioutil.WriteFile(filename, resp.AudioContent, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
