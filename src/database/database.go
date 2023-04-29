package database

import (
	types "tg-music-bot/src/telegram_types"
)

type Database interface {
	GetAllUsers() ([]types.User, error)
	GetUserById(id int) (types.User, error)
	AddUser(user types.User) error
}
