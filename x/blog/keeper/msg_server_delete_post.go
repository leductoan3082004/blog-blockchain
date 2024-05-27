package keeper

import (
	"context"
	"fmt"

	"blog/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DeletePost(goCtx context.Context, msg *types.MsgDeletePost) (*types.MsgDeletePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	post, found := k.GetPost(ctx, msg.GetId())
	if !found {
		return nil, fmt.Errorf("post %d not found", msg.GetId())
	}

	if post.Creator != msg.Creator {
		return nil, fmt.Errorf("creator %s does not match post %s", msg.Creator, post.Creator)
	}

	k.RemovePost(ctx, msg.GetId())
	return &types.MsgDeletePostResponse{}, nil
}
