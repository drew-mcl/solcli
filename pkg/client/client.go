package client

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"path"

	"golang.org/x/exp/slog"
)

type Client struct {
	http.Client
	Headers  map[string]string
	Username string
	Password string
	BaseURL  *url.URL
}

func NewClient() *Client {
	return &Client{
		Client:  http.Client{},
		BaseURL: &url.URL{},
	}
}

func (c *Client) InitalizeClient(headers map[string]string, username string, password string, rawurl string) error {

	slog.Info("Initalizing client...")

	slog.Debug("Attempting to initalize client with the following parameters:",
		slog.Any("headers", headers),
		slog.String("username", username),
		slog.String("password", password),
		slog.String("rawurl", rawurl),
	)

	parsedURL, err := url.Parse(rawurl)
	if err != nil {
		return err
	}

	err = validateURL(parsedURL)
	if err != nil {
		return err
	}

	c.Username = username
	c.Password = password
	c.Headers = headers
	c.BaseURL = parsedURL

	slog.Info("Client initalized successfully!")
	return nil
}

func validateURL(u *url.URL) error {
	if u.Scheme == "" {
		return fmt.Errorf("scheme cannot be empty")
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return fmt.Errorf("invalid scheme")
	}

	host, port, err := net.SplitHostPort(u.Host)
	if err != nil {
		return fmt.Errorf("invalid host")
	}

	if host == "" {
		return fmt.Errorf("host cannot be empty")
	}

	if port == "" {
		return fmt.Errorf("port cannot be empty")
	}

	p, err := net.LookupPort(u.Scheme, port)
	if err != nil {
		return fmt.Errorf("invalid port")
	}

	if p < 1 && p > 65535 {
		return fmt.Errorf("invalid port")
	}

	slog.Debug("URL validated successfully",
		slog.Group("url",
			slog.String("scheme", u.Scheme),
			slog.String("host", host),
			slog.String("port", port),
		))

	return nil
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}

	if c.Username != "" && c.Password != "" {
		req.SetBasicAuth(c.Username, c.Password)
	} else {
		slog.Warn("Username and password not set, not setting basic auth")
	}
	return c.Client.Do(req)
}

func (c *Client) AppendPath(appendPath string) error {
	slog.Debug("Client has been passed additonal path to append base URL",
		slog.Any("current_base_url", c.BaseURL),
		slog.String("path", appendPath),
	)

	parsedPath, err := url.Parse(appendPath)
	if err != nil {
		return err
	}

	c.BaseURL.Path = path.Join(c.BaseURL.Path, parsedPath.Path)
	slog.Debug("Client base URL has been updated", slog.Any("new_base_url", c.BaseURL))
	return nil
}

func (c *Client) MakeRequest(method string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, c.BaseURL.String(), nil)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Body = nil
		req.Body = http.NoBody
		req.ContentLength = 0
	}

	return req, nil
}

func (c *Client) SendRequest(req *http.Request) (*http.Response, error) {
	slog.Debug("Sending request to Solace PubSub+ Event Broker",
		slog.Any("request", req),
	)

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	slog.Debug("Request sent successfully",
		slog.Any("response", resp),
	)

	return resp, nil
}
