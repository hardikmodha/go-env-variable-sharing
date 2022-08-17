# go-env-variable-sharing
Go environment variable sharing between process and sub-process


To test the behavior, run the `./run_env_sharing_test.sh`. It will internally do the following...

1. Run the interceptor/main.go that will periodically read the environment variables to check if variables set via `os.SetEnv` from other go process are accessible to the different process or not.
2. Run the parent/main.go that internally spawns the child process. Child process will read the environment variables to check if it can read the variables set by the parent process. 


## Observation

Once the script exits, it will print the following output to the terminal. 

```
[3] 66021
Child gave some output: All args: [./child/child arg1 arg2]
All Environment Variables: [__parent_child_secret__ 67890]
Child found the parent-child secret environment variable
Child found the parent-child secret environment variable
All Environment Variables: [__parent_child_secret__ 67890]
Child found the parent-child secret environment variable
Child found the parent-child secret environment variable
All Environment Variables: [__parent_child_secret__ 67890]
Child found the parent-child secret environment variable
Child found the parent-child secret environment variable
All Environment Variables: [__parent_child_secret__ 67890]
Child found the parent-child secret environment variable
Child found the parent-child secret environment variable
All Environment Variables: [__parent_child_secret__ 67890]
Child found the parent-child secret environment variable
Child found the parent-child secret environment variable
Child is exiting now
[3]  + 66021 terminated  go run interceptor/main.go > interceptor_output.txt 2>&1
```

It will also log all the interceptor output in the `interceptor_output.txt` file. It will have the following content:

```
Checking if I can steal the secret...
Checking if I can steal the secret...
Checking if I can steal the secret...
Checking if I can steal the secret...
Checking if I can steal the secret...
Checking if I can steal the secret...
.
.
.
```

With that, we can observe the following points: 

1. Environment variables set via `os.SetEnv` in the parent process are not directly available to the child process.
2. Any environment variables are set via `cmd.Env` before running the child process are available to the child process.
3. Any environment variable set inside process boundary using `os.SetEnv` are not available to any other process running on the same machine.
