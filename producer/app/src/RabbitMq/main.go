package main

import (
	"../../bin/amqp"
	"../RabbitMq/Utill"
	"fmt"
	"runtime"
)

const (
	QName = "TestingQ"
	ExchangeName = "TestingExchange"
	ExchangeType ="fanout"
	RoutingKey ="" //It will ignore as ExchangeTYpe is fanout
)



func main() {
	runtime.GOMAXPROCS(4) //TO Use multiCore Cpus
	server()
	var a string
	fmt.Scanln(&a)
}



func server() {

	rabbitConn := Utill.GetRabbitConnection()
	// defer rabbitConn.Close()
	channel := Utill.CreateChannel(rabbitConn)
	//defer channel.Close()
	Utill.ExchangeDeclare(channel, ExchangeName, ExchangeType)
	qDetails := Utill.QDeclare(channel, QName)
	Utill.QBindWithExchange(channel, qDetails, RoutingKey, ExchangeName)

	go publishingMessage(channel,qDetails)
}

 func publishingMessage(channel *amqp.Channel,qDetails amqp.Queue){
	for {

			Utill.PulishMessage(channel,ExchangeName,qDetails)
		}
}

