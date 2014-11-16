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
	"flag"
	"io/ioutil"
)

func main() {

	var filename string
	var title string
	var bootstrap_theme string

	flag.StringVar(&filename, "filename", "etc/sample.md", "The name of the input file.")
	flag.StringVar(&title, "title", "Generated with genstrap", "The HTML title of the generated page")
	flag.StringVar(&bootstrap_theme, "theme", "cerulean", "The Boostrap theme")

	flag.Parse()

	gs := NewGenStrapper(filename, title, bootstrap_theme, ioutil.ReadFile, ioutil.WriteFile)

	// write output file
	err := gs.WriteFile()
	if err != nil {
		panic(err)
	}
}
