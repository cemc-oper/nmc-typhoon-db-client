package nmc_typhoon_db_client

type Record struct {
	Xuhao       int          `csv:"xuhao"`
	Tfxh        DataString   `csv:"tfxh"`
	Tfbh        DataString   `csv:"tfbh"`
	Tfbhbabj    DataString   `csv:"tfbhbabj"`
	Engname     DataString   `csv:"engname"`
	Chnname     DataString   `csv:"chnname"`
	Center      string       `csv:"center"`
	Bwtype      DataString   `csv:"bwtype"`
	FcstType    string       `db:"FCSTType" csv:"FCSTType"`
	Datetime    NullDateTime `csv:"datetime"`
	Handletime  NullDateTime `csv:"handletime"`
	Validtime   NullDateTime `csv:"validtime"`
	Fcsthour    int          `csv:"fcsthour"`
	Strength    DataString   `csv:"strength"`
	Zone        DataString   `csv:"zone"`
	Lon         DataFloat64  `csv:"lon"`
	Lat         DataFloat64  `csv:"lat"`
	Windclass   DataInt32    `csv:"windclass"`
	Windv       DataFloat64  `csv:"windv"`
	Gusts       DataFloat64  `csv:"gusts"`
	Pressure    DataInt32    `csv:"pressure"`
	Movedir     DataString   `csv:"movedir"`
	Movespeed   DataFloat64  `csv:"movespeed"`
	Wind6v1     DataFloat64  `csv:"wind6v1"`
	Wind6v2     DataFloat64  `csv:"wind6v2"`
	Wind6v3     DataFloat64  `csv:"wind6v3"`
	Wind6v4     DataFloat64  `csv:"wind6v4"`
	Wind7v1     DataFloat64  `csv:"wind7v1"`
	Wind7v2     DataFloat64  `csv:"wind7v2"`
	Wind7v3     DataFloat64  `csv:"wind7v3"`
	Wind7v4     DataFloat64  `csv:"wind7v4"`
	Wind10v1    DataFloat64  `csv:"wind10v1"`
	Wind10v2    DataFloat64  `csv:"wind10v2"`
	Wind10v3    DataFloat64  `csv:"wind10v3"`
	Wind10v4    DataFloat64  `csv:"wind10v4"`
	Wind12v1    DataFloat64  `csv:"wind12v1"`
	Wind12v2    DataFloat64  `csv:"wind12v2"`
	Wind12v3    DataFloat64  `csv:"wind12v3"`
	Wind12v4    DataFloat64  `csv:"wind12v4"`
	Wind7class  DataInt32    `csv:"wind7class"`
	Wind10class DataInt32    `csv:"wind10class"`
	Wind12class DataInt32    `csv:"wind12class"`
	Wind6class  DataInt32    `csv:"wind6class"`
	Memo        DataString   `csv:"memo"`
}
