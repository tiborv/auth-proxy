package models

import (
	"bytes"
	"net/http"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var requestCollection *mgo.Collection
var responseCollection *mgo.Collection

type Request struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Timestamp time.Time
	Token     string
	Service   string
	Data      string
}

type Response struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Duration time.Duration
	Token    string
	Service  string
	Data     string
}

func (r Request) Save() {
	r.ID = bson.NewObjectId()
	requestCollection.Insert(r)
}

func (r Response) Save() {
	r.ID = bson.NewObjectId()
	responseCollection.Insert(r)
}

func LogRequest(t, s string, ts time.Time, h http.Header) {
	go func() {
		if !mongoStatsEnabled {
			return
		}
		buff := bytes.Buffer{}
		h.Write(&buff)
		Request{
			Timestamp: ts,
			Data:      buff.String(),
			Token:     t,
			Service:   s,
		}.Save()
	}()
}

func LogResponse(t, s string, duration time.Duration, w http.ResponseWriter) {
	go func() {
		if !mongoStatsEnabled {
			return
		}
		buff := bytes.Buffer{}
		w.Header().Write(&buff)
		Response{
			Duration: duration,
			Data:     buff.String(),
			Token:    t,
			Service:  s,
		}.Save()
	}()
}

type ResponseStats struct {
	Avg   float64 `json:"avg"`
	Count int     `json:"count"`
}

func (s *Service) GetResponseStats() (ResponseStats, error) {
	result := ResponseStats{}
	pipe := responseCollection.Pipe([]bson.M{
		{"$match": bson.M{"service": s.Slug}},
		{"$group": bson.M{
			"_id":   "null",
			"avg":   bson.M{"$avg": "$duration"},
			"count": bson.M{"$sum": 1},
		}},
	})
	err := pipe.One(&result)
	s.Stats = result
	return result, err
}

type RequestStats struct {
	Count int `json:"count"`
}

func (c *Client) GetRequestStats() (RequestStats, error) {
	result := RequestStats{}
	pipe := requestCollection.Pipe([]bson.M{
		{"$match": bson.M{"token": c.Token}},
		{"$group": bson.M{
			"_id":   "null",
			"count": bson.M{"$sum": 1},
		}},
	})
	err := pipe.One(&result)
	c.Stats = result
	return result, err
}
