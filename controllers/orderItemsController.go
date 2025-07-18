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

var orderItemCollection *mongo.Collection = database.OpenCollection(database.Client,"order_item")
var orderCollection *mongo.Collection = database.OpenCollection(database.Client,"order")
var validate = validator.New()

func GetOrderItems() gin.HandlerFunc{
	return func(c *gin.Context) {}
}

func GetOrderItemsByOrder() gin.HandlerFunc{
	return func(c *gin.Context) {}
}

func ItemsByOrder(id string) (OrderItems []primitive.M, err error){}

func GetOrderItem() gin.HandlerFunc{
	return func(c *gin.Context) {}
}

func CreateOrderItem() gin.HandlerFunc{
	return func(c *gin.Context) {}
}

func UpdateOrderItem() gin.HandlerFunc{
	return func(c *gin.Context) {}
}