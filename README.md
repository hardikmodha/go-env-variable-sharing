# go-env-variable-sharing
Go environment variable sharing between process and sub-process


To test the behavior, run the `./run_env_sharing_test.sh`. It will internally do the following...

1. Run the interceptor/main.go that will periodically read the environment variables to check if variables set via `os.SetEnv` from other go process are accessible to the different process or not.
2. Run the parent/main.go that internally spawns the child process. Child process will read the environment variables to check if it can read the variables set by the parent process. 
