<p align="center"><a href="#readme"><img src=".github/images/card.svg"/></a></p>

<p align="center">
  <a href="https://kaos.sh/y/mdtoc"><img src="https://kaos.sh/y/cdf1fc4eca5b405d9ca8d703d195532a.svg" alt="Codacy badge" /></a>
  <a href="https://kaos.sh/w/mdtoc/ci"><img src="https://kaos.sh/w/mdtoc/ci.svg" alt="GitHub Actions CI Status" /></a>
  <a href="https://kaos.sh/w/mdtoc/codeql"><img src="https://kaos.sh/w/mdtoc/codeql.svg" alt="GitHub Actions CodeQL Status" /></a>
  <a href="#license"><img src=".github/images/license.svg"/></a>
</p>

<p align="center"><a href="#installation">Installation</a> • <a href="#usage">Usage</a> • <a href="#ci-status">CI Status</a> • <a href="#contributing">Contributing</a> • <a href="#license">License</a></p>

<br/>

`mdtoc` is simple utility for generating table of contents for markdown files.

### Installation

#### From source

To build the MDToc from scratch, make sure you have a working [Go 1.23+](https://github.com/essentialkaos/.github/blob/master/GO-VERSION-SUPPORT.md) workspace ([instructions](https://go.dev/doc/install)), then:

```
go install github.com/essentialkaos/mdtoc@latest
```

#### Prebuilt binaries

You can download prebuilt binaries for Linux and macOS from [EK Apps Repository](https://apps.kaos.st/mdtoc/latest):

```
bash <(curl -fsSL https://apps.kaos.st/get) mdtoc
```

### Usage

<img src=".github/images/usage.svg"/>

### CI Status

| Branch | Status |
|------------|--------|
| `master` | [![CI](https://kaos.sh/w/mdtoc/ci.svg?branch=master)](https://kaos.sh/w/mdtoc/ci?query=branch:master) |
| `develop` | [![CI](https://kaos.sh/w/mdtoc/ci.svg?branch=develop)](https://kaos.sh/w/mdtoc/ci?query=branch:develop) |

### Contributing

Before contributing to this project please read our [Contributing Guidelines](https://github.com/essentialkaos/contributing-guidelines#contributing-guidelines).

### License

[Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
