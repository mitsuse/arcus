# Arcus

[![License](https://img.shields.io/badge/license-MIT-yellowgreen.svg?style=flat-square)][license]
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)][godoc]
[![Wercker](http://img.shields.io/wercker/ci/568933fe81fd0f502201d610.svg?style=flat-square)][wercker]

[license]: LICENSE.txt
[godoc]: http://godoc.org/github.com/mitsuse/arcus
[wercker]: https://app.wercker.com/project/bykey/0573d549297d92fbc2d16f744e575399

A command-line tool to send a message to devices via [Pushbullet](https://www.pushbullet.com/).

## Usage

### Configuration

The command `arcus` requires the access token to call Pushbullet API.
Set the access token to the environment variable `ARCUS_ACCESS_TOKEN`:

```bash
export ARCUS_ACCESS_TOKEN=<your access token>
```

### Push

Send a message to your devices via Pushbullet with the following command:

```bash
$ arcus send -t <title> -m <message>
```

You can send a link or a file with the message:

```bash
$ arcus send -t <title> -m <message> -l <url or path>
```

If you want to notify the specific device, use `-d` option with its *nickname*:

```bash
$ arcus send -d <nickname> -t <title> -m <message>
```

The registered devices can be listed with `arcus list`.

## Installation

### OS X

If you are an OS X user, you can install `arcus` with [Homebrew](http://brew.sh/):

```bash
brew tap mitsuse/arcus
brew install arcus
```

### Others

For other platform, the executable binaries are downloadable
from the [release page](https://github.com/mitsuse/arcus/releases).

In the case of building from the sources, try the following:

```bash
$ go get -u github.com/mitsuse/arcus
```

## License

Please read [LICENSE.txt](LICENSE.txt).
