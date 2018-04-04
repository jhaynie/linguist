# Linguist Go Wrapper [![CircleCI](https://circleci.com/gh/jhaynie/linguist.svg?style=svg)](https://circleci.com/gh/jhaynie/linguist)

This is a simple Go API around the [Linguist](https://github.com/pinpt/linguist) server.

## Install

```shell
go get -u github.com/jhaynie/linguist
```

## Use

```golang
result, err := linguist.GetLanguageDetails(context.Background(), "test.js", []byte("var a = 1"))
```

## Adding Exclusion Rules

There are a ton of common exclusion rules to exclude certain files based on a number of heuristics built-in. However, you may need to customize the exclusion rules to further refine for your own use case.

### Add exclusion by filename

To add an exclusion that matches the name of a file, use `AddExcludedFilename`:

```golang
linguist.AddExcludedFilename("foo.extension")
```

You can remove a rule with `RemoveExcludedFilename`.

### Add exclusion by file extension

To add an exclusion that matches the extension of a file, use `AddExcludedExtension`:

```golang
linguist.AddExcludedExtension(".extension")
```

You can remove a rule with `RemoveExcludedExtension`.

### Add exclusion by regular expression match rule

To add an exclusion that matches based on a regular expression rule, use `AddExcludedRule`:

```golang
rule := linguist.NewMatcher("\\.somepath$")
linguist.AddExcludedRule(rule)
```

Use `NewMatcher` with a regular expression string which should evaluate to true to exclude the file.
Use `NewNotMatcher` with a regular expression string which should evaluate to false to exclude the file.

You can remove a rule with `RemoveExcludedRule`.

## Submitting multiple files

You can submit more than one file for analysis by using the `GetLanguageDetailsMultiple` function:

```golang
files := []*linguist.File{
	linguist.NewFile("foo.properties", []byte("foo=1")),
	linguist.NewFile("foo.js", []byte("var foo=1")),
	linguist.NewFile("foo.jsx", []byte("var foo=1")),
}
results, err := linguist.GetLanguageDetailsMultiple(context.Background(), files)
```

## Vendoring

This library depends on the Golang port of Linguist from https://github.com/generaltso/linguist.  Since this library requires a go build step to train the classifier, we have vendored the built classifier file and checked it in to source.

To build, do a normal checkout of generaltso/linguist and build it locally. Then copy over the files (remember to remove .git, .gitmodules, data/linguist folder and update .gitignore)
