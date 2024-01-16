package main

import (
    "fmt"
    "sort"
    "sync"
)

func main() {
    var ids []int
    lock := sync.Mutex{}
    wg := sync.WaitGroup{}

    for i := 1; i <= 100; i++ {
        wg.Add(1)
        go func() {
            lock.Lock()
            ids = append(ids, i)
            lock.Unlock()
            wg.Done()
        }()
    }

    wg.Wait()
    sort.Ints(ids)
    fmt.Println(ids)
}
