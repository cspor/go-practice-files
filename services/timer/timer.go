package timer

import (
	"fmt"
	"time"
)

func Took(message string, timer time.Time) {
	fmt.Printf(message+" took: %s \n", time.Since(timer))
}
