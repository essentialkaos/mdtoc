# MDToc [![Build Status](https://travis-ci.org/essentialkaos/mdtoc.svg?branch=master)](https://travis-ci.org/essentialkaos/mdtoc) [![Go Report Card](https://goreportcard.com/badge/github.com/essentialkaos/mdtoc)](https://goreportcard.com/report/github.com/essentialkaos/mdtoc) [![License](https://gh.kaos.io/ekol.svg)](https://essentialkaos.com/ekol)

`MDToc` is simple utility for generating table of contents for markdown files.

* [Installation](#installation)
* [Usage](#usage)
* [Build Status](#build-status)
* [Contributing](#contributing)
* [License](#license)

## Installation

### From source

Before the initial install allows git to use redirects for [pkg.re](https://github.com/essentialkaos/pkgre) service (reason why you should do this described [here](https://github.com/essentialkaos/pkgre#git-support)):

```
git config --global http.https://pkg.re.followRedirects true
```

To build the MDToc from scratch, make sure you have a working Go 1.6+ workspace ([instructions](https://golang.org/doc/install)), then:

```
go get github.com/essentialkaos/mdtoc
```

If you want to update MDToc to latest stable release, do:

```
go get -u github.com/essentialkaos/mdtoc
```

### Prebuilt binaries

You can download prebuilt binaries for Linux and OS X from [EK Apps Repository](https://apps.kaos.io/mdtoc/latest).

## Usage

```
Usage: mdtoc {options} file

Options

  --flat, -f             Print flat (horizontal) ToC
  --html, -H             Render HTML ToC instead Markdown (works with --flat)
  --min-level, -m 1-6    Minimal header level
  --max-level, -M 1-6    Maximum header level
  --no-color, -nc        Disable colors in output
  --help, -h             Show this help message
  --version, -v          Show version

Examples

  mdtoc readme.md
  Generate table of contents for readme.md

  mdtoc -m 2 -M 4 readme.md
  Generate table of contents for readme.md with 2-4 level headers

```

## Build Status

| Branch | Status |
|------------|--------|
| `master` | [![Build Status](https://travis-ci.org/essentialkaos/mdtoc.svg?branch=master)](https://travis-ci.org/essentialkaos/mdtoc) |
| `develop` | [![Build Status](https://travis-ci.org/essentialkaos/mdtoc.svg?branch=develop)](https://travis-ci.org/essentialkaos/mdtoc) |

## Contributing

Before contributing to this project please read our [Contributing Guidelines](https://github.com/essentialkaos/contributing-guidelines#contributing-guidelines).

## License

[EKOL](https://essentialkaos.com/ekol)
