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

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")
var validate = validator.New()

func GetOrders() gin.HandlerFunc{
	return func(c *gin.Context) {}
}

func GetOrder() gin.HandlerFunc{
	return func(c *gin.Context) {
		var ctx,cancel = context.WithTimeout(context.Background(), 100*time.Second)
		orderId:=c.Param("order_id")
		var order models.Order
		err:=orderCollection.FindOne(ctx, bson.M{"order_id": orderId}).Decode(&order)
		defer cancel()
		if err!=nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the order"})
		}
		c.JSON(http.StatusOK,order)
	}
}

func CreateOrder() gin.HandlerFunc{
	return func(c *gin.Context) {
		var ctx,cancel=context.WithTimeout(context.Background(), 100*time.Second)
		var order models.Order
		if err := c.BindJSON(&order); err!=nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr:=validate.Struct(order)
		if validationErr!=nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		order.Created_at,_ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.Updated_at,_ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.ID = primitive.NewObjectID()
		order_id := order.ID.Hex()
		order.Order_id = order_id
		result, insertErr := orderCollection.InsertOne(ctx, table)
		if insertErr!= nil{
			msg := fmt.Sprintf("Order was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return;
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func UpdateOrder() gin.HandlerFunc{
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(),100*time.Second)
		var order models.Order
		orderId:=c.Param("order_id")
		if err := c.BindJSON(&order); err!=nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var updateObj primitive.D
		order.Updated_at,_=time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, json.E{"updated_at", order.Updated_at})
		upsert:=true
		filter:=bson.M{"invoice_id":invoiceId}
		opt := options.UpdateOptions{
			Upsert: &upsert
		}
		result,err:=invoiceCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{Key: "$set", Value: updateObj},
			},
			&opt,
		)
		if err!=nil{
			msg:="Order update failed"
			c.JSON(http.StatusInternalServerError,gin.H{"error":msg})
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}