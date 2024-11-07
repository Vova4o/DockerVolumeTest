package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
    // open a file if it exists or create a new one
    file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

    if err != nil {
        log.Fatal(err)
    }

    // time now
    time := time.Now()

    testString := fmt.Sprintf("Hello, World! %s\n", time)
    // write to the file
    _, err = file.WriteString(testString)
    

    // close the file
    defer file.Close()

}
