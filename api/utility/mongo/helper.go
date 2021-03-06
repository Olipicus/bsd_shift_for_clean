package mongo

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var dbName string

//Helper Struct of MongoHelper
type Helper struct {
	session *mgo.Session
}

//Init : Initial DB
func (h *Helper) Init(mongoAddress string, databaseName string) {
	session, err := mgo.Dial(mongoAddress)
	dbName = databaseName
	if err != nil {
		log.Fatal(err)
	}
	h.session = session
	//log.Println("Connect MongoDB OK!")

	// Optional. Switch the session to a monotonic behavior.
	h.session.SetMode(mgo.Monotonic, true)
}

//Close : Close DB Session
func (h *Helper) Close() {
	h.session.Close()
}

//GetOneData : Get Single Document
func (h *Helper) GetOneData(collectionName string, id string) (interface{}, error) {
	c := h.session.DB(dbName).C(collectionName)
	var obj interface{}
	err := c.FindId(bson.ObjectIdHex(id)).One(&obj)
	return obj, err
}

//GetOneDataToObj : Get Single Document
func (h *Helper) GetOneDataToObj(collectionName string, id string, obj interface{}) error {
	c := h.session.DB(dbName).C(collectionName)
	err := c.FindId(bson.ObjectIdHex(id)).One(obj)
	return err
}

//RemoveByID : Remove Data By ID
func (h *Helper) RemoveByID(collectionName string, id string) error {
	c := h.session.DB(dbName).C(collectionName)
	return c.RemoveId(bson.ObjectIdHex(id))

}

//InsertData : Insert Document to Collection
func (h *Helper) InsertData(collectionName string, obj interface{}) error {
	c := h.session.DB(dbName).C(collectionName)
	return c.Insert(obj)
}

//UpdateData : Update Document
func (h *Helper) UpdateData(collectionName string, id string, obj interface{}) error {
	c := h.session.DB(dbName).C(collectionName)
	update := bson.M{"$set": obj}
	return c.UpdateId(bson.ObjectIdHex(id), update)
}

//GetCollecitonObj : get collection from mgo
func (h *Helper) GetCollecitonObj(collectionName string) *mgo.Collection {
	return h.session.DB(dbName).C(collectionName)
}
