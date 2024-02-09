package api

import (
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"socialnetwork/models"
	"socialnetwork/repo"
	"socialnetwork/utils"
	"time"
)

var CurrentTime = time.Now()

var Timestamp = CurrentTime.Unix()

var R = repo.NewDummyRepository()

var CommentExample = &models.Comment{
	CommentId: 1,
	Body:      "suicide squad",
	CreatedAt: Timestamp,
	ImageURL:  "imageurl",
	PostId:    1,
	UpdatedAt: Timestamp,
	UserId:    1,
}

var EventExample = &models.Event{
	EventId:     1,
	CreatedAt:   Timestamp,
	DateTime:    Timestamp,
	Description: "Event",
	GroupId:     1,
	Title:       "Event",
	UpdatedAt:   Timestamp,
	UserId:      1,
}

var EventUserExample = &models.EventUser{
	EventUserId: 1,
	CreatedAt:   Timestamp,
	EventId:     1,
	UpdatedAt:   Timestamp,
	UserId:      1,
}

var GroupExample = &models.Group{
	GroupId:     1,
	CreatedAt:   Timestamp,
	CreatorId:   1,
	Description: "Group Example",
	Title:       "Group",
	UpdatedAt:   Timestamp,
}

const (
	maxFileSize = 20 << 20 // 20MB
	dirPath     = "static/uploadFiles/images"
)

var supportedFileTypes = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
	"image/gif":  true,
}

func ImageProcessing(w http.ResponseWriter, r *http.Request, file multipart.File, fileHeader multipart.FileHeader) (string, error) {
	if fileHeader.Size > maxFileSize {
		fileHeaderErr := errors.New("file is too big")
		utils.HandleError("File is too big!!", fileHeaderErr)
		return "", fileHeaderErr
	} else if !supportedFileTypes[fileHeader.Header.Get("Content-Type")] {
		supportedFileTypesErr := errors.New("file type is not supported")
		utils.HandleError("File type is not supported!!", errors.New("file type is not supported"))
		return "", supportedFileTypesErr
	}
	//create a file in the given directory with the suffix .jpg
	osFile, createTempErr := os.CreateTemp(dirPath, "*.jp*g, *.png, *.gif")
	if createTempErr != nil {
		utils.HandleError("Error creating file: ", createTempErr)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return "", createTempErr
	}
	defer osFile.Close()
	//copy the contents of the file to the file created above
	_, copyErr := io.Copy(osFile, file)
	if copyErr != nil {
		utils.HandleError("Error copying file: ", copyErr)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return "", copyErr
	}
	return osFile.Name(), nil
}
