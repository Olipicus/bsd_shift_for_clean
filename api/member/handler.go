package member

import (
	"net/http"

	"code.olipicus.com/go_rest_api/api/rest"
	"code.olipicus.com/go_rest_api/api/utility/mongo"
)

const collection string = "member"

// HandlerMember struct
type HandlerMember struct {
	rest.REST
}

//Handler ...
var Handler HandlerMember = HandlerMember{
	rest.REST{
		Collection: collection,
		OBJ:        Member{},
	},
}

//Random handler
func (handler *HandlerMember) Random(res http.ResponseWriter, req *http.Request) {
	var mgh mongo.Helper
	mgh.Init()

}
