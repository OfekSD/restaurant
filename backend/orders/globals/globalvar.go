package globals

import "lola.com/lib"

var ConnectionPool *db.ConnectionPool


func InitializeConnectionPool(){
	ConnectionPool = db.NewConnectionPool(10)
}