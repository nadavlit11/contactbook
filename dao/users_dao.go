package dao

import (
	"contactbook/database"
	"contactbook/models"
	"context"
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/cast"
	"sync"
)

var usersDaoOnce sync.Once
var usersDao UsersDao

type UsersDao interface {
	Insert(user models.User) (int, error)
}

type UsersDaoImpl struct {
	postgresClient database.PostgresClient
	redisClient    database.RedisClient
}

func NewUsersDao(
	postgresClient database.PostgresClient,
	redisClient database.RedisClient,
) UsersDao {
	usersDaoOnce.Do(func() {
		usersDao = &UsersDaoImpl{
			postgresClient: postgresClient,
			redisClient:    redisClient,
		}
	})
	return usersDao
}

func (d *UsersDaoImpl) Insert(user models.User) (int, error) {
	pConn := d.postgresClient.GetConn()

	query := `INSERT INTO users (name) VALUES ($1) RETURNING id`

	row := pConn.QueryRowContext(context.Background(), query, user.Name)

	var userId int
	err := row.Scan(&userId)
	if err != nil {
		log.Error(err.Error())
		return 0, err
	}
	return cast.ToInt(userId), nil
}
