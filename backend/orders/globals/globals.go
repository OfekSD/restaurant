package globals

import "lola.com/lib/ConnectionPool"

var ConnectionPool *db.ConnectionPool


func InitializeConnectionPool(){
	ConnectionPool = db.NewConnectionPool(10)
}