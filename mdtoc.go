package main

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"os"
	"regexp"
	"strings"

	"pkg.re/essentialkaos/ek.v5/arg"
	"pkg.re/essentialkaos/ek.v5/fmtc"
	"pkg.re/essentialkaos/ek.v5/fmtutil"
	"pkg.re/essentialkaos/ek.v5/fsutil"
	"pkg.re/essentialkaos/ek.v5/strutil"
	"pkg.re/essentialkaos/ek.v5/usage"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	APP  = "MDToc"
	VER  = "0.0.4"
	DESC = "Utility for generating table of contents for markdown files"
)

const (
	ARG_MIN_LEVEL = "m:min-level"
	ARG_MAX_LEVEL = "M:max-level"
	ARG_FLAT      = "f:flat"
	ARG_HTML      = "H:html"
	ARG_NO_COLOR  = "nc:no-color"
	ARG_HELP      = "h:help"
	ARG_VER       = "v:version"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type Header struct {
	Level int
	Text  string
	Link  string
}

// ////////////////////////////////////////////////////////////////////////////////// //

var argList = arg.Map{
	ARG_MIN_LEVEL: &arg.V{Type: arg.INT, Value: 1, Min: 1, Max: 6},
	ARG_MAX_LEVEL: &arg.V{Type: arg.INT, Value: 6, Min: 1, Max: 6},
	ARG_FLAT:      &arg.V{Type: arg.BOOL},
	ARG_HTML:      &arg.V{Type: arg.BOOL},
	ARG_NO_COLOR:  &arg.V{Type: arg.BOOL},
	ARG_HELP:      &arg.V{Type: arg.BOOL, Alias: "u:usage"},
	ARG_VER:       &arg.V{Type: arg.BOOL, Alias: "ver"},
}

var anchorRegExp = regexp.MustCompile(`[\s\d\w-]`)

// ////////////////////////////////////////////////////////////////////////////////// //

func main() {
	args, errs := arg.Parse(argList)

	if len(errs) != 0 {
		for _, err := range errs {
			printError(err.Error())
		}

		os.Exit(1)
	}

	if arg.GetB(ARG_NO_COLOR) {
		fmtc.DisableColors = true
	}

	if arg.GetB(ARG_VER) {
		showAbout()
		return
	}

	if arg.GetB(ARG_HELP) {
		showUsage()
		return
	}

	var file string

	if len(args) == 0 {
		file = findProperReadme()

		if file == "" {
			showUsage()
			return
		}

	} else {
		file = args[0]
	}

	checkFile(file)
	printTOC(file)
}

// findProperReadme try to find readme file in current directory
func findProperReadme() string {
	file := fsutil.ProperPath("FRS", []string{"README.md", "readme.md"})
	return file
}

// checkFile check markdown file before processing
func checkFile(file string) {
	if !fsutil.IsExist(file) {
		printError("Can't read file %s: file is not exist", file)
		os.Exit(1)
	}

	if !fsutil.IsRegular(file) {
		printError("Can't read file %s: is not a file", file)
		os.Exit(1)
	}

	if !fsutil.IsReadable(file) {
		printError("Can't read file %s: file is not readable", file)
		os.Exit(1)
	}

	if !fsutil.IsNonEmpty(file) {
		printError("Can't read file %s: file is empty", file)
		os.Exit(1)
	}
}

// printTOC collect headers and print ToC for given markdown file
func printTOC(file string) {
	fd, err := os.Open(file)

	if err != nil {
		printError("Can't read file: %v", err)
	}

	defer fd.Close()

	reader := bufio.NewReader(fd)
	scanner := bufio.NewScanner(reader)

	var headers []*Header

	for scanner.Scan() {
		line := scanner.Text()

		if !strings.HasPrefix(line, "#") {
			continue
		}

		headers = append(headers, parseHeader(line))
	}

	if len(headers) == 0 {
		printWarn("Headers not found in given file")
		return
	}

	var toc string

	switch {
	case !arg.GetB(ARG_FLAT):
		toc = renderTOC(headers)
	case arg.GetB(ARG_FLAT) && arg.GetB(ARG_HTML):
		toc = renderFlatHTMLTOC(headers)
	case arg.GetB(ARG_FLAT) && !arg.GetB(ARG_HTML):
		toc = renderFlatTOC(headers)
	}

	if toc == "" {
		printWarn("Suitable headers not found in given file")
		return
	}

	fmtutil.Separator(false)
	fmtc.Println(toc)
	fmtutil.Separator(false)
}

// renderTOC render headers as default (vertical) markdown ToC
func renderTOC(headers []*Header) string {
	var toc []string

	for _, header := range headers {
		if !isSuitableHeader(header) {
			continue
		}

		toc = append(toc, fmtc.Sprintf(
			"%s [%s](%s)",
			getMarkdownListPrefix(header.Level),
			header.Text, header.Link,
		))
	}

	return strings.Join(toc, "\n")
}

// renderFlatTOC render headers as flat (horizontal) markdown ToC
func renderFlatTOC(headers []*Header) string {
	var toc []string

	for _, header := range headers {
		if !isSuitableHeader(header) {
			continue
		}

		toc = append(toc, fmtc.Sprintf("[%s](%s)", header.Text, header.Link))
	}

	if len(toc) == 0 {
		return ""
	}

	return strings.Join(toc, " • ")
}

// renderFlatTOC render headers as flat (horizontal) HTML ToC
func renderFlatHTMLTOC(headers []*Header) string {
	var toc []string

	for _, header := range headers {
		if !isSuitableHeader(header) {
			continue
		}

		toc = append(toc, fmtc.Sprintf("<a href=\"%s\">%s</a>", header.Link, header.Text))
	}

	if len(toc) == 0 {
		return ""
	}

	return "<p align=\"center\">" + strings.Join(toc, " • ") + "</p>"
}

// isSuitableHeader return true if header complies defined levels
func isSuitableHeader(header *Header) bool {
	if header.Level < arg.GetI(ARG_MIN_LEVEL) || header.Level > arg.GetI(ARG_MAX_LEVEL) {
		return false
	}

	return true
}

// parseHeader parse header text and return header struct
func parseHeader(text string) *Header {
	header := &Header{}

	headerText := strings.TrimRight(text, " ")
	headerText = removeLinks(headerText)

	header.Text, header.Level = parseHeaderText(headerText)
	header.Link = makeLink(header.Text)

	return header
}

// makeLink convert header text to anchor link name
func makeLink(text string) string {
	result := text

	result = strings.Replace(result, " ", "-", -1)
	result = strings.ToLower(result)
	result = strings.Join(anchorRegExp.FindAllString(result, -1), "")

	return "#" + result
}

// parseHeaderText parse text and return level and header
func parseHeaderText(text string) (string, int) {
	var level = 0
	var header = ""

	for i, s := range text {
		if s == '#' {
			level++
			continue
		}

		header = strings.TrimLeft(strutil.Substr(text, i, 9999), " ")

		break
	}

	return header, level
}

// getMarkdownListPrefix return list prefix for given level
func getMarkdownListPrefix(level int) string {
	return strings.Repeat("  ", level-1) + "*"
}

// removeLinks remove links from text
func removeLinks(text string) string {
	result := text

	var (
		startLink int
		innerLink bool
	)

MAINLOOP:
	for {
		if !strings.Contains(result, "](") {
			break
		}

		startLink = 0

		for i, s := range result {

			if s == '[' {
				if startLink != 0 {
					continue
				}

				startLink = i
				innerLink = strutil.Substr(result, i+1, i+2) == "!"

				continue

			} else if s == ')' {
				if innerLink {
					innerLink = false
					continue
				} else {
					result = result[0:startLink] + result[i+1:]
					result = strutil.Substr(result, 0, startLink) + strutil.Substr(result, i, 9999)
					continue MAINLOOP
				}
			}
		}
	}

	return result
}

// printError prints error message to console
func printError(f string, a ...interface{}) {
	fmtc.Printf("{r}"+f+"{!}\n", a...)
}

// printWarn prints warning message to console
func printWarn(f string, a ...interface{}) {
	fmtc.Printf("{y}"+f+"{!}\n", a...)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func showUsage() {
	info := usage.NewInfo("", "file")

	info.AddOption(ARG_FLAT, "Print flat (horizontal) ToC")
	info.AddOption(ARG_HTML, "Render HTML ToC instead Markdown (works with {g}--flat{!})")
	info.AddOption(ARG_MIN_LEVEL, "Minimal header level", "1-6")
	info.AddOption(ARG_MAX_LEVEL, "Maximum header level", "1-6")
	info.AddOption(ARG_NO_COLOR, "Disable colors in output")
	info.AddOption(ARG_HELP, "Show this help message")
	info.AddOption(ARG_VER, "Show version")

	info.AddExample("readme.md", "Generate table of contents for readme.md")
	info.AddExample("-m 2 -M 4 readme.md", "Generate table of contents for readme.md with 2-4 level headers")

	info.Render()
}

func showAbout() {
	about := &usage.About{
		App:     APP,
		Version: VER,
		Desc:    DESC,
		Year:    2006,
		Owner:   "ESSENTIAL KAOS",
		License: "Essential Kaos Open Source License <https://essentialkaos.com/ekol?en>",
	}

	about.Render()
}
