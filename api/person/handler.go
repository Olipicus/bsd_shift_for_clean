package person

import (
	"code.olipicus.com/go_rest_api/api/rest"
)

const (
	collection   string = "people"
	mongoAddress string = "127.0.0.1:27017"
	dbName       string = "bsd_shift_for_clean"
)

// HandlerPerson struct
type HandlerPerson struct {
	rest.REST
}

//Handler ...
var Handler HandlerPerson = HandlerPerson{
	rest.REST{
		MongoAddress: mongoAddress,
		DBName:       dbName,
		Collection:   collection,
		OBJ:          Person{},
	},
}
