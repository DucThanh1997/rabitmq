package consumer

import (
	"encoding/json"
	"log"
	"os"
	gopher_and_rabbit "github.com/masnun/gopher-and-rabbit"
	"github.com/streadway/amqp"
)

func handleError (err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// Tạo connect
	conn, err := amqp.Dial(gopher_and_rabbit.Config.AMQPConnectionURL)
	handleError(err, "Không connect được")
	defer conn.Close()

	// Tạo channel
	ampqChannel, err := conn.Chanel()
	handleError(err, "Không tạo channel được")
	defer amqpChannel.Close()

	//Tạo queue tên là add
	queue, err := ampqChannel.QueueDeclare("add", true, false, false, false, nil)
	handleError(err, "Không tạo queue được")

	// config
	err = ampqChannel.Qos(1, 0, false)
	handleError(err, "không config được")

	messageChannel, err := amqpChannel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Could not register consumer")

	stopChan := make(chan bool)

	go func() {
		log.Printf("Consumer ready, PID: %d", os.Getpid())
		for d := range messageChannel {
			log.Printf("Received a message: %s", d.Body)

			addTask := &gopher_and_rabbit.AddTask{}

			err := json.Unmarshal(d.Body, addTask)

			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
			}

			log.Printf("Result of %d + %d is : %d", addTask.Number1, addTask.Number2, addTask.Number1+addTask.Number2)

			if err := d.Ack(false); err != nil {
				log.Printf("Error acknowledging message : %s", err)
			} else {
				log.Printf("Acknowledged message")
			}

		}
	}()

	// Stop for program termination
	<-stopChan
}