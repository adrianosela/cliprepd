# cliprepd

[![Go Report Card](https://goreportcard.com/badge/github.com/adrianosela/cliprepd)](https://goreportcard.com/report/github.com/adrianosela/cliprepd)
[![GitHub issues](https://img.shields.io/github/issues/adrianosela/cliprepd.svg)](https://github.com/adrianosela/cliprepd/issues)
[![Documentation](https://godoc.org/github.com/adrianosela/cliprepd?status.svg)](https://godoc.org/github.com/adrianosela/cliprepd)
[![license](https://img.shields.io/github/license/adrianosela/cliprepd.svg)](https://github.com/adrianosela/cliprepd/blob/master/LICENSE)

Command line client for [IPrepd](https://github.com/mozilla-services/iprepd)

### Getting started

Set your local configuration with ```repd config set```:

```
$ repd config set --url http://localhost:8080 --token "APIKey test"
```
Verify your configuration with ```repd config show```:

```
$ repd config show --path ~/.repd
+----------+-----------------------+
| HOST_URL | http://localhost:8080 |
| AUTH_TK  | APIKey test           |
+----------+-----------------------+
```