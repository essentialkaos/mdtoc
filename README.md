<p align="center"><a href="#readme"><img src="https://gh.kaos.st/mdtoc.svg"/></a></p>

<p align="center">
  <a href="https://github.com/essentialkaos/mdtoc/actions"><img src="https://github.com/essentialkaos/mdtoc/workflows/CI/badge.svg" alt="GitHub Actions Status" /></a>
  <a href="https://github.com/essentialkaos/mdtoc/actions?query=workflow%3ACodeQL"><img src="https://github.com/essentialkaos/mdtoc/workflows/CodeQL/badge.svg" /></a>
  <a href="https://goreportcard.com/report/github.com/essentialkaos/mdtoc"><img src="https://goreportcard.com/badge/github.com/essentialkaos/mdtoc"></a>
  <a href="https://codebeat.co/projects/github-com-essentialkaos-mdtoc-master"><img alt="codebeat badge" src="https://codebeat.co/badges/196d721e-00ad-4dff-9032-9b5bbd11b723" /></a>
  <a href="#license"><img src="https://gh.kaos.st/apache2.svg"></a>
</p>

<p align="center"><a href="#installation">Installation</a> • <a href="#usage">Usage</a> • <a href="#build-status">Build Status</a> • <a href="#contributing">Contributing</a> • <a href="#license">License</a></p>

<br/>

`mdtoc` is simple utility for generating table of contents for markdown files.

### Installation

#### From source

To build the MDToc from scratch, make sure you have a working Go 1.14+ workspace ([instructions](https://golang.org/doc/install)), then:

```
go get github.com/essentialkaos/mdtoc
```

If you want to update MDToc to latest stable release, do:

```
go get -u github.com/essentialkaos/mdtoc
```

#### Prebuilt binaries

You can download prebuilt binaries for Linux and OS X from [EK Apps Repository](https://apps.kaos.st/mdtoc/latest):

```
bash <(curl -fsSL https://apps.kaos.st/get) mdtoc
```

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
| `master` | [![CI](https://github.com/essentialkaos/mdtoc/workflows/CI/badge.svg?branch=master)](https://github.com/essentialkaos/mdtoc/actions) |
| `develop` | [![CI](https://github.com/essentialkaos/mdtoc/workflows/CI/badge.svg?branch=develop)](https://github.com/essentialkaos/mdtoc/actions) |

### Contributing

Before contributing to this project please read our [Contributing Guidelines](https://github.com/essentialkaos/contributing-guidelines#contributing-guidelines).

### License

[Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>