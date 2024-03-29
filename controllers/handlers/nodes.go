package handlers

import (

	//"log"

	"encoding/json"
	"net/http"
	"time"

	"github.com/flashguru-git/node-monitor-server/dao"
	"github.com/flashguru-git/node-monitor-server/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// GET latest block height
func GetLatest(w http.ResponseWriter, r *http.Request) {
	nodeMetric, err := dao.FindLatest()
	if err != nil {
		ReturnJson(w, http.StatusInternalServerError, "Database error", false)
		return
	}
	ReturnJson(w, http.StatusOK, "Latest data", map[string]interface{}{
		"blockHeight": nodeMetric.BlockHeight,
		"nodeId":      nodeMetric.NodeID,
	})
}

// GET a node by its ID
func GetNodeById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()

	from := time.Now().Add(time.Hour * time.Duration(-1))
	to := time.Now()
	from_ts := query.Get("from_ts")
	to_ts := query.Get("to_ts")
	if len(from_ts) > 0 {
		from = decodeTimestamp(query.Get("from_ts"))
	}
	if len(to_ts) > 0 {
		to = decodeTimestamp(query.Get("to_ts"))
	}

	h, m, s := getIntervals(from, to)
	nodeInfos, err := dao.FindWithOption(params["nodeId"], from, to, h, m, s)
	if err != nil {
		ReturnJson(w, http.StatusNotFound, "Invalid node ID", false)
		return
	}
	ReturnJson(w, http.StatusOK, "Node data", map[string]interface{}{
		"count": len(nodeInfos),
		"data":  nodeInfos,
	})
}

// GET all nodes
func GetAllNodes(w http.ResponseWriter, r *http.Request) {
	nodeIds, err := dao.FindAllNodeId()
	if err != nil {
		ReturnJson(w, http.StatusInternalServerError, "Can't get nodes list", false)
		return
	}
	res := map[string]interface{}{}
	for _, id := range nodeIds {
		nodeMetric, _ := dao.FindById(id)
		res[id] = nodeMetric
	}
	ReturnJson(w, http.StatusOK, "Node data", res)
}

// POST node metric
func CreateNodeMetric(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var infoData models.NodeMetricRequest
	if err := json.NewDecoder(r.Body).Decode(&infoData); err != nil {
		ReturnJson(w, http.StatusBadRequest, "Invalid request payload", false)
		return
	}

	metric := models.NodeMetric{
		ID:          bson.NewObjectId(),
		NodeID:      infoData.NodeID,
		BlockHeight: infoData.BlockHeight,
		TimeStamp:   infoData.CreatedAt.Unix(),
		CreatedAt:   infoData.CreatedAt,
		Cpu:         infoData.Cpu,
		Memory:      infoData.Memory,
		Hour:        uint64(infoData.CreatedAt.Hour()),
		Min:         uint64(infoData.CreatedAt.Minute()),
		Sec:         uint64(infoData.CreatedAt.Second()) / 5 * 5,
	}
	if err := dao.Insert(metric); err != nil {
		ReturnJson(w, http.StatusInternalServerError, err.Error(), false)
		return
	}
	ReturnJson(w, http.StatusCreated, "Node data added", metric)
}
