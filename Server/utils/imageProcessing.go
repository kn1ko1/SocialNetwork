package utils

import (
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

const (
	maxFileSize = 20 << 20 // 20MB
	dirPath     = "static/uploadedFiles/images"
)

var supportedFileTypes = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
	"image/gif":  true,
}

func ImageProcessing(w http.ResponseWriter, r *http.Request, file multipart.File, fileHeader multipart.FileHeader) (string, error) {
	if fileHeader.Size > maxFileSize {
		fileHeaderErr := errors.New("file is too big")
		HandleError("File is too big!!", fileHeaderErr)
		return "", fileHeaderErr
	} else if !supportedFileTypes[fileHeader.Header.Get("Content-Type")] {
		supportedFileTypesErr := errors.New("file type is not supported")
		HandleError("File type is not supported!!", errors.New("file type is not supported"))
		return "", supportedFileTypesErr
	}
	// Create a temporary file in the given directory with a unique name
	osFile, createTempErr := os.CreateTemp(dirPath, "upload-*.jpg")
	if createTempErr != nil {
		HandleError("Error creating file: ", createTempErr)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return "", createTempErr
	}
	defer osFile.Close()
	// Copy the contents of the file to the file created above
	_, copyErr := io.Copy(osFile, file)
	if copyErr != nil {
		HandleError("Error copying file: ", copyErr)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return "", copyErr
	}
	return osFile.Name(), nil
}
