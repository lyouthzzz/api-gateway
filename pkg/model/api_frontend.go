package model

type APIFrontendParam struct {
	// 参数名
	Name string `grom:"column:name;type:char(60)" json:"name"`
	// 位置
	Location string `gorm:"column:location;type:int;not null" json:"location"`
	// 类型
	Type string `gorm:"column:type;type:char(60);not null" json:"type"`
	// 必须
	Must int `gorm:"column:must;type:int;not null;default 0" json:"must"`
	// 默认值
	Default string `gorm:"column:default;type:char(60)" json:"default"`
}

type APIFrontend struct {
	Model
	Schema int    `gorm:"column:schema;not null"`
	Path   string `grom:"column:path;type:char(60);not null"`
	Method string `grom:"column:method;type:char(60);not null"`
	Param  []byte `gorm:"column:param;type:text"`
	Body   []byte `gorm:"column:body;type:text"`
}

func (APIFrontend) TableName() string {
	return "api_frontend"
}

func (APIFrontend) UniqueIndexes() map[string][]string {
	return map[string][]string{
		"path_method": {"path", "method"},
	}
}
