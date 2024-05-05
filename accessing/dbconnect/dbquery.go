package dbconnect

import (
	"accessing/dbtypes"
	"errors"
)

// ///////SQL QUERYS////////
const (
	insertUser           = `INSERT INTO users(username, services) VALUES($1, $2)`
	selectUser           = `SELECT 1 FROM users WHERE username=$1  LIMIT 1`
	selectUserByUsername = `SELECT username,  services FROM users WHERE username=$1 LIMIT 1`
	createTable          = `CREATE table users(
		username	text not null,
		services	jsonb
	)`
	deleteUser = `DELETE FROM users WHERE username=$1`
	updateUser = `UPDATE users SET services=$1 WHERE username=$2`
)

func CreateTable() (err error) {
	_, err = DataBase.Exec(Ctx, createTable)
	return
}

// ////////////////////////////////////////
// ADMIN FUNCTIONS //
// ////////////////////////////////////////

func InsertUser(user *dbtypes.User) error {

	// Check if user already exists
	ok, err := userAlreadyExists(user.Username)
	if err != nil {
		return err
	}

	if ok {
		return errors.New("user already exists")
	}

	_, err = DataBase.Exec(Ctx, insertUser, &user.Username, &user.Access)
	return err
}

func DeleteUser(username string) error {
	_, err := DataBase.Exec(Ctx, deleteUser, username)
	if err != nil {
		return err
	}
	return nil
}

func AddUserService(user *dbtypes.User) error {
	dbUser, err := GetUserByUsername(user.Username)
	if err != nil {
		return err
	}
	getUserWithNewServices(user, dbUser)
	_, err = DataBase.Exec(Ctx, updateUser, &dbUser.Access, &dbUser.Username)
	if err != nil {
		return err
	}
	return nil
}

func RemoveUserServices(user *dbtypes.User) error {
	dbUser, err := GetUserByUsername(user.Username)
	if err != nil {
		return err
	}
	getUserWithRemovedServices(user, dbUser)
	_, err = DataBase.Exec(Ctx, updateUser, &dbUser.Access, &dbUser.Username)
	if err != nil {
		return err
	}
	return nil
}

//////////////////////////////////////////
// USER FUNCTIONS //
//////////////////////////////////////////
