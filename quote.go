// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package quote collects pithy sayings.
package quote // import "rsc.io/quote"

import "rsc.io/quote/v2"

// Hello returns a greeting.
func Hello() string {
	return quote.HelloV2()
}

// Glass returns a useful phrase for world travelers.
func Glass() string {
	// See http://www.oocities.org/nodotus/hbglass.html.
	return quote.GlassV2()
}

// Go returns a Go proverb.
func Go() string {
	return quote.GoV2()
}

// Opt returns an optimization truth.
func Opt() string {
	// Wisdom from ken.
	return quote.OptV2()
}
