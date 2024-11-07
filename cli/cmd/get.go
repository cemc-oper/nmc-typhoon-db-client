package cmd

import (
	"io/ioutil"
	"log"

	nmc_typhoon_db_client "github.com/cemc-oper/nmc-typhoon-db-client"
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

const getLongDecription = `Get typhoon reports from NMC Typhoon Database and save to CSV files.
`

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get typhoon reports from NMC Typhoon Database",
	Long:  getLongDecription,
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

func init() {
	getCmd.PersistentFlags().StringVar(&configFile, "config", "./config.yaml", "config file path")
	getCmd.PersistentFlags().StringVar(&outputFile, "output-file", "", "output file path")
	getCmd.PersistentFlags().StringVar(&startTime, "start-time", "", "start time, YYYYMMDDHH")
	getCmd.PersistentFlags().StringVar(&endTime, "end-time", "", "end time, YYYYMMDDHH")
	getCmd.PersistentFlags().StringVar(&forecastHour, "forecast-hour", "0", "forecast hour(s), 0 or 0-120")
	getCmd.MarkPersistentFlagRequired("start-time")
	getCmd.MarkPersistentFlagRequired("output-file")

	rootCmd.AddCommand(getCmd)
}
