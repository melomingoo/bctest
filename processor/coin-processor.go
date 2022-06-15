package processor

import (
	"bc_melomingoo/message"
	"bc_melomingoo/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

func GetCoinHistoryList(db *gorm.DB, req *message.CoinHistoryRequest) (*message.CoinHistoryListResponse, error) {

	var datalist []*model.Coin
	coinTableName := model.Coin{}.TableName()

	result := db.Table(coinTableName).Where(fmt.Sprintf("last_updated between '%s' and '%s'", req.StartTime, req.EndTime))
	if len(req.Currency) > 0 {
		result = result.Where("currency = ?", req.Currency)
	}
	if len(req.Target) > 0 {
		target := strings.Split(req.Target, ",")
		result = result.Where("symbol in (?)", target)
		fmt.Println(target)
	}
	result.Scan(&datalist)
	if result.Error != nil {
		return nil, result.Error
	}

	var returnList []*message.CoinHistoryResponse

	for _, data := range datalist {
		returnList = append(returnList, getCoinHistoryResponse(data))
	}

	coinHistoryListResponse := &message.CoinHistoryListResponse{
		Items: returnList,
	}

	return coinHistoryListResponse, nil
}

func getCoinHistoryResponse(data *model.Coin) *message.CoinHistoryResponse {
	if data == nil {
		return nil
	}

	coinHistoryResponse := message.CoinHistoryResponse{
		Id:                    data.Id,
		Name:                  data.Name,
		Symbol:                data.Symbol,
		Price:                 data.Price,
		Volume24h:             data.Volume24h,
		VolumeChange24h:       data.VolumeChange24h,
		PercentChange1h:       data.PercentChange1h,
		PercentChange24h:      data.PercentChange24h,
		PercentChange7d:       data.PercentChange7d,
		PercentChange30d:      data.PercentChange90d,
		PercentChange60d:      data.PercentChange60d,
		PercentChange90d:      data.PercentChange90d,
		MarketCap:             data.MarketCap,
		MarketCapDominance:    data.MarketCapDominance,
		FullyDilutedMarketCap: data.FullyDilutedMarketCap,
		LastUpdated:           data.LastUpdated.String(),
	}

	return &coinHistoryResponse
}
