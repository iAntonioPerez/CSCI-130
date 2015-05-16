//Antonio Perez
package main

import (
	"fmt"
)

func makeOddGenerator() func() uint {
	i := unit(1)
	return func() (ret unit) {
		ret = i
		i += 2
		return
	}
}
