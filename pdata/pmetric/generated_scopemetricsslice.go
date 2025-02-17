// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pmetric

import (
	"sort"

	otlpmetrics "go.opentelemetry.io/collector/pdata/internal/data/protogen/metrics/v1"
)

// ScopeMetricsSlice logically represents a slice of ScopeMetrics.
//
// This is a reference type. If passed by value and callee modifies it, the
// caller will see the modification.
//
// Must use NewScopeMetricsSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ScopeMetricsSlice struct {
	orig *[]*otlpmetrics.ScopeMetrics
}

func newScopeMetricsSlice(orig *[]*otlpmetrics.ScopeMetrics) ScopeMetricsSlice {
	return ScopeMetricsSlice{orig}
}

// NewScopeMetricsSlice creates a ScopeMetricsSlice with 0 elements.
// Can use "EnsureCapacity" to initialize with a given capacity.
func NewScopeMetricsSlice() ScopeMetricsSlice {
	orig := []*otlpmetrics.ScopeMetrics(nil)
	return newScopeMetricsSlice(&orig)
}

// Len returns the number of elements in the slice.
//
// Returns "0" for a newly instance created with "NewScopeMetricsSlice()".
func (es ScopeMetricsSlice) Len() int {
	return len(*es.orig)
}

// At returns the element at the given index.
//
// This function is used mostly for iterating over all the values in the slice:
//
//	for i := 0; i < es.Len(); i++ {
//	    e := es.At(i)
//	    ... // Do something with the element
//	}
func (es ScopeMetricsSlice) At(ix int) ScopeMetrics {
	return newScopeMetrics((*es.orig)[ix])
}

// CopyTo copies all elements from the current slice overriding the destination.
func (es ScopeMetricsSlice) CopyTo(dest ScopeMetricsSlice) {
	srcLen := es.Len()
	destCap := cap(*dest.orig)
	if srcLen <= destCap {
		(*dest.orig) = (*dest.orig)[:srcLen:destCap]
		for i := range *es.orig {
			newScopeMetrics((*es.orig)[i]).CopyTo(newScopeMetrics((*dest.orig)[i]))
		}
		return
	}
	origs := make([]otlpmetrics.ScopeMetrics, srcLen)
	wrappers := make([]*otlpmetrics.ScopeMetrics, srcLen)
	for i := range *es.orig {
		wrappers[i] = &origs[i]
		newScopeMetrics((*es.orig)[i]).CopyTo(newScopeMetrics(wrappers[i]))
	}
	*dest.orig = wrappers
}

// EnsureCapacity is an operation that ensures the slice has at least the specified capacity.
// 1. If the newCap <= cap then no change in capacity.
// 2. If the newCap > cap then the slice capacity will be expanded to equal newCap.
//
// Here is how a new ScopeMetricsSlice can be initialized:
//
//	es := NewScopeMetricsSlice()
//	es.EnsureCapacity(4)
//	for i := 0; i < 4; i++ {
//	    e := es.AppendEmpty()
//	    // Here should set all the values for e.
//	}
func (es ScopeMetricsSlice) EnsureCapacity(newCap int) {
	oldCap := cap(*es.orig)
	if newCap <= oldCap {
		return
	}

	newOrig := make([]*otlpmetrics.ScopeMetrics, len(*es.orig), newCap)
	copy(newOrig, *es.orig)
	*es.orig = newOrig
}

// AppendEmpty will append to the end of the slice an empty ScopeMetrics.
// It returns the newly added ScopeMetrics.
func (es ScopeMetricsSlice) AppendEmpty() ScopeMetrics {
	*es.orig = append(*es.orig, &otlpmetrics.ScopeMetrics{})
	return es.At(es.Len() - 1)
}

// Sort sorts the ScopeMetrics elements within ScopeMetricsSlice given the
// provided less function so that two instances of ScopeMetricsSlice
// can be compared.
func (es ScopeMetricsSlice) Sort(less func(a, b ScopeMetrics) bool) {
	sort.SliceStable(*es.orig, func(i, j int) bool { return less(es.At(i), es.At(j)) })
}

// MoveAndAppendTo moves all elements from the current slice and appends them to the dest.
// The current slice will be cleared.
func (es ScopeMetricsSlice) MoveAndAppendTo(dest ScopeMetricsSlice) {
	if *dest.orig == nil {
		// We can simply move the entire vector and avoid any allocations.
		*dest.orig = *es.orig
	} else {
		*dest.orig = append(*dest.orig, *es.orig...)
	}
	*es.orig = nil
}

// RemoveIf calls f sequentially for each element present in the slice.
// If f returns true, the element is removed from the slice.
func (es ScopeMetricsSlice) RemoveIf(f func(ScopeMetrics) bool) {
	newLen := 0
	for i := 0; i < len(*es.orig); i++ {
		if f(es.At(i)) {
			continue
		}
		if newLen == i {
			// Nothing to move, element is at the right place.
			newLen++
			continue
		}
		(*es.orig)[newLen] = (*es.orig)[i]
		newLen++
	}
	// TODO: Prevent memory leak by erasing truncated values.
	*es.orig = (*es.orig)[:newLen]
}
