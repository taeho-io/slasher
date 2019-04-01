package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taeho-io/idl/gen/go/slasher"
	slash "github.com/xissy/slasher"
	"golang.org/x/net/context"
)

func TestSlashHandler(t *testing.T) {
	ctx := context.Background()

	inputText := "curl --help"
	expected := slash.Slasher(inputText)

	req := &slasher.SlashRequest{
		Text: inputText,
	}

	resp, err := Slash()(ctx, req)
	assert.NotNil(t, resp)
	assert.Nil(t, err)

	assert.Equal(t, resp.SlashedText, expected)
}
