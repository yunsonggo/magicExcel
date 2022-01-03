package common

type ResErrCode int64

const (
	SuccessCode ResErrCode = 1000 + iota
	ParamSuccess
	SignUpSuccess
	SignInSuccess
	FailedCode ResErrCode = 2000 + iota
	TokenFailed
	ParamFailed
	LostFailed
	AuthFailed
	UploadFailed
	SaveFileFailed
	InsertDBFailed
	HasTableName
	CreateTableFailed
	RenameTableFailed
	OpenExcelFileFailed
	ReadExcelFileFailed
	ParseDataFailed
	FindListFailed
	CaptchaFailed
	QueryDBFailed
)

var CodeMsg = map[ResErrCode]string{
	SuccessCode:         "ok",
	TokenFailed:         "token错误",
	UploadFailed:        "上传文件错误",
	SaveFileFailed:      "保存文件失败",
	InsertDBFailed:      "插入数据库失败",
	HasTableName:        "数据表名已存在",
	CreateTableFailed:   "创建数据表失败",
	RenameTableFailed:   "数据表重命名失败",
	OpenExcelFileFailed: "打开Excel文件失败",
	ReadExcelFileFailed: "读取Excel数据失败",
	ParseDataFailed:     "解析Excel数据失败",
	FindListFailed:      "获取数据列表失败",
	CaptchaFailed:       "验证码错误",
	QueryDBFailed:       "查询数据库失败",
}
