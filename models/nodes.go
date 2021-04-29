package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type NodeInfo struct {
	ID          bson.ObjectId          `bson:"_id" json:"_id"`
	NodeID      string                 `bson:"nodeId" json:"nodeId"`
	BlockHeight uint64                 `bson:"blockHeight" json:"blockHeight"`
	TimeStamp   int64                  `bson:"timestamp" json:"timestamp"`
	Cpu         map[string]interface{} `bson:"cpu" json:"cpu"`
	Memory      map[string]interface{} `bson:"memory" json:"memory"`
	Hour        uint64                 `bson:"hour" json:"-"`
	Min         uint64                 `bson:"minute" json:"-"`
	Sec         uint64                 `bson:"second" json:"-"`
}

type NodeInfoRequest struct {
	NodeID      string                 `bson:"nodeId" json:"nodeId"`
	BlockHeight uint64                 `bson:"blockHeight" json:"blockHeight"`
	TimeStamp   time.Time              `bson:"timestamp" json:"timestamp"`
	Cpu         map[string]interface{} `bson:"cpu" json:"cpu"`
	Memory      map[string]interface{} `bson:"memory" json:"memory"`
}
