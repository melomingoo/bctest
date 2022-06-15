package handler

import (
	"bc_melomingoo/common"
	"github.com/jinzhu/gorm"
)

type BaseHandler struct {
	Config *common.Config
	TestDB *gorm.DB
}
