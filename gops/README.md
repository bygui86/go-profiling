
# gops

1. Run application

    ```bash
    go run main.go
    ```

1. Use Vegeta to generate some traffic

    ```bash
    echo "GET http://localhost:8080" | vegeta attack -duration=30s
    ```

1. Use gops tool to visualise profiling

    ```bash
    go get -u github.com/google/gops
    
    gops <APP_PID>
    
    gops tree
    
    gops pprof-cpu <APP_PID>
    ```
