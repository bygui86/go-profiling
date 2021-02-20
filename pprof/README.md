
# pprof

1. Run application

    ```bash
    go run main.go
    ```

1. Use Vegeta to generate some traffic

    ```bash
    echo "GET http://localhost:8080" | vegeta attack -duration=30s
    ```

1. Use pprof tool to visualise profiling

    ```bash
    go tool pprof cpu.prof
    ```

1. Get same info running

    ```bash
    go tool pprof -top http://localhost:8080/debug/pprof/heap
    ```
