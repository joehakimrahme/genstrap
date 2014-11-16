// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.
package main

import (
	"os"
	"testing"
)

// not sure these functions should be mocked or not.
// I tend to think that they should be, but I might be wrong
func MockReader(string) ([]byte, error) {
	return []byte("sample text"), nil
}
func MockWriter(string, []byte, os.FileMode) error {
	return nil
}

// These 2 tests don't make a lot of sense to me
func TestReader(t *testing.T) {
	gs := NewGenStrapper("sample.md", "title", "theme", MockReader, MockWriter)

	txt, err := gs.ReadFile()

	if string(txt) != "sample text" {
		t.Error("Didn't get expected content: %s", txt)
	}
	if err != nil {
		t.Error(err)
	}
}

func TestWriter(t *testing.T) {
	gs := NewGenStrapper("sample.md", "title", "theme", MockReader, MockWriter)
	err := gs.WriteFile()

	if err != nil {
		t.Error("err")
	}
}
func TestGetHeader(t *testing.T) {
	expected := "<!DOCTYPE html><html><title>title</title><xmp theme=\"theme\" style=\"display:none;\">"
	gs := NewGenStrapper("sample.md", "title", "theme", MockReader, MockWriter)
	header := gs.GetHeader()

	if header != expected {
		t.Error("Bad header: %s", header)
	}
}

func TestGetFooter(t *testing.T) {
	expected := "</xmp><script src=\"http://strapdownjs.com/v/0.2/strapdown.js\"></script></html>"
	gs := NewGenStrapper("sample.md", "title", "theme", MockReader, MockWriter)
	footer := gs.GetFooter()

	if footer != expected {
		t.Error("Bad footer: %s", footer)
	}
}

func testGetStrapfile(t *testing.T) {
	expected := "<!DOCTYPE html><html><title>title</title><xmp theme=\"theme\" style=\"display:none;\">sample text</xmp><script src=\"http://strapdownjs.com/v/0.2/strapdown.js\"></script></html>"
	gs := NewGenStrapper("sample.md", "title", "theme", MockReader, MockWriter)
	strapfile := gs.GetStrapFile()

	if strapfile != expected {
		t.Error("Bad strapfile: %s", strapfile)
	}
}
