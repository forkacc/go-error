package main

import (
	"errors"
	"fmt"
	"sql"
)

type User struct {
	id   int
	name string
}

func GetUser(uid int) (User, error) {
	var user User
	row, err := db.Query("select uid, uname from users where id = ?", uid)
	if err != nil {
		// error in db Query function, return directly
		return user, err
	}
	defer row.Close()
	for row.Next() {
		err := row.Scan(&user.id, &user.name)
		if err != nil {
			break
		}
	}
	// need context so we wrap the error with uid, and let caller to handle this error
	return user, errors.Wrap(row.Err(), "failed to get user with uid: %d", uid)
}

func main() {
	// Assume this is the controller and call the dao service
	User, err := GetUser(0)
	if err != nil {
		// check sentinel error and decide what to do
		if errors.Is(err, sql.ErrNoRows) {
			// do something and log error
		} else {
			// do something and log error
		}
	}
	fmt.Printf("Get user - id: %d, name: %s", User.id, User.name)
}
