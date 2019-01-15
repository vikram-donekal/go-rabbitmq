package main

import (
	"../RabbitMq/Utill"
	"fmt"
	 "../../bin/amqp"
	"log"
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
	//defer rabbitConn.Close()
	channel := Utill.CreateChannel(rabbitConn)
	//defer channel.Close()
	Utill.ExchangeDeclare(channel, ExchangeName, ExchangeType)
	qDetails := Utill.QDeclare(channel, QName)
	Utill.QBindWithExchange(channel, qDetails, RoutingKey, ExchangeName)
	log.Printf("Ready to Accept Messages  from Q : %s",QName)


	go consumeMessages(channel,qDetails)
}

 func consumeMessages(channel *amqp.Channel,qDetails amqp.Queue){
	 incomingMessagesChannel := Utill.ConsumeMessage(channel,qDetails)
	 Utill.HandleMessages(incomingMessagesChannel)


	 var a string
	 fmt.Scanln(&a)

}

