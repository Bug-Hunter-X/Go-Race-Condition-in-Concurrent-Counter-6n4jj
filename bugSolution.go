```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var count int
        var mu sync.Mutex //Add Mutex for synchronization
        
        for i := 0; i < 10; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        mu.Lock() //Lock before accessing shared resource
                        count += i 
                        mu.Unlock() //Unlock after accessing shared resource
                }(i)
        }
        wg.Wait()
        fmt.Println("Count:", count)
}
```