package controllers

import (
	"context"
	"fmt"
	"log"
	"math"
	"net/http"
	"restraunt-go/database"
	"restraunt-go/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func GetUsers() gin.HandlerFunc{
	return func(c*gin.Context){}
}

func GetUser() gin.HandlerFunc{
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		userId:=c.Param("user_id")
		var user models.User
		err:=userCollection.FindOne(ctx, bson.M{"user_id":userId}).Decode(&user)
		defer cancel()
		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":"error occured while fetching user"})
		}
		c.JSON(http.StatusOK,user)
	}
}

func SignUp() gin.HandlerFunc{
	return func(c *gin.Context) {}
}

func Login() gin.HandlerFunc{
	return func(c *gin.Context) {}
}

func HashPassword(password string){}

func VerifyPassword(userPassword string, providedPassword string)(bool, string ){}