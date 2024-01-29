package repo

import (
	"database/sql"
	"log"
	"socialnetwork/models"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Field commented for implementation example reasons:
// Should take as a field a private DB instance
//
// The repo instance is then responsible for maintaining
// the concrete DB connection
type DummyRepository struct {
	identityDb *sql.DB
	businessDb *sql.DB
}

// Constructor function
func NewDummyRepository() *DummyRepository {
	// The DB field would be constructed properly here!
	identityDb, err := sql.Open(dbDriver, identityDbPath)
	if err != nil {
		log.Fatal(err)
	}
	businessDb, err := sql.Open(dbDriver, businessDbPath)
	if err != nil {
		log.Fatal(err)
	}
	return &DummyRepository{identityDb: identityDb, businessDb: businessDb}
}

// Sim user creation
func (r *DummyRepository) CreateUser(user models.User) (models.User, error) {
	return user, nil
}

// Sim User retrieval
func (r *DummyRepository) GetAllUsers() ([]models.User, error) {
	ctime := time.Now().UTC().UnixMilli()
	dummyUsers := []models.User{
		{
			UserId:            1,
			Bio:               "I am cool",
			CreatedAt:         ctime,
			DOB:               ctime,
			Email:             "a@b.com",
			EncryptedPassword: "really safe",
			FirstName:         "Test",
			ImageUrl:          "example/url",
			IsPublic:          false,
			LastName:          "User",
			UpdatedAt:         ctime,
			Username:          "test",
		},
		{
			UserId:            2,
			Bio:               "I am not cool",
			CreatedAt:         ctime,
			DOB:               ctime,
			Email:             "b@c.com",
			EncryptedPassword: "really safe",
			FirstName:         "Example",
			ImageUrl:          "example/url/2",
			IsPublic:          true,
			LastName:          "User",
			UpdatedAt:         ctime,
			Username:          "example",
		},
	}
	return dummyUsers, nil
}

func (r *DummyRepository) CreatePost(post models.Post) (models.Post, error) {
	return post, nil
}

func (r *DummyRepository) GetAllPosts() ([]models.Post, error) {
	ctime := time.Now().UTC().UnixMilli()
	dummyPosts := []models.Post{
		{
			PostId:    1,
			Body:      "This is fun",
			CreatedAt: ctime,
			GroupId:   0,
			ImageURL:  "example/url",
			UpdatedAt: ctime,
			UserId:    1,
		},
		{
			PostId:    2,
			Body:      "This is not fun",
			CreatedAt: ctime,
			GroupId:   0,
			ImageURL:  "example/url2",
			UpdatedAt: ctime,
			UserId:    2,
		},
	}
	return dummyPosts, nil
}

func (r *DummyRepository) CreateNotification(notification models.Notification) (models.Notification, error) {
	return notification, nil
}

func (r *DummyRepository) GetNotificationById(notificationId int) (models.Notification, error) {
	ctime := time.Now().UTC().UnixMilli()

	notification := models.Notification{
		NotificationId:   notificationId,
		CreatedAt:        ctime,
		NotificationType: "dummy type",
		ObjectId:         1,
		SenderId:         1,
		Status:           "dummy status",
		TargetId:         2,
		UpdatedAt:        ctime,
	}

	return notification, nil
}
