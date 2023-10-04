package controllers

import (
	"bytes"
	"encoding/json"
	"middleware/api/utils"
	"middleware/config"
	"net/http"
	"io/ioutil"
)

type Endpoint struct {
	Endpoint string `json:"Endpoint"`
}

type Response struct{
	Token string `json:"token"`
}

var apiKey = config.ApiKey

var data interface{}

func MiddlewarePost(w http.ResponseWriter, r *http.Request){
    body := utils.BodyParser(r)
    var endpoint Endpoint
    err := json.Unmarshal(body, &endpoint)
    if err != nil {
		utils.ToJson(w, "Json incorrect", http.StatusBadRequest)
    }

	if endpoint.Endpoint == "/connexion"{
		PostConnexion(w, r, body)
	}else if endpoint.Endpoint == "/freezebee"{
		PostFreezeBee(w, r, body)
	}else{
		utils.ToJson(w, "Wrong Endpoint", http.StatusBadRequest)
	}
}

func PostConnexion(w http.ResponseWriter, r *http.Request, body []byte){

    // URL de l'API que vous souhaitez appeler
    url := "http://connexion:9000/connexion"
	contentType := "application/json"

    client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
        utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("ApiKey", apiKey)

	resp, err := client.Do(req)
	if err != nil {
        utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
        utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	var response Response

	err = json.Unmarshal(body, &response)
    if err != nil {
        utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
        return
    }

	utils.ToJson(w, response, http.StatusOK)
}

func PostFreezeBee(w http.ResponseWriter, r *http.Request, body []byte) {
    // URL de l'API que vous souhaitez appeler
    url := "http://freezebee:9200/freezebee"
    contentType := "application/json"

    client := &http.Client{}
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
    if err != nil {
        utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
        return
    }
    req.Header.Add("Content-Type", contentType)
    req.Header.Add("ApiKey", apiKey)
	req.Header.Add("Authorization", r.Header.Get("Authorization"))

    resp, err := client.Do(req)
    if err != nil {
        utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
        return
    }
    defer resp.Body.Close()

    responseBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
        return
    }

    // Décodez la réponse JSON en tant que chaîne
    var responseString string
    err = json.Unmarshal(responseBody, &responseString)
    if err != nil {
        utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
        return
    }

    // responseBody contient maintenant la réponse JSON sous forme de chaîne
    utils.ToJson(w, responseString, resp.StatusCode)
}
