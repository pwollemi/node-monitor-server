package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"gopkg.in/mgo.v2/bson"
)

func ReturnJson(w http.ResponseWriter, code int, msg string, payload interface{}) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	if payload == false {
		json.NewEncoder(w).Encode(bson.M{"code": code, "success": false, "msg": msg, "data": nil})
		return
	}
	json.NewEncoder(w).Encode(bson.M{"code": code, "success": true, "msg": msg, "data": payload})
}

func decodeTimestamp(ts string) time.Time {
	i, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return time.Now()
	}
	res := time.Unix(i/1000, (i%1000)*1000*1000)
	return res
}

func getIntervals(from time.Time, to time.Time) ([]uint64, []uint64, []uint64) {
	h, m, s := getHMSInterval(from, to)
	hours := []uint64{}
	mins := []uint64{}
	secs := []uint64{}
	var i uint64 = 0
	for ; i < 60; i++ {
		if i < 24 && h > 0 && i%h == 0 {
			hours = append(hours, i)
		}
		if m > 0 && i%m == 0 {
			mins = append(mins, i)
		}
		if i%5 == 0 && s > 0 && i%s == 0 {
			secs = append(secs, i)
		}
	}
	return hours, mins, secs
}

func getHMSInterval(from time.Time, to time.Time) (uint64, uint64, uint64) {
	const defaultRecordCnt uint64 = 600

	cur := 5
	secTicks := []uint64{5}
	for i := 1; i < 60; i++ {
		if i > cur && 60%i == 0 && i%5 == 0 {
			cur = i
		}
		secTicks = append(secTicks, uint64(cur))
	}
	cur = 1
	minTicks := []uint64{0}
	for i := 1; i < 60; i++ {
		if i > cur && 60%i == 0 {
			cur = i
		}
		minTicks = append(minTicks, uint64(cur))
	}
	cur = 1
	hourTicks := []uint64{0}
	for i := 1; i < 24; i++ {
		if i > cur && 24%i == 0 {
			cur = i
		}
		hourTicks = append(hourTicks, uint64(cur))
	}

	interval := uint64(to.Sub(from).Seconds()) / defaultRecordCnt / 5 * 5
	if interval < 60 {
		return 0, 0, secTicks[interval]
	}
	interval = interval / 60
	if interval < 60 {
		return 0, minTicks[interval], 0
	}
	interval = interval / 60 % 24
	return hourTicks[interval], 0, 0
}
