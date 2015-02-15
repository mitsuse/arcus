# bullet

[![License](https://img.shields.io/badge/license-MIT-yellowgreen.svg?style=flat-square)](LICENSE.txt)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/mitsuse/bullet)
[![Wercker](http://img.shields.io/wercker/ci/54de05bd3e143292231627b6.svg?style=flat-square)](https://app.wercker.com/project/bykey/ce611c227b5097606b6f3dc589c80ff5)
[![Coverage](https://img.shields.io/coveralls/mitsuse/bullet.svg?style=flat-square)](https://coveralls.io/r/mitsuse/bullet)

A command-line tool to send a message or a file to [Pushbullet](https://www.pushbullet.com/).

## Usage

### Configuration

First, bullet requires the access token to call Pushbullet API.
Execute the below command and register your access token:

```bash
$ bullet auth
```

### Push

Send a message to your devices via Pushbullet with the following command:

```bash
$ bullet send -t <title> -m <message>
```

You can send a link or a file with the message:

```bash
$ bullet send -t <title> -m <message> -l <url or path>
```

## Installation

The executable binaries have not been distributed yet.
Please compile sources into the binary:

```bash
$ go get github.com/mitsuse/bullet/cmd/bullet
```

## License

The MIT License (MIT)

Copyright (c) 2015 Tomoya Kose.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
