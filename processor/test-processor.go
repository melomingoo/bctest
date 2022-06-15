package processor

import (
	"bc_melomingoo/message"
	"bc_melomingoo/model"
	"fmt"
	"github.com/jinzhu/gorm"
)

func GetTestList(db *gorm.DB) (*message.TestListResponse, error) {

	var datalist []*model.Test
	testTableName := model.Test{}.TableName()
	result := db.Table(testTableName).Find(&datalist)

	if result.Error != nil {
		return nil, result.Error
	}

	var returnList []*message.TestResponse
	fmt.Println(datalist)
	for _, data := range datalist {
		returnList = append(returnList, getTestResponse(data))
	}

	testListResponse := &message.TestListResponse{
		Items: returnList,
	}

	return testListResponse, nil
}

func getTestResponse(data *model.Test) *message.TestResponse {
	if data == nil {
		return nil
	}

	testResponse := message.TestResponse{
		ID:     data.ID,
		TestCd: data.Check,
	}

	return &testResponse
}
