package dao

import (
	"github.com/shopspring/decimal"
	"time"
)

// 账户表
type Account struct {
	Id             int    `gorm:"size:11;primary_key;auto_increment;not null" json:"id"`
	Name           string `gorm:"size:255;not null;comment:'账户名称'"`
	AccountNo      string `gorm:"size:255;not null;comment:'账户编号'"`
	UserId         int    `gorm:"size:11;not null;comment:'用户ID'"`
	UserName       string `gorm:"size:255;not null;comment:'用户名称'"`
	Type           int8   `gorm:"size:1;not null;comment:'账户类型'"`
	Balance        decimal.Decimal  `gorm:"not null;comment:'余额'"`
	Status         int8   `gorm:"size:1;not null;comment:'账户状态'"`
	CreatedAt      time.Time `gorm:"not null;comment:'创建时间';autoCreateTime"`
	UpdatedAt      time.Time `gorm:"not null;comment:'修改时间';autoUpdateTime"`
}

// 账户流水表
type AccountLog struct {
	Id               int    `gorm:"size:11;primary_key;auto_increment;not null" json:"id"`
	TradeNo          string `gorm:"size:255;not null;comment:'交易单号'"`
	LogNo            string `gorm:"size:255;not null;comment:'流水编号'"`
	AccountNo        string `gorm:"size:255;not null;comment:'账户编号'"`
	TargetAccountNo  string  `gorm:"size:255;not null;comment:'目标账户编号'"`
	UserId           int  `gorm:"size:11;not null;comment:'用户ID'"`
	TargetUserId     int  `gorm:"size:11;not null;comment:'目标用户ID'"`
	Amount           decimal.Decimal  `gorm:"not null;comment:'交易金额'"`
	Balance          decimal.Decimal  `gorm:"not null;comment:'余额'"`
	Type             int8 `gorm:"size:1;not null;comment:'交易类型'"`
	Flag             int8 `gorm:"size:1;not null;comment:'交易变化标示'"`
	Status           int8 `gorm:"size:1;not null;comment:'交易状态'"`
	CreatedAt        time.Time `gorm:"not null;comment:'创建时间';autoCreateTime"`
	UpdatedAt        time.Time `gorm:"not null;comment:'修改时间';autoUpdateTime"`
}
