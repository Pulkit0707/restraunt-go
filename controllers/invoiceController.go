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

var invoiceCollection *mongo.Collection = database.OpenCollection(database.Client, "invoice")
var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")
var validate = validator.New()

func GetInvoices() gin.HandlerFunc{
	return func(c *gin.Context) {}
}

func GetInvoice() gin.HandlerFunc{
	return func(c *gin.Context) {
		var ctx,cancel = context.WithTimeout(context.Background(), 100*time.Second)
		invoiceId:=c.Param("invoice_id")
		var invoice models.Invoice
		err:=invoiceCollection.FindOne(ctx, bson.M{"food_id": invoiceId}).Decode(&invoice)
		defer cancel()
		if err!=nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the invoice"})
		}
		c.JSON(http.StatusOK,invoice)
	}
}

func CreateInvoice() gin.HandlerFunc{
	return func(c *gin.Context) {
		var ctx,cancel=context.WithTimeout(context.Background(), 100*time.Second)
		var invoice models.Invoice
		if err := c.BindJSON(&invoice); err!=nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr:=validate.Struct(invoice)
		if validationErr!=nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		invoice.Created_at,_ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		invoice.Updated_at,_ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		invoice.ID = primitive.NewObjectID()
		invoice_id := invoice.ID.Hex()
		invoice.Invoice_id = invoice_id
		result, insertErr := invoiceCollection.InsertOne(ctx, table)
		if insertErr!= nil{
			msg := fmt.Sprintf("Invoice was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return;
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func UpdateInvoice() gin.HandlerFunc{
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(),100*time.Second)
		var invoice models.Invoice
		var order models.Order
		invoiceId:=c.Param("invoice_id")
		if err := c.BindJSON(&invoice); err!=nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var updateObj primitive.D
		if invoice.Payment_due_date!=nil{
			updateObj = append(updateObj, bson.E{"payment_due_date":invoice.Payment_due_date})
		}
		if invoice.Payment_method!=nil{
			updateObj = append(updateObj, bson.E{"payment_method":invoice.Payment_method})
		}
		if invoice.Payment_status!=nil{
			updateObj = append(updateObj, bson.E{"payment_status": invoice.Payment_status})
		}
		if invoice.Order_id!=nil{
			err:= orderCollection.FindOne(ctx, bson.M{"order_id":invoice.Order_id}).Decode(&order)
			defer cancel()
			if err!=nil{
				msg:=fmt.Sprintf("Order was not found")
				c.JSON(http.StatusInternalServerError, gin.H{"error":msg})
			}
		}
		invoice.Updated_at,_=time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, json.E{"updated_at", invoice.Updated_at})
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
			msg:="Invoice update failed"
			c.JSON(http.StatusInternalServerError,gin.H{"error":msg})
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}