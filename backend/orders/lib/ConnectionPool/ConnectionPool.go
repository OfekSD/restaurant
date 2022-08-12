package db

import (
	"database/sql"
	"sync"
	"fmt"
)

const (
	dbname	  = "restaurant"
	user	  = "postgres"
	password  = "Aa123456"
	host 	  = "localhost"
	port      = 5432
)

type ConnectionPool struct{
	connections []*sql.DB
	size int
	lock sync.Mutex
}

func (c *ConnectionPool) GetConnection() *sql.DB{
	for len(c.connections) == 0 {}
	c.lock.Lock()
	con := c.connections[1]
	c.connections = (c.connections[1:])
	c.lock.Unlock()
	return con
}

func (c *ConnectionPool) ReturnConnection(con *sql.DB){
	c.lock.Lock()
	c.connections = append(c.connections, con)
	c.lock.Unlock()
}

func (c *ConnectionPool) Close(){
	c.lock.Lock()
	for _,con := range(c.connections){
		con.Close()
	}
	c.lock.Unlock()
}

func NewConnectionPool(size int) *ConnectionPool{

	connections := make([]*sql.DB,0)
	for i:=0;i<size;i++{
		conn := new_con()
		connections = append(connections, conn)
	}

	return &ConnectionPool{connections:connections, size: size}
	
}



func new_con() *sql.DB{
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil{
		panic(err)
	}
	return db
	
}