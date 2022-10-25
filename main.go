package main

import (
    "bufio"
    "fmt"
    "net"
    "strings"
)

func handleConnection(c net.Conn) {
    scanner := bufio.NewScanner(c)
    
    for {
        var parameters[]string = nil
    
        for scanner.Scan() {
            line := scanner.Text()
            line = strings.TrimSpace(line)
        
            if len(line) == 0 {
                break
            } else {
//                fmt.Println(line)
                parameters = append(parameters, line)
            }  
        }

        if scanner.Err() != nil {
            fmt.Printf(" > Failed!: %v\n", scanner.Err())
            break
        }
            
        pa := preProcessParameters(parameters)
        result := processCommand(pa)
        
        c.Write([]byte(string(result)))
    
        if pa.keepAlive == false {
            break
        }
    }  

    c.Close()
}

func main() {
    config := get_config()
    
    l, err := net.Listen("tcp4", config.Server.Host + ":" + config.Server.Port)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer l.Close()

    fmt.Println("Executing server")
    for {
        c, err := l.Accept()
        if err != nil {
            fmt.Println(err)
            return
        }
        go handleConnection(c)
    }
}
