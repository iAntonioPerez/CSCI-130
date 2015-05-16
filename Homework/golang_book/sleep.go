//Antonio Perez
package main

import (
	"fmt"
	"time"
)

func sleep(t int) {
	for i := 0; i < n; i++ {
		<-time.After(time.Second)
	}
}
