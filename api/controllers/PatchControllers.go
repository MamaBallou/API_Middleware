package controllers

import (
	"bytes"
	"encoding/json"
	"middleware/api/utils"
	"net/http"
	"io/ioutil"
)

func MiddlewarePatch(w http.ResponseWriter, r *http.Request){
    body := utils.BodyParser(r)
    var endpoint Endpoint
    err := json.Unmarshal(body, &endpoint)
    if err != nil {
		utils.ToJson(w, "Json incorrect", http.StatusBadRequest)
    }

 	if endpoint.Endpoint == "/freezebee"{
		PatchFreezeBee(w, r, body)
	}else{
		utils.ToJson(w, "Wrong Endpoint", http.StatusBadRequest)
	}
}

func PatchFreezeBee(w http.ResponseWriter, r *http.Request, body []byte) {
    // URL de l'API que vous souhaitez appeler
    url := "http://freezebee:9200/freezebee"
    contentType := "application/json"

    client := &http.Client{}
    req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(body))
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