package main

import (
    "fmt"
    "net"
    "os"
    "strconv"
    "time"
)

func main() {
    var openPorts []int
    host := os.Args[1]
    ports, _ := strconv.Atoi(os.Args[2])

    start := time.Now()
    for i := 0; i <= ports; i++ {
        address := fmt.Sprintf("%s:%d", host, i)
        conn, err := net.DialTimeout("tcp", address, time.Duration(1)*time.Second)
        if err != nil {
            continue
        }
        conn.Close()
        openPorts = append(openPorts, i)
    }
    end := time.Now()

    for _, p := range openPorts {
        fmt.Printf("Opening %d port\n", p)
    }
    fmt.Printf("Done!! Took %f seconds\n", (end.Sub(start)).Seconds())
}
