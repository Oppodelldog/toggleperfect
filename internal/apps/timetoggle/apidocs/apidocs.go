package apidocs

import (
	"github.com/go-playground/statics/static"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

var swaggerUIFiles *static.Files

func init() {
	var err error

	cfg := &static.Config{
		UseStaticFiles: true,
		FallbackToDisk: false,
	}

	swaggerUIFiles, err = newStaticDist(cfg)
	if err != nil {
		panic(err)
	}
}

func NewHandler() http.Handler {
	return http.FileServer(replacementFS{wrappedFS: swaggerUIFiles.FS()})
}

type replacementFS struct {
	wrappedFS http.FileSystem
}

func (r replacementFS) Open(name string) (http.File, error) {
	f, err := r.wrappedFS.Open(name)
	if err != nil {
		return nil, err
	}
	if strings.HasSuffix(name, "swagger.yml") {
		if host, ok := os.LookupEnv("TOGGLE_PERFECT_API_DOCS_HOST"); ok {
			return replaceFileContent(f, "${TOGGLE_PERFECT_API_DOCS_HOST}", host)
		}
	}
	return f, err
}

func replaceFileContent(f http.File, what, with string) (http.File, error) {
	content, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	originalFileInfo, err := f.Stat()
	if err != nil {
		return nil, err
	}
	replacedContent := strings.Replace(string(content), what, with, -1)
	return ContentReplacedHttpFile{
		originalFile:        f,
		originalFileContent: content,
		replacedBuffer:      strings.NewReader(replacedContent),
		StatWrap:            StatWrap{original: originalFileInfo, newContentLength: int64(len(replacedContent))},
	}, nil
}

type StatWrap struct {
	original         os.FileInfo
	newContentLength int64
}

func (s StatWrap) Name() string {
	return s.original.Name()
}

func (s StatWrap) Size() int64 {
	return s.newContentLength
}

func (s StatWrap) Mode() os.FileMode {
	return s.original.Mode()
}

func (s StatWrap) ModTime() time.Time {
	return s.original.ModTime()
}

func (s StatWrap) IsDir() bool {
	return s.original.IsDir()
}

func (s StatWrap) Sys() interface{} {
	return s.original.Sys()
}

type ContentReplacedHttpFile struct {
	originalFile        http.File
	originalFileContent []byte
	replacedBuffer      *strings.Reader
	StatWrap            os.FileInfo
}

func (r ContentReplacedHttpFile) Close() error {
	return r.originalFile.Close()
}

func (r ContentReplacedHttpFile) Read(p []byte) (n int, err error) {
	return r.replacedBuffer.Read(p)
}

func (r ContentReplacedHttpFile) Seek(offset int64, whence int) (int64, error) {
	return r.replacedBuffer.Seek(offset, whence)
}

func (r ContentReplacedHttpFile) Readdir(count int) ([]os.FileInfo, error) {
	return r.originalFile.Readdir(count)
}

func (r ContentReplacedHttpFile) Stat() (os.FileInfo, error) {
	return r.StatWrap, nil
}
