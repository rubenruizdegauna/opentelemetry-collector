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
	"testing"

	"github.com/stretchr/testify/assert"

	otlpmetrics "go.opentelemetry.io/collector/pdata/internal/data/protogen/metrics/v1"
)

func TestHistogramDataPointSlice(t *testing.T) {
	es := NewHistogramDataPointSlice()
	assert.Equal(t, 0, es.Len())
	es = newHistogramDataPointSlice(&[]*otlpmetrics.HistogramDataPoint{})
	assert.Equal(t, 0, es.Len())

	es.EnsureCapacity(7)
	emptyVal := newHistogramDataPoint(&otlpmetrics.HistogramDataPoint{})
	testVal := generateTestHistogramDataPoint()
	assert.Equal(t, 7, cap(*es.orig))
	for i := 0; i < es.Len(); i++ {
		el := es.AppendEmpty()
		assert.Equal(t, emptyVal, el)
		fillTestHistogramDataPoint(el)
		assert.Equal(t, testVal, el)
	}
}

func TestHistogramDataPointSlice_CopyTo(t *testing.T) {
	dest := NewHistogramDataPointSlice()
	// Test CopyTo to empty
	NewHistogramDataPointSlice().CopyTo(dest)
	assert.Equal(t, NewHistogramDataPointSlice(), dest)

	// Test CopyTo larger slice
	generateTestHistogramDataPointSlice().CopyTo(dest)
	assert.Equal(t, generateTestHistogramDataPointSlice(), dest)

	// Test CopyTo same size slice
	generateTestHistogramDataPointSlice().CopyTo(dest)
	assert.Equal(t, generateTestHistogramDataPointSlice(), dest)
}

func TestHistogramDataPointSlice_EnsureCapacity(t *testing.T) {
	es := generateTestHistogramDataPointSlice()
	// Test ensure smaller capacity.
	const ensureSmallLen = 4
	expectedEs := make(map[*otlpmetrics.HistogramDataPoint]bool)
	for i := 0; i < es.Len(); i++ {
		expectedEs[es.At(i).orig] = true
	}
	assert.Equal(t, es.Len(), len(expectedEs))
	es.EnsureCapacity(ensureSmallLen)
	assert.Less(t, ensureSmallLen, es.Len())
	foundEs := make(map[*otlpmetrics.HistogramDataPoint]bool, es.Len())
	for i := 0; i < es.Len(); i++ {
		foundEs[es.At(i).orig] = true
	}
	assert.Equal(t, expectedEs, foundEs)

	// Test ensure larger capacity
	const ensureLargeLen = 9
	oldLen := es.Len()
	expectedEs = make(map[*otlpmetrics.HistogramDataPoint]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		expectedEs[es.At(i).orig] = true
	}
	assert.Equal(t, oldLen, len(expectedEs))
	es.EnsureCapacity(ensureLargeLen)
	assert.Equal(t, ensureLargeLen, cap(*es.orig))
	foundEs = make(map[*otlpmetrics.HistogramDataPoint]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		foundEs[es.At(i).orig] = true
	}
	assert.Equal(t, expectedEs, foundEs)
}

func TestHistogramDataPointSlice_MoveAndAppendTo(t *testing.T) {
	// Test MoveAndAppendTo to empty
	expectedSlice := generateTestHistogramDataPointSlice()
	dest := NewHistogramDataPointSlice()
	src := generateTestHistogramDataPointSlice()
	src.MoveAndAppendTo(dest)
	assert.Equal(t, generateTestHistogramDataPointSlice(), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo empty slice
	src.MoveAndAppendTo(dest)
	assert.Equal(t, generateTestHistogramDataPointSlice(), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo not empty slice
	generateTestHistogramDataPointSlice().MoveAndAppendTo(dest)
	assert.Equal(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.Equal(t, expectedSlice.At(i), dest.At(i))
		assert.Equal(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}
}

func TestHistogramDataPointSlice_RemoveIf(t *testing.T) {
	// Test RemoveIf on empty slice
	emptySlice := NewHistogramDataPointSlice()
	emptySlice.RemoveIf(func(el HistogramDataPoint) bool {
		t.Fail()
		return false
	})

	// Test RemoveIf
	filtered := generateTestHistogramDataPointSlice()
	pos := 0
	filtered.RemoveIf(func(el HistogramDataPoint) bool {
		pos++
		return pos%3 == 0
	})
	assert.Equal(t, 5, filtered.Len())
}

func generateTestHistogramDataPointSlice() HistogramDataPointSlice {
	tv := NewHistogramDataPointSlice()
	fillTestHistogramDataPointSlice(tv)
	return tv
}

func fillTestHistogramDataPointSlice(tv HistogramDataPointSlice) {
	*tv.orig = make([]*otlpmetrics.HistogramDataPoint, 7)
	for i := 0; i < 7; i++ {
		(*tv.orig)[i] = &otlpmetrics.HistogramDataPoint{}
		fillTestHistogramDataPoint(newHistogramDataPoint((*tv.orig)[i]))
	}
}
