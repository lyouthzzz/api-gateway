package model

type Project struct {
	Model
	Name string `gorm:"column:name;type:char(60);not null"`
	Desc string `gorm:"column:desc;type:char(255)"`
}

func (Project) TableName() string {
	return "project"
}

func (Project) UniqueIndexes() map[string][]string {
	return map[string][]string{
		"name": {"name"},
	}
}
