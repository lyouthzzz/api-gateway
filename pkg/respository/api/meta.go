package api

import "gopkg.in/guregu/null.v3"

type MetaParam struct {
	Name            null.String `json:"name"`
	Domain          null.String `json:"domain"`
	AuthType        null.Int    `json:"auth_type"`
	QPS             null.Int    `json:"qps"`
	Status          int         `json:"status"`
	GroupId         null.Int    `json:"group_id"`
	GroupName       null.String `json:"group_name"`
	FrontendContent []byte      `json:"frontend_content"`
	BackendContent  []byte      `json:"backend_content"`
	ResultContent   []byte      `json:"result_content"`
}

type FilterParam struct {
	Id        uint32 `json:"id"`
	GroupId   uint32 `json:"group_id"`
	GroupName string `json:"group_name"`
}
