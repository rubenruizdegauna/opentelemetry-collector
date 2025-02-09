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

package ptrace

import (
	"testing"

	"github.com/stretchr/testify/assert"

	otlptrace "go.opentelemetry.io/collector/pdata/internal/data/protogen/trace/v1"
)

func TestScopeSpansSlice(t *testing.T) {
	es := NewScopeSpansSlice()
	assert.Equal(t, 0, es.Len())
	es = newScopeSpansSlice(&[]*otlptrace.ScopeSpans{})
	assert.Equal(t, 0, es.Len())

	es.EnsureCapacity(7)
	emptyVal := newScopeSpans(&otlptrace.ScopeSpans{})
	testVal := generateTestScopeSpans()
	assert.Equal(t, 7, cap(*es.orig))
	for i := 0; i < es.Len(); i++ {
		el := es.AppendEmpty()
		assert.Equal(t, emptyVal, el)
		fillTestScopeSpans(el)
		assert.Equal(t, testVal, el)
	}
}

func TestScopeSpansSlice_CopyTo(t *testing.T) {
	dest := NewScopeSpansSlice()
	// Test CopyTo to empty
	NewScopeSpansSlice().CopyTo(dest)
	assert.Equal(t, NewScopeSpansSlice(), dest)

	// Test CopyTo larger slice
	generateTestScopeSpansSlice().CopyTo(dest)
	assert.Equal(t, generateTestScopeSpansSlice(), dest)

	// Test CopyTo same size slice
	generateTestScopeSpansSlice().CopyTo(dest)
	assert.Equal(t, generateTestScopeSpansSlice(), dest)
}

func TestScopeSpansSlice_EnsureCapacity(t *testing.T) {
	es := generateTestScopeSpansSlice()
	// Test ensure smaller capacity.
	const ensureSmallLen = 4
	expectedEs := make(map[*otlptrace.ScopeSpans]bool)
	for i := 0; i < es.Len(); i++ {
		expectedEs[es.At(i).orig] = true
	}
	assert.Equal(t, es.Len(), len(expectedEs))
	es.EnsureCapacity(ensureSmallLen)
	assert.Less(t, ensureSmallLen, es.Len())
	foundEs := make(map[*otlptrace.ScopeSpans]bool, es.Len())
	for i := 0; i < es.Len(); i++ {
		foundEs[es.At(i).orig] = true
	}
	assert.Equal(t, expectedEs, foundEs)

	// Test ensure larger capacity
	const ensureLargeLen = 9
	oldLen := es.Len()
	expectedEs = make(map[*otlptrace.ScopeSpans]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		expectedEs[es.At(i).orig] = true
	}
	assert.Equal(t, oldLen, len(expectedEs))
	es.EnsureCapacity(ensureLargeLen)
	assert.Equal(t, ensureLargeLen, cap(*es.orig))
	foundEs = make(map[*otlptrace.ScopeSpans]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		foundEs[es.At(i).orig] = true
	}
	assert.Equal(t, expectedEs, foundEs)
}

func TestScopeSpansSlice_MoveAndAppendTo(t *testing.T) {
	// Test MoveAndAppendTo to empty
	expectedSlice := generateTestScopeSpansSlice()
	dest := NewScopeSpansSlice()
	src := generateTestScopeSpansSlice()
	src.MoveAndAppendTo(dest)
	assert.Equal(t, generateTestScopeSpansSlice(), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo empty slice
	src.MoveAndAppendTo(dest)
	assert.Equal(t, generateTestScopeSpansSlice(), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo not empty slice
	generateTestScopeSpansSlice().MoveAndAppendTo(dest)
	assert.Equal(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.Equal(t, expectedSlice.At(i), dest.At(i))
		assert.Equal(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}
}

func TestScopeSpansSlice_RemoveIf(t *testing.T) {
	// Test RemoveIf on empty slice
	emptySlice := NewScopeSpansSlice()
	emptySlice.RemoveIf(func(el ScopeSpans) bool {
		t.Fail()
		return false
	})

	// Test RemoveIf
	filtered := generateTestScopeSpansSlice()
	pos := 0
	filtered.RemoveIf(func(el ScopeSpans) bool {
		pos++
		return pos%3 == 0
	})
	assert.Equal(t, 5, filtered.Len())
}

func generateTestScopeSpansSlice() ScopeSpansSlice {
	tv := NewScopeSpansSlice()
	fillTestScopeSpansSlice(tv)
	return tv
}

func fillTestScopeSpansSlice(tv ScopeSpansSlice) {
	*tv.orig = make([]*otlptrace.ScopeSpans, 7)
	for i := 0; i < 7; i++ {
		(*tv.orig)[i] = &otlptrace.ScopeSpans{}
		fillTestScopeSpans(newScopeSpans((*tv.orig)[i]))
	}
}
