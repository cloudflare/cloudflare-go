package apijson

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func makeIPSlice(n int) []string {
	out := make([]string, n)
	for i := 0; i < n; i++ {
		out[i] = fmt.Sprintf("10.%d.%d.%d", (i>>16)&0xff, (i>>8)&0xff, i&0xff)
	}
	return out
}

func TestArrayMarshalCorrectness(t *testing.T) {
	for _, n := range []int{0, 1, 2, 100} {
		in := makeIPSlice(n)
		got, err := Marshal(in)
		if err != nil {
			t.Fatalf("Marshal(n=%d) error: %v", n, err)
		}
		want, _ := json.Marshal(in)
		if !strings.EqualFold(string(got), string(want)) {
			t.Fatalf("n=%d: got=%s want=%s", n, got, want)
		}
	}
}

func benchmarkArrayMarshal(b *testing.B, n int) {
	in := makeIPSlice(n)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := Marshal(in)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkArrayMarshal_100(b *testing.B)   { benchmarkArrayMarshal(b, 100) }
func BenchmarkArrayMarshal_1000(b *testing.B)  { benchmarkArrayMarshal(b, 1000) }
func BenchmarkArrayMarshal_5000(b *testing.B)  { benchmarkArrayMarshal(b, 5000) }
func BenchmarkArrayMarshal_10000(b *testing.B) { benchmarkArrayMarshal(b, 10000) }
func BenchmarkArrayMarshal_30000(b *testing.B) { benchmarkArrayMarshal(b, 30000) }
