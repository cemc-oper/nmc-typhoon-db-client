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

const queryTimeFormat = "2006-01-02 15:04:05"
const optionTimeFormat = "2006010215"

func GetRecords(
	conditions QueryConditions,
	config DatabaseConfig,
) ([]Record, error) {
	datetimeQuery, err := generateDateTimeQuery(conditions)
	if err != nil {
		return nil, fmt.Errorf("get datetime query has error: %v", err)
	}

	forecastHourQuery, err := generateForecastHourQuery(conditions)
	if err != nil {
		return nil, fmt.Errorf("parse forecast hour query has error: %v", err)
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

	querySQL := fmt.Sprintf("SELECT %s FROM %s WHERE %s AND %s",
		queryColumnsString,
		tableName,
		datetimeQuery,
		forecastHourQuery,
	)

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

func generateDateTimeQuery(
	conditions QueryConditions,
) (string, error) {
	startTime, err := time.Parse(optionTimeFormat, conditions.StartTime)
	if err != nil {
		return "", fmt.Errorf("parse start time has error: %v", err)
	}

	if len(conditions.EndTime) == 0 {
		query := fmt.Sprintf("datetime='%s'", startTime.Format(queryTimeFormat))
		return query, nil
	}

	endTime, err := time.Parse(optionTimeFormat, conditions.EndTime)
	if err != nil {
		return "", fmt.Errorf("parse end time has error: %v", err)
	}

	query := fmt.Sprintf(
		"datetime BETWEEN '%s' and '%s'",
		startTime.Format(queryTimeFormat),
		endTime.Format(queryTimeFormat),
	)
	return query, nil
}

func generateForecastHourQuery(
	conditions QueryConditions,
) (string, error) {
	i := strings.Index(conditions.ForecastHour, "-")
	if i == -1 {
		forecastHour, err := strconv.Atoi(conditions.ForecastHour)
		if err != nil {
			return "", fmt.Errorf("parse forecsat hour has error: %v", err)
		}

		return fmt.Sprintf("fcsthour=%d", forecastHour), nil
	}

	startForecastHour, err := strconv.Atoi(conditions.ForecastHour[:i])
	if err != nil {
		return "", fmt.Errorf("parse forecsat hour has error: %v", err)
	}
	endForecastHour, err := strconv.Atoi(conditions.ForecastHour[i+1:])
	if err != nil {
		return "", fmt.Errorf("parse forecsat hour has error: %v", err)
	}
	return fmt.Sprintf("fcsthour BETWEEN %d AND %d", startForecastHour, endForecastHour), nil
}
