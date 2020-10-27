package nmc_typhoon_db_client

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type QueryConditions struct {
	StartTime    string
	EndTime      string
	ForecastHour string
}

func GetRecords(
	conditions QueryConditions,
	config DatabaseConfig,
) ([]Record, error) {
	startTime, err := time.Parse("2006010215", conditions.StartTime)
	if err != nil {
		return nil, fmt.Errorf("parse start time has error: %v", err)
	}
	// endTime := nil
	forecastHour, err := strconv.Atoi(conditions.ForecastHour)
	if err != nil {
		return nil, fmt.Errorf("parse forecsat hour has error: %v", err)
	}

	var db *sqlx.DB

	conn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local",
		config.Auth.User,
		config.Auth.Password,
		config.Host,
		config.DatabaseName)

	db, err = sqlx.Open("mysql", conn)

	if err != nil {
		log.Fatal("open db connection has error:", err)
		return nil, err
	}

	defer db.Close()

	queryColumnsString := strings.Join(QueryColumns, ",")

	tableName := config.TableName

	querySQL := fmt.Sprintf("SELECT %s FROM %s WHERE datetime='%s' AND fcsthour=%d",
		queryColumnsString,
		tableName,
		startTime.Format("2006-01-02 15:04:05"),
		forecastHour,
	)

	log.Println(querySQL)

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
