package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	nmc_typhoon_db_client "github.com/nwpc-oper/nmc-typhoon-db-client"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var (
	configFile   string
	outputFile   string
	startTime    string
	endTime      string
	forecastHour string
)

var rootCmd = &cobra.Command{
	Use:   "nmc-typhoon-db-client",
	Short: "Get typhoon report from NMC Typhoon Database",
	Long:  `Get typhoon report from NMC Typhoon Database`,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := ioutil.ReadFile(configFile)
		if err != nil {
			panic(err)
		}

		config := nmc_typhoon_db_client.Config{}

		err = yaml.Unmarshal(data, &config)
		if err != nil {
			panic(err)
		}

		log.Println("get records...")

		conditions := nmc_typhoon_db_client.QueryConditions{
			StartTime:    startTime,
			EndTime:      endTime,
			ForecastHour: forecastHour,
		}

		records, err := nmc_typhoon_db_client.GetRecords(conditions, config.Database)
		if err != nil {
			log.Fatal("get records has error:", err)
		}
		log.Println("get records...done")

		log.Println("save records...")
		err = nmc_typhoon_db_client.WriteToCSV(records, outputFile)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("save records...done")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "./config.yaml", "config file path")
	rootCmd.PersistentFlags().StringVar(&outputFile, "output-file", "", "output file path")
	rootCmd.PersistentFlags().StringVar(&startTime, "start-time", "", "start time, YYYYMMDDHH")
	rootCmd.PersistentFlags().StringVar(&endTime, "end-time", "", "end time, YYYYMMDDHH")
	rootCmd.PersistentFlags().StringVar(&forecastHour, "forecast-hour", "0", "forecast hour, 0 or 0-120")
	rootCmd.MarkPersistentFlagRequired("start-time")
	rootCmd.MarkPersistentFlagRequired("output-file")
}
