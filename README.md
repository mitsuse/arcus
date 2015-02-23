# bullet

[![License](https://img.shields.io/badge/license-MIT-yellowgreen.svg?style=flat-square)](LICENSE.txt)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/mitsuse/bullet)
[![Wercker](http://img.shields.io/wercker/ci/54e273b67e0e958634000915.svg?style=flat-square)](https://app.wercker.com/project/bykey/c73697ea3e91260dfed3ab032341f2a1)

A command-line tool to send a message to devices via [Pushbullet](https://www.pushbullet.com/).

## Usage

### Configuration

The command "bullet" requires the access token to call Pushbullet API.
Set the access token to the environment variable `BULLET_ACCESS_TOKEN`:

```bash
export BULLET_ACCESS_TOKEN=<your access token>
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

If you want to notify the specific device, use `-d` option with its *nickname*:

```bash
$ bullet send -d <nickname> -t <title> -m <message>
```

The registered devices can be listed with `bullet list`.

## Installation

### OSX

If you are an OSX user, you can install "bullet" with [Homebrew](http://brew.sh/):

```bash
brew tap mitsuse/bullet
brew install bullet-cmd
```

### ArchLinux

The [PKGBUILD file](https://aur.archlinux.org/packages/bullet-cmd/) is available on AUR.

### Others

For other platform, the executable binaries are downloadable
from the [release page](https://github.com/mitsuse/bullet/releases).

In the case of building from the sources, try the following:

```bash
$ go get -u github.com/mitsuse/bullet
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
