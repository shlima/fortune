package blockchain

import (
	"net/http"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

type Setup struct {
	*Cli
}

func MustSetup(t *testing.T) *Setup {
	cli := New()
	httpmock.ActivateNonDefault(cli.client.GetClient())
	return &Setup{Cli: cli}
}

func (s *Setup) ExpectJSONGet(url string, status int, body string) {
	res := httpmock.NewStringResponder(status, body)
	res = res.HeaderSet(http.Header{"content-type": []string{"application/json"}})
	httpmock.RegisterResponder("GET", url, res)
}

func MustTextFixture(t *testing.T, filename string) string {
	got, err := os.ReadFile(filename)
	require.NoError(t, err)
	return string(got)
}
