package grpcclient

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

type ClientConstructor interface {
	NewClient(conn *grpc.ClientConn) interface{}
}

type Client struct {
	conn        *grpc.ClientConn
	client      interface{}
	constructor ClientConstructor
	timeout     time.Duration
}

func NewClient(timeout time.Duration, constructor ClientConstructor) *Client {
	return &Client{
		timeout:     timeout,
		constructor: constructor,
	}
}

func (c *Client) Connect(ctx context.Context, target string) error {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		target,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		return err
	}

	c.conn = conn
	c.client = c.constructor.NewClient(conn)

	return nil
}

func (c *Client) AutoReconnect(ctx context.Context, target string) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				if !c.IsReady() {
					if err := c.Connect(ctx, target); err != nil {
						log.Printf("Connection failed: %v", err)
						time.Sleep(3 * time.Second)
						continue
					}
					log.Println("Successfully connected to UserService")
				}
				time.Sleep(5 * time.Second)
			}
		}
	}()
}

func (c *Client) IsReady() bool {
	if c.conn == nil {
		return false
	}
	return c.conn.GetState() == connectivity.Ready
}

func (c *Client) GetClient() interface{} {
	return c.client
}

func (c *Client) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}
