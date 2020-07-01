package auth

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httputil"
	"testing"
)

func TestAppAuth(t *testing.T) {
	appAuth := NewAppAuth("1111", "2222")

	req, _ := http.NewRequest(http.MethodPost, "http://127.0.0.1:9000", nil)
	err := appAuth.Add(req)
	require.NoError(t, err)

	dump, err := httputil.DumpRequest(req, false)
	require.NoError(t, err)
	fmt.Println("[===http===]\n" + string(dump))

	err = appAuth.Verify(req)
	require.NoError(t, err)
}
