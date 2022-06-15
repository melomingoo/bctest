package common

import (
	"bc_melomingoo/message"
	"bc_melomingoo/model"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

var ClientStoreManager storeManager

type storeManager interface {
	Start()
}

type StoreManager struct {
	config *Config
	db     *gorm.DB
}

func NewStoreManager(db *gorm.DB, config *Config) storeManager {
	storeManager := &StoreManager{
		db:     db,
		config: config,
	}
	storeManager.Start()
	ClientStoreManager = storeManager
	return ClientStoreManager
}

func (sm *StoreManager) Start() {
	//updateIntervalMin := 10 * time.Second // 갱신 간격 (시간)
	updateIntervalMin := 1 * time.Hour // 갱신 간격 (시간)
	ticker := time.NewTicker(updateIntervalMin)
	go func() {
		for {
			select {
			case <-ticker.C:
				sm.UpdateCoinData()
			}
		}
	}()
}

func (sm *StoreManager) UpdateCoinData() {
	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/cryptocurrency/listings/latest", sm.config.CoinMarket.Host), nil)
	if err != nil {

	}
	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "5000")
	q.Add("convert", "USD")

	request.Header.Set("Accepts", "application/json")
	request.Header.Add("X-CMC_PRO_API_KEY", sm.config.CoinMarket.ClientID)
	request.URL.RawQuery = q.Encode()

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}
	respBody, _ := ioutil.ReadAll(response.Body)
	CoinData := message.CoinList{}
	err = json.Unmarshal([]byte(string(respBody)), &CoinData)

	if err != nil {
		fmt.Println(err)
	}

	codeTableName := model.Coin{}.TableName()
	db := sm.db.Table(codeTableName).Debug()
	now := time.Now()
	for _, item := range CoinData.Data {
		tee, _ := time.Parse(time.RFC3339, item.LastUpdated)
		tee.Format("2006-01-02 15:04:05")
		coinData := model.Coin{
			UUID:                          uuid.New(),
			Currency:                      "USD",
			Id:                            item.Id,
			CreatedAt:                     now,
			UpdatedAt:                     now,
			Name:                          item.Name,
			Symbol:                        item.Symbol,
			Slug:                          item.Slug,
			NumMarketPairs:                item.NumMarketPairs,
			DateAdded:                     item.DateAdded,
			MaxSupply:                     item.MaxSupply,
			CirculatingSupply:             item.CirculatingSupply,
			TotalSupply:                   item.TotalSupply,
			CmcRank:                       item.CmcRank,
			SelfReportedCirculatingSupply: item.SelfReportedCirculatingSupply,
			SelfReportedMarketCap:         item.SelfReportedMarketCap,
			LastUpdated:                   tee,
			Price:                         item.Quote.USD.Price,
			Volume24h:                     item.Quote.USD.Volume24h,
			VolumeChange24h:               item.Quote.USD.VolumeChange24h,
			PercentChange1h:               item.Quote.USD.PercentChange1h,
			PercentChange24h:              item.Quote.USD.PercentChange24h,
			PercentChange7d:               item.Quote.USD.PercentChange7d,
			PercentChange30d:              item.Quote.USD.PercentChange30d,
			PercentChange60d:              item.Quote.USD.PercentChange60d,
			PercentChange90d:              item.Quote.USD.PercentChange90d,
			MarketCap:                     item.Quote.USD.MarketCap,
			MarketCapDominance:            item.Quote.USD.MarketCapDominance,
			FullyDilutedMarketCap:         item.Quote.USD.FullyDilutedMarketCap,
		}
		result := db.Create(&coinData)
		if result.Error != nil {
			return
		}
	}

}
