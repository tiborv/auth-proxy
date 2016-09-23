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
