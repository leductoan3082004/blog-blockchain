package keeper

import (
	"context"
	"fmt"

	"blog/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdatePost(goCtx context.Context, msg *types.MsgUpdatePost) (*types.MsgUpdatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx
	post := types.Post{
		Creator: msg.Creator,
		Id:      msg.Id,
		Title:   msg.Title,
		Body:    msg.Body,
	}

	post, found := k.GetPost(ctx, msg.GetId())
	if !found {
		return nil, fmt.Errorf("post %d not found", msg.GetId())
	}

	if msg.Creator != post.Creator {
		return nil, fmt.Errorf("creator %s does not match post %s", msg.Creator, post.Creator)
	}

	post.Title = msg.Title
	post.Body = msg.Body

	k.SetPost(ctx, post)

	return &types.MsgUpdatePostResponse{}, nil
}
