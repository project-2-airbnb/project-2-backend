package handler

import (
	"context"
	"io"
	"os"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

type RoomRequest struct {
	RoomPicture     string `json:"room_picture" form:"room_picture"`
	RoomName        string `json:"room_name" form:"room_name"`
	Description     string `json:"description" form:"description"`
	Location        string `json:"location" form:"location"`
	QuantityGuest   int    `json:"quantity_guest" form:"quantity_guest"`
	QuantityBedroom int    `json:"quantity_bedroom" form:"quantity_bedroom"`
	QuantityBed     int    `json:"quantity_bed" form:"quantity_bed"`
	Price           int    `json:"price" form:"price"`
}

func (r RoomRequest) uploadToCloudinary(file io.Reader, filename string) (string, error) {
	// Konfigurasi Cloudinary
	cloudinaryURL := os.Getenv("CLOUDINARY_URL")
	cld, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		return "", err
	}

	// Upload file ke Cloudinary
	uploadParams := uploader.UploadParams{
		Folder:   "user_pictures",
		PublicID: filename,
	}
	uploadResult, err := cld.Upload.Upload(context.Background(), file, uploadParams)
	if err != nil {
		return "", err
	}

	// Ambil URL publik dari hasil unggah
	publicURL := uploadResult.URL
	return publicURL, nil
}
