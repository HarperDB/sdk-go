package harper

import (
	"crypto/tls"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
)

const (
	DEFAULT_ENDPOINT = "http://localhost:9925"
	DEFAULT_USERNAME = "HDB_ADMIN"
	DEFAULT_PASSWORD = "password"
)

var (
	c *Client
)

func init() {
	c = createClient()
}

func createClient() *Client {
	return NewClient(DEFAULT_ENDPOINT, DEFAULT_USERNAME, DEFAULT_PASSWORD)
}

func wait() {
	// Some operations are asynchronously propagated between processes, so
	// while the objects are created they are not immediately visible
	time.Sleep(200 * time.Millisecond)
}

func randomID() string {
	return fmt.Sprintf("id_%s", strings.ReplaceAll(uuid.NewString(), "-", "_"))
}

func TestNewClient(t *testing.T) {
	_, err := c.RegistrationInfo()
	if err != nil {
		t.Fatal(err)
	}

}

func TestGetFingerprint(t *testing.T) {
	_, err := c.GetFingerprint()
	if err != nil {
		t.Fatal(err)
	}

}

func ExampleNewClient() {
	// Connecting to a Harper instance
	c := NewClient("http://localhost:9925", "HDB_ADMIN", "password")

	// (optional) set some proxy
	c.HttpClient.SetProxy("http://localhost:8888")
}

func ExampleNewClient_https() {
	// Connecting to a Harper instance
	c := NewClient("https://localhost:31283", "HDB_ADMIN", "password")

	// with a self-signed certificate
	// will most likely fail so disable security check (https)
	c.HttpClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
}
