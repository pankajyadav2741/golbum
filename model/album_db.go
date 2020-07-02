package model

import (
	"fmt"
	"net/http"

	"github.com/pankajyadav2741/golbum/utils"
)

//Albums : Slice of struct
var Albums []Album

//AlbumExists : Check if an album exists
func AlbumExists(albName string) (bool, *utils.ApplicationError) {
	for _, item := range Albums {
		if item.Name == albName {
			return true, &utils.ApplicationError{
				Message:    fmt.Sprintf("Album %v found", albName),
				StatusCode: http.StatusConflict,
			}
		}
	}
	return false, &utils.ApplicationError{
		Message:    fmt.Sprintf("Album %v not found", albName),
		StatusCode: http.StatusNotFound,
	}
}

//ImageExists : Check if an image exists in an album
func ImageExists(albName, imgName string) (bool, *utils.ApplicationError) {
	for idx := 0; idx < len(Albums); idx++ {
		if Albums[idx].Name == albName {
			for i := 0; i < len(Albums[idx].Image); i++ {
				if Albums[idx].Image[i].Name == imgName {
					return true, &utils.ApplicationError{
						Message:    fmt.Sprintf("Image %v found in album %v", imgName, albName),
						StatusCode: http.StatusConflict,
					}
				}
			}
		}
	}
	return false, &utils.ApplicationError{
		Message:    fmt.Sprintf("Image %v not found in album %v", imgName, albName),
		StatusCode: http.StatusNotFound,
	}
}

//ShowAlbum : Show all albums
func ShowAlbum() []Album {
	return Albums
}

//AddAlbum : Create a new album
func AddAlbum(albName string) *utils.ApplicationError {
	if ok, err := AlbumExists(albName); ok != false {
		return err
	}
	Albums = append(Albums, Album{Name: albName})
	return nil
}

//DeleteAlbum : Delete an existing album
func DeleteAlbum(albName string) *utils.ApplicationError {
	if ok, err := AlbumExists(albName); ok != true {
		return err
	}
	for idx, item := range Albums {
		if item.Name == albName {
			Albums = append(Albums[:idx], Albums[idx+1:]...)
			break
		}
	}
	return nil
}

//ShowImagesInAlbum : Show all images in an album
func ShowImagesInAlbum(albName string) ([]Image, *utils.ApplicationError) {
	if ok, err := AlbumExists(albName); ok != true {
		return nil, err
	}

	for _, item := range Albums {
		if item.Name == albName {
			return item.Image, nil
		}
	}
	return nil, nil
}

//ShowImage : Show a particular image inside an album
func ShowImage(albName, imgName string) (*Image, *utils.ApplicationError) {
	if ok, err := AlbumExists(albName); ok != true {
		return nil, err
	}

	if ok, err := ImageExists(albName, imgName); ok != true {
		return nil, err
	}

	for _, item := range Albums {
		if item.Name == albName {
			for i := 0; i < len(item.Image); i++ {
				if item.Image[i].Name == imgName {
					return &item.Image[i], nil
				}
			}
		}
	}
	return nil, nil
}

//AddImage : Create an image in an album
func AddImage(albName, imgName string) *utils.ApplicationError {
	if ok, err := AlbumExists(albName); ok != true {
		return err
	}

	if ok, err := ImageExists(albName, imgName); ok != true {
		return err
	}
	image := Image{Name: imgName}
	for idx, item := range Albums {
		if item.Name == albName {
			Albums[idx].Image = append(Albums[idx].Image, image)
			break
		}
	}
	return nil
}

//DeleteImage : Delete an image in an album
func DeleteImage(albName, imgName string) *utils.ApplicationError {
	if ok, err := AlbumExists(albName); ok != true {
		return err
	}

	if ok, err := ImageExists(albName, imgName); ok != false {
		return err
	}
	var alb []Album
	for idx, item := range Albums {
		if item.Name == albName {
			for i := 0; i < len(item.Image); i++ {
				if item.Image[i].Name == imgName {
					item.Image = append(item.Image[:i], item.Image[i+1:]...)
					alb = append(Albums[:idx], Album{Name: albName, Image: item.Image})
					Albums = append(alb, Albums[idx+1:]...)
					break
				}
			}
		}
	}
	return nil
}
