// The MIT License (MIT)
//
// Copyright (c) 2016 aerth aerth@riseup.net
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// command webhookd
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

var (
	hookdir    = "./hooks"
	addr       = "127.0.0.1:8008"
	versionstr = "webhookd v0.0.1"
	version    string
)

func main() {
	hooks, err := ioutil.ReadDir(hookdir)
	if err != nil {
		errorf("Fatal: %s\n", err)
		os.Exit(1)
	}

	if hooks == nil || len(hooks) == 0 {
		errorf("Fatal: no hooks in %q directory\n", hookdir)
		os.Exit(1)
	}

	go fmt.Println("Listening on:", addr)
	err = http.ListenAndServe(addr, http.HandlerFunc(hookd))
	if err != nil {
		errorf("Fatal: %s\n", err)
		os.Exit(1)
	}
}

func hookd(w http.ResponseWriter, r *http.Request) {
	s := strings.Split(r.URL.Path, "/")
	last := s[len(s)-1]
	if _, err := ioutil.ReadFile("hooks/" + last); err != nil {
		fmt.Println(time.Now(), r.RemoteAddr, err)
		http.NotFound(w, r)
		return
	}
	fmt.Println(r.URL.Path)
	w.Write([]byte(versionstr))
	cmd := exec.Command("sh", "hooks/"+last)
	b, err := cmd.CombinedOutput()
  if err != nil {
    fmt.Println("err")
  }
  fmt.Println(string(b))
}

func errorf(f string, i ...interface{}) {
	fmt.Fprintf(os.Stderr, f, i...)
}
