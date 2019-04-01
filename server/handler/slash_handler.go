package handler

import (
	"github.com/taeho-io/idl/gen/go/slasher"
	slash "github.com/xissy/slasher"
	"golang.org/x/net/context"
)

type SlashHandlerFunc func(context.Context, *slasher.SlashRequest) (*slasher.SlashResponse, error)

func Slash() SlashHandlerFunc {
	return func(ctx context.Context, req *slasher.SlashRequest) (*slasher.SlashResponse, error) {
		return &slasher.SlashResponse{
			SlashedText: slash.Slasher(req.Text),
		}, nil
	}
}
