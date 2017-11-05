package reverseproxy

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/favclip/testerator"
	"google.golang.org/appengine/aetest"
)

var inst aetest.Instance

func TestMain(m *testing.M) {
	var err error
	defer func() {
		if err = testerator.SpinDown(); err != nil {
			panic(err)
		}
	}()

	inst, _, err = testerator.SpinUp()
	if err != nil {
		panic(err)
	}
	m.Run()
}

func TestReverseProxy(t *testing.T) {
	message := "hello 世界"
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(message))
	}))
	backendUrl, err := url.Parse(backend.URL)
	if err != nil {
		t.Fatal(err)
	}
	frontend := NewSingleHostReverseProxy(backendUrl)
	req, err := inst.NewRequest("POST", backend.URL, bytes.NewBuffer([]byte{}))
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	frontend.ServeHTTP(rec, req)
	if have, want := rec.Code, http.StatusOK; have != want {
		t.Errorf("invalid status code, want = %d, have = %d", want, have)
	}
	if have, want := rec.Body.String(), message; have != want {
		t.Errorf("invalid body, want = '%s', have = '%s'", want, have)
	}
}
