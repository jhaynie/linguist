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
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "foo.js" {
		t.Fatalf("expected Path to be foo.js, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "JavaScript" {
		t.Fatalf("expected Language.Name to be JavaScript, was %v", r.Result.Language.Name)
	}
}

func TestEmptyJavascript(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "test/foo.js", nil)
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "test/foo.js" {
		t.Fatalf("expected Path to be test/foo.js, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "JavaScript" {
		t.Fatalf("expected Language.Name to be JavaScript, was %v", r.Result.Language.Name)
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
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "foo.go" {
		t.Fatalf("expected Path to be foo.go, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "Go" {
		t.Fatalf("expected Language.Name to be Go, was %v", r.Result.Language.Name)
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
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "foo.swift" {
		t.Fatalf("expected Path to be foo.swift, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "Swift" {
		t.Fatalf("expected Language.Name to be Swift, was %v", r.Result.Language.Name)
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
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "Makefile" {
		t.Fatalf("expected Path to be Makefile, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "Makefile" {
		t.Fatalf("expected Language.Name to be Makefile, was %v", r.Result.Language.Name)
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
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "foo.json" {
		t.Fatalf("expected Path to be foo.json, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "JSON" {
		t.Fatalf("expected Language.Name to be JSON, was %v", r.Result.Language.Name)
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
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "foo.yaml" {
		t.Fatalf("expected Path to be foo.yaml, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "YAML" {
		t.Fatalf("expected Language.Name to be YAML, was %v", r.Result.Language.Name)
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
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "foo.yml" {
		t.Fatalf("expected Path to be foo.yml, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "YAML" {
		t.Fatalf("expected Language.Name to be YAML, was %v", r.Result.Language.Name)
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
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "foo.ejs" {
		t.Fatalf("expected Path to be foo.ejs, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "EJS" {
		t.Fatalf("expected Language.Name to be EJS, was %v", r.Result.Language.Name)
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
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "foo.html" {
		t.Fatalf("expected Path to be foo.html, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "HTML" {
		t.Fatalf("expected Language.Name to be HTML, was %v", r.Result.Language.Name)
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
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "foo.css" {
		t.Fatalf("expected Path to be foo.css, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "CSS" {
		t.Fatalf("expected Language.Name to be CSS, was %v", r.Result.Language.Name)
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
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "foo.scss" {
		t.Fatalf("expected Path to be foo.scss, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "SCSS" {
		t.Fatalf("expected Language.Name to be SCSS, was %v", r.Result.Language.Name)
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
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "foo.md" {
		t.Fatalf("expected Path to be foo.md, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "Markdown" {
		t.Fatalf("expected Language.Name to be Markdown, was %v", r.Result.Language.Name)
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
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "foo.sh" {
		t.Fatalf("expected Path to be foo.sh, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "Shell" {
		t.Fatalf("expected Language.Name to be Shell, was %v", r.Result.Language.Name)
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
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "foo.jsx" {
		t.Fatalf("expected Path to be foo.jsx, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "JSX" {
		t.Fatalf("expected Language.Name to be JSX, was %v", r.Result.Language.Name)
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
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "foo.m" {
		t.Fatalf("expected Path to be foo.m, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "Objective-C" {
		t.Fatalf("expected Language.Name to be Objective-C, was %v", r.Result.Language.Name)
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
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "foo.mm" {
		t.Fatalf("expected Path to be foo.mm, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "Objective-C++" {
		t.Fatalf("expected Language.Name to be Objective-C++, was %v", r.Result.Language.Name)
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
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "foo.c" {
		t.Fatalf("expected Path to be foo.c, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "C" {
		t.Fatalf("expected Language.Name to be C, was %v", r.Result.Language.Name)
	}
}

func TestLanguageOptimizationsCHeader(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.h", []byte("#include <stdlib.h>\nint c;\n"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "foo.h" {
		t.Fatalf("expected Path to be foo.h, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "C" {
		t.Fatalf("expected Language.Name to be C, was %v", r.Result.Language.Name)
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
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "foo.cpp" {
		t.Fatalf("expected Path to be foo.cpp, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "C++" {
		t.Fatalf("expected Language.Name to be C++, was %v", r.Result.Language.Name)
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
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "foo.cc" {
		t.Fatalf("expected Path to be foo.cc, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "C++" {
		t.Fatalf("expected Language.Name to be C++, was %v", r.Result.Language.Name)
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
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "foo.json5" {
		t.Fatalf("expected Path to be foo.json5, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "JSON5" {
		t.Fatalf("expected Language.Name to be JSON5 was %v", r.Result.Language.Name)
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
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "foo.proto" {
		t.Fatalf("expected Path to be foo.proto, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "Protocol Buffer" {
		t.Fatalf("expected Language.Name to be Protocol Buffer was %v", r.Result.Language.Name)
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
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "foo.py" {
		t.Fatalf("expected Path to be foo.py, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "Python" {
		t.Fatalf("expected Language.Name to be Python was %v", r.Result.Language.Name)
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
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "foo.rb" {
		t.Fatalf("expected Path to be foo.rb, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "Ruby" {
		t.Fatalf("expected Language.Name to be Ruby was %v", r.Result.Language.Name)
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
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "foo.java" {
		t.Fatalf("expected Path to be foo.java, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "Java" {
		t.Fatalf("expected Language.Name to be Java was %v", r.Result.Language.Name)
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
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "foo.cs" {
		t.Fatalf("expected Path to be foo.cs, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "C#" {
		t.Fatalf("expected Language.Name to be C# was %v", r.Result.Language.Name)
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
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "foo.xml" {
		t.Fatalf("expected Path to be foo.xml, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "XML" {
		t.Fatalf("expected Language.Name to be XML was %v", r.Result.Language.Name)
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
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "foo.hbs" {
		t.Fatalf("expected Path to be foo.hbs, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "Handlebars" {
		t.Fatalf("expected Language.Name to be Handlebars was %v", r.Result.Language.Name)
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
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "foo.lua" {
		t.Fatalf("expected Path to be foo.lua, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "Lua" {
		t.Fatalf("expected Language.Name to be Lua was %v", r.Result.Language.Name)
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
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "Dockerfile" {
		t.Fatalf("expected Path to be Dockerfile, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "Dockerfile" {
		t.Fatalf("expected Language.Name to be Dockerfile was %v", r.Result.Language.Name)
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
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "foo.txt" {
		t.Fatalf("expected Path to be foo.txt, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "Text" {
		t.Fatalf("expected Language.Name to be Text was %v", r.Result.Language.Name)
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
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "foo.sql" {
		t.Fatalf("expected Path to be foo.sql, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "SQL" {
		t.Fatalf("expected Language.Name to be SQL was %v", r.Result.Language.Name)
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
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "foo.ts" {
		t.Fatalf("expected Path to be foo.ts, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "TypeScript" {
		t.Fatalf("expected Language.Name to be TypeScript was %v", r.Result.Language.Name)
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
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "foo.coffee" {
		t.Fatalf("expected Path to be foo.coffee, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "CoffeeScript" {
		t.Fatalf("expected Language.Name to be CoffeeScript was %v", r.Result.Language.Name)
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
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.Result.Path != "foo.properties" {
		t.Fatalf("expected Path to be foo.properties, was %s", r.Result.Path)
	}
	if r.Result.Type != "text" {
		t.Fatalf("expected Type to be text, was %v", r.Result.Type)
	}
	if r.Result.Language.Name != "INI" {
		t.Fatalf("expected Language.Name to be INI was %v", r.Result.Language.Name)
	}
}

func TestCacheStats(t *testing.T) {
	cacheCounterReset()
	if CacheHits() != 0 {
		t.Fatalf("expected cache hits to be 0, was %d", CacheHits())
	}
	if CacheMisses() != 0 {
		t.Fatalf("expected cache misses to be 0, was %d", CacheMisses())
	}
	popular := MostPopular()
	if popular.Language.Name != "YAML" {
		t.Fatalf("expected popular.Language to be YAML, was %v", popular.Language.Name)
	}
	for i := 0; i < 101; i++ {
		GetLanguageDetails(context.Background(), "foo.go", []byte("package test\nvar a string\n"))
	}
	popular = MostPopular()
	if popular.Language.Name != "Go" {
		t.Fatalf("expected popular.Language to be Go, was %v", popular.Language.Name)
	}
	if CacheHits() != 101 {
		t.Fatalf("expected cache hits to be 101, was %d", CacheHits())
	}
	if CacheMisses() != 0 {
		t.Fatalf("expected cache misses to be 0, was %d", CacheMisses())
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
	if r.Result != nil {
		t.Fatal("expected results to be nil")
	}
	if !r.IsExcluded {
		t.Fatal("expected IsExcluded to be true")
	}
	if !r.IsBinary {
		t.Fatal("expected IsBinary to be true")
	}
	if r.IsLarge {
		t.Fatal("expected IsLarge to be false")
	}
}

func TestIgnoreLargeBuffer(t *testing.T) {
	var buf bytes.Buffer
	for i := 0; i < MaxBufferSize+1; i++ {
		buf.Write([]byte("x"))
	}
	r, err := GetLanguageDetails(context.Background(), "foo.js", buf.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	if !r.Success {
		t.Fatal("expected success to be true")
	}
	if r.Result != nil {
		t.Fatal("expected results to be nil")
	}
	if !r.IsExcluded {
		t.Fatal("expected IsExcluded to be true")
	}
	if r.IsBinary {
		t.Fatal("expected IsBinary to be false")
	}
	if !r.IsLarge {
		t.Fatal("expected IsLarge to be true")
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
	if r.Result != nil {
		t.Fatal("expected results to be nil")
	}
	if !r.IsExcluded {
		t.Fatal("expected IsExcluded to be true")
	}
	if r.IsBinary {
		t.Fatal("expected IsBinary to be false")
	}
	if r.IsLarge {
		t.Fatal("expected IsLarge to be false")
	}
}

func TestIgnoreExcludedFilename(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "npm-debug.log", []byte("foo\n"))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Success {
		t.Fatal("expected success to be true")
	}
	if r.Result != nil {
		t.Fatal("expected results to be nil")
	}
	if !r.IsExcluded {
		t.Fatal("expected IsExcluded to be true")
	}
	if r.IsBinary {
		t.Fatal("expected IsBinary to be false")
	}
	if r.IsLarge {
		t.Fatal("expected IsLarge to be false")
	}
}

func TestMutation(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.js", []byte("var foo\n"))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Success {
		t.Fatal("expected success to be true")
	}
	if r.Result == nil {
		t.Fatal("expected results to not be nil")
	}
	if r.IsExcluded {
		t.Fatal("expected IsExcluded to be false")
	}
	if r.IsBinary {
		t.Fatal("expected IsBinary to be false")
	}
	if r.IsLarge {
		t.Fatal("expected IsLarge to be false")
	}
	if r.Result.Language.Name != "JavaScript" {
		t.Fatalf("expected language to be JavaScript, but was %s", r.Result.Language.Name)
	}
	r.Result.Language.Name = "foo"
	r, err = GetLanguageDetails(context.Background(), "foo.js", []byte("var foo\n"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Result.Language.Name != "JavaScript" {
		t.Fatalf("expected language to be JavaScript, but was %s", r.Result.Language.Name)
	}
}

func TestGoVendored(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "vendor/github.com/jhaynie/foo/foo.go", []byte("package foo\n"))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Success {
		t.Fatal("expected success to be true")
	}
	if r.Result != nil {
		t.Fatal("expected results to be nil")
	}
	if !r.IsExcluded {
		t.Fatal("expected IsExcluded to be true")
	}
	if r.IsBinary {
		t.Fatal("expected IsBinary to be false")
	}
	if r.IsLarge {
		t.Fatal("expected IsLarge to be false")
	}
	r, err = GetLanguageDetails(context.Background(), "Godeps/github.com/jhaynie/foo/foo.go", []byte("package foo\n"))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Success {
		t.Fatal("expected success to be true")
	}
	if r.Result != nil {
		t.Fatal("expected results to be nil")
	}
	if !r.IsExcluded {
		t.Fatal("expected IsExcluded to be true")
	}
	if r.IsBinary {
		t.Fatal("expected IsBinary to be false")
	}
	if r.IsLarge {
		t.Fatal("expected IsLarge to be false")
	}
}

func TestNodeVendored(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "node_modules/foo/foo.js", []byte("var a = 1\n"))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Success {
		t.Fatal("expected success to be true")
	}
	if r.Result != nil {
		t.Fatal("expected results to be nil")
	}
	if !r.IsExcluded {
		t.Fatal("expected IsExcluded to be true")
	}
	if r.IsBinary {
		t.Fatal("expected IsBinary to be false")
	}
	if r.IsLarge {
		t.Fatal("expected IsLarge to be false")
	}
}

func TestMinimizedJS(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.min.js", []byte("var a = 1\n"))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Success {
		t.Fatal("expected success to be true")
	}
	if r.Result != nil {
		t.Fatal("expected results to be nil")
	}
	if !r.IsExcluded {
		t.Fatal("expected IsExcluded to be true")
	}
	if r.IsBinary {
		t.Fatal("expected IsBinary to be false")
	}
	if r.IsLarge {
		t.Fatal("expected IsLarge to be false")
	}
}

func TestDistJS(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "dist/foo.js", []byte("var a = 1\n"))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Success {
		t.Fatal("expected success to be true")
	}
	if r.Result != nil {
		t.Fatal("expected results to be nil")
	}
	if !r.IsExcluded {
		t.Fatal("expected IsExcluded to be true")
	}
	if r.IsBinary {
		t.Fatal("expected IsBinary to be false")
	}
	if r.IsLarge {
		t.Fatal("expected IsLarge to be false")
	}
}

func TestJSSourceMap(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.js.map", []byte("var a = 1\n"))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Success {
		t.Fatal("expected success to be true")
	}
	if r.Result != nil {
		t.Fatal("expected results to be nil")
	}
	if !r.IsExcluded {
		t.Fatal("expected IsExcluded to be true")
	}
	if r.IsBinary {
		t.Fatal("expected IsBinary to be false")
	}
	if r.IsLarge {
		t.Fatal("expected IsLarge to be false")
	}
}

func TestPreoptimizationAPI(t *testing.T) {
	r := CheckPreoptimizationCache("test.js")
	if !r.Success {
		t.Fatal("expected success to be true")
	}
	if r.Result == nil {
		t.Fatal("expected results to be non-nil")
	}
	if r.IsExcluded {
		t.Fatal("expected IsExcluded to be false")
	}
	if r.IsBinary {
		t.Fatal("expected IsBinary to be false")
	}
	if r.IsLarge {
		t.Fatal("expected IsLarge to be false")
	}
	if r.Result.Language.Name != "JavaScript" {
		t.Fatalf("expected r.Result.Language.Name to be JavaScript, was %v", r.Result.Language.Name)
	}
}

func TestAddCustomExtensionRule(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.jeff", []byte("var a = 1\n"))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Success {
		t.Fatal("expected success to be true")
	}
	if r.Result == nil {
		t.Fatal("expected results to not be nil")
	}
	if r.IsExcluded {
		t.Fatal("expected IsExcluded to be false")
	}
	AddExcludedExtension(".jeff")
	defer RemoveExcludedExtension(".jeff")
	r, err = GetLanguageDetails(context.Background(), "foo.jeff", []byte("var a = 1\n"))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Success {
		t.Fatal("expected success to be true")
	}
	if r.Result != nil {
		t.Fatal("expected results to be nil")
	}
	if !r.IsExcluded {
		t.Fatal("expected IsExcluded to be true")
	}
	RemoveExcludedExtension(".jeff")
	r, err = GetLanguageDetails(context.Background(), "foo.jeff", []byte("var a = 1\n"))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Success {
		t.Fatal("expected success to be true")
	}
	if r.Result == nil {
		t.Fatal("expected results to not be nil")
	}
	if r.IsExcluded {
		t.Fatal("expected IsExcluded to be false")
	}
}

func TestAddCustomFilenameRule(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.jeff", []byte("var a = 1\n"))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Success {
		t.Fatal("expected success to be true")
	}
	if r.Result == nil {
		t.Fatal("expected results to not be nil")
	}
	if r.IsExcluded {
		t.Fatal("expected IsExcluded to be false")
	}
	AddExcludedFilename("foo.jeff")
	defer RemoveExcludedFilename("foo.jeff")
	r, err = GetLanguageDetails(context.Background(), "foo.jeff", []byte("var a = 1\n"))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Success {
		t.Fatal("expected success to be true")
	}
	if r.Result != nil {
		t.Fatal("expected results to be nil")
	}
	if !r.IsExcluded {
		t.Fatal("expected IsExcluded to be true")
	}
	RemoveExcludedFilename("foo.jeff")
	r, err = GetLanguageDetails(context.Background(), "foo.jeff", []byte("var a = 1\n"))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Success {
		t.Fatal("expected success to be true")
	}
	if r.Result == nil {
		t.Fatal("expected results to not be nil")
	}
	if r.IsExcluded {
		t.Fatal("expected IsExcluded to be false")
	}
}

func TestAddCustomMatchRule(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.jeff", []byte("var a = 1\n"))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Success {
		t.Fatal("expected success to be true")
	}
	if r.Result == nil {
		t.Fatal("expected results to not be nil")
	}
	if r.IsExcluded {
		t.Fatal("expected IsExcluded to be false")
	}
	match := NewMatcher("\\.jeff$")
	AddExcludedRule(match)
	defer RemoveExcludedRule(match)
	r, err = GetLanguageDetails(context.Background(), "foo.jeff", []byte("var a = 1\n"))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Success {
		t.Fatal("expected success to be true")
	}
	if r.Result != nil {
		t.Fatal("expected results to be nil")
	}
	if !r.IsExcluded {
		t.Fatal("expected IsExcluded to be true")
	}
	RemoveExcludedRule(match)
	r, err = GetLanguageDetails(context.Background(), "foo.jeff", []byte("var a = 1\n"))
	if err != nil {
		t.Fatal(err)
	}
	if !r.Success {
		t.Fatal("expected success to be true")
	}
	if r.Result == nil {
		t.Fatal("expected results to not be nil")
	}
	if r.IsExcluded {
		t.Fatal("expected IsExcluded to be false")
	}
}

func TestMulti(t *testing.T) {
	files := []*File{
		NewFile("foo.properties", []byte("foo=1")),
		NewFile("foo.js", []byte("var foo=1")),
		NewFile("foo.jsx", []byte("var foo=1")),
		NewFile("foo.go", []byte("package foo\n")),
	}
	results, err := GetLanguageDetailsMultiple(context.Background(), files)
	if err != nil {
		t.Fatal(err)
	}
	if len(results) != len(files) {
		t.Fatalf("expected length to be %d, was %d", len(files), len(results))
	}
	var expected = []struct {
		language string
		thetype  string
		path     string
	}{
		{"INI", "text", "foo.properties"},
		{"JavaScript", "text", "foo.js"},
		{"JSX", "text", "foo.jsx"},
		{"Go", "text", "foo.go"},
	}
	for i, e := range expected {
		if !results[i].Success {
			t.Fatalf("expected %d to be Success but was not", i)
		}
		if results[i].Result == nil {
			t.Fatalf("expected %d Result to not be nil", i)
		}
		if results[i].Result.Language.Name != e.language {
			t.Fatalf("expected %d language to be %s was %s", i, e.language, results[i].Result.Language.Name)
		}
		if results[i].Result.Type != e.thetype {
			t.Fatalf("expected %d language type to be %s was %s", i, e.thetype, results[i].Result.Type)
		}
		if results[i].Result.Path != e.path {
			t.Fatalf("expected %d language path to be %s was %s", i, e.path, results[i].Result.Path)
		}
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

func TestExplicitPreoptimizationCache(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.js", []byte("a = 'bar'"), false)
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if !r.IsCached {
		t.Fatal("expected IsCached to be true")
	}
}

func TestSkipPreoptimizationCache(t *testing.T) {
	r, err := GetLanguageDetails(context.Background(), "foo.js", []byte("a = 'bar'"), true)
	if err != nil {
		t.Fatal(err)
	}
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if r.Result == nil {
		t.Fatal("expected a result but was nil")
	}
	if r.IsCached {
		t.Fatal("expected IsCached to be false")
	}
}

func TestPreoptimizationExcluded(t *testing.T) {
	r := CheckPreoptimizationCache("package.json")
	if r.Success == false {
		t.Fatal("expected result.success to be true")
	}
	if r.Result != nil {
		t.Fatal("expected result to be nil")
	}
	if r.IsCached {
		t.Fatal("expected IsCached to be false")
	}
	if !r.IsExcluded {
		t.Fatal("expected IsExcluded to be true")
	}
}

func TestPreoptimizationExcludedRules(t *testing.T) {
	for _, name := range []string{
		"tests/vendor/bundle/ruby/2.0.0/gems/page-object-0.9.2/spec/page-object/platforms/selenium_webdriver/selenium_page_object_spec.rb",
		"vendor/bundle/ruby/2.2.0/gems/actionpack-4.2.3/lib/action_dispatch/routing.rb",
		"node_modules/some/test.js",
		"vendor/github.com/pinpt/worker/main.go",
	} {
		r := CheckPreoptimizationCache(name)
		if r.Success == false {
			t.Fatal("expected result.success to be true", name)
		}
		if r.Result != nil {
			t.Fatal("expected result to be nil", name)
		}
		if r.IsCached {
			t.Fatal("expected IsCached to be false", name)
		}
		if !r.IsExcluded {
			t.Fatal("expected IsExcluded to be true", name)
		}
	}
}
