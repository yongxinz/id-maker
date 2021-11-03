package integration_test

import (
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	. "github.com/Eun/go-hit"
)

const (
	// Attempts connection
	host       = "app:8080"
	healthPath = "http://" + host + "/healthz"
	attempts   = 20

	// HTTP REST
	basePath = "http://" + host + "/v1"
)

func TestMain(m *testing.M) {
	err := healthCheck(attempts)
	if err != nil {
		log.Fatalf("Integration tests: host %s is not available: %s", host, err)
	}

	log.Printf("Integration tests: host %s is available", host)

	code := m.Run()
	os.Exit(code)
}

func healthCheck(attempts int) error {
	var err error

	for attempts > 0 {
		err = Do(Get(healthPath), Expect().Status().Equal(http.StatusOK))
		if err == nil {
			return nil
		}

		log.Printf("Integration tests: url %s is not available, attempts left: %d", healthPath, attempts)

		time.Sleep(time.Second)

		attempts--
	}

	return err
}

// HTTP POST: /tag.
func TestHTTPCreateTag(t *testing.T) {
	body := `{
		"biz_tag": "test",
		"step": 100,
	}`
	Test(t,
		Description("CreateTag Success"),
		Post(basePath+"/tag"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusOK),
		// Expect().Body().JSON().JQ(".translation").Equal("text for translation"),
	)

	body = `{
		"biz_tag": "test",
	}`
	Test(t,
		Description("CreateTag Fail"),
		Post(basePath+"/tag"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusBadRequest),
		Expect().Body().JSON().JQ(".error").Equal("invalid request body"),
	)
}

// HTTP GET: /id/test.
func TestHTTPGetId(t *testing.T) {
	Test(t,
		Description("GetId Success"),
		Get(basePath+"/id/test"),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().String().Contains(`{"id":`),
	)
}
