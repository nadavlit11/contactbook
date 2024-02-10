package dao

import (
	"contactbook/database"
	"contactbook/models"
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/cast"
	"sync"
)

var contactsCacheDaoOnce sync.Once
var contactsCacheDao ContactsCacheDao

type ContactsCacheDao interface {
	CacheContacts(userId int, contacts []models.Contact) error
	GetContacts(userId int) ([]models.Contact, error)
	AddContact(userId int, contact models.Contact) error
}

type ContactsCacheDaoImpl struct {
	redisClient database.RedisClient
}

func NewContactsCacheDao(
	redisClient database.RedisClient,
) ContactsCacheDao {
	contactsCacheDaoOnce.Do(func() {
		contactsCacheDao = &ContactsCacheDaoImpl{
			redisClient: redisClient,
		}
	})
	return contactsCacheDao
}

func (d *ContactsCacheDaoImpl) CacheContacts(userId int, contacts []models.Contact) error {
	rConn := d.redisClient.GetConn()

	contactsJson, err := json.Marshal(contacts)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	err = rConn.Set(context.Background(), cast.ToString(userId), cast.ToString(contactsJson), 0).Err()
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (d *ContactsCacheDaoImpl) GetContacts(userId int) ([]models.Contact, error) {
	rConn := d.redisClient.GetConn()

	contactsJson, err := rConn.Get(context.Background(), cast.ToString(userId)).Result()
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	var contacts []models.Contact
	err = json.Unmarshal([]byte(contactsJson), &contacts)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return contacts, nil
}

func (d *ContactsCacheDaoImpl) AddContact(userId int, contact models.Contact) error {
	contacts, err := d.GetContacts(userId)
	if err != nil {
		log.Error(err.Error())
	}

	contacts = append(contacts, contact)
	err = d.CacheContacts(userId, contacts)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}
