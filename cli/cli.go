package cli

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/essentialkaos/ek/v13/fmtc"
	"github.com/essentialkaos/ek/v13/fmtutil"
	"github.com/essentialkaos/ek/v13/fsutil"
	"github.com/essentialkaos/ek/v13/mathutil"
	"github.com/essentialkaos/ek/v13/options"
	"github.com/essentialkaos/ek/v13/strutil"
	"github.com/essentialkaos/ek/v13/support"
	"github.com/essentialkaos/ek/v13/support/deps"
	"github.com/essentialkaos/ek/v13/terminal"
	"github.com/essentialkaos/ek/v13/terminal/tty"
	"github.com/essentialkaos/ek/v13/usage"
	"github.com/essentialkaos/ek/v13/usage/completion/bash"
	"github.com/essentialkaos/ek/v13/usage/completion/fish"
	"github.com/essentialkaos/ek/v13/usage/completion/zsh"
	"github.com/essentialkaos/ek/v13/usage/man"
	"github.com/essentialkaos/ek/v13/usage/update"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// App info
const (
	APP  = "MDToc"
	VER  = "1.2.10"
	DESC = "Utility for generating table of contents for markdown files"
)

// Options
const (
	OPT_MIN_LEVEL = "m:min-level"
	OPT_MAX_LEVEL = "M:max-level"
	OPT_FLAT      = "f:flat"
	OPT_HTML      = "H:html"
	OPT_NO_COLOR  = "nc:no-color"
	OPT_HELP      = "h:help"
	OPT_VER       = "v:version"

	OPT_VERB_VER     = "vv:verbose-version"
	OPT_COMPLETION   = "completion"
	OPT_GENERATE_MAN = "generate-man"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Header contains info about header
type Header struct {
	Level int    // Header level 1-7
	Text  string // Header text
	Link  string // Link
}

// ////////////////////////////////////////////////////////////////////////////////// //

var optMap = options.Map{
	OPT_MIN_LEVEL: {Type: options.INT, Value: 1, Min: 1, Max: 6},
	OPT_MAX_LEVEL: {Type: options.INT, Value: 6, Min: 1, Max: 6},
	OPT_FLAT:      {Type: options.BOOL},
	OPT_HTML:      {Type: options.BOOL},
	OPT_NO_COLOR:  {Type: options.BOOL},
	OPT_HELP:      {Type: options.BOOL},
	OPT_VER:       {Type: options.MIXED},

	OPT_VERB_VER:     {Type: options.BOOL},
	OPT_COMPLETION:   {},
	OPT_GENERATE_MAN: {Type: options.BOOL},
}

var (
	anchorRegExp = regexp.MustCompile(`[\s\d\w-]`)
	badgeRegExp  = regexp.MustCompile(`\[!\[[^\]]*\]\((.*?)\s*("(?:.*[^"])")?\s*\)\]\((.*?)\s*("(?:.*[^"])")?\s*\)`)
	linkRegExp   = regexp.MustCompile(`\[([^\]]+)\]\([^\)]+\)`)
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Init is main function
func Init(gitRev string, gomod []byte) {
	preConfigureUI()

	args, errs := options.Parse(optMap)

	if !errs.IsEmpty() {
		terminal.Error("Options parsing errors:")
		terminal.Error(errs.Error(" - "))
		os.Exit(1)
	}

	configureUI()

	switch {
	case options.Has(OPT_COMPLETION):
		os.Exit(printCompletion())
	case options.Has(OPT_GENERATE_MAN):
		printMan()
		os.Exit(0)
	case options.GetB(OPT_VER):
		genAbout(gitRev).Print(options.GetS(OPT_VER))
		os.Exit(0)
	case options.GetB(OPT_VERB_VER):
		support.Collect(APP, VER).
			WithRevision(gitRev).
			WithDeps(deps.Extract(gomod)).
			Print()
		os.Exit(0)
	case options.GetB(OPT_HELP):
		genUsage().Print()
		os.Exit(0)
	}

	var file string

	if len(args) == 0 {
		file = findProperReadme()

		if file == "" {
			genUsage().Print()
			os.Exit(0)
		}
	} else {
		file = args.Get(0).Clean().String()
	}

	checkFile(file)
	process(file)
}

// preConfigureUI preconfigures UI based on information about user terminal
func preConfigureUI() {
	if !tty.IsTTY() {
		fmtc.DisableColors = true
	}
}

// configureUI configures user interface
func configureUI() {
	if options.GetB(OPT_NO_COLOR) {
		fmtc.DisableColors = true
	}

	fmtutil.SeparatorFullscreen = true
	fmtutil.SeparatorSymbol = "–"
	fmtutil.SeparatorColorTag = "{s-}"
}

// findProperReadme tries to find readme file in current directory
func findProperReadme() string {
	file := fsutil.ProperPath("FRS", []string{"README.md", "readme.md"})
	return file
}

// checkFile checks markdown file before processing
func checkFile(file string) {
	err := fsutil.ValidatePerms("FRS", file)

	if err != nil {
		printErrorAndExit(err.Error())
	}
}

// process starts file processing
func process(file string) {
	headers := extractHeaders(file)

	if len(headers) == 0 {
		terminal.Warn("Headers not found in given file")
		return
	}

	printTOC(headers)
}

// extractHeaders extracts headers from markdown file
func extractHeaders(file string) []*Header {
	fd, err := os.Open(file)

	if err != nil {
		printErrorAndExit("File reading error: %v", err)
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

	return headers
}

// printTOC collects headers and print ToC for given markdown file
func printTOC(headers []*Header) {
	var toc string

	switch {
	case !options.GetB(OPT_FLAT):
		toc = renderTOC(headers)
	case options.GetB(OPT_FLAT) && options.GetB(OPT_HTML):
		toc = renderFlatHTMLTOC(headers)
	case options.GetB(OPT_FLAT) && !options.GetB(OPT_HTML):
		toc = renderFlatTOC(headers)
	}

	if toc == "" {
		terminal.Warn("Suitable headers not found in given file")
		return
	}

	fmtutil.Separator(false)
	fmtc.Println(toc)
	fmtutil.Separator(false)
}

// renderTOC renders headers as default (vertical) markdown ToC
func renderTOC(headers []*Header) string {
	var toc []string

	minLevel := getMinLevel(headers)

	for _, header := range headers {
		if !isSuitableHeader(header) {
			continue
		}

		toc = append(toc, fmtc.Sprintf(
			"%s [%s](%s)",
			getMarkdownListPrefix(header.Level, minLevel),
			header.Text, header.Link,
		))
	}

	return strings.Join(toc, "\n")
}

// renderFlatTOC renders headers as flat (horizontal) markdown ToC
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

// renderFlatTOC renders headers as flat (horizontal) HTML ToC
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

// isSuitableHeader returns true if header complies defined levels
func isSuitableHeader(header *Header) bool {
	if header.Level < options.GetI(OPT_MIN_LEVEL) || header.Level > options.GetI(OPT_MAX_LEVEL) {
		return false
	}

	return true
}

// parseHeader parses header text and return header struct
func parseHeader(text string) *Header {
	header := &Header{}

	headerText := strings.TrimRight(text, " ")
	headerText = formatHeader(headerText)

	header.Text, header.Level = parseHeaderText(headerText)
	header.Link = makeLink(headerText)

	return header
}

// makeLink converts header text to anchor link name
func makeLink(text string) string {
	result := text

	result = strings.TrimLeft(result, "# ")
	result = strings.ReplaceAll(result, " ", "-")
	result = strings.ToLower(result)
	result = strings.Join(anchorRegExp.FindAllString(result, -1), "")

	return "#" + result
}

// parseHeaderText parses text and return level and header
func parseHeaderText(text string) (string, int) {
	level := strutil.PrefixSize(text, '#')
	header := strings.TrimLeft(text, "# ")
	header = strings.TrimRight(header, " ")
	header = removeMarkdownTags(header)

	return header, level
}

// removeMarkdownTags removes markdown tags from header
func removeMarkdownTags(header string) string {
	for _, r := range "`_*~" {
		if strings.Count(header, string(r))%2 == 0 {
			header = strings.ReplaceAll(header, string(r), "")
		}
	}

	return header
}

// getMarkdownListPrefix returns list prefix for given level
func getMarkdownListPrefix(level, minLevel int) string {
	return strings.Repeat("  ", level-minLevel) + "*"
}

// getMinLevel returns minimal header level
func getMinLevel(headers []*Header) int {
	result := 6

	for _, header := range headers {
		if !isSuitableHeader(header) {
			continue
		}

		result = mathutil.Min(result, header.Level)
	}

	return result
}

// formatHeader formats header text removing badges and links
func formatHeader(text string) string {
	text = badgeRegExp.ReplaceAllString(text, "")

	for _, repl := range linkRegExp.FindAllStringSubmatch(text, -1) {
		text = strings.ReplaceAll(text, repl[0], repl[1])
	}

	return removeMarkdownTags(text)
}

// printErrorAndExit prints error message and exit with exit code 1
func printErrorAndExit(f string, a ...interface{}) {
	terminal.Error(f, a...)
	os.Exit(1)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// printCompletion prints completion for given shell
func printCompletion() int {
	info := genUsage()

	switch options.GetS(OPT_COMPLETION) {
	case "bash":
		fmt.Print(bash.Generate(info, "mdtoc"))
	case "fish":
		fmt.Print(fish.Generate(info, "mdtoc"))
	case "zsh":
		fmt.Print(zsh.Generate(info, optMap, "mdtoc"))
	default:
		return 1
	}

	return 0
}

// printMan prints man page
func printMan() {
	fmt.Println(man.Generate(genUsage(), genAbout("")))
}

// genUsage generates usage info
func genUsage() *usage.Info {
	info := usage.NewInfo("", "file")

	info.AddOption(OPT_FLAT, "Print flat (horizontal) ToC")
	info.AddOption(OPT_HTML, "Render HTML ToC instead Markdown (works with {g}--flat{!})")
	info.AddOption(OPT_MIN_LEVEL, "Minimal header level", "1-6")
	info.AddOption(OPT_MAX_LEVEL, "Maximum header level", "1-6")
	info.AddOption(OPT_NO_COLOR, "Disable colors in output")
	info.AddOption(OPT_HELP, "Show this help message")
	info.AddOption(OPT_VER, "Show version")

	info.AddExample("readme.md", "Generate table of contents for readme.md")
	info.AddExample("-m 2 -M 4 readme.md", "Generate table of contents for readme.md with 2-4 level headers")

	return info
}

// genAbout generates info about version
func genAbout(gitRev string) *usage.About {
	about := &usage.About{
		App:     APP,
		Version: VER,
		Desc:    DESC,
		Year:    2006,
		Owner:   "ESSENTIAL KAOS",
		License: "Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>",
	}

	if gitRev != "" {
		about.Build = "git:" + gitRev
		about.UpdateChecker = usage.UpdateChecker{"essentialkaos/mdtoc", update.GitHubChecker}
	}

	return about
}
