// Package slicediff is a utility to determine the additions and deletions that happened to a sorted slice after each update.
//
// All the slices are assumed to be sorted!
package slicediff

import "container/list"

// SliceDiff stores the current state of the sorted slice
type SliceDiff struct {
	s *list.List // s for slice
}

// New creates a new SliceDiff
func New() *SliceDiff {
	return &SliceDiff{
		list.New(),
	}
}

// Append appends a slice at the end of the SliceDiff
//
// sa is assumed to be sorted
func (sd *SliceDiff) Append(sa []string) {
	for _, s := range sa {
		sd.s.PushBack(s)
	}
}

// SortedDiff compares the updated slice with l and returns the additions and deletions performed
//
// updated is assumed to be sorted
func (sd *SliceDiff) SortedDiff(updated []string) (additions, deletions []string) {
	e := sd.s.Front()

	additions = make([]string, 0, len(updated))
	deletions = make([]string, 0, sd.s.Len())

	for _, s := range updated {
		// Delete the small elements at the beginning of the list
		for e != nil && e.Value.(string) < s {
			deletions = append(deletions, e.Value.(string))
			e = sd.removeAndGetNext(e)
		}

		if e == nil {
			// End of list is empty: simply push it
			additions = append(additions, s)
			sd.s.PushBack(s)
		} else if s == e.Value.(string) {
			// Same as current element: skip it
			e = e.Next()
		} else {
			// Smaller than current element: insert it
			additions = append(additions, s)
			sd.s.InsertBefore(s, e)
		}
	}

	// delete end of the list
	for e != nil {
		deletions = append(deletions, e.Value.(string))
		e = sd.removeAndGetNext(e)
	}

	return additions, deletions
}

func (sd *SliceDiff) removeAndGetNext(e *list.Element) (next *list.Element) {
	next = e.Next()
	sd.s.Remove(e)
	return next
}
