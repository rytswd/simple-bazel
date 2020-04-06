package main

import (
	"context"
	"fmt"

	_ "cloud.google.com/go/language/apiv1"
	"cloud.google.com/go/spanner"

	"github.com/rytswd/simple-bazel/timesvc"
)

func main() {
	ctx := context.Background()
	_ = spanner.Statement{}
	_ = timesvc.GetNowResponse{}
	_ = ctx
	fmt.Println("hello")
}
