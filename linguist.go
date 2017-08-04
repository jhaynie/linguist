package linguist

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
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
)

var authtoken, linguisturl string

func init() {
	authtoken = getEnv("PP_LINGUIST_AUTH", "1234")
	urlprefix := getEnv("PP_LINGUIST_URL", "https://linguist:25032")
	linguisturl = fmt.Sprintf("%s/detect", urlprefix)
}

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
	Loc                    int       `json:"loc,omitempty"`
	Sloc                   int       `json:"sloc,omitempty"`
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
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Results []Detection `json:"results"`
}

type preoptimization struct {
	Pattern   *regexp.Regexp
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
	client = &http.Client{Transport: transport, Timeout: time.Second * 30}
	mutex  = sync.Mutex{}
)

func preoptimize(re string, filename string, body string) {
	result, err := getLanguageDetails(context.Background(), filename, []byte(body))
	if err == nil && result.Success {
		p := &preoptimization{
			Pattern: regexp.MustCompile(re),
			Result:  result,
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

// MostPopular returns the most popular language based on cache hits since the worker has started
func MostPopular() Detection {
	resort()
	// for i, r := range preoptimizations {
	// 	fmt.Printf("%d %d %v\n", i, r.CacheHits, r.Result.Results[0].Language.Name)
	// }
	return preoptimizations[0].Result.Results[0]
}

// Initialize will warm up the preoptimization cache
func Initialize() {
	preoptimizeInit()
}

// initialize a pre-optimization cache for well-known languages to speed up
// calculating predictable language results
func preoptimizeInit() {
	if preoptimized == false {
		preoptimized = true
		preoptimize("\\.js$", "test.js", "var a")
		preoptimize("\\.ts$", "test.ts", "interface Foo {\n}")
		preoptimize("\\.ejs$", "test.ejs", "<% if (names.length) { %>foo<% } %>")
		preoptimize("\\.go$", "test.go", "package main\nfunc main(){\n}\n")
		preoptimize("Makefile$", "Makefile", ".phony foo\n")
		preoptimize("\\.ya?ml$", "test.yml", "---\nfoo: 1\n")
		preoptimize("\\.json$", "test.json", "{\"a\":1}")
		preoptimize("\\.swift$", "test.swift", "let a=0")
		preoptimize("\\.c(\\+\\+|pp|c)$", "test.cpp", "class Foo{\n};\n")
		preoptimize("\\.hbs$", "test.hbs", "<div>{{foo}}</div>")
		preoptimize("\\.html$", "test.html", "<div>hi</div>")
		preoptimize("\\.css$", "test.css", ".rule {color:red}")
		preoptimize("\\.scss$", "test.scss", ".rule {color:red}")
		preoptimize("\\.(ba|z)?sh$", "test.sh", "#!/bin/sh\n")
		preoptimize("\\.md$", "test.md", "# Foo\n")
		preoptimize("\\.json5$", "test.json5", "{a:1}")
		preoptimize("\\.jsx$", "test.jsx", "import a from 'foo'\n")
		preoptimize("\\.m$", "test.m", "@implementation Foo\n@end\n")
		preoptimize("\\.mm$", "test.mm", "@implementation Foo\n@end\n")
		preoptimize("\\.(c|h)$", "test.c", "void main(){\n}\n")
		preoptimize("\\.rb$", "test.rb", "print \"hello\"")
		preoptimize("\\.py$", "test.py", "def foo\nend\n")
		preoptimize("\\.proto$", "test.proto", "package foo\nmessage Bar\n{\n}\n")
		preoptimize("\\.java$", "test.java", "package foo\npublic class Bar\n{\n}\n")
		preoptimize("\\.cs$", "test.cs", "class Bar\n{\n}\n")
		preoptimize("\\.xml$", "test.xml", "<a>foo</a>")
		preoptimize("\\.lua$", "test.lua", "x=0")
		preoptimize("\\.txt$", "test.txt", "hi")
		preoptimize("\\.sql$", "test.sql", "delete from foo")
		preoptimize("\\.coffee$", "test.coffee", "a = 1")
		preoptimize("\\.properties$", "test.properties", "a=1")
		preoptimize("Dockerfile(\\.*)$", "Dockerfile", "FROM nodejs\n")
		preoptimize("LICENSE$", "LICENSE", "MIT License\n")
		// reset after loading.
		atomic.StoreInt32(&cacheMisses, 0)
		atomic.StoreInt32(&cacheHits, 0)
	}
}

func checkPreoptimizationCache(filename string) Result {
	mutex.Lock()
	for _, p := range preoptimizations {
		if p.Pattern.MatchString(filename) {
			p.Result.Results[0].Path = filename
			p.Result.Results[0].Sloc = 0
			p.Result.Results[0].Loc = 0
			atomic.AddInt32(&p.CacheHits, 1)
			mutex.Unlock()
			return p.Result
		}
	}
	mutex.Unlock()
	return Result{Success: false}
}

var concurrent int32

var noResult = Result{true, "", nil}

// GetLanguageDetails returns the linguist results for a given file
func GetLanguageDetails(ctx context.Context, filename string, body []byte) (Result, error) {
	if isExcluded(filename, body) {
		return noResult, nil
	}
	if preop := checkPreoptimizationCache(filename); preop.Success {
		hits := atomic.AddInt32(&cacheHits, 1)
		// log.Debug("linguist cache hit [%05d/%05d]", hits, CacheMisses())
		// every N hits, resort so that the most popular stays
		// at the top of the heap for faster access and less popular go to bottom
		if hits%100 == 0 {
			resort()
		}
		// since we mutate this from a separate thread we need to make a copy
		// so that we can return and not have a race when the caller accesses
		c := make([]Detection, len(preop.Results))
		mutex.Lock()
		copy(c, preop.Results)
		mutex.Unlock()
		return Result{Success: true, Results: c}, nil
	}
	// count := atomic.AddInt32(&concurrent, 1)
	atomic.AddInt32(&concurrent, 1)
	defer func() { atomic.AddInt32(&concurrent, -1) }()
	// log.Debug("linguist cache miss [%05d/%05d]", CacheHits(), CacheMisses(), log.WithFields("file", filename, "concurrent", count))
	result, err := getLanguageDetails(ctx, filename, body)
	if result.Success {
		atomic.AddInt32(&cacheMisses, 1)
	}
	// since we mutate this from a separate thread we need to make a copy
	// so that we can return and not have a race when the caller accesses
	c := make([]Detection, len(result.Results))
	mutex.Lock()
	copy(c, result.Results)
	mutex.Unlock()
	return Result{Success: result.Success, Results: c, Message: result.Message}, err
}

var failed = Result{Success: false}

func attempt(ctx context.Context, jsonbuf string, url string, authtoken string, attempts int) (Result, error) {
	if attempts > 10 {
		return failed, fmt.Errorf("error attempting to load %s after %d attempts", url, attempts)
	}
	_req, err := http.NewRequest("POST", url, strings.NewReader(jsonbuf))
	if err != nil {
		return failed, err
	}
	req := _req.WithContext(ctx)
	if authtoken != "" {
		req.Header.Set("Authorization", authtoken)
	}
	resp, err := client.Do(req)
	if err != nil {
		es := err.Error()
		if strings.Contains(es, "connection reset by peer") || strings.Contains(es, "EOF") {
			time.Sleep(time.Millisecond * time.Duration(50*attempts+1))
			return attempt(ctx, jsonbuf, url, authtoken, attempts+1)
		}
		return failed, err
	}
	defer resp.Body.Close()
	result := Result{}
	d := json.NewDecoder(resp.Body)
	d.UseNumber() // prevent numbers from getting converted
	err = d.Decode(&result)
	if err != nil {
		return failed, err
	}
	if result.Success {
		return result, nil
	}
	return failed, errors.New(result.Message)
}

func isLikelyBinary(body []byte) bool {
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
	return false
}

const maxBufferSize = 10000

func isLargeBuffer(body []byte) bool {
	return len(body) > maxBufferSize
}

var excludeExtensions = map[string]bool{
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
}

var excludedFilenames = map[string]bool{
	"npm-debug.log": true,
	"LICENSE":       true,
	"LICENSE.md":    true,
}

func isFilenameExcluded(name string) bool {
	return excludedFilenames[filepath.Base(name)] || excludeExtensions[filepath.Ext(name)]
}

func isExcluded(filename string, body []byte) bool {
	if isLikelyBinary(body) || isLargeBuffer(body) || isFilenameExcluded(filename) {
		return true
	}
	return false
}

func getLanguageDetails(ctx context.Context, filename string, body []byte) (Result, error) {
	jsonbody := []interface{}{map[string]string{
		"name": filename,
		"body": string(body),
	}}
	return attempt(ctx, stringify(jsonbody), linguisturl, authtoken, 1)
}
