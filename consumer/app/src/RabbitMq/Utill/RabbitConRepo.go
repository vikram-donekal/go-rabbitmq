package Utill

import (
	"../../../bin/amqp"
	"../DTO"
	"../DataBase"
	"fmt"
	"log"
)

func GetRabbitConnection() (*amqp.Connection) {

	conn, err := amqp.Dial("amqp://guest@localhost:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	return  conn
}

func CreateChannel(conn *amqp.Connection)  (channel *amqp.Channel){

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	return  ch
}


func  ExchangeDeclare(channel *amqp.Channel, exchangeName string,exchangeType string)  {

	var error  = channel.ExchangeDeclare(
		exchangeName,   // name
		exchangeType, // type
		false,     // durable
		true,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if error != nil {
		failOnError(error, "Failed to Declare Exchange")
	}
}

func QDeclare(channel *amqp.Channel, qName string) (q amqp.Queue){
	qDeclate, err := channel.QueueDeclare(
		qName, // name
		false,      // durable
		true,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		failOnError(err, "Failed to Declare Q")
	}
	return  qDeclate
}

func QBindWithExchange(channel *amqp.Channel ,QDeatils amqp.Queue,routingKey string ,exchangeName string){
	var err = channel.QueueBind(QDeatils.Name, routingKey, exchangeName, false, nil)

	if err != nil {
		failOnError(err, "Failed to Bind Exchange with  Q")
	}
}

func ConsumeMessage(channel *amqp.Channel,QDeatils amqp.Queue)(incomingMessages <-chan amqp.Delivery){


	incomingMessages, err := channel.Consume(
		QDeatils.Name, // name
		"",       // consumer tag
		true,    // auto ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait,
		nil,
	)
	if err != nil {
		failOnError(err, "Failed to Consume Messages")
	}

	return
}

func HandleMessages(incomingMessages <-chan amqp.Delivery) {
	log.Println("Entering HandleMessages")
	for MessageString := range incomingMessages {
		// Process message here.
		log.Printf( " \nMessage Recieved ----------> %s  \n\n\n\n  ",MessageString.Body )

		//Persiting Messages to DataBase.
		dto:= &DTO.MessageDTO{}
		dto.Data = string(MessageString.Body)
		DataBase.SaveMessages(dto)
	}
}



func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}