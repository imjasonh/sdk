/*
Copyright 2023 Chainguard, Inc.
SPDX-License-Identifier: Apache-2.0
*/

package v1

import (
	"context"
	"fmt"
	"net/url"
	"time"

	delegate "chainguard.dev/go-grpc-kit/pkg/options"
	"google.golang.org/grpc"
	"knative.dev/pkg/logging"

	"chainguard.dev/sdk/pkg/auth"
)

type Clients interface {
	Registry() RegistryClient

	Close() error
}

func NewClients(ctx context.Context, addr string, token string) (Clients, error) {
	uri, err := url.Parse(addr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse tenant service address, must be a url: %w", err)
	}

	target, opts := delegate.GRPCOptions(*uri)

	// TODO: we may want to require transport security at some future point.
	if cred := auth.NewFromToken(ctx, token, false); cred != nil {
		opts = append(opts, grpc.WithPerRPCCredentials(cred))
	} else {
		logging.FromContext(ctx).Warn("No authentication provided, this may end badly.")
	}

	var cancel context.CancelFunc
	if _, timeoutSet := ctx.Deadline(); !timeoutSet {
		ctx, cancel = context.WithTimeout(ctx, 300*time.Second)
		defer cancel()
	}
	conn, err := grpc.DialContext(ctx, target, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the iam server: %w", err)
	}

	return &clients{
		registry: NewRegistryClient(conn),

		conn: conn,
	}, nil
}

func NewClientsFromConnection(conn *grpc.ClientConn) Clients {
	return &clients{
		registry: NewRegistryClient(conn),
		// conn is not set, this client struct does not own closing it.
	}
}

type clients struct {
	registry RegistryClient

	conn *grpc.ClientConn
}

func (c *clients) Registry() RegistryClient {
	return c.registry
}

func (c *clients) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}
