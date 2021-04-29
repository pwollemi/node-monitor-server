package dao

import (
	"time"

	"github.com/flashguru-git/node-monitor-server/models"
	"gopkg.in/mgo.v2/bson"
)

func FindAll() ([]models.NodeInfo, error) {
	var infos []models.NodeInfo
	err := db.C(COLLECTION).Find(bson.M{}).All(&infos)
	return infos, err
}

func FindById(id string) (models.NodeInfo, error) {
	var nodeInfo models.NodeInfo
	err := db.C(COLLECTION).Find(bson.M{"nodeId": id}).Sort("-timestamp").One(&nodeInfo)
	return nodeInfo, err
}

func FindWithOption(id string, from time.Time, to time.Time, h []uint64, m []uint64, s []uint64) ([]models.NodeInfo, error) {
	filter := bson.M{"nodeId": id, "timestamp": bson.M{"$gte": from.Unix(), "$lte": to.Unix()}}
	if len(h) > 0 {
		filter["hour"] = bson.M{"$in": h}
	}
	if len(m) > 0 {
		filter["minute"] = bson.M{"$in": m}
	}
	if len(s) > 0 {
		filter["second"] = bson.M{"$in": s}
	}

	var res []models.NodeInfo
	err := db.C(COLLECTION).Find(filter).Sort("-timestamp").All(&res)
	return res, err
}

func FindAllNodeId() ([]string, error) {
	var res []string
	err := db.C(COLLECTION).Find(bson.M{}).Distinct("nodeId", &res)
	return res, err
}

func FindLatest() (models.NodeInfo, error) {
	var nodeInfo models.NodeInfo
	err := db.C(COLLECTION).Find(bson.M{}).Sort("-blockHeight").One(&nodeInfo)
	return nodeInfo, err
}

func Insert(nodeInfo models.NodeInfo) error {
	err := db.C(COLLECTION).Insert(&nodeInfo)
	return err
}

func Delete(nodeInfo models.NodeInfo) error {
	err := db.C(COLLECTION).Remove(&nodeInfo)
	return err
}
