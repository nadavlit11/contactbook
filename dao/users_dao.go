package dao

import (
	"contactbook/models"
	"fmt"
	"sync"
)

var usersDaoOnce sync.Once
var usersDao UsersDao

type UsersDao interface {
	Connect()
	Insert(user models.User) (int, error)
}

type UsersDaoImpl struct {
	usersDB   []models.User
	mu        sync.Mutex
	autoIncId int
}

func NewUsersDao() UsersDao {
	usersDaoOnce.Do(func() {
		usersDao = &UsersDaoImpl{}
	})
	return usersDao
}

func (d *UsersDaoImpl) Connect() {
	d.usersDB = make([]models.User, 0)
	fmt.Println("Connected with UsersDao")
}

func (d *UsersDaoImpl) Insert(user models.User) (int, error) {
	d.mu.Lock()

	d.autoIncId++
	user.ID = d.autoIncId
	d.usersDB = append(d.usersDB, user)

	d.mu.Unlock()
	return d.autoIncId, nil
}
