package po

import (
	dbutils "github.com/airdb/sailor/dbutil"
	"github.com/jinzhu/gorm"
)

// AgentIp
type AgentIp struct {
	gorm.Model
	IP            string `gorm:"type:varchar(32)"`
	Port          string `gorm:"type:varchar(32)"`
	ProxyType     string `gorm:"type:varchar(8)"`
	Anonymity     string `gorm:"type:varchar(8)"`
	Country       string `gorm:"type:varchar(4096)"`
	City          string `gorm:"type:varchar(4096)"`
	Speed         string `gorm:"type:varchar(4096)"`
	Origin        string `gorm:"type:varchar(4096)"`
	Operator	  string `gorm:"type:varchar(10);comment:运营商;"`
	Actived       *bool  `gorm:"actived"`
	Seen          *bool  `gorm:"seen"`
	LastCheckedAt uint   `json:"last_checked_at"`

}

// BatchFindIp
func BatchFindIp(ipArr []string) []*AgentIp {
	var data []*AgentIp
	dbutils.ReadDB(dbMinaAPIRead).Debug().Where("ip in ?", ipArr).Find(&data)
	return data
}

// BatchInsert
func BatchAgentInsert(data []AgentIp) error {
	return dbutils.ReadDB(dbMinaAPIRead).Debug().Create(&data).Error
}
// UpdateAgent
func UpdateAgent(data *AgentIp) error {
	return dbutils.ReadDB(dbMinaAPIRead).Save(data).Error
}
