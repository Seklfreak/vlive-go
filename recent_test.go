package vlive_go

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVLive_Recents(t *testing.T) {
	client := NewVLive(http.DefaultClient)

	recents, err := client.Recents()
	assert.NoError(t, err)
	assert.Len(t, recents, 50)
}
