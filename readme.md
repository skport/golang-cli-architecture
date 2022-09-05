# golang-webfetcher

## About

These sources are published as a record of my Golang learning.

It's a simple program that runs on the command line and displays a summary of the target website.

## Usage

Move to root directory and run go cmd.

Basic:

```Shell
go run ./app/cli/main.go summary [Target URL]
```

Response:

```
title : [string]
H1 : [string]

## Test Commands

```
go test ./app/cli/cmd -v
```