package main

import (
    "fmt"
    "net"
    "os"
    "strconv"
    "sync"
    "time"
)

func scanner(buffer chan int, host string, openPorts chan int) {
    // ❺
    for p := range buffer {
        conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, p), time.Duration(1)*time.Second)
        if err != nil {
            continue
        }
        conn.Close()
        openPorts <- p
    }
}

func main() {
    var results []int

    host := os.Args[1]
    maxPort, _ := strconv.Atoi(os.Args[2])
    worker, _ := strconv.Atoi(os.Args[3])

    buffer := make(chan int, worker) // ❶
    openPorts := make(chan int)      // ❷

    var wg sync.WaitGroup
    wg.Add(worker)

    for i := 0; i < worker; i++ {
        go func() {
            scanner(buffer, host, openPorts) // ❸
            wg.Done()
        }()
    }

    go func() {
        wg.Wait()
        close(openPorts) // ❻
    }()

    start := time.Now()
    go func() {
        for p := 0; p <= maxPort; p++ {
            buffer <- p // ❹
        }
        close(buffer)
    }()

    for p := range openPorts {
        results = append(results, p)
    }

    end := time.Now()

    for _, p := range results {
        fmt.Printf("Opening %d port\n", p)
    }

    fmt.Printf("Done!! Took %f seconds\n", (end.Sub(start)).Seconds())
}
