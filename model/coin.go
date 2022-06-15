package model

import "time"

import (
	"github.com/google/uuid"
)

type Coin struct {
	UUID                          uuid.UUID `gorm:"column:user_id;primary_key"`
	Currency                      string    `gorm:"column:currency;type:varchar(40);not null;index"`
	Id                            int32     `gorm:"column:ID;type:int(36);index"`
	CreatedAt                     time.Time ``
	UpdatedAt                     time.Time ``
	DeletedAt                     time.Time ``
	Name                          string    `gorm:"column:name;type:text;not null;index"`
	Symbol                        string    `gorm:"column:symbol;type:varchar(40)"`
	Slug                          string    `gorm:"column:slug;type:text"`
	NumMarketPairs                float64   `gorm:"column:num_market_pairs;type:varchar(40)"`
	DateAdded                     string    `gorm:"column:date_added;type:varchar(40)"`
	MaxSupply                     float64   `gorm:"column:max_supply;size:11"`
	CirculatingSupply             float64   `gorm:"column:circulating_supply;size:11"`
	TotalSupply                   float64   `gorm:"column:total_supply;size:11"`
	CmcRank                       float64   `gorm:"column:cmc_rank;size:11"`
	SelfReportedCirculatingSupply float64   `gorm:"column:self_reported_circulating_supply;size:11"`
	SelfReportedMarketCap         float64   `gorm:"column:self_reported_market_cap;size:11"`
	LastUpdated                   time.Time `gorm:"column:last_updated;"`
	Price                         float64   `gorm:"column:price;size:11"`
	Volume24h                     float64   `gorm:"column:volume_24h;size:11"`
	VolumeChange24h               float64   `gorm:"column:volume_change_24h;size:11"`
	PercentChange1h               float64   `gorm:"column:percent_change_1h;size:11"`
	PercentChange24h              float64   `gorm:"column:percent_change_24h;size:11"`
	PercentChange7d               float64   `gorm:"column:percent_change_7d;size:11"`
	PercentChange30d              float64   `gorm:"column:percent_change_30d;size:11"`
	PercentChange60d              float64   `gorm:"column:percent_change_60d;size:11"`
	PercentChange90d              float64   `gorm:"column:percent_change_90d;size:11"`
	MarketCap                     float64   `gorm:"column:market_cap;size:11"`
	MarketCapDominance            float64   `gorm:"column:market_cap_dominance;size:11"`
	FullyDilutedMarketCap         float64   `gorm:"column:fully_diluted_market_cap;size:11"`
}

// TableName 테이블이름
func (Coin) TableName() string {
	return "coin"
}
