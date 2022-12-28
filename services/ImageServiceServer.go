package services

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/where-is-my-brick/api/grpc/image_service"
	"log"
	"os"
)

type ImageServiceServer struct {
	PathPrefix string
	image_service.UnimplementedImageServiceServer
}

// saveImage saves the image to the path specified by the category on the file system
func saveImage(category string, imageFormat image_service.ImageFormat, imageData []byte) error {
	// generate a unique id for the image
	var imageName = uuid.New().String()

	// TODO: convert image to the right format
	if imageFormat != image_service.ImageFormat_JPEG {
		const errorMessage = "does not support this image format"
		log.Fatalf(errorMessage)
		return errors.New(errorMessage)
	}

	// save the image to the file system
	err := os.WriteFile(category+"/"+imageName+".jpeg", imageData, 0644)
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
		return err
	}
	return nil
}

func (s *ImageServiceServer) UploadImage(ctx context.Context, req *image_service.UploadImageRequest) (*image_service.UploadImageResponse, error) {
	// Save the image to a database or file system
	imagePath := s.PathPrefix + req.Category
	err := saveImage(imagePath, req.ImageFormat, req.ImageData)
	if err != nil {
		log.Fatal(err)
		return &image_service.UploadImageResponse{
			Success: false,
		}, err
	}

	return &image_service.UploadImageResponse{
		Success: true,
	}, nil
}
