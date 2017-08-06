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

## Environment

Make sure that you set environment variable `PP_LINGUIST_URL` to the url of your linguist server. Defaults to `https://linguist:25032`.

You can also set the authorization token by setting the environment variable `PP_LINGUIST_AUTH` which defaults to `1234`.

## Running Linguist

The easiest way to run linguist is via Docker:

```shell
docker run -d -p 25032:25032 pinpt/linguist
```

## Adding Exclusion Rules

There are a ton of common exclusion rules to exclude certain files based on a number of heuristics built-in. However, you may need to customize the exclusion rules to further refine for your own use case.

### Add exclusion by filename

To add an exclusion that matches the name of a file, use `AddExcludedFilename`:

```golang
AddExcludedFilename("foo.extension")
```

You can remove a rule with `RemoveExcludedFilename`.

### Add exclusion by file extension

To add an exclusion that matches the extension of a file, use `AddExcludedExtension`:

```golang
AddExcludedExtension(".extension")
```

You can remove a rule with `RemoveExcludedExtension`.

### Add exclusion by regular expression match rule

To add an exclusion that matches based on a regular expression rule, use `AddExcludedRule`:

```golang
rule := NewMatcher("\\.somepath$")
AddExcludedRule(rule)
```

Use `NewMatcher` with a regular expression string which should evaluate to true to exclude the file.
Use `NewNotMatcher` with a regular expression string which should evaluate to false to exclude the file.

You can remove a rule with `RemoveExcludedRule`.
