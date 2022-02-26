package credentials

import (
	"context"
	"fmt"
	"github.com/Cleopha/ifttt-like-common/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn *grpc.ClientConn
	clt  protos.CredentialServiceClient
}

func NewClient(port string) (*Client, error) {
	conn, err := grpc.Dial(":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %w", err)
	}

	conn.Connect()
	clt := protos.NewCredentialServiceClient(conn)

	return &Client{
		conn,
		clt,
	}, nil
}

func (c *Client) InsertCredential(ctx context.Context, in *protos.InsertCredentialRequest, opts ...grpc.CallOption) (*protos.Credential, error) {
	return c.clt.InsertCredential(ctx, in, opts...)
}

func (c *Client) GetCredential(ctx context.Context, in *protos.GetCredentialRequest, opts ...grpc.CallOption) (*protos.Credential, error) {
	return c.clt.GetCredential(ctx, in, opts...)
}

func (c *Client) UpdateCredential(ctx context.Context, in *protos.UpdateCredentialRequest, opts ...grpc.CallOption) (*protos.Credential, error) {
	return c.clt.UpdateCredential(ctx, in, opts...)
}

func (c *Client) RemoveCredential(ctx context.Context, in *protos.RemoveCredentialRequest, opts ...grpc.CallOption) (*protos.Credential, error) {
	return c.clt.RemoveCredential(ctx, in, opts...)
}
