//Antonio Perez
package main

import (
	"fmt"
)

func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp

}
