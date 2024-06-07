module application.go

replace github.com/aws/aws-xray-sdk-go => ../../

require (
	github.com/aws/aws-sdk-go v1.47.9
	golang.org/x/net v0.23.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
)

go 1.19
