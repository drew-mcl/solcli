package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"solcli/pkg/client"
	"solcli/pkg/models"

	"golang.org/x/exp/slog"
)

type APIHanlder struct {
	Client *client.Client
}

func NewAPIHandler(c *client.Client, apiPath string) *APIHanlder {
	err := c.AppendPath(apiPath)
	if err != nil {
		slog.Error("Failed to append API path to client!",
			slog.Any("error", err),
		)
	}

	return &APIHanlder{
		Client: c,
	}
}

func (h *APIHanlder) APIRequest(method string, body io.Reader) (*http.Response, error) {
	req, err := h.Client.MakeRequest(method, body)
	if err != nil {
		return nil, err
	}

	resp, err := h.Client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	//slog.Debug("Attempting to decode response from Solace PubSub+ Event Broker",
	//	slog.Any("response", resp),
	//)

	return resp, nil
}

func (h *APIHanlder) MsgVpnAPIRequest(method string, body io.Reader) (interface{}, error) {
	var data interface{}

	resp, err := h.APIRequest(method, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errResp models.ErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&errResp)
		if err == nil {
			return nil, fmt.Errorf("received non-200 response code: %d, description: %s", resp.StatusCode, errResp.Meta.Error.Description)
		}
		return nil, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	switch method {
	case "GET":
		data = &models.MsgVpnList{}
	case "POST", "PUT", "DELETE":
		data = &models.MsgVpnSingle{}
	default:
		return nil, fmt.Errorf("unsupported HTTP method: %s", method)
	}

	err = json.NewDecoder(resp.Body).Decode(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (h *APIHanlder) QueueAPIRequest(method string, body io.Reader) (interface{}, error) {
	var data interface{}

	resp, err := h.APIRequest(method, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errResp models.ErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&errResp)
		if err == nil {
			return nil, fmt.Errorf("received non-200 response code: %d, description: %s", resp.StatusCode, errResp.Meta.Error.Description)
		}
		return nil, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	switch method {
	case "GET":
		data = &models.QueueList{}
	case "POST", "PUT", "DELETE":
		data = &models.QueueSingle{}
	default:
		return nil, fmt.Errorf("unsupported HTTP method: %s", method)
	}

	err = json.NewDecoder(resp.Body).Decode(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
