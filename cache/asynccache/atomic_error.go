// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package asynccache

import "sync/atomic"

// Error is an atomic type-safe wrapper around Value for errors
type Error struct{ v atomic.Value }

// errorHolder is non-nil holder for error object.
// atomic.Value panics on saving nil object, so err object needs to be
// wrapped with valid object first.
type errorHolder struct{ err error }

// Load atomically loads the wrapped error
func (e *Error) Load() error {
	v := e.v.Load()
	if v == nil {
		return nil
	}

	eh := v.(errorHolder)
	return eh.err
}

// Store atomically stores error.
// NOTE: a holder object is allocated on each Store call.
func (e *Error) Store(err error) {
	e.v.Store(errorHolder{err: err})
}
