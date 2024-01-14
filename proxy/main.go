package main

import "fmt"

func main() {
	fmt.Println("Proxy Pattern")

	nginxServer := newNginxServer()
	appStatusUrl := "/app/status"
	createuserUrl := "/create/user"

	httpCode, body := nginxServer.handleRequest(appStatusUrl, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusUrl, httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusUrl, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusUrl, httpCode, body)

	httpCode, body = nginxServer.handleRequest(createuserUrl, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", createuserUrl, httpCode, body)

	httpCode, body = nginxServer.handleRequest(createuserUrl, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", createuserUrl, httpCode, body)
}
