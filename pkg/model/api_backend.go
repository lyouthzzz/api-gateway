package model

// 后端请求常量参数
type APIBackendConstParam struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Location string `json:"location"`
	Type     string `json:"type"`
}

// 后端请求系统参数
type APIBackendSystemParam struct {
	Name string `json:"name"`
	//系统参数名称
	SystemName string `json:"system_name"`
	// 参数位置
	Location string `json:"location"`
}

type APIBackend struct {
	Model
	Schema      int    `gorm:"column:schema;type:int;not null"`
	Host        string `gorm:"column:host;type:char(60);not null"`
	Path        string `grom:"column:path;type:char(60);not null"`
	Method      string `gorm:"column:method;char(10);not null"`
	Timeout     int64  `gorm:"column:timeout;type:int;not null"`
	ConstParam  []byte `gorm:"column:const_param;type:text"`
	SystemParam []byte `gorm:"column:system_param;type:text"`
}

func (APIBackend) TableName() string {
	return "api_backend"
}

func (APIBackend) UniqueIndexes() map[string][]string {
	return map[string][]string{
		"host_port_path": {"host", "port", "path"},
	}
}
