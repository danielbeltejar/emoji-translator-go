package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"strconv"
	"time"
)

type RequestBody struct {
	InputText string `json:"input-text"`
}

type ResponseBody struct {
	OutputText  string `json:"output_text"`
	BackendTime string `json:"backend_time"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/v1/translator/go", translateRealtime).Methods("POST")

	corsHandler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":8000", corsHandler))
}

func translateRealtime(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	inputText := requestBody.InputText

	if inputText == "" {
		response := ResponseBody{
			OutputText:  "",
			BackendTime: "0",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	outputText := Translator(inputText)

	elapsedTimeMs := float64(time.Since(startTime).Microseconds()) / 1000

	response := ResponseBody{
		OutputText:  outputText.String(),
		BackendTime: strconv.FormatFloat(elapsedTimeMs, 'f', 2, 64),
	}
	json.NewEncoder(w).Encode(response)
}
