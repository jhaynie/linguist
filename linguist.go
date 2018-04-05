package linguist

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	generaltso "github.com/jhaynie/linguist/generaltso/linguist"
)

func getEnv(name, def string) string {
	v := os.Getenv(name)
	if v == "" {
		return def
	}
	return v
}

func stringify(v interface{}) string {
	buf, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("<error:%v>", err)
	}
	return string(buf)
}

// Language represents the language details that were detected
type Language struct {
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
	Group       string `json:"group,omitempty"`
	AceMode     string `json:"ace_mode,omitempty"`
	IsPopular   bool   `json:"is_popular,omitempty"`
	IsUnpopular bool   `json:"is_unpopular,omitempty"`
}

// Detection represents a language detection result
type Detection struct {
	Path                   string    `json:"path,omitempty"`
	Type                   string    `json:"type,omitempty"`
	ExtName                string    `json:"extname,omitempty"`
	MimeType               string    `json:"mime_type,omitempty"`
	ContentType            string    `json:"content_type,omitempty"`
	Disposition            string    `json:"disposition,omitempty"`
	IsDocumentation        bool      `json:"is_documentation,omitempty"`
	IsLarge                bool      `json:"is_large,omitempty"`
	IsGenerated            bool      `json:"is_generated,omitempty"`
	IsText                 bool      `json:"is_text,omitempty"`
	IsImage                bool      `json:"is_image,omitempty"`
	IsBinary               bool      `json:"is_binary,omitempty"`
	IsVendored             bool      `json:"is_vendored,omitempty"`
	IsHighRatioOfLongLines bool      `json:"is_high_ratio_of_long_lines,omitempty"`
	IsViewable             bool      `json:"is_viewable,omitempty"`
	IsSafeToColorize       bool      `json:"is_safe_to_colorize,omitempty"`
	Language               *Language `json:"language,omitempty"`
}

// Result is the result details of a detection
type Result struct {
	Success    bool       `json:"success"`
	Message    string     `json:"message,omitempty"`
	Result     *Detection `json:"result"`
	IsBinary   bool       `json:"binary"`
	IsLarge    bool       `json:"large"`
	IsExcluded bool       `json:"excluded"`
	IsCached   bool       `json:"cached"`
}

// String returns a string representation
func (r Result) String() string {
	return fmt.Sprintf("Result<success:%v,message:%v,result:%v,binary:%v,large:%v,excluded:%v,cached:%v>", r.Success, r.Message, r.Result, r.IsBinary, r.IsLarge, r.IsExcluded, r.IsCached)
}

// LResult is the result that comes back from linguist
type LResult struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Results []Detection `json:"results"`
}

type preoptimization struct {
	Matchers  []Match
	Result    Result
	CacheHits int32
}

var (
	preoptimizations = make([]*preoptimization, 0)
	cacheMisses      int32
	cacheHits        int32
	preoptimized     bool
	transport        = &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		TLSHandshakeTimeout: 5 * time.Second,
		MaxIdleConnsPerHost: 50,
	}
	client   = &http.Client{Transport: transport, Timeout: time.Second * 30}
	mutex    = sync.RWMutex{}
	noResult = Result{}
)

func preoptimize(re Match, filename string, body string, rules ...Match) {
	result, err := getLanguageDetails(context.Background(), filename, []byte(body))
	if err == nil && result.Success {
		p := &preoptimization{
			Matchers: []Match{re},
			Result:   result,
		}
		if len(rules) > 0 {
			for _, r := range rules {
				p.Matchers = append(p.Matchers, r)
			}
		}
		preoptimizations = append(preoptimizations, p)
	}
}

func resort() {
	mutex.Lock()
	sort.Slice(preoptimizations, func(i, j int) bool {
		return preoptimizations[j].CacheHits < preoptimizations[i].CacheHits
	})
	mutex.Unlock()
}

// CacheHits returns the number of cache hits
func CacheHits() int32 {
	return atomic.LoadInt32(&cacheHits)
}

// CacheMisses returns the number of cache misses
func CacheMisses() int32 {
	return atomic.LoadInt32(&cacheMisses)
}

func cacheCounterReset() {
	atomic.StoreInt32(&cacheHits, 0)
	atomic.StoreInt32(&cacheMisses, 0)
	resort()
}

// MostPopular returns the most popular language based on cache hits since the worker has started
func MostPopular() Detection {
	resort()
	// for i, r := range preoptimizations {
	// 	fmt.Printf("%d %d %v\n", i, r.CacheHits, r.Result.Results[0].Language.Name)
	// }
	return *preoptimizations[0].Result.Result
}

// Initialize will warm up the preoptimization cache
func Initialize() {
	preoptimizeInit()
}

// Match is a simple struct for describing a match rule
type Match struct {
	re     *regexp.Regexp
	invert bool
}

// String is the string representation of Match object
func (m Match) String() string {
	return fmt.Sprintf("Match<%s,%v>", m.re, m.invert)
}

// MatchString will return true if the expression matches s
func (m Match) MatchString(s string) bool {
	matched := m.re.MatchString(s)
	if !m.invert && matched {
		return true
	} else if m.invert && !matched {
		return true
	}
	return false
}

// NewMatcher will create a Matcher with a regular expression
func NewMatcher(s string) Match {
	return Match{regexp.MustCompile(s), false}
}

// NewNotMatcher will create a Matcher that will NOT match the regular expression
func NewNotMatcher(s string) Match {
	return Match{regexp.MustCompile(s), true}
}

// initialize a pre-optimization cache for well-known languages to speed up
// calculating predictable language results
func preoptimizeInit() {
	if preoptimized == false {
		preoptimized = true
		noVendorMatcher := NewNotMatcher("^(node_modules|vendor|Godeps)/")
		preoptimize(NewMatcher("\\.js$"), "test.js", "var a", noVendorMatcher)
		preoptimize(NewMatcher("\\.ts$"), "test.ts", "interface Foo {\n}", noVendorMatcher)
		preoptimize(NewMatcher("\\.ejs$"), "test.ejs", "<% if (names.length) { %>foo<% } %>", noVendorMatcher)
		preoptimize(NewMatcher("\\.go$"), "test.go", "package main\nfunc main(){\n}\n", noVendorMatcher)
		preoptimize(NewMatcher("Makefile$"), "Makefile", ".phony foo\n", noVendorMatcher)
		preoptimize(NewMatcher("\\.ya?ml$"), "test.yml", "---\nfoo: 1\n", noVendorMatcher)
		preoptimize(NewMatcher("\\.json$"), "test.json", "{\"a\":1}", noVendorMatcher)
		preoptimize(NewMatcher("\\.swift$"), "test.swift", "let a=0")
		preoptimize(NewMatcher("\\.c(\\+\\+|pp|c)$"), "test.cpp", "class Foo{\n};\n", noVendorMatcher)
		preoptimize(NewMatcher("\\.hbs$"), "test.hbs", "<div>{{foo}}</div>", noVendorMatcher)
		preoptimize(NewMatcher("\\.html$"), "test.html", "<div>hi</div>", noVendorMatcher)
		preoptimize(NewMatcher("\\.css$"), "test.css", ".rule {color:red}", noVendorMatcher)
		preoptimize(NewMatcher("\\.scss$"), "test.scss", ".rule {color:red}", noVendorMatcher)
		preoptimize(NewMatcher("\\.(ba|z)?sh$"), "test.sh", "#!/bin/sh\n", noVendorMatcher)
		preoptimize(NewMatcher("\\.(md|markdown)$"), "test.md", "# Foo\n## Hello\nthis is a markdown file\n", noVendorMatcher)
		preoptimize(NewMatcher("\\.json5$"), "test.json5", "{a:1}", noVendorMatcher)
		preoptimize(NewMatcher("\\.jsx$"), "test.jsx", "import a from 'foo'\n", noVendorMatcher)
		preoptimize(NewMatcher("\\.ts$"), "test.ts", "import a from 'foo'\n", noVendorMatcher)
		preoptimize(NewMatcher("\\.tsx$"), "test.tsx", "import a from 'foo'\n", noVendorMatcher)
		preoptimize(NewMatcher("\\.m$"), "test.m", "@implementation Foo\n@end\n", noVendorMatcher)
		preoptimize(NewMatcher("\\.mm$"), "test.mm", "@implementation Foo\n@end\n", noVendorMatcher)
		preoptimize(NewMatcher("\\.(c|h)$"), "test.c", "void main(){\n}\n", noVendorMatcher)
		preoptimize(NewMatcher("\\.rb$"), "test.rb", "print \"hello\"")
		preoptimize(NewMatcher("\\.py$"), "test.py", "def foo\nend\n")
		preoptimize(NewMatcher("\\.proto$"), "test.proto", "package foo\nmessage Bar\n{\n}\n", noVendorMatcher)
		preoptimize(NewMatcher("\\.java$"), "test.java", "package foo\npublic class Bar\n{\n}\n")
		preoptimize(NewMatcher("\\.cs$"), "test.cs", "class Bar\n{\n}\n")
		preoptimize(NewMatcher("\\.xml$"), "test.xml", "<a>foo</a>", noVendorMatcher)
		preoptimize(NewMatcher("\\.lua$"), "test.lua", "x=0")
		preoptimize(NewMatcher("\\.txt$"), "test.txt", "hi", noVendorMatcher)
		preoptimize(NewMatcher("\\.sql$"), "test.sql", "-- test\ndelete from `foo`;\nCREATE TABLE IF NOT EXISTS `foo` (l int(11));\n", noVendorMatcher)
		preoptimize(NewMatcher("\\.coffee$"), "test.coffee", "a = 1", noVendorMatcher)
		preoptimize(NewMatcher("\\.properties$"), "test.properties", "a=1", noVendorMatcher)
		preoptimize(NewMatcher("Dockerfile(\\.*)$"), "Dockerfile", "FROM nodejs\n")
		preoptimize(NewMatcher("LICENSE$"), "LICENSE", "MIT License\n", noVendorMatcher)
		// reset after loading.
		atomic.StoreInt32(&cacheMisses, 0)
		atomic.StoreInt32(&cacheHits, 0)
	}
}

// CheckPreoptimizationCache will return a potential Result for a filename match based on the preoptimization cache
func CheckPreoptimizationCache(filename string) Result {
	mutex.RLock()
	for _, p := range preoptimizations {
		var matched bool
		for _, matcher := range p.Matchers {
			if matcher.MatchString(filename) {
				matched = true
			} else {
				break
			}
		}
		if matched {
			ex, r := IsExcluded(filename, nil)
			if ex {
				return *r
			}
			// make a copy so that the result can't be mutated
			l := Language(*p.Result.Result.Language)
			result := Result{
				Success:  true,
				IsCached: true,
				Result: &Detection{
					Path:                   filename,
					Type:                   p.Result.Result.Type,
					ExtName:                p.Result.Result.ExtName,
					MimeType:               p.Result.Result.MimeType,
					ContentType:            p.Result.Result.ContentType,
					Disposition:            p.Result.Result.Disposition,
					IsDocumentation:        p.Result.Result.IsDocumentation,
					IsLarge:                p.Result.Result.IsLarge,
					IsGenerated:            p.Result.Result.IsGenerated,
					IsText:                 p.Result.Result.IsText,
					IsImage:                p.Result.Result.IsImage,
					IsBinary:               p.Result.Result.IsBinary,
					IsVendored:             p.Result.Result.IsVendored,
					IsHighRatioOfLongLines: p.Result.Result.IsHighRatioOfLongLines,
					IsViewable:             p.Result.Result.IsViewable,
					IsSafeToColorize:       p.Result.Result.IsSafeToColorize,
					Language:               &l,
				},
				IsBinary:   p.Result.Result.IsBinary,
				IsLarge:    p.Result.Result.IsLarge,
				IsExcluded: p.Result.Result.IsBinary,
			}
			atomic.AddInt32(&p.CacheHits, 1)
			mutex.RUnlock()
			return result
		}
	}
	mutex.RUnlock()
	return noResult
}

// GetLanguageDetails returns the linguist results for a given file
func GetLanguageDetails(ctx context.Context, filename string, body []byte, skip ...bool) (Result, error) {
	if ex, r := IsExcluded(filename, body); ex {
		return *r, nil
	}
	if len(skip) == 0 || !skip[0] {
		if preop := CheckPreoptimizationCache(filename); preop.Success {
			hits := atomic.AddInt32(&cacheHits, 1)
			// every N hits, resort so that the most popular stays
			// at the top of the heap for faster access and less popular go to bottom
			if hits%100 == 0 {
				resort()
			}
			return preop, nil
		}
	}
	result, err := getLanguageDetails(ctx, filename, body)
	if result.Success {
		atomic.AddInt32(&cacheMisses, 1)
	}
	return result, err
}

// File is a wrapper around a file name and body
type File struct {
	filename string
	body     []byte
}

// NewFile will return a File struct
func NewFile(filename string, body []byte) *File {
	return &File{filename, body}
}

// Filereq is used internally
type Filereq struct {
	Name  string
	Body  []byte
	Index int
}

// GetLanguageDetailsMultiple returns the linguist results for one or more files
func GetLanguageDetailsMultiple(ctx context.Context, files []*File, skipCache ...bool) ([]Result, error) {
	results := make([]Result, 0)
	jsonbody := make([]Filereq, 0)
	var skip bool
	if len(skipCache) != 0 && skipCache[0] {
		skip = true
	}
	for i, file := range files {
		if ex, r := IsExcluded(file.filename, file.body); ex {
			results = append(results, *r)
			continue
		}
		if !skip {
			if preop := CheckPreoptimizationCache(file.filename); preop.Success {
				hits := atomic.AddInt32(&cacheHits, 1)
				// every N hits, resort so that the most popular stays
				// at the top of the heap for faster access and less popular go to bottom
				if hits%100 == 0 {
					resort()
				}
				results = append(results, preop)
				continue
			}
		}
		jsonbody = append(jsonbody, Filereq{file.filename, file.body, i})
		results = append(results, Result{})
	}
	if len(jsonbody) == 0 {
		return results, nil
	}
	for _, j := range jsonbody {
		r, err := getLanguageDetails(ctx, j.Name, j.Body)
		if err != nil {
			return nil, err
		}
		results[j.Index] = r
	}
	return results, nil
}

// IsLikelyBinary returns true if the body is likely a binary buffer
func IsLikelyBinary(body []byte) bool {
	ct := http.DetectContentType(body)
	if strings.HasPrefix(ct, "image/") || strings.HasPrefix(ct, "video/") {
		return true
	}
	switch ct {
	case "application/octet-stream", "application/pdf", "application/ogg",
		"application/x-rar-compressed", "application/zip", "application/x-gzip":
		{
			return true
		}
	}
	return generaltso.IsBinary(body)
}

// MaxBufferSize is the large size in bytes that a buffer can be before it's considered "large"
const MaxBufferSize = 100000

// IsLargeBuffer returns true if the size is larger than MaxBufferSize
func IsLargeBuffer(size int) bool {
	return size > MaxBufferSize
}

func isFilenameExcluded(name string) bool {
	if excludedFilenames[filepath.Base(name)] || excludeExtensions[filepath.Ext(name)] {
		return true
	}
	for _, rule := range excludedRules {
		if rule.MatchString(name) {
			return true
		}
	}
	return false
}

var (
	languageOverrides = map[string]map[string]string{
		"PLpgSQL": map[string]string{".sql": "SQL"},
	}
	excludeExtensions = map[string]bool{
		".swp":           true,
		".DS_Store":      true,
		".winmd":         true,
		".node":          true,
		".dll":           true,
		".a":             true,
		".lib":           true,
		".dylib":         true,
		".exe":           true,
		".gif":           true,
		".png":           true,
		".webp":          true,
		".svg":           true,
		".sketch":        true,
		".eps":           true,
		".pdf":           true,
		".psd":           true,
		".tif":           true,
		".tiff":          true,
		".bmp":           true,
		".ico":           true,
		".raw":           true,
		".wav":           true,
		".mpg":           true,
		".mpeg":          true,
		".mp3":           true,
		".mp4":           true,
		".3gp":           true,
		".aac":           true,
		".m4a":           true,
		".ogg":           true,
		".wma":           true,
		".avi":           true,
		".ppt":           true,
		".doc":           true,
		".docx":          true,
		".zip":           true,
		".zipx":          true,
		".cab":           true,
		".7z":            true,
		".bkf":           true,
		".dmg":           true,
		".lz":            true,
		".rar":           true,
		".iso":           true,
		".lzma":          true,
		".tar":           true,
		".tgz":           true,
		".bz2":           true,
		".gz":            true,
		".gzip":          true,
		".jar":           true,
		".ear":           true,
		".aar":           true,
		".class":         true,
		".pbxproj":       true,
		".xcworkspace":   true,
		".nib":           true,
		".xib":           true,
		".plist":         true,
		".pyc":           true,
		".gitignore":     true,
		".gitmodules":    true,
		".gitattributes": true,
		".npmignore":     true,
		".lock":          true,
		".npmrc":         true,
		".babelrc":       true,
		".jshintrc":      true,
		".eslintrc":      true,
		".eslintrc.json": true,
		".eslintignore":  true,
		".editorconfig":  true,
		".flowconfig":    true,
	}
	excludedFilenames = map[string]bool{
		".travis.yml":                true,
		"npm-debug.log":              true,
		"package-lock.json":          true,
		"package.json":               true,
		".eslintrc.js":               true,
		"postcss.config.js":          true,
		"jest.config.json":           true,
		"jest-preset.json":           true,
		"webpack.js":                 true,
		"webpack.config.js":          true,
		"webpack.config.dev.js":      true,
		"webpack.config.prod.js":     true,
		"webpackDevServer.config.js": true,
		"bower.json":                 true,
		"AUTHORS":                    true,
		"AUTHORS.md":                 true,
		"PATENTS":                    true,
		"license":                    true,
		"LICENSE":                    true,
		"LICENSE.md":                 true,
		"VERSION":                    true,
		"PULL_REQUEST_TEMPLATE.md":   true,
		"glide.yaml":                 true,
		"Gopkg.lock":                 true,
		"Gopkg.toml":                 true,
		"rollup.config.js":           true,
		"appveyor.yml":               true,
	}
	excludedRules = []Match{
		NewMatcher("^(\\.github|\\.vscode)\\/"),
		NewMatcher("(node_modules|vendor|Godeps)\\/"),
		NewMatcher("\\.min\\.js$"),     // minimized JS
		NewMatcher("\\.js\\.map$"),     // JS sourcemap
		NewMatcher("^dist/(.*)\\.js$"), // generated JS files
	}
	binaryResult   = &Result{true, "", nil, true, false, true, false}
	largeResult    = &Result{true, "", nil, false, true, true, false}
	excludedResult = &Result{true, "", nil, false, false, true, false}
)

// AddExcludedRule will add a rule to the exclusions list
func AddExcludedRule(match Match) {
	excludedRules = append(excludedRules, match)
}

// AddExcludedFilename will add a filename rule to be excluded
func AddExcludedFilename(filename string) {
	excludedFilenames[filename] = true
}

// AddExcludedExtension will add extension to the exclusion list
func AddExcludedExtension(ext string) {
	excludeExtensions[ext] = true
}

// RemoveExcludedExtension will remove the extension as an exclusion rule
func RemoveExcludedExtension(ext string) {
	delete(excludeExtensions, ext)
}

// RemoveExcludedFilename will remove the filename as an exclusion rule
func RemoveExcludedFilename(filename string) {
	delete(excludedFilenames, filename)
}

// RemoveExcludedRule will remove the added match from the exclusion rule
func RemoveExcludedRule(match Match) {
	for i, m := range excludedRules {
		if match == m {
			excludedRules[i] = excludedRules[len(excludedRules)-1]
			excludedRules = excludedRules[:len(excludedRules)-1]
			break
		}
	}
}

// IsExcluded returns true if the filename and optional body is excluded. If nil body, will only check for filename
func IsExcluded(filename string, body []byte) (bool, *Result) {
	if body != nil {
		if IsLikelyBinary(body) {
			return true, binaryResult
		}
		if IsLargeBuffer(len(body)) {
			return true, largeResult
		}
	}
	if isFilenameExcluded(filename) {
		return true, excludedResult
	}
	return false, nil
}

func getLanguageDetails(ctx context.Context, filename string, body []byte) (Result, error) {
	hints := generaltso.LanguageHints(filename)
	language := generaltso.LanguageByContents(body, hints)
	// see if we have any language rule overrides
	kv := languageOverrides[language]
	if kv != nil {
		l := kv[filepath.Ext(filename)]
		if l != "" {
			language = l
		}
	}
	binary := IsLikelyBinary(body)
	vendored := generaltso.IsVendored(filename)
	generated := false
	large := IsLargeBuffer(len(body))
	excluded := binary || vendored || generated
	return Result{
		Success:    true,
		IsBinary:   binary,
		IsExcluded: excluded,
		IsLarge:    large,
		Result: &Detection{
			Path: filename,
			Type: "text",
			Language: &Language{
				Name: language,
			},
			IsLarge:     large,
			IsBinary:    binary,
			IsGenerated: generated,
			IsVendored:  vendored,
		},
	}, nil
}
