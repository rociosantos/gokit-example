package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/rociosantos/gokit-example/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rociosantos/gokit-example/repo"
	"github.com/rociosantos/gokit-example/transport"
	"github.com/rociosantos/gokit-example/endpoint"
	
	awssession "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/sirupsen/logrus"
)

func main()  {
	var httpAddr =  flag.String("http", ":8085", "http listen address")

	// logger setup
	logger, err := createLogger()
	if err != nil || logger == nil {
		log.Fatal("creating logger: %w", err)
	}

	logger.Info("service started")
	defer logger.Info("service ended")

	AWSDynamoDBClient, err := createAWSDynamoDBClient(logger)
	if err != nil {
		logger.
			WithError(err).
			Fatal("creating DynamoDB client")
	}

	flag.Parse()
	ctx := context.Background()
	repository := repo.NewRepo(AWSDynamoDBClient, logger)
	svc := service.NewService(repository, logger)

	errs := make(chan error)

	go func(){
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s" ,<- c)
	}()

	e := endpoint.MakeEndpoints(svc)

	go func() {
		fmt.Println("listening on port 8085")
		handler := transport.NewHTTPServer(ctx, e)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	logger.Debug("exit", <-errs)
}

func createLogger() (*logrus.Logger, error) {
	logLevel := "debug"
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"log_level": logLevel,
		}).Error("parsing log_level")

		return nil, err
	}

	logger := logrus.New()
	logger.SetLevel(level)
	logger.Out = os.Stdout
	logger.Formatter = &logrus.JSONFormatter{}
	return logger, nil
}

func createAWSDynamoDBClient(logger *logrus.Logger) (*repo.DynamoClient, error) {
	// Apparently, Must() will panic if the session fails, unlike NewSession()
	awsSession, err := createAWSSession()
	if err != nil {
		return nil, err
	}
	tables := map[string]string{
		"users": "users",
	}
	d := repo.NewDynamoClient(tables, awsSession, logger)

	return d, nil
}

func createAWSSession() (*awssession.Session, error) {
	conf := aws.NewConfig().
		WithRegion("us-east-1").
		WithEndpoint("http://localhost:4569")

	return awssession.NewSession(conf)
}

