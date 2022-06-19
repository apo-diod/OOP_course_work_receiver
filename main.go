package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const PORT = "6120"
const TRANSFORM_PORT = "6121"
const HOST = "127.0.0.1"

func main() {
	pairs := make(map[string]string)

	r := gin.Default()

	r.POST("/callback", func(ctx *gin.Context) {
		body, _ := ioutil.ReadAll(ctx.Request.Body)
		var req map[string]interface{}
		json.Unmarshal(body, &req)
		log.Println(req["id"])
		var resp map[string]interface{} = make(map[string]interface{})
		resp["id"] = pairs[req["id"].(string)]
		resp["data"] = req["data"]
		toSendbts, _ := json.Marshal(resp)
		log.Println(string(toSendbts))
		http.Post("http://"+HOST+":"+TRANSFORM_PORT+"/use_module", "application/json", bytes.NewBuffer(toSendbts))
		ctx.JSON(http.StatusOK, nil)
	})

	r.POST("/add_module", func(ctx *gin.Context) {
		body, _ := ioutil.ReadAll(ctx.Request.Body)
		var req map[string]interface{}
		json.Unmarshal(body, &req)
		module := req["module"].(string)
		settingsmap := req["settings"]
		settingsbts, _ := json.Marshal(settingsmap)
		settings := string(settingsbts)
		id := newModule(module, settings)
		if id == "0" {
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"id": id})
	})

	r.POST("/link", func(ctx *gin.Context) {
		body, _ := ioutil.ReadAll(ctx.Request.Body)
		var req map[string]interface{}
		json.Unmarshal(body, &req)
		first := req["first"].(string)
		second := req["second"].(string)
		pairs[first] = second
		ctx.JSON(http.StatusOK, nil)
	})

	r.Run(":" + PORT)
}
