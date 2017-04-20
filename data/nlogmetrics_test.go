package data_test

import (
	"log"
	"os"
	"runtime"
	"testing"

	_ "github.com/denisenkom/go-mssqldb"

	"github.com/danesparza/nlog-metrics/data"
)

//	Gets the database connection information from the environment
func getMSSQLDBConnection() data.MSSqlDB {

	//	Set this information from environment variables?
	return data.MSSqlDB{
		Server:   os.Getenv("nlm_test_mssql_server"), /* Ex: test-server:3306 If this is blank, it assumes a local database on port 3306 */
		Database: os.Getenv("nlm_test_mssql_database"),
		User:     os.Getenv("nlm_test_mssql_user"),
		Password: os.Getenv("nlm_test_mssql_password")}
}

//	MSSQL get should return successfully even if the item doesn't exist
func TestMssql_GetMetrics_ItemDoesntExist_Successful(t *testing.T) {

	if runtime.GOOS != "windows" {
		t.Skip("Skipping MSSQL tests: Not on Windows")
	}

	//	Arrange
	db := getMSSQLDBConnection()
	t.Logf("Using - server: %v / database: %v / user: %v / password: %v", db.Server, db.Database, db.User, db.Password)

	//	Act
	response, err := db.GetMetrics()

	//	Assert
	if err != nil {
		t.Errorf("GetMetrics failed: %v", err)
	}

	for _, nlogItem := range response {
		log.Printf("[DEBUG] NLogItem found: %v %v %v", nlogItem.Application, nlogItem.LogLevel, nlogItem.Count)
	}
}
