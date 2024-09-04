package internal

//
//import (
//	"fmt"
//	amqp "github.com/rabbitmq/amqp091-go"
//	"main/utils"
//)
//
//type RabbitMQIns struct {
//}
//
//func NewRabbitMQInstance() *RabbitMQIns {
//	return &RabbitMQIns{}
//}
//
//func (r *RabbitMQIns) ConnectRabbitMQ(config utils.RabbitMQConfig) *amqp.Connection {
//	connectionString := fmt.Sprintf("amqp://%v:%v@%v:%v/", config.Username, config.Password, config.Host, config.Port)
//	conn, err := amqp.Dial(connectionString)
//
//	if err != nil {
//		panic(err)
//	}
//	utils.ShowInfoLogs(fmt.Sprintf("Connected to rabbitMQ %s", config.Host))
//	return conn
//}
