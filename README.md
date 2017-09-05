# slack2md

[![Build Status](https://travis-ci.org/ohtake/slack2md.svg?branch=master)](https://travis-ci.org/ohtake/slack2md)
[![Code Climate](https://codeclimate.com/github/ohtake/slack2md/badges/gpa.svg)](https://codeclimate.com/github/ohtake/slack2md)
[![codecov](https://codecov.io/gh/ohtake/slack2md/branch/master/graph/badge.svg)](https://codecov.io/gh/ohtake/slack2md)

You can convert [exported Slack history json](https://get.slack.help/hc/en-us/articles/201658943-Exporting-your-team-s-Slack-history) into Markdown.

## Example output

See <https://github.com/ohtake/slack2md/blob/example/output/index.md>.

## Usage

1. [Export Slack history](https://my.slack.com/services/export) and wait its completion.
1. Download and extract the zip file into `slack_export` directory.
1. Run `slack2md` and you will get Markdown files at `output` directory.
   * You can fetch pre-built binaries from [releases](https://github.com/ohtake/slack2md/releases).
   * You can also build binaries using [Go](https://golang.org/doc/install): `go build`.

## Options

Type `./slack2md -help`.
