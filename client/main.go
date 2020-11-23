package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	routes "santara.co.id/set/market-common-service-grpc/client/route"
)

func main() {
	// log.Println("Client running ...")

	// conn, err := grpc.Dial(":50054", grpc.WithInsecure(), grpc.WithBlock())
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer conn.Close()

	// client := common.NewCommonServiceClient(conn)

	// requestBearer := &common.RequestGetToken{Token: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjUxMiwiZGF0YSI6eyJpZCI6NTEyLCJ1dWlkIjoiNDQwODljYjQtNmYyNy0xMWU4LTg2NDEtMzM4NGZiNjBlZDg1IiwiZW1haWwiOiJwZWp1YW5nLnNhbnRhcmFAc2FudGFyYS5jby5pZCIsImNyZWF0ZWRfYXQiOiIyMDE4LTA1LTE5IDA5OjU1OjU3IiwidXBkYXRlZF9hdCI6IjIwMjAtMDgtMDMgMDM6MTI6NTAiLCJkZWxldGVkX2F0IjpudWxsLCJyb2xlX2lkIjoyLCJpc192ZXJpZmllZCI6MSwidHdvX2ZhY3Rvcl9hdXRoIjowLCJ0d29fZmFjdG9yX3NlY3JldCI6Im51bGwiLCJpc19sb2dnZWRfaW4iOjAsImlzX2RlbGV0ZWQiOjAsImNyZWF0ZWRfYnkiOm51bGwsInVwZGF0ZWRfYnkiOm51bGwsImlzX290cCI6MSwiYXR0ZW1wdCI6MCwiYXR0ZW1wdF9lbWFpbCI6MCwiZmluZ2VyX3ByaW50IjowLCJhdHRlbXB0X290cCI6bnVsbH0sImlhdCI6MTU5NjU5NTcyNn0.AQa2IZRifdUDD9vfIo6r4V8uEUocCQreRi2xDDHXLLg"}

	// requestPin := &common.RequestGetPin{
	// 	Token: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjUxMiwiZGF0YSI6eyJpZCI6NTEyLCJ1dWlkIjoiNDQwODljYjQtNmYyNy0xMWU4LTg2NDEtMzM4NGZiNjBlZDg1IiwiZW1haWwiOiJwZWp1YW5nLnNhbnRhcmFAc2FudGFyYS5jby5pZCIsImNyZWF0ZWRfYXQiOiIyMDE4LTA1LTE5IDA5OjU1OjU3IiwidXBkYXRlZF9hdCI6IjIwMjAtMDgtMDMgMDM6MTI6NTAiLCJkZWxldGVkX2F0IjpudWxsLCJyb2xlX2lkIjoyLCJpc192ZXJpZmllZCI6MSwidHdvX2ZhY3Rvcl9hdXRoIjowLCJ0d29fZmFjdG9yX3NlY3JldCI6Im51bGwiLCJpc19sb2dnZWRfaW4iOjAsImlzX2RlbGV0ZWQiOjAsImNyZWF0ZWRfYnkiOm51bGwsInVwZGF0ZWRfYnkiOm51bGwsImlzX290cCI6MSwiYXR0ZW1wdCI6MCwiYXR0ZW1wdF9lbWFpbCI6MCwiZmluZ2VyX3ByaW50IjowLCJhdHRlbXB0X290cCI6bnVsbH0sImlhdCI6MTU5NjU5NTcyNn0.AQa2IZRifdUDD9vfIo6r4V8uEUocCQreRi2xDDHXLLg",
	// 	Pin:   "123456",
	// }

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	// defer cancel()

	// responseBearer, err := client.GetBearer(ctx, requestBearer)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// responsePin, err := client.GetPin(ctx, requestPin)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// log.Println("Response Bearer:", responseBearer)
	// log.Println("Response PIN:", responsePin)
	// if os.Getenv("ENV") == "development" {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error getting env")
	}
	// }
	port := os.Getenv("PORTS")
	if port == "" {
		port = "1111"
	}
	r := routes.Init()
	// running
	r.Run(":" + port)
}
