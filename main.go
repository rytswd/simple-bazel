package main

import (
	"context"
	"fmt"

	_ "cloud.google.com/go/language/apiv1"
	"cloud.google.com/go/spanner"
)

func main() {
	ctx := context.Background()
	_ = spanner.Statement{}
	_ = ctx
	fmt.Println("hello")
}
