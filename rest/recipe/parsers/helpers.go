// Copyright 2013 Joshua Marsh. All rights reserved.  Use of this
// source code is governed by a BSD-style license that can be found in
// the LICENSE file.

package parsers

import (
	"html"
	"regexp"
	"strings"
)

// CleanField is a helper function that cleans up a string from an
// HTML source. It is trimmed and unescaped.
func cleanField(s string) string {
	return strings.Trim(html.UnescapeString(s), " ")
}

// getFirstH1 returns the contents of the first <h1> element on the
// page.
func getFirstH1(data []byte) string {
	re := regexp.MustCompile(`(?sU)<h1.*>(.*)</h1>`)

	// Return the submatch or "".
	found := re.FindSubmatch(data)
	if len(found) > 1 {
		return cleanField(string(found[1]))
	}

	return ""
}
