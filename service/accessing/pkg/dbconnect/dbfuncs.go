package dbconnect

import (
	"accessing/pkg/dbtypes"
	"errors"

	pgx "github.com/jackc/pgx/v5"
)

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

func GetUserByUsername(username string) (*dbtypes.User, error) {
	row := DataBase.QueryRow(Ctx, selectUserByUsername, username)
	user := dbtypes.User{}
	err := row.Scan(&user.Username, &user.Access)
	if err != nil {
		return nil, err
	}
	if user.Username == "" {
		return nil, errors.New("user not found")
	}
	return &user, nil
}
