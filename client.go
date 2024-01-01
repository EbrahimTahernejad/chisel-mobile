package chisel

import (
	"net/http"
	"strconv"

	chclient "github.com/jpillora/chisel/client"
	"github.com/jpillora/chisel/share/cos"

	_ "github.com/sagernet/gomobile/bind"
)

type Client struct {
	client *chclient.Client
}

func NewClient(server string, port int) (*Client, error) {
	config := chclient.Config{Headers: http.Header{}}
	config.Server = server
	remote := strconv.Itoa(port) + ":socks"
	config.Remotes = []string{remote}
	c, err := chclient.NewClient(&config)
	if err != nil {
		return nil, err
	}
	return &Client{
		client: c,
	}, nil
}

func (c *Client) Start() error {
	go cos.GoStats()
	ctx := cos.InterruptContext()
	if err := c.client.Start(ctx); err != nil {
		return err
	}
	if err := c.client.Wait(); err != nil {
		return err
	}
	return nil
}

func (c *Client) Stop() error {
	return c.client.Close()
}
