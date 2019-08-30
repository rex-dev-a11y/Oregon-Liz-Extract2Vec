package clients

import (
	"context"
	mocket "github.com/Selvatico/go-mocket"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"database/sql"
)

func testingHTTPClient(handler http.Handler) (*http.Client, func()){
	s := httptest.NewServer(handler)

	cli := &http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, network, _ string) (net.Conn, error) {
				return net.Dial(network, s.Listener.Addr().String())
			},
		},
	}
	return cli, s.Close
}

func TestDBClient(t *testing.T) {
	mocket.Catcher.Register()

	t.Run("Running local db...", func(t *testing.T) {
		_, err := sql.Open(mocket.DriverName, "don't_matter")
		if err != nil {
			t.Errorf("Failed to load fake db driver: %s", err)
		}

		return
	})
}


