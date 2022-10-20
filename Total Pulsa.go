package main

import "fmt"

func main () {
    var n, nominalPulsa, total int
    
    fmt.Scan(&n)
    
    total = 0
    
    for i:= 1; i<=n; i++ {
        fmt.Scan(&nominalPulsa)
        total += nominalPulsa
    }
    fmt.Println(total)
}
