# bullet

[![License](https://img.shields.io/badge/license-MIT-yellowgreen.svg?style=flat-square)](LICENSE.txt)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/mitsuse/bullet)
[![Wercker](http://img.shields.io/wercker/ci/54de05bd3e143292231627b6.svg?style=flat-square)](https://app.wercker.com/project/bykey/ce611c227b5097606b6f3dc589c80ff5)
[![Coverage](https://img.shields.io/coveralls/mitsuse/bullet.svg?style=flat-square)](https://coveralls.io/r/mitsuse/bullet)

A command-line tool to send a message or a file to [Pushbullet](https://www.pushbullet.com/).

## Usage

Send a message to your devices via Pushbullet with the following command:

```bash
$ bullet send -t <title> -m <message>
```

You can send a link with the message:

```bash
$ bullet send -t <title> -m <message> -l <url>
```

## Install

The executable binaries have not been distributed yet.
Please compile sources into the binary:

```bash
$ go get github.com/mitsuse/bullet/cmd/bullet
```
