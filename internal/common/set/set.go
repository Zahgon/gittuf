// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package set

import (
	"cmp"
)

// Set implements a generic set data structure for use in gittuf metadata and
// workflows.
type Set[T cmp.Ordered] struct {
	contents map[T]struct{}
}

// NewSet creates a new instance of a set for the specified type that fulfils
// the cmp.Ordered constraint.
func NewSet[T cmp.Ordered]() *Set[T] { _ = "STUB: not implemented"; return nil }

// NewSetFromItems creates a new instance of a set and populates it with the
// items provided.
func NewSetFromItems[T cmp.Ordered](items ...T) *Set[T] { _ = "STUB: not implemented"; return nil }

// MarshalJSON is used to serialize the instance of the set into JSON.
func (s *Set[T]) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON is used to load a set from the JSON representation.
func (s *Set[T]) UnmarshalJSON(jsonBytes []byte) error { _ = "STUB: not implemented"; return nil }

// Contents returns the objects present in the set.
func (s *Set[T]) Contents() []T { _ = "STUB: not implemented"; return nil }

// Add inserts an item into the set.
func (s *Set[T]) Add(item T) { _ = "STUB: not implemented"; return }

// Remove deletes the item from the set.
func (s *Set[T]) Remove(item T) { _ = "STUB: not implemented"; return }

// Extend adds all of the items in the passed set, resulting in a union
// operation.
func (s *Set[T]) Extend(set *Set[T]) { _ = "STUB: not implemented"; return }

// Has returns true if the set has the corresponding item.
func (s Set[T]) Has(item T) bool { _ = "STUB: not implemented"; return false }

// Len returns the number of objects in the set.
func (s *Set[T]) Len() int { _ = "STUB: not implemented"; return 0 }

// Intersection returns a new set consisting of the items present in both sets.
func (s *Set[T]) Intersection(set *Set[T]) *Set[T] { _ = "STUB: not implemented"; return nil }

// Minus returns a new set consisting of the items present only in the current
// set.
func (s *Set[T]) Minus(set *Set[T]) *Set[T] { _ = "STUB: not implemented"; return nil }

// Equal returns true if both sets have the same items.
func (s *Set[T]) Equal(set *Set[T]) bool { _ = "STUB: not implemented"; return false }
