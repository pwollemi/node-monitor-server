package dao

import (
	"time"

	"github.com/flashguru-git/node-monitor-server/models"
	"gopkg.in/mgo.v2/bson"
)

func FindAll() ([]models.NodeMetric, error) {
	var infos []models.NodeMetric
	err := db.C(COLLECTION).Find(bson.M{}).All(&infos)
	return infos, err
}

func FindById(id string) (models.NodeMetric, error) {
	var nodeMetric models.NodeMetric
	err := db.C(COLLECTION).Find(bson.M{"nodeId": id}).Sort("-timestamp").One(&nodeMetric)
	return nodeMetric, err
}

func FindWithOption(id string, from time.Time, to time.Time, h []uint64, m []uint64, s []uint64) ([]models.NodeMetric, error) {
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

	var res []models.NodeMetric
	err := db.C(COLLECTION).Find(filter).Sort("-timestamp").All(&res)
	return res, err
}

func FindAllNodeId() ([]string, error) {
	var res []string
	err := db.C(COLLECTION).Find(bson.M{}).Distinct("nodeId", &res)
	return res, err
}

func FindLatest() (models.NodeMetric, error) {
	var nodeMetric models.NodeMetric
	err := db.C(COLLECTION).Find(bson.M{}).Sort("-blockHeight").One(&nodeMetric)
	return nodeMetric, err
}

func Insert(nodeMetric models.NodeMetric) error {
	err := db.C(COLLECTION).Insert(&nodeMetric)
	return err
}

func Delete(nodeMetric models.NodeMetric) error {
	err := db.C(COLLECTION).Remove(&nodeMetric)
	return err
}
