package handler

import (
	"context"
	"io"
	"os"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

type UserRequest struct {
	FullName       string `json:"fullname" form:"fullname"`
	Email          string `json:"email" form:"email"`
	Password       string `json:"password" form:"password"`
	RetypePassword string `json:"retype_password" form:"retype_password"`
	Address        string `json:"address" form:"address"`
	PhoneNumber    string `json:"phone_number" form:"phone_number"`
	PictureProfile string `json:"picture_profile" form:"picture_profile"`
	UserType       string `json:"user_type" form:"user_type"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	UserType string `json:"user_type" form:"user_type"`
}

func (u UserRequest) uploadToCloudinary(file io.Reader, filename string) (string, error) {
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
