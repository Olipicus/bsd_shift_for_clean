package member

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"code.olipicus.com/go_rest_api/api/rest"
	"code.olipicus.com/go_rest_api/api/utility/mongo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	collection   string = "member"
	mongoAddress string = "127.0.0.1:27017"
	dbName       string = "bsd_shift_for_clean"
)

// HandlerMember struct
type HandlerMember struct {
	rest.REST
}

//Handler ...
var Handler HandlerMember = HandlerMember{
	rest.REST{
		MongoAddress: mongoAddress,
		DBName:       dbName,
		Collection:   collection,
		OBJ:          Member{},
	},
}

//Result handler
func (handler *HandlerMember) Result(res http.ResponseWriter, req *http.Request) {
	var mgh mongo.Helper

	mgh.Init(Handler.MongoAddress, Handler.DBName)
	defer mgh.Close()

	dayColor := map[int]string{
		1: "yellow",
		2: "pink",
		3: "green",
		4: "orange",
		5: "blue",
	}

	var listResult []ResultDay
	//for day, color := range dayColor {
	for i := 1; i <= len(dayList); i++ {
		var result ResultDay
		result.Day = dayList[i]
		result.Color = dayColor[i]
		resultCollection := mgh.GetCollecitonObj(handler.Collection)
		err := resultCollection.Find(bson.M{"day": dayList[i]}).All(&result.Members)

		if err != nil {
			panic(err)
		}

		listResult = append(listResult, result)

	}

	handler.REST.ResponseDataResult(res, rest.Result{StatusCode: 200, Description: "success"}, listResult)

}

//Random handler
func (handler *HandlerMember) Random(res http.ResponseWriter, req *http.Request) {
	var mgh mongo.Helper

	mgh.Init(Handler.MongoAddress, Handler.DBName)
	defer mgh.Close()

	vars := mux.Vars(req)
	id := vars["id"]

	var member Member
	err := mgh.GetOneDataToObj(handler.Collection, id, &member)
	log.Print(id)

	if err == mgo.ErrNotFound {
		handler.REST.ResponseResult(res, rest.Result{StatusCode: 404, Description: "Data Not Found"})
	}

	objResult := AssignDay(id, &mgh)

	var listMember []Member
	resultCollection := mgh.GetCollecitonObj(handler.Collection)
	err = resultCollection.Find(bson.M{"day": objResult.Day}).All(&listMember)

	if err != nil {
		panic(err)
	}

	handler.REST.ResponseDataResult(res, rest.Result{StatusCode: 200, Description: "success"}, listMember)

}
