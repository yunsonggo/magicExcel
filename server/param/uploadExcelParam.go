package param

import "mime/multipart"

type UploadExcelParam struct {
	DataTag     string                `json:"dataTag" form:"dataTag"`
	DataOption  string                `json:"dataOption" form:"dataOption"`
	MonthString string                `json:"monthString" form:"monthString"`
	MonthPicker string                `json:"monthPicker" form:"monthPicker"`
	File        *multipart.FileHeader `form:"file"`
}

type ReTableNameParam struct {
	TableName   string `json:"table_name"`
	TableOption string `json:"table_option"`
	FilePath    string `json:"file_path"`
}
