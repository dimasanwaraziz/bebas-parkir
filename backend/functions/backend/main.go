package handler

import (
	"fmt"
	"os"

	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/open-runtimes/types-for-go/v4/openruntimes"
)

type Response struct {
	Motto       string `json:"motto"`
	Learn       string `json:"learn"`
	Connect     string `json:"connect"`
	GetInspired string `json:"getInspired"`
}

func Main(Context openruntimes.Context) openruntimes.Response {
	var _ = appwrite.NewClient(
		appwrite.WithProject(os.Getenv("APPWRITE_FUNCTION_PROJECT_ID")),
		appwrite.WithKey(Context.Req.Headers["x-appwrite-key"]),
	)

	fmt.Println("Hello, Logs!")

	fmt.Fprintln(os.Stderr, "Error:", "Hello, Errors!")

	if Context.Req.Method == "GET" {
		return Context.Res.Text("Hello, World!")
	}

	return Context.Res.Json(
		Response{
			Motto:       "Build like a team of hundreds_",
			Learn:       "https://appwrite.io/docs",
			Connect:     "https://appwrite.io/discord",
			GetInspired: "https://builtwith.appwrite.io",
		})
}
