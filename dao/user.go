package dao

import "time"

// 用户表
type User struct {
	Id        int       `gorm:"size:11;primary_key;auto_increment;not null" json:"id"`
	AliasName string    `gorm:"size:10;not null;comment:'别名'"`
	UserName  string    `gorm:"size:255;not null;comment:'用户名'"`
	Sex       string    `gorm:"size:2;not null;comment:'用户性别'"`
	HeadUrl   string    `gorm:"size:255;not null;comment:'头像'"`
	IdNo      string    `gorm:"size:25;not null;comment:'身份ID'"`
	Region    string    `gorm:"size:255;not null;comment:'地区'"`
	Signature string    `gorm:"size:255;not null;comment:'签名'"`
	CreatedAt time.Time `gorm:"not null;comment:'创建时间';autoCreateTime"`
	UpdatedAt time.Time `gorm:"not null;comment:'修改时间';autoUpdateTime"`
}

