package dao

import (
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
