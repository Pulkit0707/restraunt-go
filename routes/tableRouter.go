package routes

import(
	"github.com/gin-gonic/gin"
	"restraunt-go/controllers"
)

func TableRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/tables", controllers.GetTables())
	incomingRoutes.GET("/tables/:table_id", controllers.GetTable())
	incomingRoutes.POST("/tables",controllers.CreateTable())
	incomingRoutes.PATCH("/tables/:table_id",controllers.UpdateTable())
}