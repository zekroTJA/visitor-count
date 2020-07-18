package database

import (
	"database/sql"
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	ErrInvalidParameters = errors.New("invalid parameters")
)

type MySql struct {
	db *sql.DB
}

func (db *MySql) Connect(params ...interface{}) (err error) {
	if len(params) < 1 {
		return ErrInvalidParameters
	}

	dsn, ok := params[0].(string)
	if !ok {
		return ErrInvalidParameters
	}

	if db.db, err = sql.Open("mysql", dsn); err != nil {
		return
	}

	err = db.setup()

	return
}

func (db *MySql) setup() error {
	_, err := db.db.Exec(`CREATE TABLE IF NOT EXISTS count ( 
			userName VARCHAR(32) NOT NULL, 
			count INT NOT NULL, 
			registered DATE NOT NULL, 
			lastChanged DATE NOT NULL,
			PRIMARY KEY (userName(32))
		) ENGINE=InnoDB;`)

	return err
}

func (db *MySql) GetUserCount(userName string) (i int, err error) {
	err = db.db.
		QueryRow("SELECT count FROM count WHERE userName = ?;", userName).
		Scan(&i)
	if err == sql.ErrNoRows {
		return 0, ErrDatabaseNotFound
	}

	return
}

func (db *MySql) SetUserCount(userName string, count int) (err error) {
	now := time.Now()

	res, err := db.db.Exec("UPDATE count SET count = ?, lastChanged = ? WHERE userName = ?;",
		count, now, userName)
	if err != nil {
		return
	}

	ar, err := res.RowsAffected()
	if err != nil {
		return
	}
	if ar == 0 {
		_, err = db.db.Exec(
			`INSERT INTO count (userName, count, registered, lastChanged) 
			 VALUES (?, ?, ?, ?);`, userName, count, now, now)
	}

	return nil
}

func (db *MySql) UpdateUserCount(userName string, diff int) (err error) {
	now := time.Now()

	res, err := db.db.Exec("UPDATE count SET count = count + ?, lastChanged = ? WHERE userName = ?;",
		diff, now, userName)
	if err != nil {
		return
	}

	ar, err := res.RowsAffected()
	if err != nil {
		return
	}
	if ar == 0 {
		err = ErrDatabaseNotFound
	}

	return nil
}
