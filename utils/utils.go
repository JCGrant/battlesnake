package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/JCGrant/battlesnake/api"
)

func Respond(res http.ResponseWriter, obj interface{}) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(obj)
	res.Write([]byte("\n"))
}

func DecodeSnakeRequest(req *http.Request) (api.SnakeRequest, error) {
	result := api.SnakeRequest{}
	err := json.NewDecoder(req.Body).Decode(&result)
	return result, err
}

func Dump(obj interface{}) {
	data, err := json.MarshalIndent(obj, "", "  ")
	if err == nil {
		log.Printf(string(data))
	}
}
