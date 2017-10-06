package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var pend sync.WaitGroup
	for i := 0; i < 128; i++ {
		pend.Add(1)
		go func() {
			defer pend.Done()
			for {
				res, err := http.Get("http://t.sidekickopen65.com/e1t/o/5/f18dQhb0S7ks8dDMPbW2n0x6l2B9gXrN7sKj6v4LGDdVRs66P4XyXhWW5v0JCY3LvrVvW4w60PY1k1H6H0?si=6298133195390976&amp;pi=177d31d3-786a-43f2-95c1-5e33fb450333")
				if err == nil {
					res.Body.Close()
				} else {
					fmt.Println(err)
					return
				}
				fmt.Printf(".")
			}
		}()
	}
	pend.Wait()
}
