
# multi-profile

1. Run application

    ```bash
    go run main.go
    ```

1. Use Vegeta to generate some traffic

    ```bash
    echo "GET http://localhost:8080" | vegeta attack -duration=30s
    ```

1. Use pprof tool to visualise profiling

   PDF
    ```bash
    go tool pprof --pdf cpu.prof > cpu.pdf
    go tool pprof --pdf mem.prof > mem.pdf
    go tool pprof --pdf goroutine.prof > goroutine.pdf
    ```

   browser
    ```bash
    go tool pprof --web cpu.prof
    go tool pprof --web mem.prof
    go tool pprof --web goroutine.prof
    ```
