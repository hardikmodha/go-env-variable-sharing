package main

import (
	"envtest/params"
	"fmt"
	"os"
	"time"
)

func main() {
	vars := os.Environ()
	fmt.Printf("All args: %+v\n", os.Args)

	for i := 0; i < 5; i++ {
		time.Sleep(5 * time.Second)

		fmt.Printf("All Environment Variables: %+v\n", vars)

		for _, item := range vars {
			if item == params.ParentChildSecretEnvKey || item == params.ParentChildSecretEnvValue {
				fmt.Printf("Child found the parent-child secret environment variable\n")
			}

			if item == params.ParentSecretEnvKey || item == params.ParentSecretEnvValue {
				fmt.Printf("Child found the parents' secret environment variable\n")
			}

			if item == params.InterceptorEnvKey || item == params.InterceptorEnvValue {
				fmt.Printf("Child found the interceptor environment variable\n")
			}
		}
	}

	fmt.Println("Child is exiting now")
}
