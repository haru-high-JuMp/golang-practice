package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	assertCorrentMessae := func(t testing.TB, got, want string){
		t.Helper()
		if got != want {
			t.Errorf("expected %q but got %q", got, want)
		}
	}

	t.Run("Count 5", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrentMessae(t, repeated, expected)
	})

	t.Run("Count 9", func(t *testing.T) {
		repeated := Repeat("a", 9)
		expected := "aaaaaaaaa"
		assertCorrentMessae(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat(){
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}