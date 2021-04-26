package handlers

import (

	//"log"

	"encoding/json"
	"net/http"

	"github.com/flashguru-git/node-monitor-server/dao"
	"github.com/flashguru-git/node-monitor-server/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// GET latest block height
func GetLatest(w http.ResponseWriter, r *http.Request) {
	nodeInfo, err := dao.FindLatest()
	if err != nil {
		ReturnJson(w, http.StatusInternalServerError, "Database error", false)
		return
	}
	ReturnJson(w, http.StatusOK, "Latest data", map[string]interface{}{
		"blockHeight": nodeInfo.BlockHeight,
		"nodeId":      nodeInfo.NodeID,
	})
}

// GET a node by its ID
func GetNodeById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	nodeInfo, err := dao.FindById(params["nodeId"])
	if err != nil {
		ReturnJson(w, http.StatusNotFound, "Invalid node ID", false)
		return
	}
	ReturnJson(w, http.StatusOK, "Node data", map[string]interface{}{
		"blockHeight": nodeInfo.BlockHeight,
		"nodeId":      nodeInfo.NodeID,
		"cpu":         nodeInfo.Cpu,
		"memory":      nodeInfo.Memory,
	})
}

// POST node info
func PostNodeInfo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var info models.NodeInfo
	if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
		ReturnJson(w, http.StatusBadRequest, "Invalid request payload", false)
		return
	}

	info.ID = bson.NewObjectId()
	if err := dao.Insert(info); err != nil {
		ReturnJson(w, http.StatusInternalServerError, err.Error(), false)
		return
	}
	ReturnJson(w, http.StatusCreated, "Node data added", info)
}
