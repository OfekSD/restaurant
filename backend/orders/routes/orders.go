package routes

import "github.com/gin-gonic/gin"
import "lola.com/globals"
import "fmt"

func Orders(router *gin.RouterGroup){
	{
		router.GET("/",getOrders)
	}
}

type Food struct{
	name string 
	price float32
}

// var connectionPool *database.ConnectionPool = globalvars.ConnectionPool


func getOrders(ctx *gin.Context){
	con := globals.ConnectionPool.GetConnection()
	// rows, _ := con.Query(`select name from foods`)

	// println(err.Error())
	rows, err := con.Query(`select name, price from foods`)
	if err != nil{panic(err)}
	defer rows.Close()
	results := make([]string,0)
	
	for rows.Next(){
	// 	// var name string
		s := Food{}
	// 	// var name string
	// 	// var price float32
		
		_ = rows.Scan(&s.name,&s.price)
		fmt.Printf("%s %.2f\n",s.name,s.price)
		results = append(results,s.name)
	}
	globals.ConnectionPool.ReturnConnection(con)

	ctx.JSON(200,gin.H{"results":results})
}