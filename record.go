package nmc_typhoon_db_client

type Record struct {
	Xuhao     int          `csv:"xuhao"`
	Center    string       `csv:"center"`
	Bwtype    DataString   `csv:"bwtype"`
	FcstType  string       `db:"FCSTType" csv:"FCSTType"`
	Zone      DataString   `csv:"zone"`
	Tfbh      DataString   `csv:"tfbh"`
	Tfbhbabj  DataString   `csv:"tfbhbabj"`
	Engname   DataString   `csv:"engname"`
	Datetime  NullDateTime `csv:"datetime"`
	Fcsthour  int          `csv:"fcsthour"`
	Lat       DataFloat64  `csv:"lat"`
	Lon       DataFloat64  `csv:"lon"`
	Pressure  DataInt32    `csv:"pressure"`
	Windv     DataFloat64  `csv:"windv"`
	Gusts     DataFloat64  `csv:"gusts"`
	Strength  DataString   `csv:"strength"`
	Windclass DataInt32    `csv:"windclass"`
	Movedir   DataString   `csv:"movedir"`
	Movespeed DataFloat64  `csv:"movespeed"`
	Wind7v1   DataFloat64  `csv:"wind7v1"`
	Wind7v2   DataFloat64  `csv:"wind7v2"`
	Wind7v3   DataFloat64  `csv:"wind7v3"`
	Wind7v4   DataFloat64  `csv:"wind7v4"`
	Wind10v1  DataFloat64  `csv:"wind10v1"`
	Wind10v2  DataFloat64  `csv:"wind10v2"`
	Wind10v3  DataFloat64  `csv:"wind10v3"`
	Wind10v4  DataFloat64  `csv:"wind10v4"`
	Wind12v1  DataFloat64  `csv:"wind12v1"`
	Wind12v2  DataFloat64  `csv:"wind12v2"`
	Wind12v3  DataFloat64  `csv:"wind12v3"`
	Wind12v4  DataFloat64  `csv:"wind12v4"`
}
