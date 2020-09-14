package main

import (
    "fmt"
    "log"

    "github.com/beevik/ntp"
)

func main() {

    time, err := ntp.Time("ru.pool.ntp.org")
    if err != nil {
	log.Fatalf("Fatal with error message: %s", err.Error())
    }

    hour, min, sec := time.Clock()
    fmt.Printf("Current time: %d:%d:%d\n", hour, min, sec)
}