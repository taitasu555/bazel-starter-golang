package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	stringutil "github.com/taitasu555/bazel-starter-golang/pkg/string"
)

func main() {
	uuidStr := uuid.NewString()
	reverceStr := stringutil.Reverse(uuidStr)
	fmt.Printf("uuid: %s\nreverce: %s\n", uuidStr, reverceStr)

	command := &cobra.Command{
		Use:   "server",
		Short: "Start the server",
		Run: func(cmd *cobra.Command, args []string) {
			//local server start
			h1 := func(w http.ResponseWriter, _ *http.Request) {
				io.WriteString(w, "Hello from a HandleFunc #1!\n")
			}
			h2 := func(w http.ResponseWriter, _ *http.Request) {
				io.WriteString(w, "Hello from a HandleFunc #2!\n")
			}

			http.HandleFunc("/", h1)
			http.HandleFunc("/endpoint", h2)

			log.Fatal(http.ListenAndServe(":8080", nil))
		},
	}

	command.Execute()

}
