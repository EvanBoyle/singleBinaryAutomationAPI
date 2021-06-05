package main

import (
	"context"
	"fmt"
	"os"

	"github.com/pulumi/pulumi/auto/v3/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {

	deployFunc := func(ctx *pulumi.Context) error {
		ctx.Export("numBinaries", pulumi.String("just one"))
		return nil
	}

	ctx := context.Background()

	projectName := "singleBinary"
	stackName := "demo"

	s, err := auto.UpsertStackInlineSource(ctx, stackName, projectName, deployFunc)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Created/Selected stack %q\n", stackName)
	fmt.Println("Starting update")

	// run the update
	_, err = s.Up(ctx)
	if err != nil {
		fmt.Printf("Failed to update stack: %v\n\n", err)
		os.Exit(1)
	}

	fmt.Println("Update succeeded!")
}
