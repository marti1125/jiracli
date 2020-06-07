package jira

func GetServerInfo() string {
	return Request("GET", "/rest/api/3/serverInfo", nil)
}
