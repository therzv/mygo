package floor

import (
	"fmt"
)

func ShowAdultCompany(s string, n int) {
	for i := 0; i < n; i++ {
		n += i
	}
	fmt.Println(s)
}
