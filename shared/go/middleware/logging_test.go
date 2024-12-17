package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLoggingMiddleware(t *testing.T) {

	//want := `{"aa": "baaa"}`

	logger, buffer := makeLoggerAndBuffer()

	body := bytes.NewReader(
		[]byte(
			`{"name":"jeff"}`,
		),
	)

	rq := httptest.NewRequest(http.MethodGet, "/", body)
	rq.Header.Add(RequestIDHeader, "fake-request-id")
	rs := httptest.NewRecorder()

	LoggingMiddleware(logger, nil)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello")
		logger.Info("hmm?")
		logger.InfoContext(r.Context(), "ok")
	})).ServeHTTP(rs, rq)

	got := strings.Split(strings.TrimRight(buffer.String(), "\n"), "\n")

	for _, str := range got {
		fmt.Println(str)

		if str == "" {
			fmt.Println("?????")
			continue
		}

		var val map[string]interface{}
		if err := json.NewDecoder(strings.NewReader(str)).Decode(&val); err != nil {
			t.Errorf("failed decoding json from log buffer: %v", err)
		}

		if val["msg"] == "Starting Request" {
			fmt.Println("a")
			continue
		}
		if val["msg"] == "Finished Request" {
			fmt.Println("b")
			continue
		}

		fmt.Println("c")
	}

}
