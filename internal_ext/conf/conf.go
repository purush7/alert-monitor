package conf

func GetDBConn() string {
	return "postgres://alert:monitor@alertdb:5432/alert?sslmode=disable"
}
