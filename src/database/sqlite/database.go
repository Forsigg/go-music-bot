package sqlite

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	types "tg-music-bot/src/telegram_types"
)

const AddUserError = "unknown error while insterting user in database"

type SQLiteDB struct {
	db *sql.DB
}

// NewSQLiteDB constructor of SQLiteDB
func NewSQLiteDB(dbFile string) (*SQLiteDB, error) {
	sqlDB, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}

	if _, err = sqlDB.Exec(schemaSQL); err != nil {
		return nil, err
	}

	db := SQLiteDB{db: sqlDB}
	return &db, nil
}

func (db *SQLiteDB) GetAllUsers() ([]types.User, error) {
	rows, err := db.db.Query(allUsersSQL)
	if err != nil {
		return nil, err
	}

	var users []types.User
	for rows.Next() {
		u := types.User{}
		err = rows.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Username)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (db *SQLiteDB) GetUserById(id int) (types.User, error) {
	var u types.User
	row, err := db.db.Query(userByIdSQL, id)
	if err != nil {
		return u, err
	}

	err = row.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Username)
	if err != nil {
		return u, err
	}

	return u, nil
}

func (db *SQLiteDB) AddUser(user types.User) error {
	stmt, err := db.db.Prepare(addUserSQL)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(user.Id, user.FirstName, user.LastName, user.Username)
	if err != nil {
		return err
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowCount != 1 {
		return fmt.Errorf(AddUserError)
	}
	return nil
}
