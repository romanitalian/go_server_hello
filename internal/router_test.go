package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	router "github.com/romanitalian/go_server_hello/internal"
	"github.com/stretchr/testify/require"
)

func TestNotFoud(t *testing.T) {
	rt := router.New()
	srv := httptest.NewServer(rt.RootHandler())
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/unknown")
	require.NoError(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}