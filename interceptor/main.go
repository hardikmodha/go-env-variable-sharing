package main

import (
	"envtest/params"
	"fmt"
	"os"
	"time"
)

func main() {
	for {
		fmt.Println("Checking if I can steal the secret...")

		for _, item := range os.Environ() {
			if item == params.ParentChildSecretEnvKey || item == params.ParentChildSecretEnvValue {
				fmt.Println("Alert! Interceptor found the environment parent child secret")
			}

			if item == params.ParentSecretEnvKey || item == params.ParentSecretEnvValue {
				fmt.Println("Alert! Interceptor found the environment parent secret")
			}
		}

		os.Setenv(params.InterceptorEnvKey, params.InterceptorEnvValue)

		time.Sleep(2 * time.Second)
	}
}
