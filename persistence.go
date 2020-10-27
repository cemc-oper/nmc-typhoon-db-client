package nmc_typhoon_db_client

import (
	"fmt"
	"os"

	"github.com/jszwec/csvutil"
)

func WriteToCSV(records []Record, filePath string) error {
	b, err := csvutil.Marshal(records)
	if err != nil {
		fmt.Println("error:", err)
		return err
	}

	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println("crate file error:", err)
		return err
	}
	f.Write(b)
	f.Close()
	return nil
}
