package Utill

import (
	"fmt"
	"../../../bin/amqp"
	"log"
	"math/rand"

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

func PulishMessage(channel *amqp.Channel,exchangeName string,QDeatils amqp.Queue){

	var MessageString = randomString()
	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(MessageString),
	}
	log.Printf( " \nMessage ----------> %s  \nExchangeName  ---------> %s \nQName  ---------> %s \n\n\n  ",MessageString ,exchangeName,QDeatils.Name)
	err:= channel.Publish(exchangeName, QDeatils.Name, false, false, msg)
	if err!= nil{
		failOnError(err,"Failed to Publish message")
	}
}

func  randomString() string {
	//string Generator with range of min and max characters
	return  StringGenerator(rand.Intn(200 - 150) + 15)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}