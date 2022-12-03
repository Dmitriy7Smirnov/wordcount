package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    src := os.Args[1]    
    words := strings.Fields(src) 
    m := make(map[string]int) 
    for _, word := range words { 
        m[word] += 1 
    }
    fmt.Println(len(m))
}
