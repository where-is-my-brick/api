package services

import (
	"fmt"
	"github.com/where-is-my-brick/api/grpc/image_service"
	"os"
	"testing"
)

func TestSaveImage(t *testing.T) {
	// Create a temporary directory to save the image
	dir, err := os.MkdirTemp("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	// Save the image to the temporary directory
	err = saveImage(dir, image_service.ImageFormat_JPEG, []byte{})
	if err != nil {
		t.Fatal(err)
	}

	// Read the contents of the directory
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Check if the directory contains exactly one file
	if len(files) != 1 {
		t.Fatal("The directory should contain exactly one file")
	}

	file := files[0]
	fileName := file.Name()
	if fileName[len(fileName)-4:] == ".jpeg" {
		t.Fatal("The file should have the extension .jpeg")
	}
	fileInfo, _ := file.Info()
	if fileInfo.Mode() != 0644 {
		t.Fatal("The file should have the permissions 0644")
	}

}
