# slack2md

[![Code Climate](https://codeclimate.com/github/ohtake/slack2md/badges/gpa.svg)](https://codeclimate.com/github/ohtake/slack2md)

You can convert [exported Slack history json](https://get.slack.help/hc/en-us/articles/201658943-Exporting-your-team-s-Slack-history) into Markdown.

## Requirement

* [Go](https://golang.org/doc/install)

## Export

1. [Export Slack history](https://my.slack.com/services/export) and wait its completion.
1. Download and extract the zip file into `slack_export` directory.

## Build and publish

```bash
git checkout -B example
go build
./slack2md
git add -f output
git commit -m "Convert slack json into markdown"
git push origin example -f
git checkout -
```
