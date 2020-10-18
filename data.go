package metaz

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Logger interface {
	LogServiceMetadata(Data)
}

type Servicer interface {
	_ServiceMetadata() (metadata Data, err error)
}

const ServiceName = "_metadata"

type ReadService interface {
	Read() (metadata Data, err error)
}

// Data stores metadata for the service
type Data struct {
	Name      string `json:"name" yaml:"name" ini:"name"`
	Version   string `json:"version" yaml:"version" ini:"version"`
	Commit    string `json:"commit" yaml:"commit" ini:"commit"`
	UpdatedAt string `json:"updatedAt" yaml:"updatedAt ini:"updated-at"`

	// httpContentType for ServeHTTP. Options text or json (default json).
	httpContentType string
}

func (d Data) Log(logger Logger) {
	logger.LogServiceMetadata(d)
}

type ServeAsFlag int

const (
	PlainText ServeAsFlag = iota + 1
	JSON
)

// ServeAs set content type for ServeHTTP
func (d Data) ServeAs(as ServeAsFlag) Data {
	if as == PlainText {
		d.httpContentType = "text/plain"
	} else {
		d.httpContentType = "application/json"
	}

	return d
}

func (d Data) fprintf(w io.Writer) {
	fmt.Fprintf(w,
		"Name: %s\nVersion: %s\nCommit: %s\nUpdated at: %s\n",
		d.Name, d.Version, d.Commit, d.UpdatedAt,
	)
}

func (d Data) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if d.httpContentType == "" || d.httpContentType == "application/json" {
		w.Header().Add("Content-Type", d.httpContentType)
		_ = json.NewEncoder(w).Encode(d)
		return
	}

	d.fprintf(w)
}

func (d Data) _ServiceMetadata() (metadata Data, err error) {
	return d, nil
}

func (d Data) Read() (metadata Data, err error) {
	return d, nil
}

func (d Data) Print() {
	d.fprintf(os.Stdout)
}

func (d Data) String() string {
	buf := &bytes.Buffer{}
	d.fprintf(buf)
	return buf.String()
}
