package main

import (
	"fmt"
	"os"
	"time"

	ld "gopkg.in/launchdarkly/go-client.v4"
)

const (
	ENV_KEY  = "LD_SDK_KEY"
	ENV_USER = "LD_USER"
)

func main() {
	/* from env variable LD_SDK_KEY */
	sdk_key := os.Getenv(ENV_KEY)
	if sdk_key == "" {
		panic("Environment variable '" + ENV_KEY + "' required for sample.")
	}
	user := os.Getenv(ENV_USER)
	if user == "" {
		panic("Environment variable '" + ENV_USER + "' required for sample.")
	}

	/* The LD doc doesn't have any reference to the second return, error */
	ld_client, _ := ld.MakeClient(sdk_key, 5*time.Second)

	show_feature, _ := ld_client.BoolVariation("test_in_go", ld.NewUser(user), false)
	if show_feature {
		fmt.Printf("The feature is enabled!\n")
	} else {
		fmt.Printf("The feature is disabled\n")
	}

	ld_client.Close()
}
