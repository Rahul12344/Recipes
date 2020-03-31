package controllers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Rahul12344/Recipes/util/parsing"
)

// ImageController controls image actions
type ImageController struct {
}

//NewImageController creates new image controller
func NewImageController() *ImageController {
	return &ImageController{}
}

//Setup mounts
func (ic *ImageController) Setup(r *mux.Router) {
	r.HandleFunc("/find", ic.ParseImage).Methods("GET")
}

//ParseImage parses image for text
func (ic *ImageController) ParseImage(w http.ResponseWriter, r *http.Request) {
	imageParser := parsing.NewParser()
	imageParser.Detect(r.URL.Query().Get("filename"))
}
