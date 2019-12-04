<p align="center"><a href="#readme"><img src="https://gh.kaos.st/mdtoc.svg"/></a></p>

<p align="center">
  <a href="https://travis-ci.com/essentialkaos/mdtoc"><img src="https://travis-ci.com/essentialkaos/mdtoc.svg"></a>
  <a href="https://goreportcard.com/report/github.com/essentialkaos/mdtoc"><img src="https://goreportcard.com/badge/github.com/essentialkaos/mdtoc"></a>
  <a href="https://codebeat.co/projects/github-com-essentialkaos-mdtoc-master"><img alt="codebeat badge" src="https://codebeat.co/badges/196d721e-00ad-4dff-9032-9b5bbd11b723" /></a>
  <a href="https://essentialkaos.com/ekol"><img src="https://gh.kaos.st/ekol.svg"></a>
</p>

<p align="center"><a href="#installation">Installation</a> • <a href="#usage">Usage</a> • <a href="#build-status">Build Status</a> • <a href="#contributing">Contributing</a> • <a href="#license">License</a></p>

<br/>

`mdtoc` is simple utility for generating table of contents for markdown files.

### Installation

#### From source

Before the initial install allows git to use redirects for [pkg.re](https://github.com/essentialkaos/pkgre) service (reason why you should do this described [here](https://github.com/essentialkaos/pkgre#git-support)):

```
git config --global http.https://pkg.re.followRedirects true
```

To build the MDToc from scratch, make sure you have a working Go 1.8+ workspace ([instructions](https://golang.org/doc/install)), then:

```
go get github.com/essentialkaos/mdtoc
```

If you want to update MDToc to latest stable release, do:

```
go get -u github.com/essentialkaos/mdtoc
```

#### Prebuilt binaries

You can download prebuilt binaries for Linux and OS X from [EK Apps Repository](https://apps.kaos.st/mdtoc/latest).

### Usage

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

### Build Status

| Branch | Status |
|------------|--------|
| `master` | [![Build Status](https://travis-ci.com/essentialkaos/mdtoc.svg?branch=master)](https://travis-ci.com/essentialkaos/mdtoc) |
| `develop` | [![Build Status](https://travis-ci.com/essentialkaos/mdtoc.svg?branch=develop)](https://travis-ci.com/essentialkaos/mdtoc) |

### Contributing

Before contributing to this project please read our [Contributing Guidelines](https://github.com/essentialkaos/contributing-guidelines#contributing-guidelines).

### License

[EKOL](https://essentialkaos.com/ekol)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
