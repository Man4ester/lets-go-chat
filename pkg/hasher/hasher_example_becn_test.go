package hasher

import (
	"testing"
	"fmt"
)

func BenchmarkHashPassword(b *testing.B) {
	for n := 0; n < b.N; n++ {
		HashPassword(password)
	}
}

func ExampleHashPassword() {
	res, _ := HashPassword(password)
	fmt.Println(res)
	// Output: 5f4dcc3b5aa765d61d8327deb882cf99
}
