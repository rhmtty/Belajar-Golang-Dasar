package database

var connection string

func init() {
	connection = "Postgres Connection"
}

func GetDatabase() string {
	return connection
}
