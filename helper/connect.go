package helper

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

func ConnectRabbitMQServer() *amqp.Connection {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panicln("Error loading .env file")
	}

	rabbitMQServer := os.Getenv("RABBITMQ_SERVER")
	rabbitMQPort := os.Getenv("RABBITMQ_PORT")
	rabbitMQUsername := os.Getenv("RABBITMQ_USERNAME")
	rabbitMQPassword := os.Getenv("RABBITMQ_PASSWORD")

	constring := fmt.Sprintf(`amqp://%s:%s@%s:%s/`, rabbitMQUsername, rabbitMQPassword, rabbitMQServer, rabbitMQPort)

	conn, err2 := amqp.Dial(constring)

	if err2 != nil {
		log.Panicln("Error connecting to RabbitMQ " + err2.Error())

		return nil
	}

	return conn
}
