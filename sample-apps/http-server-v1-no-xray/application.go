package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"golang.org/x/net/context/ctxhttp"
)

func webServer() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("healthcheck"))
	}))

	// test http instrumentation
	http.Handle("/outgoing-http-call", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := ctxhttp.Get(r.Context(), http.DefaultClient, "https://aws.amazon.com")
		if err != nil {
			log.Println(err)
			return
		}
		_, _ = w.Write([]byte("Tracing http call!"))
	}))

	// test aws sdk instrumentation
	http.Handle("/aws-sdk-call", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		testAWSCalls(r.Context())
		_, _ = w.Write([]byte("Tracing aws sdk call!"))
	}))

	listenAddress := os.Getenv("LISTEN_ADDRESS")
	if listenAddress == "" {
		listenAddress = "127.0.0.1:5000"
	}

	_ = http.ListenAndServe(listenAddress, nil)
	log.Printf("App is listening on %s !", listenAddress)
}

func testAWSCalls(ctx context.Context) {
	awsSess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")}))

	s3Client := s3.New(awsSess)
	if _, err := s3Client.ListBucketsWithContext(ctx, nil); err != nil {
		log.Println(err)
		return
	}

	log.Println("Successfully traced aws sdk call")
}

func main() {
	webServer()
}
