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

