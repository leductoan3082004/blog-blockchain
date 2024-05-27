package keeper

import (
	"context"

	"blog/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	post := types.Post{
		Title:   msg.GetTitle(),
		Body:    msg.GetBody(),
		Creator: msg.GetCreator(),
		Id:      0,
	}

	post.Id = k.AppendPost(ctx, post)
	return &types.MsgCreatePostResponse{
		Id: post.Id,
	}, nil
}
