package nmc_typhoon_db_client

import (
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func GetRecords(
	config DatabaseConfig,
) ([]Record, error) {
	var db *sqlx.DB

	conn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local",
		config.Auth.User,
		config.Auth.Password,
		config.Host,
		config.DatabaseName)

	db, err := sqlx.Open("mysql", conn)

	if err != nil {
		log.Fatal("open db connection has error:", err)
		return nil, err
	}

	defer db.Close()

	queryColumnsString := strings.Join(QueryColumns, ",")

	tableName := config.TableName

	querySQL := "SELECT " + queryColumnsString + " " +
		"FROM " + tableName + " " +
		"WHERE datetime='2020-10-25' AND fcsthour=0"

	// log.Println(querySQL)

	rows, err := db.Queryx(querySQL)

	if err != nil {
		log.Fatal("query db has error:", err)
		return nil, err
	}

	defer rows.Close()

	records := []Record{}
	for rows.Next() {
		var record Record
		err = rows.StructScan(&record)

		if err != nil {
			log.Fatal("scan row has error:", err)
			return nil, err
		}
		// log.Println(record.xuhao, record.center, record.datetime)
		records = append(records, record)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal("query rows has error:", err)
		return nil, err
	}

	return records, nil
}
