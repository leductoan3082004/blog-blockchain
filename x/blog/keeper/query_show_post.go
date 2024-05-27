package keeper

import (
	"context"
	"fmt"

	"blog/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowPost(goCtx context.Context, req *types.QueryShowPostRequest) (*types.QueryShowPostResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	post, found := k.GetPost(ctx, req.GetId())
	if !found {
		return nil, fmt.Errorf("no post with id %d", req.GetId())
	}

	return &types.QueryShowPostResponse{
		Post: &post,
	}, nil
}
