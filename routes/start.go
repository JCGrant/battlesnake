package main

import (
	"log"
	"net/http"

	"github.com/JCGrant/battlesnake/bot"
	"github.com/JCGrant/battlesnake/utils"
)

func Start(res http.ResponseWriter, req *http.Request) {
	decoded, err := utils.DecodeSnakeRequest(req)
	if err != nil {
		log.Printf("Bad request: %v", err)
	}
	utils.Dump(decoded)
	action := bot.Start(decoded)
	utils.Respond(res, action)
}
