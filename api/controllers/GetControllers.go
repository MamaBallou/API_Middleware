package controllers

import (
	"bytes"
	"encoding/json"
	"middleware/api/utils"
	"net/http"
	"io/ioutil"
)

func MiddlewareGet(w http.ResponseWriter, r *http.Request){
    endpoint := r.URL.Query().Get("Endpoint")
    table := r.URL.Query().Get("Table")

    requestData := struct {
        Endpoint string `json:"Endpoint"`
        Table    string `json:"Table"`
    }{
        Endpoint: endpoint,
        Table:    table,
    }

    // Convertir l'objet en JSON
    requestDataJSON, err := json.Marshal(requestData)
    if err != nil {
        http.Error(w, "Erreur de création du JSON", http.StatusInternalServerError)
        return
    }

	GetFreezeBee(w, r, requestDataJSON)
	
}

func GetFreezeBee(w http.ResponseWriter, r *http.Request, body []byte) {
    // URL de l'API que vous souhaitez appeler
    url := "http://freezebee:9200/freezebee"
    contentType := "application/json"

    client := &http.Client{}
    req, err := http.NewRequest("GET", url, bytes.NewBuffer(body))
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

    // Lecture de la réponse en tant que tableau de bytes
    responseBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
        return
    }

    // Vous pouvez maintenant essayer de décoder la réponse en tant que différents types
    var responseSlice []map[string]interface{}
    err = json.Unmarshal(responseBody, &responseSlice)
    if err == nil {
        // Si la réponse est un tableau JSON, vous avez réussi à la décoder
        // responseSlice contiendra les données sous forme de []map[string]interface{}
        utils.ToJson(w, responseSlice, resp.StatusCode)
        return
    }

    // Si la réponse n'est pas un tableau JSON, essayez de le décoder en tant qu'objet JSON
    var responseObject map[string]interface{}
    err = json.Unmarshal(responseBody, &responseObject)
    if err == nil {
        // Si la réponse est un objet JSON, vous avez réussi à la décoder
        // responseObject contiendra les données sous forme de map[string]interface{}
        utils.ToJson(w, responseObject, resp.StatusCode)
        return
    }

    // Si la réponse ne peut pas être décoder comme un tableau ou un objet JSON, vous pouvez renvoyer une erreur
    utils.ToJson(w, "Réponse JSON invalide", http.StatusUnprocessableEntity)
}