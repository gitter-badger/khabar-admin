package models

import "github.com/bulletind/khabar-admin/Godeps/_workspace/src/gopkg.in/mgo.v2/bson"

type Article struct {
	Id        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string        `json:"title" form:"title" binding:"required" bson:"title"`
	Body      string        `json:"body" form:"body" binding:"required" bson:"body"`
	CreatedOn int64         `json:"created_on" bson:"created_on"`
	// User      bson.ObjectId `json:"user"`
}
