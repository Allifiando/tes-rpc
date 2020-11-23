package controller

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"santara.co.id/set/market-common-service-grpc/pkg/proto/common"
)

// User ...
type User struct{}

// GetBearer ...
func (U *User) GetBearer(c *gin.Context) {
	conn, err := grpc.Dial(":50054", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := common.NewCommonServiceClient(conn)

	requestBearer := &common.RequestGetToken{Token: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjUxMiwiZGF0YSI6eyJpZCI6NTEyLCJ1dWlkIjoiNDQwODljYjQtNmYyNy0xMWU4LTg2NDEtMzM4NGZiNjBlZDg1IiwiZW1haWwiOiJwZWp1YW5nLnNhbnRhcmFAc2FudGFyYS5jby5pZCIsImNyZWF0ZWRfYXQiOiIyMDE4LTA1LTE5IDA5OjU1OjU3IiwidXBkYXRlZF9hdCI6IjIwMjAtMDgtMDMgMDM6MTI6NTAiLCJkZWxldGVkX2F0IjpudWxsLCJyb2xlX2lkIjoyLCJpc192ZXJpZmllZCI6MSwidHdvX2ZhY3Rvcl9hdXRoIjowLCJ0d29fZmFjdG9yX3NlY3JldCI6Im51bGwiLCJpc19sb2dnZWRfaW4iOjAsImlzX2RlbGV0ZWQiOjAsImNyZWF0ZWRfYnkiOm51bGwsInVwZGF0ZWRfYnkiOm51bGwsImlzX290cCI6MSwiYXR0ZW1wdCI6MCwiYXR0ZW1wdF9lbWFpbCI6MCwiZmluZ2VyX3ByaW50IjowLCJhdHRlbXB0X290cCI6bnVsbH0sImlhdCI6MTU5NjU5NTcyNn0.AQa2IZRifdUDD9vfIo6r4V8uEUocCQreRi2xDDHXLLg"}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	responseBearer, err := client.GetBearer(ctx, requestBearer)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Response Bearer:", responseBearer)

	c.JSON(200, responseBearer)
}
