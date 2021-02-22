package dao

import (
	"github.com/shopspring/decimal"
	"time"
)

type EnvelopeOrder struct {
	Id               int    `gorm:"size:11;primary_key;auto_increment;not null" json:"id"`
	EnvelopeNo       string  `gorm:"size:255;not null;comment:'红包编号'"`
	EnvelopeType     int8   `gorm:"size:1;not null;comment:'红包类型'"`
	UserId           int    `gorm:"size:11;not null;comment:'用户ID'"`
	UserName         string  `gorm:"size:255;not null;comment:'用户名称'"`
	Blessing         string  `gorm:"size:255;not null;comment:'祝福语'"`
	Amount           decimal.Decimal  `gorm:"not null;comment:'红包总金额'"`
	AmountOne        decimal.Decimal  `gorm:"not null;comment:'普通单个红包金额'"`
	Quantity         int    `gorm:"size:4;not null;comment:'红包数量'"`
	ExpiredAt        time.Time  `gorm:"not null;comment:'过期时间'"`
	RemainAmount     decimal.Decimal  `gorm:"not null;comment:'红包剩余金额'"`
	RemainQuantity   int    `gorm:"size:4;not null;comment:'红包剩余数量'"`
	OrderStatus      int8   `gorm:"size:1;not null;comment:'红包订单状态'"`
	OrderType        int8   `gorm:"size:1;not null;comment:'红包订单类型'"`
	PayStatus        int8   `gorm:"size:1;not null;comment:'红包支付状态'"`
	CreatedAt        time.Time `gorm:"not null;comment:'创建时间';autoCreateTime"`
	UpdatedAt        time.Time `gorm:"not null;comment:'修改时间';autoUpdateTime"`
}

type EnvelopeItem struct {
	Id               int    `gorm:"size:11;primary_key;auto_increment;not null" json:"id"`
	ItemNo           string  `gorm:"size:255;not null;comment:'详情编号'"`
	EnvelopeNo       string  `gorm:"size:255;not null;comment:'红包编号'"`
	UserId           int    `gorm:"size:11;not null;comment:'收红包用户ID'"`
	UserName         string  `gorm:"size:255;not null;comment:'收红包用户名称'"`
	Amount           decimal.Decimal  `gorm:"not null;comment:'收红包金额'"`
	Quantity         int    `gorm:"size:4;not null;comment:'收红包数量'"`
	RemainAmount     decimal.Decimal  `gorm:"not null;comment:'红包剩余金额'"`
	AccountNo        string  `gorm:"size:255;not null;comment:'账户编号'"`
	PayStatus        int8   `gorm:"size:1;not null;comment:'红包支付状态'"`
	CreatedAt        time.Time `gorm:"not null;comment:'创建时间';autoCreateTime"`
	UpdatedAt        time.Time `gorm:"not null;comment:'修改时间';autoUpdateTime"`
}