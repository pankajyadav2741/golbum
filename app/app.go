package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pankajyadav2741/golbum/controller"
)

//StartApp : Start Application
func StartApp() {
	//Initialize Router
	myRouter := mux.NewRouter().StrictSlash(true)

	//Show all albums
	myRouter.HandleFunc("/", controller.ShowAlbum).Methods(http.MethodGet)
	//Create a new album
	myRouter.HandleFunc("/{album}", controller.AddAlbum).Methods(http.MethodPost)
	//Delete an existing album
	myRouter.HandleFunc("/{album}", controller.DeleteAlbum).Methods(http.MethodDelete)

	//Show all images in an album
	myRouter.HandleFunc("/{album}", controller.ShowImagesInAlbum).Methods(http.MethodGet)
	//Show a particular image inside an album
	myRouter.HandleFunc("/{album}/{image}", controller.ShowImage).Methods(http.MethodGet)
	//Create an image in an album
	myRouter.HandleFunc("/{album}/{image}", controller.AddImage).Methods(http.MethodPost)
	//Delete an image in an album
	myRouter.HandleFunc("/{album}/{image}", controller.DeleteImage).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":5000", myRouter))
}
