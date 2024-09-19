package main

import (
	"context"
	stdSQL "database/sql"
	"errors"
	driver "github.com/go-sql-driver/mysql"
	"log"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-sql/v3/pkg/sql"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
)

var (
	logger = watermill.NewStdLogger(false, false)
)

func main() {

	router, err := message.NewRouter(message.RouterConfig{
		CloseTimeout: 10 * time.Second,
	}, logger)
	if err != nil {
		panic(err)
	}

	// The SQL publisher and subscriber don't immediately send the message to the dead letter queue (bug?)
	db := createDB()
	publisher, subscriber, err := getSqlPublisherSubscriber(db)

	// This works, it sends the message to the dead letter queue and stops retrying after 3 tries
	//pubSub := gochannel.NewGoChannel(gochannel.Config{}, logger)
	//publisher := pubSub
	//subscriber := pubSub

	poisonQueueMiddleware, err := middleware.PoisonQueue(publisher, "dead_letter_queue")
	if err != nil {
		panic(err)
	}

	router.AddMiddleware(
		poisonQueueMiddleware,
		middleware.Retry{
			MaxRetries:      3,
			InitialInterval: time.Millisecond * 100,
			Logger:          logger,
		}.Middleware,
		middleware.CorrelationID,
	)

	// Registering a handler for messages_topic
	router.AddNoPublisherHandler(
		"struct_handler", // handler name, must be unique
		"messages_topic", // topic from which we will read events
		subscriber,
		structHandler{}.Handler,
	)

	// Registering a handler for the DLQ. This is for debugging purposes.
	router.AddNoPublisherHandler(
		"dead_letter_handler", // handler name, must be unique
		"dead_letter_queue",   // topic from which we will read events
		subscriber,
		deadLetterHandler{}.Handler,
	)

	// Producing an incoming message in background
	go publishMessages(publisher)

	// Now that all handlers are registered, we're running the Router.
	// Run is blocking while the router is running.
	ctx := context.Background()
	if err := router.Run(ctx); err != nil {
		panic(err)
	}
}

func getSqlPublisherSubscriber(db *stdSQL.DB) (*sql.Publisher, *sql.Subscriber, error) {
	publisher, err := sql.NewPublisher(
		db,
		sql.PublisherConfig{
			SchemaAdapter: sql.DefaultMySQLSchema{},
		},
		logger,
	)
	if err != nil {
		return nil, nil, err
	}

	subscriber, err := sql.NewSubscriber(
		db,
		sql.SubscriberConfig{
			SchemaAdapter:    sql.DefaultMySQLSchema{},
			OffsetsAdapter:   sql.DefaultMySQLOffsetsAdapter{},
			InitializeSchema: true,
		},
		logger,
	)
	if err != nil {
		return nil, nil, err
	}
	return publisher, subscriber, nil
}

func publishMessages(publisher message.Publisher) {
	time.Sleep(2 * time.Second)

	msg := message.NewMessage(watermill.NewUUID(), []byte(`{"message": "Hello, world!"}`))
	middleware.SetCorrelationID(watermill.NewUUID(), msg)

	log.Printf("sending message %s, correlation id: %s\n", msg.UUID, middleware.MessageCorrelationID(msg))

	if err := publisher.Publish("messages_topic", msg); err != nil {
		panic(err)
	}
}

type structHandler struct {
}

func (s structHandler) Handler(msg *message.Message) error {
	log.Println("structHandler received message", msg.UUID)

	// simulate a long-running process
	time.Sleep(11 * time.Second)

	log.Println("structHandler returning error for message", msg.UUID)
	return errors.New("error from structHandler")
}

type deadLetterHandler struct {
}

func (s deadLetterHandler) Handler(msg *message.Message) error {
	log.Println("deadLetterHandler received message", msg.UUID)

	return nil
}

func createDB() *stdSQL.DB {
	conf := driver.NewConfig()
	conf.Net = "tcp"
	conf.User = "root"
	conf.Addr = "mysql"
	conf.DBName = "watermill"

	db, err := stdSQL.Open("mysql", conf.FormatDSN())
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
