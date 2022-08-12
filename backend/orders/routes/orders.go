package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"lola.com/globals"
	"lola.com/models"
)

func Orders(router *gin.RouterGroup){
	{
		router.GET("/",getOrders)
		router.POST("/",createOrder)
		router.GET("/:id",getOrder)
		router.DELETE("/:id",deleteOrder)
	}
}



func getOrders(ctx *gin.Context){
	con := globals.ConnectionPool.GetConnection()
	defer globals.ConnectionPool.ReturnConnection(con)
	
	rows, err := con.Query(`SELECT id, dishes, orderer, order_time, delivered from orders`)
	if err != nil{panic(err)}
	defer rows.Close()
	results := make([]models.Order,0)
	
	for rows.Next(){
		o := models.Order{}
		_ = rows.Scan(&o.Id,pq.Array(&o.Dishes),&o.Orderer,&o.OrderTime,&o.Delivered)
		results = append(results,o)
	}
	// fmt.Println(results[0])
	
	ctx.JSON(200,results)
}

func createOrder(ctx *gin.Context){
	var order models.Order
	ctx.BindJSON(&order)
	fmt.Println(order)	
	fmt.Printf("%s\n",order.Dishes)
	con := globals.ConnectionPool.GetConnection()
	defer globals.ConnectionPool.ReturnConnection(con)

	rows, err := con.Query(`insert into orders(orderer,dishes) values($1,$2)
	 returning id,dishes,orderer,order_time,delivered`,order.Orderer,pq.Array(order.Dishes))
	defer rows.Close()	
	if err != nil{
		fmt.Println(err)
	}
	rows.Next()
	rows.Scan(&order.Id,&order.Dishes,&order.Orderer,&order.OrderTime,&order.Delivered)

	ctx.JSON(200,order)
}

func getOrder(ctx *gin.Context){

	id := ctx.Param("id")
	order:= models.Order{}
	
	con := globals.ConnectionPool.GetConnection()
	defer globals.ConnectionPool.ReturnConnection(con)
	err := con.QueryRow(`select id, dishes, orderer, order_time, delivered from orders where id=$1`,id).Scan(&order.Id,pq.Array(&order.Dishes),&order.Orderer,&order.OrderTime,&order.Delivered)
	if err != nil{
		println(err.Error())
		ctx.JSON(404,gin.H{})
		return
	}
	ctx.JSON(200,order)
	

}


func deleteOrder(ctx *gin.Context){
	id := ctx.Param("id")
	order:= models.Order{}
	
	con := globals.ConnectionPool.GetConnection()
	defer globals.ConnectionPool.ReturnConnection(con)
	err := con.QueryRow(`select id, dishes, orderer, order_time, delivered from orders where id=$1`,id).Scan(&order.Id,pq.Array(&order.Dishes),&order.Orderer,&order.OrderTime,&order.Delivered)
	if err != nil{
		println(err.Error())
		ctx.JSON(404,gin.H{})
		return
	}
	fmt.Println("lola")
	con.Query(`delete from orders where id=$1`,id)
	ctx.JSON(200,gin.H{})
	

}


