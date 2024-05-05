package dbconnect

import (
	"authentification/dbtypes"
	"errors"

	"github.com/jackc/pgx/v5"
)

const (
	insertUser           = `INSERT INTO users(username, password, services) VALUES($1, $2, $3)`
	selectUser           = `SELECT 1 FROM users WHERE username=$1 AND password=$2 LIMIT 1`
	selectUserByUsername = `SELECT username, password, services FROM users WHERE username=$1 LIMIT 1`
	createTable          = `CREATE table users(
		username	text not null,
		password	text not null,
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

	_, err = DataBase.Exec(Ctx, insertUser, &user.Username, &user.Password, &user.Access)
	return err
}

func GetUserByUsername(username string) (*dbtypes.User, error) {
	row := DataBase.QueryRow(Ctx, selectUserByUsername, username)
	user := dbtypes.User{}
	err := row.Scan(&user.Username, &user.Password, &user.Access)
	if err != nil {
		return nil, err
	}
	if user.Username == "" || user.Password == "" {
		return nil, errors.New("user not found")
	}
	return &user, nil
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

/////////////////////////////////////////////////////////////////

func getUserWithRemovedServices(user, dbUser *dbtypes.User) {
	recordMap := make(map[string]struct{})
	newRecord := make([]string, 0)
	for _, record := range user.Access.Archive.Records {
		if _, ok := recordMap[record]; !ok {
			recordMap[record] = struct{}{}
		}
	}

	for _, dbrecord := range dbUser.Access.Archive.Records {
		if _, ok := recordMap[dbrecord]; !ok {
			newRecord = append(newRecord, dbrecord)
			recordMap[dbrecord] = struct{}{}
		}
	}

	dbUser.Access.Archive.Records = newRecord

	agentMap := make(map[string]struct{})
	newAgent := make([]string, 0)

	for _, agent := range user.Access.Task.Agents {
		if _, ok := agentMap[agent]; !ok {
			agentMap[agent] = struct{}{}
		}
	}

	for _, dbagent := range dbUser.Access.Task.Agents {
		if _, ok := agentMap[dbagent]; !ok {
			newAgent = append(newAgent, dbagent)
			agentMap[dbagent] = struct{}{}
		}
	}

	dbUser.Access.Task.Agents = newAgent
}

func getUserWithNewServices(user, dbUser *dbtypes.User) {
	recordMap := make(map[string]struct{})
	newRecord := make([]string, 0)
	for _, dbrecord := range dbUser.Access.Archive.Records {
		if _, ok := recordMap[dbrecord]; !ok {
			newRecord = append(newRecord, dbrecord)
			recordMap[dbrecord] = struct{}{}
		}
	}
	for _, record := range user.Access.Archive.Records {
		if _, ok := recordMap[record]; !ok {
			newRecord = append(newRecord, record)
			recordMap[record] = struct{}{}
		}
	}
	dbUser.Access.Archive.Records = newRecord

	agentMap := make(map[string]struct{})
	newAgent := make([]string, 0)
	for _, dbagent := range dbUser.Access.Task.Agents {
		if _, ok := agentMap[dbagent]; !ok {
			newAgent = append(newAgent, dbagent)
			agentMap[dbagent] = struct{}{}
		}
	}
	for _, agent := range user.Access.Task.Agents {
		if _, ok := agentMap[agent]; !ok {
			newAgent = append(newAgent, agent)
			agentMap[agent] = struct{}{}
		}
	}
	dbUser.Access.Task.Agents = newAgent
}

func userAlreadyExists(username string) (bool, error) {
	var x int64
	err := DataBase.QueryRow(Ctx, selectUserByUsername, username).Scan(&x)
	if err == pgx.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
