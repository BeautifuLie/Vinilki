package mongostorage_test

import (
	"errors"
	"program/model"
	"program/storage/mongostorage"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func initConnection() *mongostorage.DatabaseStorage {

	var ms, _ = mongostorage.NewDatabaseStorage("mongodb://localhost:27017")
	return ms
}
func TestCreateUser(t *testing.T) {
	db := initConnection()
	exist := errors.New("this username already exists")
	user1 := model.User{
		Username:    "Denys",
		Email:       "1233",
		Create_time: time.Now(),
	}
	user2 := model.User{
		Username:    "Denys",
		Email:       "1233",
		Create_time: time.Now(),
	}
	user3 := model.User{
		Username:    "Denys",
		Email:       "email1",
		Create_time: time.Now(),
	}
	user4 := model.User{
		Username:    "Denys",
		Email:       "email2",
		Create_time: time.Now(),
	}
	t.Run("NoUsers", func(t *testing.T) {
		userID, err := db.AssignID()
		require.NoError(t, err)
		assert.Equal(t, userID.User_id, 0)

	})
	t.Run("CorrectUser", func(t *testing.T) {
		err := db.CreateUser(user1)
		require.NoError(t, err)

	})
	t.Run("SameUser", func(t *testing.T) {
		err := db.CreateUser(user2)
		require.Error(t, err)
		assert.Equal(t, err, exist)

	})
	t.Run("CorrectID", func(t *testing.T) {
		userID, err := db.AssignID()
		require.NoError(t, err)
		user3.User_id = userID.User_id + 1
		err = db.CreateUser(user3)
		require.NoError(t, err)
		userID2, err := db.AssignID()
		require.NoError(t, err)
		user4.User_id = userID2.User_id + 1
		err = db.CreateUser(user4)
		require.NoError(t, err)
		assert.NotEqual(t, user3.User_id, user4.User_id)

	})

}
