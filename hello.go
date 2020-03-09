package main
import "fmt"
import "github.com/onsi/gomega"
import "github.com/aws/aws-lambda-go/lambda"
func hello() (string, error) {
	return "Hello ƛ!", nil
}
func main() {
	fmt.Println("Hello, world.")
	lambda.Start(hello)
}
