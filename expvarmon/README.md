
# ExpvarMon

1. Run application

    ```bash
    go run main.go
    ```

1. Visit `localhost:8080/debug/vars`

1. Run expvarmon

   ```bash
   go get github.com/divan/expvarmon
   
   expvarmon -ports="8080"
   ```
