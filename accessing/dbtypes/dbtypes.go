package dbtypes

type AccessServices struct {
	Archive ArchiveManager `json:"archive_manager",omitempty`
	Task    TaskManager    `json:"task_manager",omitempty`
}

type ArchiveManager struct {
	Records []string `json:"records",omitempty`
}

type TaskManager struct {
	Agents []string `json:"agents,omitempty"`
}

type User struct {
	Username string         `json:"username"`
	Access   AccessServices `json:"services",omitempty`
}

type ClientRequest struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
