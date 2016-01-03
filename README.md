# bullet

[![License](https://img.shields.io/badge/license-MIT-yellowgreen.svg?style=flat-square)][license]
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)][godoc]
[![Wercker](http://img.shields.io/wercker/ci/54e273b67e0e958634000915.svg?style=flat-square)][wercker]

[license]: LICENSE.txt
[godoc]: http://godoc.org/github.com/mitsuse/bullet
[wercker]: https://app.wercker.com/project/bykey/c73697ea3e91260dfed3ab032341f2a1

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

### Others

For other platform, the executable binaries are downloadable
from the [release page](https://github.com/mitsuse/bullet/releases).

In the case of building from the sources, try the following:

```bash
$ go get -u github.com/mitsuse/bullet
```

## License

Please read [LICENSE.txt](LICENSE.txt).
