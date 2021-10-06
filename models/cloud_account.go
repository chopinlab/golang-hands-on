package models

import "time"

type CloudAccount struct {
	CloudAccountId   string    `gorm:"primary_key;column:cloud_account_id;type:varchar;size:50;"      json:"cloud_account_id"    example:"39285a58-9362-xxxx-xxxx-afda095bb612"`
	CloudAccountName string    `gorm:"column:cloud_account_name;type:varchar;size:50;"                json:"cloud_account_name"  example:"testKey"`
	Provider         string    `gorm:"column:provider;type:varchar;size:10;"                          json:"provider"            example:"aws"`
	DisplayName      string    `gorm:"column:display_name;type:varchar;size:50;"                      json:"display_name;"       example:"ec2"`
	UseFlag          int8      `gorm:"column:use_flag;type:tinyint;default:0;"                        json:"use_flag"      example:"1"`
	Created          time.Time `gorm:"column:created;type:timestamp;default:CURRENT_TIMESTAMP;"       json:"created"       example:"2021-01-01T00:00:00Z"`
	Modified         time.Time `gorm:"column:modified;type:timestamp;"                                json:"modified"      example:"2021-01-01T00:00:00Z"`
}

func (m *CloudAccount) TableName() string {
	return "cloud_account"
}