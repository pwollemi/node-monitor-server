package dao

import (
	"log"

	. "github.com/flashguru-git/node-monitor-server/config"
	"gopkg.in/mgo.v2"
)

var config = Config{}
var db *mgo.Database

const (
	COLLECTION = "nodes"
)

// Parse the configuration file 'config.toml', and establish a connection to DB
func Connect() {
	config.Read()
	session, err := mgo.Dial(config.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(config.Database)
}

// Ensure Indexes on keys
func PopulateIndex() {
	for _, key := range []string{"nodeId"} {
		index := mgo.Index{
			Key: []string{key},
		}
		db.C(COLLECTION).EnsureIndex(index)
	}
}
