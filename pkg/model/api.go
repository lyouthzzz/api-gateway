package model

type APIStatus uint32

const (
	APIStatusOff APIStatus = 0 // 未发布
	APIStatusOn  APIStatus = 1 // 发布
)

type APIAuthStrategy uint32

type API struct {
	Model
	Name string `gorm:"column:name;type:char(60);not null"`

	Domain   string    `gorm:"column:domain;type:char(60);not null"`
	AuthType int       `grom:"column:auth_type:int;not null;default 0"`
	QPS      int64     `grom:"column:qps;type:int"`
	Status   APIStatus `grom:"column:status;type:int"`

	GroupID   uint32 `gorm:"column:group_id;type:int unsigned;not null"`
	GroupName string `gorm:"column:group_name;char(60);not null"`

	// 前端定义配置ID
	FrontendID uint32 `gorm:"column:frontend_id;type:int unsigned;not null"`
	// 后端定义配置ID
	BackendID uint32 `gorm:"column:backend_id;type:int unsigned;not null"`
	// 后端返回结果
	ResultID uint32 `gorm:"column:result_id;type:int unsigned;not null"`
}

func (API) TableName() string {
	return "api"
}

func (API) Indexes() map[string][]string {
	return map[string][]string{
		"name":       {"name"},
		"group_name": {"group_name"},
	}
}

func (API) UniqueIndexes() map[string][]string {
	return map[string][]string{
		"domain": {"domain"},
	}
}

func (API) ForeignKeys() map[string]string {
	return map[string]string{
		"group_id":    "`group`(id)",
		"frontend_id": "api_frontend(id)",
		"backend_id":  "api_backend(id)",
		"result_id":   "api_result(id)",
	}
}
