package models

import "gopkg.in/mgo.v2/bson"

type Request struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
}
