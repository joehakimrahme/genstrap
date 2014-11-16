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
	"fmt"
	"os"
)

type FileReader func(filename string) ([]byte, error)
type FileWriter func(filename string, text []byte, mode os.FileMode) error

type GenStrapper struct {
	html_filename        string
	html_title           string
	html_bootstrap_theme string

	reader FileReader
	writer FileWriter
}

func NewGenStrapper(filename string, title string, theme string, filereader FileReader, filewriter FileWriter) *GenStrapper {
	return &GenStrapper{
		html_filename:        filename,
		html_title:           title,
		html_bootstrap_theme: theme,
		reader:               filereader,
		writer:               filewriter,
	}
}

func (gs *GenStrapper) ReadFile() ([]byte, error) {
	return gs.reader(gs.html_filename)
}

func (gs *GenStrapper) WriteFile() error {
	return gs.writer(gs.html_filename+".html", []byte(gs.GetStrapFile()), 0644)
}

func (gs *GenStrapper) GetHeader() string {
	return fmt.Sprintf("<!DOCTYPE html><html><title>%s</title><xmp theme=\"%s\" style=\"display:none;\">", gs.html_title, gs.html_bootstrap_theme)
}

func (gs *GenStrapper) GetFooter() string {
	return "</xmp><script src=\"http://strapdownjs.com/v/0.2/strapdown.js\"></script></html>"
}

func (gs *GenStrapper) GetStrapFile() string {
	txt, err := gs.ReadFile()

	if err != nil {
		panic(err)
	}

	return gs.GetHeader() + string(txt) + gs.GetFooter()
}
