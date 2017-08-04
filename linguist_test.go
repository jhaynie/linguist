package linguist

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"sync"
	"testing"
)

func TestMain(m *testing.M) {
	Initialize()
	os.Exit(m.Run())
}

func TestLanguageOptimizationsJavaScript(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.js", []byte("a = 'bar'"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if len(r.Results) != 1 {
		t.Fatalf("expected len of results to be 1, was %d", len(r.Results))
	}
	l := r.Results[0]
	if l.Path != "foo.js" {
		t.Fatalf("expected Path to be foo.js, was %s", l.Path)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Loc to be 0, was %v", l.Loc)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Sloc to be 0, was %v", l.Loc)
	}
	if l.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", l.Type)
	}
	if l.Language.Name != "JavaScript" {
		t.Fatalf("expected Language.Name to be JavaScript, was %v", l.Language.Name)
	}
}

func TestLanguageOptimizationsGolang(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.go", []byte("package test\nvar a string\n"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if len(r.Results) != 1 {
		t.Fatalf("expected len of results to be 1, was %d", len(r.Results))
	}
	l := r.Results[0]
	if l.Path != "foo.go" {
		t.Fatalf("expected Path to be foo.go, was %s", l.Path)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Loc to be 0, was %v", l.Loc)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Sloc to be 0, was %v", l.Loc)
	}
	if l.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", l.Type)
	}
	if l.Language.Name != "Go" {
		t.Fatalf("expected Language.Name to be Go, was %v", l.Language.Name)
	}
}

func TestLanguageOptimizationsSwift(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.swift", []byte("let a = 0"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if len(r.Results) != 1 {
		t.Fatalf("expected len of results to be 1, was %d", len(r.Results))
	}
	l := r.Results[0]
	if l.Path != "foo.swift" {
		t.Fatalf("expected Path to be foo.swift, was %s", l.Path)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Loc to be 0, was %v", l.Loc)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Sloc to be 0, was %v", l.Loc)
	}
	if l.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", l.Type)
	}
	if l.Language.Name != "Swift" {
		t.Fatalf("expected Language.Name to be Swift, was %v", l.Language.Name)
	}
}

func TestLanguageOptimizationsMakefile(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "Makefile", []byte(".phony a\n"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if len(r.Results) != 1 {
		t.Fatalf("expected len of results to be 1, was %d", len(r.Results))
	}
	l := r.Results[0]
	if l.Path != "Makefile" {
		t.Fatalf("expected Path to be Makefile, was %s", l.Path)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Loc to be 0, was %v", l.Loc)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Sloc to be 0, was %v", l.Loc)
	}
	if l.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", l.Type)
	}
	if l.Language.Name != "Makefile" {
		t.Fatalf("expected Language.Name to be Makefile, was %v", l.Language.Name)
	}
}

func TestLanguageOptimizationsJSON(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.json", []byte("{\"a\":1}"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if len(r.Results) != 1 {
		t.Fatalf("expected len of results to be 1, was %d", len(r.Results))
	}
	l := r.Results[0]
	if l.Path != "foo.json" {
		t.Fatalf("expected Path to be foo.json, was %s", l.Path)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Loc to be 0, was %v", l.Loc)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Sloc to be 0, was %v", l.Loc)
	}
	if l.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", l.Type)
	}
	if l.Language.Name != "JSON" {
		t.Fatalf("expected Language.Name to be JSON, was %v", l.Language.Name)
	}
}

func TestLanguageOptimizationsYaml(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.yaml", []byte("---\ninvoice: 34843\n"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if len(r.Results) != 1 {
		t.Fatalf("expected len of results to be 1, was %d", len(r.Results))
	}
	l := r.Results[0]
	if l.Path != "foo.yaml" {
		t.Fatalf("expected Path to be foo.yaml, was %s", l.Path)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Loc to be 0, was %v", l.Loc)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Sloc to be 0, was %v", l.Loc)
	}
	if l.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", l.Type)
	}
	if l.Language.Name != "YAML" {
		t.Fatalf("expected Language.Name to be YAML, was %v", l.Language.Name)
	}
}

func TestLanguageOptimizationsYaml2(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.yml", []byte("---\ninvoice: 34843\n"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if len(r.Results) != 1 {
		t.Fatalf("expected len of results to be 1, was %d", len(r.Results))
	}
	l := r.Results[0]
	if l.Path != "foo.yml" {
		t.Fatalf("expected Path to be foo.yml, was %s", l.Path)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Loc to be 0, was %v", l.Loc)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Sloc to be 0, was %v", l.Loc)
	}
	if l.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", l.Type)
	}
	if l.Language.Name != "YAML" {
		t.Fatalf("expected Language.Name to be YAML, was %v", l.Language.Name)
	}
}

func TestLanguageOptimizationsEJS(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.ejs", []byte("<div><%= foo %></div>"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if len(r.Results) != 1 {
		t.Fatalf("expected len of results to be 1, was %d", len(r.Results))
	}
	l := r.Results[0]
	if l.Path != "foo.ejs" {
		t.Fatalf("expected Path to be foo.ejs, was %s", l.Path)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Loc to be 0, was %v", l.Loc)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Sloc to be 0, was %v", l.Loc)
	}
	if l.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", l.Type)
	}
	if l.Language.Name != "EJS" {
		t.Fatalf("expected Language.Name to be EJS, was %v", l.Language.Name)
	}
}

func TestLanguageOptimizationsHTML(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.html", []byte("<html></html>"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if len(r.Results) != 1 {
		t.Fatalf("expected len of results to be 1, was %d", len(r.Results))
	}
	l := r.Results[0]
	if l.Path != "foo.html" {
		t.Fatalf("expected Path to be foo.html, was %s", l.Path)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Loc to be 0, was %v", l.Loc)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Sloc to be 0, was %v", l.Loc)
	}
	if l.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", l.Type)
	}
	if l.Language.Name != "HTML" {
		t.Fatalf("expected Language.Name to be HTML, was %v", l.Language.Name)
	}
}

func TestLanguageOptimizationsCSS(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.css", []byte(".foo { color: red; }"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if len(r.Results) != 1 {
		t.Fatalf("expected len of results to be 1, was %d", len(r.Results))
	}
	l := r.Results[0]
	if l.Path != "foo.css" {
		t.Fatalf("expected Path to be foo.css, was %s", l.Path)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Loc to be 0, was %v", l.Loc)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Sloc to be 0, was %v", l.Loc)
	}
	if l.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", l.Type)
	}
	if l.Language.Name != "CSS" {
		t.Fatalf("expected Language.Name to be CSS, was %v", l.Language.Name)
	}
}

func TestLanguageOptimizationsSCSS(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.scss", []byte(".foo { color: red; }"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if len(r.Results) != 1 {
		t.Fatalf("expected len of results to be 1, was %d", len(r.Results))
	}
	l := r.Results[0]
	if l.Path != "foo.scss" {
		t.Fatalf("expected Path to be foo.scss, was %s", l.Path)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Loc to be 0, was %v", l.Loc)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Sloc to be 0, was %v", l.Loc)
	}
	if l.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", l.Type)
	}
	if l.Language.Name != "SCSS" {
		t.Fatalf("expected Language.Name to be SCSS, was %v", l.Language.Name)
	}
}

func TestLanguageOptimizationsMarkdown(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.md", []byte("# heading\n"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if len(r.Results) != 1 {
		t.Fatalf("expected len of results to be 1, was %d", len(r.Results))
	}
	l := r.Results[0]
	if l.Path != "foo.md" {
		t.Fatalf("expected Path to be foo.md, was %s", l.Path)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Loc to be 0, was %v", l.Loc)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Sloc to be 0, was %v", l.Loc)
	}
	if l.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", l.Type)
	}
	if l.Language.Name != "Markdown" {
		t.Fatalf("expected Language.Name to be Markdown, was %v", l.Language.Name)
	}
}

func TestLanguageOptimizationsShell(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.sh", []byte("#!/bin/sh\n"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if len(r.Results) != 1 {
		t.Fatalf("expected len of results to be 1, was %d", len(r.Results))
	}
	l := r.Results[0]
	if l.Path != "foo.sh" {
		t.Fatalf("expected Path to be foo.sh, was %s", l.Path)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Loc to be 0, was %v", l.Loc)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Sloc to be 0, was %v", l.Loc)
	}
	if l.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", l.Type)
	}
	if l.Language.Name != "Shell" {
		t.Fatalf("expected Language.Name to be Shell, was %v", l.Language.Name)
	}
}

func TestLanguageOptimizationsJSX(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.jsx", []byte("import { a } from './foo'\n"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if len(r.Results) != 1 {
		t.Fatalf("expected len of results to be 1, was %d", len(r.Results))
	}
	l := r.Results[0]
	if l.Path != "foo.jsx" {
		t.Fatalf("expected Path to be foo.jsx, was %s", l.Path)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Loc to be 0, was %v", l.Loc)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Sloc to be 0, was %v", l.Loc)
	}
	if l.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", l.Type)
	}
	if l.Language.Name != "JSX" {
		t.Fatalf("expected Language.Name to be JSX, was %v", l.Language.Name)
	}
}

func TestLanguageOptimizationsObjectiveC(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.m", []byte("@implementation Foo\n@end\n"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if len(r.Results) != 1 {
		t.Fatalf("expected len of results to be 1, was %d", len(r.Results))
	}
	l := r.Results[0]
	if l.Path != "foo.m" {
		t.Fatalf("expected Path to be foo.m, was %s", l.Path)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Loc to be 0, was %v", l.Loc)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Sloc to be 0, was %v", l.Loc)
	}
	if l.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", l.Type)
	}
	if l.Language.Name != "Objective-C" {
		t.Fatalf("expected Language.Name to be Objective-C, was %v", l.Language.Name)
	}
}

func TestLanguageOptimizationsObjectiveCPlusPlus(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.mm", []byte("@implementation Foo\n@end\n"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if len(r.Results) != 1 {
		t.Fatalf("expected len of results to be 1, was %d", len(r.Results))
	}
	l := r.Results[0]
	if l.Path != "foo.mm" {
		t.Fatalf("expected Path to be foo.mm, was %s", l.Path)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Loc to be 0, was %v", l.Loc)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Sloc to be 0, was %v", l.Loc)
	}
	if l.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", l.Type)
	}
	if l.Language.Name != "Objective-C++" {
		t.Fatalf("expected Language.Name to be Objective-C++, was %v", l.Language.Name)
	}
}

func TestLanguageOptimizationsC(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.c", []byte("void foo() {\n}\n"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if len(r.Results) != 1 {
		t.Fatalf("expected len of results to be 1, was %d", len(r.Results))
	}
	l := r.Results[0]
	if l.Path != "foo.c" {
		t.Fatalf("expected Path to be foo.c, was %s", l.Path)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Loc to be 0, was %v", l.Loc)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Sloc to be 0, was %v", l.Loc)
	}
	if l.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", l.Type)
	}
	if l.Language.Name != "C" {
		t.Fatalf("expected Language.Name to be C, was %v", l.Language.Name)
	}
}

func TestLanguageOptimizationsCHeader(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.h", []byte("#include <stdlib.h>\n"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if len(r.Results) != 1 {
		t.Fatalf("expected len of results to be 1, was %d", len(r.Results))
	}
	l := r.Results[0]
	if l.Path != "foo.h" {
		t.Fatalf("expected Path to be foo.h, was %s", l.Path)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Loc to be 0, was %v", l.Loc)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Sloc to be 0, was %v", l.Loc)
	}
	if l.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", l.Type)
	}
	if l.Language.Name != "C" {
		t.Fatalf("expected Language.Name to be C, was %v", l.Language.Name)
	}
}

func TestLanguageOptimizationsCPlusPlus(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.cpp", []byte("#include <stdlib>\nclass Foo {\n};\n"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if len(r.Results) != 1 {
		t.Fatalf("expected len of results to be 1, was %d", len(r.Results))
	}
	l := r.Results[0]
	if l.Path != "foo.cpp" {
		t.Fatalf("expected Path to be foo.cpp, was %s", l.Path)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Loc to be 0, was %v", l.Loc)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Sloc to be 0, was %v", l.Loc)
	}
	if l.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", l.Type)
	}
	if l.Language.Name != "C++" {
		t.Fatalf("expected Language.Name to be C++, was %v", l.Language.Name)
	}
}

func TestLanguageOptimizationsCPlusPlus2(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.cc", []byte("#include <stdlib>\nclass Foo {\n};\n"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if len(r.Results) != 1 {
		t.Fatalf("expected len of results to be 1, was %d", len(r.Results))
	}
	l := r.Results[0]
	if l.Path != "foo.cc" {
		t.Fatalf("expected Path to be foo.cc, was %s", l.Path)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Loc to be 0, was %v", l.Loc)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Sloc to be 0, was %v", l.Loc)
	}
	if l.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", l.Type)
	}
	if l.Language.Name != "C++" {
		t.Fatalf("expected Language.Name to be C++, was %v", l.Language.Name)
	}
}

func TestLanguageOptimizationsJSON5(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.json5", []byte("{foo:'bar'}"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if len(r.Results) != 1 {
		t.Fatalf("expected len of results to be 1, was %d", len(r.Results))
	}
	l := r.Results[0]
	if l.Path != "foo.json5" {
		t.Fatalf("expected Path to be foo.json5, was %s", l.Path)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Loc to be 0, was %v", l.Loc)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Sloc to be 0, was %v", l.Loc)
	}
	if l.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", l.Type)
	}
	if l.Language.Name != "JSON5" {
		t.Fatalf("expected Language.Name to be JSON5 was %v", l.Language.Name)
	}
}

func TestLanguageOptimizationsProtobuf(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.proto", []byte("package foo\nmessage A {\n}\n"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if len(r.Results) != 1 {
		t.Fatalf("expected len of results to be 1, was %d", len(r.Results))
	}
	l := r.Results[0]
	if l.Path != "foo.proto" {
		t.Fatalf("expected Path to be foo.proto, was %s", l.Path)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Loc to be 0, was %v", l.Loc)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Sloc to be 0, was %v", l.Loc)
	}
	if l.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", l.Type)
	}
	if l.Language.Name != "Protocol Buffer" {
		t.Fatalf("expected Language.Name to be Protocol Buffer was %v", l.Language.Name)
	}
}

func TestLanguageOptimizationsPython(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.py", []byte("def foo\nend\n"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if len(r.Results) != 1 {
		t.Fatalf("expected len of results to be 1, was %d", len(r.Results))
	}
	l := r.Results[0]
	if l.Path != "foo.py" {
		t.Fatalf("expected Path to be foo.py, was %s", l.Path)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Loc to be 0, was %v", l.Loc)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Sloc to be 0, was %v", l.Loc)
	}
	if l.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", l.Type)
	}
	if l.Language.Name != "Python" {
		t.Fatalf("expected Language.Name to be Python was %v", l.Language.Name)
	}
}

func TestLanguageOptimizationsRuby(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.rb", []byte("print \"hey\""))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if len(r.Results) != 1 {
		t.Fatalf("expected len of results to be 1, was %d", len(r.Results))
	}
	l := r.Results[0]
	if l.Path != "foo.rb" {
		t.Fatalf("expected Path to be foo.rb, was %s", l.Path)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Loc to be 0, was %v", l.Loc)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Sloc to be 0, was %v", l.Loc)
	}
	if l.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", l.Type)
	}
	if l.Language.Name != "Ruby" {
		t.Fatalf("expected Language.Name to be Ruby was %v", l.Language.Name)
	}
}

func TestLanguageOptimizationsJava(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.java", []byte("package foo\npublic class Foo\n{\n}\n"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if len(r.Results) != 1 {
		t.Fatalf("expected len of results to be 1, was %d", len(r.Results))
	}
	l := r.Results[0]
	if l.Path != "foo.java" {
		t.Fatalf("expected Path to be foo.java, was %s", l.Path)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Loc to be 0, was %v", l.Loc)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Sloc to be 0, was %v", l.Loc)
	}
	if l.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", l.Type)
	}
	if l.Language.Name != "Java" {
		t.Fatalf("expected Language.Name to be Java was %v", l.Language.Name)
	}
}

func TestLanguageOptimizationsCSharp(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.cs", []byte("public class Hello\n{\n}"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if len(r.Results) != 1 {
		t.Fatalf("expected len of results to be 1, was %d", len(r.Results))
	}
	l := r.Results[0]
	if l.Path != "foo.cs" {
		t.Fatalf("expected Path to be foo.cs, was %s", l.Path)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Loc to be 0, was %v", l.Loc)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Sloc to be 0, was %v", l.Loc)
	}
	if l.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", l.Type)
	}
	if l.Language.Name != "C#" {
		t.Fatalf("expected Language.Name to be C# was %v", l.Language.Name)
	}
}

func TestLanguageOptimizationsXML(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.xml", []byte("<?xml version=\"1.0\"?>\n<a></a>\n"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if len(r.Results) != 1 {
		t.Fatalf("expected len of results to be 1, was %d", len(r.Results))
	}
	l := r.Results[0]
	if l.Path != "foo.xml" {
		t.Fatalf("expected Path to be foo.xml, was %s", l.Path)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Loc to be 0, was %v", l.Loc)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Sloc to be 0, was %v", l.Loc)
	}
	if l.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", l.Type)
	}
	if l.Language.Name != "XML" {
		t.Fatalf("expected Language.Name to be XML was %v", l.Language.Name)
	}
}

func TestLanguageOptimizationsHandlebars(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.hbs", []byte("<div>{{hello}}</div>"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if len(r.Results) != 1 {
		t.Fatalf("expected len of results to be 1, was %d", len(r.Results))
	}
	l := r.Results[0]
	if l.Path != "foo.hbs" {
		t.Fatalf("expected Path to be foo.hbs, was %s", l.Path)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Loc to be 0, was %v", l.Loc)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Sloc to be 0, was %v", l.Loc)
	}
	if l.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", l.Type)
	}
	if l.Language.Name != "Handlebars" {
		t.Fatalf("expected Language.Name to be Handlebars was %v", l.Language.Name)
	}
}

func TestLanguageOptimizationsLua(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.lua", []byte("x = 0"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if len(r.Results) != 1 {
		t.Fatalf("expected len of results to be 1, was %d", len(r.Results))
	}
	l := r.Results[0]
	if l.Path != "foo.lua" {
		t.Fatalf("expected Path to be foo.lua, was %s", l.Path)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Loc to be 0, was %v", l.Loc)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Sloc to be 0, was %v", l.Loc)
	}
	if l.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", l.Type)
	}
	if l.Language.Name != "Lua" {
		t.Fatalf("expected Language.Name to be Lua was %v", l.Language.Name)
	}
}

func TestLanguageOptimizationsDockerfile(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "Dockerfile", []byte("FROM nodejs\n"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if len(r.Results) != 1 {
		t.Fatalf("expected len of results to be 1, was %d", len(r.Results))
	}
	l := r.Results[0]
	if l.Path != "Dockerfile" {
		t.Fatalf("expected Path to be Dockerfile, was %s", l.Path)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Loc to be 0, was %v", l.Loc)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Sloc to be 0, was %v", l.Loc)
	}
	if l.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", l.Type)
	}
	if l.Language.Name != "Dockerfile" {
		t.Fatalf("expected Language.Name to be Dockerfile was %v", l.Language.Name)
	}
}

func TestLanguageOptimizationsText(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.txt", []byte("yo\n"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if len(r.Results) != 1 {
		t.Fatalf("expected len of results to be 1, was %d", len(r.Results))
	}
	l := r.Results[0]
	if l.Path != "foo.txt" {
		t.Fatalf("expected Path to be foo.txt, was %s", l.Path)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Loc to be 0, was %v", l.Loc)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Sloc to be 0, was %v", l.Loc)
	}
	if l.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", l.Type)
	}
	if l.Language.Name != "Text" {
		t.Fatalf("expected Language.Name to be Text was %v", l.Language.Name)
	}
}

func TestLanguageOptimizationsSQL(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.sql", []byte("insert into foo values(1);"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if len(r.Results) != 1 {
		t.Fatalf("expected len of results to be 1, was %d", len(r.Results))
	}
	l := r.Results[0]
	if l.Path != "foo.sql" {
		t.Fatalf("expected Path to be foo.sql, was %s", l.Path)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Loc to be 0, was %v", l.Loc)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Sloc to be 0, was %v", l.Loc)
	}
	if l.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", l.Type)
	}
	if l.Language.Name != "SQL" {
		t.Fatalf("expected Language.Name to be SQL was %v", l.Language.Name)
	}
}

func TestLanguageOptimizationsTypescript(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.ts", []byte("interface Bar{}"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if len(r.Results) != 1 {
		t.Fatalf("expected len of results to be 1, was %d", len(r.Results))
	}
	l := r.Results[0]
	if l.Path != "foo.ts" {
		t.Fatalf("expected Path to be foo.ts, was %s", l.Path)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Loc to be 0, was %v", l.Loc)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Sloc to be 0, was %v", l.Loc)
	}
	if l.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", l.Type)
	}
	if l.Language.Name != "TypeScript" {
		t.Fatalf("expected Language.Name to be TypeScript was %v", l.Language.Name)
	}
}

func TestLanguageOptimizationsCoffeescript(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.coffee", []byte("foo = 1"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if len(r.Results) != 1 {
		t.Fatalf("expected len of results to be 1, was %d", len(r.Results))
	}
	l := r.Results[0]
	if l.Path != "foo.coffee" {
		t.Fatalf("expected Path to be foo.coffee, was %s", l.Path)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Loc to be 0, was %v", l.Loc)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Sloc to be 0, was %v", l.Loc)
	}
	if l.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", l.Type)
	}
	if l.Language.Name != "CoffeeScript" {
		t.Fatalf("expected Language.Name to be CoffeeScript was %v", l.Language.Name)
	}
}

func TestLanguageOptimizationsProperties(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.properties", []byte("foo=1"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if len(r.Results) != 1 {
		t.Fatalf("expected len of results to be 1, was %d", len(r.Results))
	}
	l := r.Results[0]
	if l.Path != "foo.properties" {
		t.Fatalf("expected Path to be foo.properties, was %s", l.Path)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Loc to be 0, was %v", l.Loc)
	}
	if l.Loc != 0 {
		t.Fatalf("expected Sloc to be 0, was %v", l.Loc)
	}
	if l.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", l.Type)
	}
	if l.Language.Name != "INI" {
		t.Fatalf("expected Language.Name to be INI was %v", l.Language.Name)
	}
}

func TestCacheStats(t *testing.T) {
	if CacheHits() != 35 {
		t.Fatalf("expected cache hits to be 34, was %d", CacheHits())
	}
	if CacheMisses() != 0 {
		t.Fatalf("expected cache misses to be 0, was %d", CacheMisses())
	}
	popular := MostPopular()
	if popular.Language.Name != "YAML" {
		t.Fatalf("expected popular.Language to be YAML, was %v", popular.Language.Name)
	}
	GetLanguageDetails(context.Background(), "foo.go", []byte("package test\nvar a string\n"))
	GetLanguageDetails(context.Background(), "foo.go", []byte("package test\nvar a string\n"))
	GetLanguageDetails(context.Background(), "foo.go", []byte("package test\nvar a string\n"))
	GetLanguageDetails(context.Background(), "foo.go", []byte("package test\nvar a string\n"))
	GetLanguageDetails(context.Background(), "foo.go", []byte("package test\nvar a string\n"))
	GetLanguageDetails(context.Background(), "foo.go", []byte("package test\nvar a string\n"))
	GetLanguageDetails(context.Background(), "foo.go", []byte("package test\nvar a string\n"))
	GetLanguageDetails(context.Background(), "foo.go", []byte("package test\nvar a string\n"))
	popular = MostPopular()
	if popular.Language.Name != "Go" {
		t.Fatalf("expected popular.Language to be Go, was %v", popular.Language.Name)
	}
}

func TestConcurrency(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			r, err := GetLanguageDetails(context.Background(), "foo.go", []byte("package test\nvar a string\n"))
			if err != nil {
				t.Fatal(err)
			}
			if !r.Success {
				t.Fatal("should have been successful")
			}
		}()
	}
	wg.Wait()
}

func TestIgnoreImage(t *testing.T) {
	f, err := os.Open("./testdata/image.png")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}
	r, err := GetLanguageDetails(context.Background(), "image.png", buf)
	if err != nil {
		t.Fatal(err)
	}
	if !r.Success {
		t.Fatal("expected success to be true")
	}
	if len(r.Results) > 0 {
		t.Fatal("expected results to be empty")
	}
}

func TestIgnoreLargeBuffer(t *testing.T) {
	var buf bytes.Buffer
	for i := 0; i < maxBufferSize+1; i++ {
		buf.Write([]byte("x"))
	}
	r, err := GetLanguageDetails(context.Background(), "foo.js", buf.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	if !r.Success {
		t.Fatal("expected success to be true")
	}
	if len(r.Results) > 0 {
		t.Fatal("expected results to be empty")
	}
}

func TestIgnoreExcludedExtension(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), ".npmrc", []byte("foo\n"))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Success {
		t.Fatal("expected success to be true")
	}
	if len(r.Results) > 0 {
		t.Fatal("expected results to be empty")
	}
}

func TestIgnoreExcludedFilename(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo/LICENSE", []byte("foo\n"))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Success {
		t.Fatal("expected success to be true")
	}
	if len(r.Results) > 0 {
		t.Fatal("expected results to be empty")
	}
}

func BenchmarkLinguist(b *testing.B) {
	buf := []byte("package test\nvar a string\n")
	ctx := context.Background()
	name := "foo.foogo"
	for i := 0; i < b.N; i++ {
		_, err := GetLanguageDetails(ctx, name, buf)
		if err != nil {
			b.Fatal(err)
		}
	}
}
