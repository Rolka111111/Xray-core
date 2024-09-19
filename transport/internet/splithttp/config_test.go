package splithttp_test

import (
	"testing"

	. "github.com/luckyluke-a/xray-core/transport/internet/splithttp"
)

func Test_GetNormalizedPath(t *testing.T) {
	c := Config{
		Path: "/?world",
	}

	path := c.GetNormalizedPath()
	if path != "/" {
		t.Error("Unexpected: ", path)
	}
}
