package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"solcli/pkg/client"

	"golang.org/x/exp/slog"
)

type APIResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Links []Links     `json:"links,omitempty"`
	Meta  Meta        `json:"meta,omitempty"`
}

type Links struct {
	ACLProfilesURI                  string `json:"aclProfilesUri,omitempty"`
	AuthenticationOauthProfilesURI  string `json:"authenticationOauthProfilesUri,omitempty"`
	AuthenticationOauthProvidersURI string `json:"authenticationOauthProvidersUri,omitempty"`
	AuthorizationGroupsURI          string `json:"authorizationGroupsUri,omitempty"`
	BridgesURI                      string `json:"bridgesUri,omitempty"`
	CertMatchingRulesURI            string `json:"certMatchingRulesUri,omitempty"`
	ClientProfilesURI               string `json:"clientProfilesUri,omitempty"`
	ClientUsernamesURI              string `json:"clientUsernamesUri,omitempty"`
	DistributedCachesURI            string `json:"distributedCachesUri,omitempty"`
	DmrBridgesURI                   string `json:"dmrBridgesUri,omitempty"`
	JndiConnectionFactoriesURI      string `json:"jndiConnectionFactoriesUri,omitempty"`
	JndiQueuesURI                   string `json:"jndiQueuesUri,omitempty"`
	JndiTopicsURI                   string `json:"jndiTopicsUri,omitempty"`
	KafkaReceiversURI               string `json:"kafkaReceiversUri,omitempty"`
	KafkaSendersURI                 string `json:"kafkaSendersUri,omitempty"`
	MqttRetainCachesURI             string `json:"mqttRetainCachesUri,omitempty"`
	MqttSessionsURI                 string `json:"mqttSessionsUri,omitempty"`
	ProxiesURI                      string `json:"proxiesUri,omitempty"`
	QueueTemplatesURI               string `json:"queueTemplatesUri,omitempty"`
	QueuesURI                       string `json:"queuesUri,omitempty"`
	ReplayLogsURI                   string `json:"replayLogsUri,omitempty"`
	ReplicatedTopicsURI             string `json:"replicatedTopicsUri,omitempty"`
	RestDeliveryPointsURI           string `json:"restDeliveryPointsUri,omitempty"`
	SequencedTopicsURI              string `json:"sequencedTopicsUri,omitempty"`
	TelemetryProfilesURI            string `json:"telemetryProfilesUri,omitempty"`
	TopicEndpointTemplatesURI       string `json:"topicEndpointTemplatesUri,omitempty"`
	TopicEndpointsURI               string `json:"topicEndpointsUri,omitempty"`
	URI                             string `json:"uri,omitempty"`
}
type Error struct {
	Code        int    `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status,omitempty"`
}
type Paging struct {
	CursorQuery string `json:"cursorQuery,omitempty"`
	NextPageURI string `json:"nextPageUri,omitempty"`
}
type Request struct {
	Method string `json:"method,omitempty"`
	URI    string `json:"uri,omitempty"`
}
type Meta struct {
	Count        int     `json:"count,omitempty"`
	Error        Error   `json:"error,omitempty"`
	Paging       Paging  `json:"paging,omitempty"`
	Request      Request `json:"request,omitempty"`
	ResponseCode int     `json:"responseCode,omitempty"`
}

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

func (h *APIHanlder) APIRequest(method string, body io.Reader, result *APIResponse) error {
	req, err := h.Client.MakeRequest(method, body)
	if err != nil {
		return err
	}

	resp, err := h.Client.SendRequest(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.Debug("Received non-200 response code from Solace PubSub+ Event Broker",
			slog.Any("response", resp),
		)
		return fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	//slog.Debug("Attempting to decode response from Solace PubSub+ Event Broker",
	//	slog.Any("response", resp),
	//)

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return err
	}

	return nil
}
