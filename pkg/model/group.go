package model

type Group struct {
	Model
	Name      string `gorm:"column:name;type:char(60);not null"`
	Desc      string `gorm:"column:desc;type:char(255)"`
	ProjectID string `gorm:"column:project_id;type:int unsigned;not null"`
}

func (Group) TableName() string {
	return "group"
}

func (Group) UniqueIndexes() map[string][]string {
	return map[string][]string{
		"Name": {"Name"},
	}
}

func (Group) ForeignKeys() map[string]string {
	return map[string]string{
		"project_id": "project(id)",
	}
}
