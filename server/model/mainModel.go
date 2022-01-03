package model

type MainModel struct {
	ID               int64   `json:"id"`
	Class            string  `json:"class"`
	CarName          string  `json:"car_name"`
	OilDateString    string  `json:"oil_date_string"`
	BackupNum        string  `json:"backup_num"`
	NowNum           string  `json:"now_num"`
	OilType          string  `json:"oil_type"`
	OilNum           float64 `json:"oil_num"`
	OilPay           float64 `json:"oil_pay"`
	OilStatus        string  `json:"oil_status"`
	RepairPay        float64 `json:"repair_pay"`
	RepairStatus     string  `json:"repair_status"`
	RepairDateString string  `json:"repair_date_string"`
}

type MainDataModel struct {
	DateString       string  `json:"date_string"`
	BackupNum        string  `json:"backup_num"`
	NowNum           string  `json:"now_num"`
	OilType          string  `json:"oil_type"`
	OilNum           float64 `json:"oil_num"`
	Pay              float64 `json:"pay"`
	Status           string  `json:"status"`
	RepairDateString string  `json:"repair_date_string"`
	RepairPay        float64 `json:"repair_pay"`
	RepairStatus     string  `json:"repair_status"`
}
