package member

import (
	"net/http"

	"github.com/gorilla/mux"

	"code.olipicus.com/go_rest_api/api/rest"
	"code.olipicus.com/go_rest_api/api/utility/mongo"
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

//Random handler
func (handler *HandlerMember) Random(res http.ResponseWriter, req *http.Request) {
	var mgh mongo.Helper
	mgh.Init(Handler.MongoAddress, Handler.DBName)

	defer mgh.Close()

	vars := mux.Vars(req)
	id := vars["id"]

	objMember := AssignDay(id, &mgh)

	handler.REST.ResponseDataResult(res, rest.Result{200, "success"}, objMember)

}
