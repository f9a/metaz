package metaz

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestData(t *testing.T) {
	m := Data{
		Name:      "chunky-bacon",
		UpdatedAt: "2020-04-01T43:24:06Z",
		Commit:    "XXX",
		Version:   "1.0.0-yummy",
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	m.ServeHTTP(w, r)

	wantJSON := "{\"name\":\"chunky-bacon\",\"version\":\"1.0.0-yummy\",\"commit\":\"XXX\",\"updatedAt\":\"2020-04-01T43:24:06Z\"}\n"
	assert.Equal(t, wantJSON, w.Body.String())

	m = m.ServeAs(PlainText)

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/", nil)
	m.ServeHTTP(w, r)

	want := "Name: chunky-bacon\nVersion: 1.0.0-yummy\nCommit: XXX\nUpdated at: 2020-04-01T43:24:06Z\n"
	assert.Equal(t, want, w.Body.String())
	assert.Equal(t, want, m.String())

	m = m.ServeAs(JSON)

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/", nil)
	m.ServeHTTP(w, r)

	assert.Equal(t, wantJSON, w.Body.String())
}
