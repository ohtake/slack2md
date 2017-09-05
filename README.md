# slack2md

[![Build Status](https://travis-ci.org/ohtake/slack2md.svg?branch=master)](https://travis-ci.org/ohtake/slack2md)
[![Code Climate](https://codeclimate.com/github/ohtake/slack2md/badges/gpa.svg)](https://codeclimate.com/github/ohtake/slack2md)
[![codecov](https://codecov.io/gh/ohtake/slack2md/branch/master/graph/badge.svg)](https://codecov.io/gh/ohtake/slack2md)

You can convert [exported Slack history json](https://get.slack.help/hc/en-us/articles/201658943-Exporting-your-team-s-Slack-history) into Markdown.

## Example output

See <https://github.com/ohtake/slack2md/blob/example/output/index.md>.

## Requirement

* [Go](https://golang.org/doc/install)

## Usage

1. [Export Slack history](https://my.slack.com/services/export) and wait its completion.
1. Download and extract the zip file into `slack_export` directory.
1. Execute `go build` and you will get `slack2md` executable file.
1. Run `slack2md` and you will get Markdown files at `output` directory.

## Options

`./slack2md -help`
