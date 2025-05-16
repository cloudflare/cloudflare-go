//go:build go1.23

package pagination

import "iter"

// Range return an iterator suitable for use with the Go 1.23+ range-functions
// feature. Iterating over an auto-pager then enables the simplified
// `for n := pager.Range() {}` syntax.
func (r *V4PagePaginationArrayAutoPager[T]) Range() iter.Seq[T] {
	return func(yield func(T) bool) {
		for r.Next() {
			if !yield(r.Current()) {
				return
			}
		}
	}
}
