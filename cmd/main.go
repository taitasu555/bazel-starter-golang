package main

import (
	"fmt"

	"github.com/google/uuid"
	stringutil "github.com/taitasu555/bazel-starter-golang/pkg/string"
)

func main() {
	uuidStr := uuid.NewString()
	reverceStr := stringutil.Reverse(uuidStr)
	fmt.Printf("uuid: %s\nreverce: %s\n", uuidStr, reverceStr)

}
