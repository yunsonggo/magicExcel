package param

// 导出数据参数
type OutputParam struct {
	QueryTableArray []QueryTableParam `json:"query_data"`
}

type ClassParam struct {
	Text  string `json:"text"`
	Value string `json:"value"`
}

type QueryTableParam struct {
	Class            string `json:"class"`
	Car              string `json:"car"`
	DateString       string `json:"lastDate"`
	BackupNum        string `json:"backupNum"`
	NowNum           string `json:"nowNum"`
	OilType          string `json:"oilType"`
	OilNum           string `json:"oilNum"`
	Pay              string `json:"pay"`
	Status           string `json:"status"`
	OilPer           string `json:"per"`
	RepairPay        string `json:"repair_pay"`
	RepairStatus     string `json:"repair_status"`
	RepairDateString string `json:"repair_date_string"`
}
