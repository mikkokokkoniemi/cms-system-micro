package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mikkokokkoniemi/cms-system-micro/engine/database"
	hashids "github.com/speps/go-hashids"
)

var dao = database.MongoDAO{}
var hashid = hashids.HashID{}

func init() {
	dao.Server = "localhost"
	dao.Database = "cms_engine"
	dao.Connect()

	hd := hashids.NewData()
	hd.Salt = "secret_salt"
	hd.MinLength = 8
	h, _ := hashids.NewWithData(hd)
	hashid = *h
}

func respondError(w http.ResponseWriter, code int, msg string) {
	respondJSON(w, code, map[string]string{"error": msg})
}

func respondJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.WriteHeader(code)
	w.Write(response)
}
