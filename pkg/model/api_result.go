package model

type APIError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Desc string `json:"desc"`
}

type APIResult struct {
	Model
	// 返回ContentType
	ContentType string `gorm:"column:content_type;type:char(60);not null"`
	// 返回结果示例
	ExampleSuccess []byte `gorm:"column:example_success;type:text"`
	// 失败返回结果示例
	ExampleFailed []byte `gorm:"column:example_failed;type:text"`
	// 后端错误码说明
	Errors []byte `gorm:"column:errors;type:text"`
}

func (APIResult) TableName() string {
	return "api_result"
}
