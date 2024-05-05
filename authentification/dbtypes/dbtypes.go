package dbtypes

type AccessServices struct {
	Archive ArchiveManager `json:"archive_manager",omitempty`
	Task    TaskManager    `json:"task_manager",omitempty`
}

type ArchiveManager struct {
	Records []string `json:"records",omitempty`
	Token   string   `json:"token",omitempty`
}

type TaskManager struct {
	Agents []string `json:"agents,omitempty"`
	Token  string   `json:"token",omitempty`
}

type User struct {
	Username string         `json:"username"`
	Password string         `json:"password"`
	Access   AccessServices `json:"services",omitempty`
}

func NewUser() *User {
	return &User{}
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		Agents: nil,
		Token:  "",
	}
}
