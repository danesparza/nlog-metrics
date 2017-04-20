package data

import (
	"database/sql"
	"fmt"

	// SQL server driver
	_ "github.com/denisenkom/go-mssqldb"
)

// NLogMetric represents a single metric
type NLogMetric struct {
	Application string `sql:"log_application"`
	LogLevel    string `sql:"log_level"`
	Count       int    `sql:"items"`
}

// MSSqlDB represetns the MSSQL database information
type MSSqlDB struct {
	Server   string
	Database string
	User     string
	Password string
}

// GetMetrics retrieves metrics data from the database
func (store MSSqlDB) GetMetrics() ([]NLogMetric, error) {
	//	Our return item:
	retval := []NLogMetric{}

	//	Open the database:
	db, err := sql.Open("mssql", fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s", store.Server, store.Database, store.User, store.Password))
	defer db.Close()
	if err != nil {
		return retval, err
	}

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("select log_application, log_level, items = count(*) from system_logging where entered_date > DateADD(ss, -30, getdate()) group by log_application, log_level order by log_application")

	defer rows.Close()
	if err != nil {
		return retval, err
	}

	for rows.Next() {
		var logApplication string
		var logLevel string
		var items int

		//	Scan the row into our variables
		err = rows.Scan(&logApplication, &logLevel, &items)

		if err != nil {
			return retval, err
		}

		//	Append to return values
		retval = append(retval, NLogMetric{
			Application: logApplication,
			LogLevel:    logLevel,
			Count:       items})
	}

	//	Return what we found:
	return retval, nil
}
