package member

import (
	"net/http"

	"code.olipicus.com/go_rest_api/api/rest"
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

}
