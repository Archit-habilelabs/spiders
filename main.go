package main

import (
    "encoding/json"
    "fmt"
    "os"

    "rsc.io/quote"
)

func main() {
    text := quote.Go()
    fmt.Printf("[dep-validation][go] rsc.io/quote produced %q\n", text)
    msg := map[string]interface{}{
        "ipc":  true,
        "type": "data",
        "payload": map[string]interface{}{
            "spider":     "go",
            "dependency": "rsc.io/quote",
            "quote":      text,
            "task_id":    os.Getenv("SE_TASK_ID"),
        },
    }
    b, _ := json.Marshal(msg)
    fmt.Println(string(b))
}