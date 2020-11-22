package main

import (
	"context"
	"fmt"

	_ "cloud.google.com/go/language/apiv1"
	"cloud.google.com/go/spanner"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	// _ "google.golang.org/genproto/protobuf/field_mask"
	// _ "google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/rytswd/simple-bazel/timesvc"
)

func main() {
	ctx := context.Background()
	_ = spanner.Statement{}
	_ = timesvc.GetNowResponse{}
	_ = ctx
	fmt.Println("hello")
}
