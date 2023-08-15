package api

import (
	"io"
)

var VPNAPIPath = "/SEMP/v2/monitor/msgVpns"

type EventConnectionCountThreshold struct {
	ClearPercent int `json:"clearPercent,omitempty"`
	ClearValue   int `json:"clearValue,omitempty"`
	SetPercent   int `json:"setPercent,omitempty"`
	SetValue     int `json:"setValue,omitempty"`
}
type EventEgressFlowCountThreshold struct {
	ClearPercent int `json:"clearPercent,omitempty"`
	ClearValue   int `json:"clearValue,omitempty"`
	SetPercent   int `json:"setPercent,omitempty"`
	SetValue     int `json:"setValue,omitempty"`
}
type EventEgressMsgRateThreshold struct {
	ClearValue int `json:"clearValue,omitempty"`
	SetValue   int `json:"setValue,omitempty"`
}
type EventEndpointCountThreshold struct {
	ClearPercent int `json:"clearPercent,omitempty"`
	ClearValue   int `json:"clearValue,omitempty"`
	SetPercent   int `json:"setPercent,omitempty"`
	SetValue     int `json:"setValue,omitempty"`
}
type EventIngressFlowCountThreshold struct {
	ClearPercent int `json:"clearPercent,omitempty"`
	ClearValue   int `json:"clearValue,omitempty"`
	SetPercent   int `json:"setPercent,omitempty"`
	SetValue     int `json:"setValue,omitempty"`
}
type EventIngressMsgRateThreshold struct {
	ClearValue int `json:"clearValue,omitempty"`
	SetValue   int `json:"setValue,omitempty"`
}
type EventMsgSpoolUsageThreshold struct {
	ClearPercent int `json:"clearPercent,omitempty"`
	ClearValue   int `json:"clearValue,omitempty"`
	SetPercent   int `json:"setPercent,omitempty"`
	SetValue     int `json:"setValue,omitempty"`
}
type EventServiceAmqpConnectionCountThreshold struct {
	ClearPercent int `json:"clearPercent,omitempty"`
	ClearValue   int `json:"clearValue,omitempty"`
	SetPercent   int `json:"setPercent,omitempty"`
	SetValue     int `json:"setValue,omitempty"`
}
type EventServiceMqttConnectionCountThreshold struct {
	ClearPercent int `json:"clearPercent,omitempty"`
	ClearValue   int `json:"clearValue,omitempty"`
	SetPercent   int `json:"setPercent,omitempty"`
	SetValue     int `json:"setValue,omitempty"`
}
type EventServiceRestIncomingConnectionCountThreshold struct {
	ClearPercent int `json:"clearPercent,omitempty"`
	ClearValue   int `json:"clearValue,omitempty"`
	SetPercent   int `json:"setPercent,omitempty"`
	SetValue     int `json:"setValue,omitempty"`
}
type EventServiceSmfConnectionCountThreshold struct {
	ClearPercent int `json:"clearPercent,omitempty"`
	ClearValue   int `json:"clearValue,omitempty"`
	SetPercent   int `json:"setPercent,omitempty"`
	SetValue     int `json:"setValue,omitempty"`
}
type EventServiceWebConnectionCountThreshold struct {
	ClearPercent int `json:"clearPercent,omitempty"`
	ClearValue   int `json:"clearValue,omitempty"`
	SetPercent   int `json:"setPercent,omitempty"`
	SetValue     int `json:"setValue,omitempty"`
}
type EventSubscriptionCountThreshold struct {
	ClearPercent int `json:"clearPercent,omitempty"`
	ClearValue   int `json:"clearValue,omitempty"`
	SetPercent   int `json:"setPercent,omitempty"`
	SetValue     int `json:"setValue,omitempty"`
}
type EventTransactedSessionCountThreshold struct {
	ClearPercent int `json:"clearPercent,omitempty"`
	ClearValue   int `json:"clearValue,omitempty"`
	SetPercent   int `json:"setPercent,omitempty"`
	SetValue     int `json:"setValue,omitempty"`
}
type EventTransactionCountThreshold struct {
	ClearPercent int `json:"clearPercent,omitempty"`
	ClearValue   int `json:"clearValue,omitempty"`
	SetPercent   int `json:"setPercent,omitempty"`
	SetValue     int `json:"setValue,omitempty"`
}
type MsgVpnObject struct {
	Alias                                                   string                                           `json:"alias,omitempty"`
	AuthenticationBasicEnabled                              bool                                             `json:"authenticationBasicEnabled,omitempty"`
	AuthenticationBasicProfileName                          string                                           `json:"authenticationBasicProfileName,omitempty"`
	AuthenticationBasicRadiusDomain                         string                                           `json:"authenticationBasicRadiusDomain,omitempty"`
	AuthenticationBasicType                                 string                                           `json:"authenticationBasicType,omitempty"`
	AuthenticationClientCertAllowAPIProvidedUsernameEnabled bool                                             `json:"authenticationClientCertAllowApiProvidedUsernameEnabled,omitempty"`
	AuthenticationClientCertCertificateMatchingRulesEnabled bool                                             `json:"authenticationClientCertCertificateMatchingRulesEnabled,omitempty"`
	AuthenticationClientCertEnabled                         bool                                             `json:"authenticationClientCertEnabled,omitempty"`
	AuthenticationClientCertMaxChainDepth                   int                                              `json:"authenticationClientCertMaxChainDepth,omitempty"`
	AuthenticationClientCertRevocationCheckMode             string                                           `json:"authenticationClientCertRevocationCheckMode,omitempty"`
	AuthenticationClientCertUsernameSource                  string                                           `json:"authenticationClientCertUsernameSource,omitempty"`
	AuthenticationClientCertValidateDateEnabled             bool                                             `json:"authenticationClientCertValidateDateEnabled,omitempty"`
	AuthenticationKerberosAllowAPIProvidedUsernameEnabled   bool                                             `json:"authenticationKerberosAllowApiProvidedUsernameEnabled,omitempty"`
	AuthenticationKerberosEnabled                           bool                                             `json:"authenticationKerberosEnabled,omitempty"`
	AuthenticationOauthDefaultProfileName                   string                                           `json:"authenticationOauthDefaultProfileName,omitempty"`
	AuthenticationOauthDefaultProviderName                  string                                           `json:"authenticationOauthDefaultProviderName,omitempty"`
	AuthenticationOauthEnabled                              bool                                             `json:"authenticationOauthEnabled,omitempty"`
	AuthorizationLdapGroupMembershipAttributeName           string                                           `json:"authorizationLdapGroupMembershipAttributeName,omitempty"`
	AuthorizationLdapTrimClientUsernameDomainEnabled        bool                                             `json:"authorizationLdapTrimClientUsernameDomainEnabled,omitempty"`
	AuthorizationProfileName                                string                                           `json:"authorizationProfileName,omitempty"`
	AuthorizationType                                       string                                           `json:"authorizationType,omitempty"`
	BridgingTLSServerCertEnforceTrustedCommonNameEnabled    bool                                             `json:"bridgingTlsServerCertEnforceTrustedCommonNameEnabled,omitempty"`
	BridgingTLSServerCertMaxChainDepth                      int                                              `json:"bridgingTlsServerCertMaxChainDepth,omitempty"`
	BridgingTLSServerCertValidateDateEnabled                bool                                             `json:"bridgingTlsServerCertValidateDateEnabled,omitempty"`
	BridgingTLSServerCertValidateNameEnabled                bool                                             `json:"bridgingTlsServerCertValidateNameEnabled,omitempty"`
	DistributedCacheManagementEnabled                       bool                                             `json:"distributedCacheManagementEnabled,omitempty"`
	DmrEnabled                                              bool                                             `json:"dmrEnabled,omitempty"`
	Enabled                                                 bool                                             `json:"enabled,omitempty"`
	EventConnectionCountThreshold                           EventConnectionCountThreshold                    `json:"eventConnectionCountThreshold,omitempty"`
	EventEgressFlowCountThreshold                           EventEgressFlowCountThreshold                    `json:"eventEgressFlowCountThreshold,omitempty"`
	EventEgressMsgRateThreshold                             EventEgressMsgRateThreshold                      `json:"eventEgressMsgRateThreshold,omitempty"`
	EventEndpointCountThreshold                             EventEndpointCountThreshold                      `json:"eventEndpointCountThreshold,omitempty"`
	EventIngressFlowCountThreshold                          EventIngressFlowCountThreshold                   `json:"eventIngressFlowCountThreshold,omitempty"`
	EventIngressMsgRateThreshold                            EventIngressMsgRateThreshold                     `json:"eventIngressMsgRateThreshold,omitempty"`
	EventLargeMsgThreshold                                  int                                              `json:"eventLargeMsgThreshold,omitempty"`
	EventLogTag                                             string                                           `json:"eventLogTag,omitempty"`
	EventMsgSpoolUsageThreshold                             EventMsgSpoolUsageThreshold                      `json:"eventMsgSpoolUsageThreshold,omitempty"`
	EventPublishClientEnabled                               bool                                             `json:"eventPublishClientEnabled,omitempty"`
	EventPublishMsgVpnEnabled                               bool                                             `json:"eventPublishMsgVpnEnabled,omitempty"`
	EventPublishSubscriptionMode                            string                                           `json:"eventPublishSubscriptionMode,omitempty"`
	EventPublishTopicFormatMqttEnabled                      bool                                             `json:"eventPublishTopicFormatMqttEnabled,omitempty"`
	EventPublishTopicFormatSmfEnabled                       bool                                             `json:"eventPublishTopicFormatSmfEnabled,omitempty"`
	EventServiceAmqpConnectionCountThreshold                EventServiceAmqpConnectionCountThreshold         `json:"eventServiceAmqpConnectionCountThreshold,omitempty"`
	EventServiceMqttConnectionCountThreshold                EventServiceMqttConnectionCountThreshold         `json:"eventServiceMqttConnectionCountThreshold,omitempty"`
	EventServiceRestIncomingConnectionCountThreshold        EventServiceRestIncomingConnectionCountThreshold `json:"eventServiceRestIncomingConnectionCountThreshold,omitempty"`
	EventServiceSmfConnectionCountThreshold                 EventServiceSmfConnectionCountThreshold          `json:"eventServiceSmfConnectionCountThreshold,omitempty"`
	EventServiceWebConnectionCountThreshold                 EventServiceWebConnectionCountThreshold          `json:"eventServiceWebConnectionCountThreshold,omitempty"`
	EventSubscriptionCountThreshold                         EventSubscriptionCountThreshold                  `json:"eventSubscriptionCountThreshold,omitempty"`
	EventTransactedSessionCountThreshold                    EventTransactedSessionCountThreshold             `json:"eventTransactedSessionCountThreshold,omitempty"`
	EventTransactionCountThreshold                          EventTransactionCountThreshold                   `json:"eventTransactionCountThreshold,omitempty"`
	ExportSubscriptionsEnabled                              bool                                             `json:"exportSubscriptionsEnabled,omitempty"`
	JndiEnabled                                             bool                                             `json:"jndiEnabled,omitempty"`
	MaxConnectionCount                                      int                                              `json:"maxConnectionCount,omitempty"`
	MaxEgressFlowCount                                      int                                              `json:"maxEgressFlowCount,omitempty"`
	MaxEndpointCount                                        int                                              `json:"maxEndpointCount,omitempty"`
	MaxIngressFlowCount                                     int                                              `json:"maxIngressFlowCount,omitempty"`
	MaxMsgSpoolUsage                                        int                                              `json:"maxMsgSpoolUsage,omitempty"`
	MaxSubscriptionCount                                    int                                              `json:"maxSubscriptionCount,omitempty"`
	MaxTransactedSessionCount                               int                                              `json:"maxTransactedSessionCount,omitempty"`
	MaxTransactionCount                                     int                                              `json:"maxTransactionCount,omitempty"`
	MqttRetainMaxMemory                                     int                                              `json:"mqttRetainMaxMemory,omitempty"`
	MsgVpnName                                              string                                           `json:"msgVpnName,omitempty"`
	ReplicationAckPropagationIntervalMsgCount               int                                              `json:"replicationAckPropagationIntervalMsgCount,omitempty"`
	ReplicationBridgeAuthenticationBasicClientUsername      string                                           `json:"replicationBridgeAuthenticationBasicClientUsername,omitempty"`
	ReplicationBridgeAuthenticationBasicPassword            string                                           `json:"replicationBridgeAuthenticationBasicPassword,omitempty"`
	ReplicationBridgeAuthenticationClientCertContent        string                                           `json:"replicationBridgeAuthenticationClientCertContent,omitempty"`
	ReplicationBridgeAuthenticationClientCertPassword       string                                           `json:"replicationBridgeAuthenticationClientCertPassword,omitempty"`
	ReplicationBridgeAuthenticationScheme                   string                                           `json:"replicationBridgeAuthenticationScheme,omitempty"`
	ReplicationBridgeCompressedDataEnabled                  bool                                             `json:"replicationBridgeCompressedDataEnabled,omitempty"`
	ReplicationBridgeEgressFlowWindowSize                   int                                              `json:"replicationBridgeEgressFlowWindowSize,omitempty"`
	ReplicationBridgeRetryDelay                             int                                              `json:"replicationBridgeRetryDelay,omitempty"`
	ReplicationBridgeTLSEnabled                             bool                                             `json:"replicationBridgeTlsEnabled,omitempty"`
	ReplicationBridgeUnidirectionalClientProfileName        string                                           `json:"replicationBridgeUnidirectionalClientProfileName,omitempty"`
	ReplicationEnabled                                      bool                                             `json:"replicationEnabled,omitempty"`
	ReplicationEnabledQueueBehavior                         string                                           `json:"replicationEnabledQueueBehavior,omitempty"`
	ReplicationQueueMaxMsgSpoolUsage                        int                                              `json:"replicationQueueMaxMsgSpoolUsage,omitempty"`
	ReplicationQueueRejectMsgToSenderOnDiscardEnabled       bool                                             `json:"replicationQueueRejectMsgToSenderOnDiscardEnabled,omitempty"`
	ReplicationRejectMsgWhenSyncIneligibleEnabled           bool                                             `json:"replicationRejectMsgWhenSyncIneligibleEnabled,omitempty"`
	ReplicationRole                                         string                                           `json:"replicationRole,omitempty"`
	ReplicationTransactionMode                              string                                           `json:"replicationTransactionMode,omitempty"`
	RestTLSServerCertEnforceTrustedCommonNameEnabled        bool                                             `json:"restTlsServerCertEnforceTrustedCommonNameEnabled,omitempty"`
	RestTLSServerCertMaxChainDepth                          int                                              `json:"restTlsServerCertMaxChainDepth,omitempty"`
	RestTLSServerCertValidateDateEnabled                    bool                                             `json:"restTlsServerCertValidateDateEnabled,omitempty"`
	RestTLSServerCertValidateNameEnabled                    bool                                             `json:"restTlsServerCertValidateNameEnabled,omitempty"`
	SempOverMsgBusAdminClientEnabled                        bool                                             `json:"sempOverMsgBusAdminClientEnabled,omitempty"`
	SempOverMsgBusAdminDistributedCacheEnabled              bool                                             `json:"sempOverMsgBusAdminDistributedCacheEnabled,omitempty"`
	SempOverMsgBusAdminEnabled                              bool                                             `json:"sempOverMsgBusAdminEnabled,omitempty"`
	SempOverMsgBusEnabled                                   bool                                             `json:"sempOverMsgBusEnabled,omitempty"`
	SempOverMsgBusShowEnabled                               bool                                             `json:"sempOverMsgBusShowEnabled,omitempty"`
	ServiceAmqpMaxConnectionCount                           int                                              `json:"serviceAmqpMaxConnectionCount,omitempty"`
	ServiceAmqpPlainTextEnabled                             bool                                             `json:"serviceAmqpPlainTextEnabled,omitempty"`
	ServiceAmqpPlainTextListenPort                          int                                              `json:"serviceAmqpPlainTextListenPort,omitempty"`
	ServiceAmqpTLSEnabled                                   bool                                             `json:"serviceAmqpTlsEnabled,omitempty"`
	ServiceAmqpTLSListenPort                                int                                              `json:"serviceAmqpTlsListenPort,omitempty"`
	ServiceMqttAuthenticationClientCertRequest              string                                           `json:"serviceMqttAuthenticationClientCertRequest,omitempty"`
	ServiceMqttMaxConnectionCount                           int                                              `json:"serviceMqttMaxConnectionCount,omitempty"`
	ServiceMqttPlainTextEnabled                             bool                                             `json:"serviceMqttPlainTextEnabled,omitempty"`
	ServiceMqttPlainTextListenPort                          int                                              `json:"serviceMqttPlainTextListenPort,omitempty"`
	ServiceMqttTLSEnabled                                   bool                                             `json:"serviceMqttTlsEnabled,omitempty"`
	ServiceMqttTLSListenPort                                int                                              `json:"serviceMqttTlsListenPort,omitempty"`
	ServiceMqttTLSWebSocketEnabled                          bool                                             `json:"serviceMqttTlsWebSocketEnabled,omitempty"`
	ServiceMqttTLSWebSocketListenPort                       int                                              `json:"serviceMqttTlsWebSocketListenPort,omitempty"`
	ServiceMqttWebSocketEnabled                             bool                                             `json:"serviceMqttWebSocketEnabled,omitempty"`
	ServiceMqttWebSocketListenPort                          int                                              `json:"serviceMqttWebSocketListenPort,omitempty"`
	ServiceRestIncomingAuthenticationClientCertRequest      string                                           `json:"serviceRestIncomingAuthenticationClientCertRequest,omitempty"`
	ServiceRestIncomingAuthorizationHeaderHandling          string                                           `json:"serviceRestIncomingAuthorizationHeaderHandling,omitempty"`
	ServiceRestIncomingMaxConnectionCount                   int                                              `json:"serviceRestIncomingMaxConnectionCount,omitempty"`
	ServiceRestIncomingPlainTextEnabled                     bool                                             `json:"serviceRestIncomingPlainTextEnabled,omitempty"`
	ServiceRestIncomingPlainTextListenPort                  int                                              `json:"serviceRestIncomingPlainTextListenPort,omitempty"`
	ServiceRestIncomingTLSEnabled                           bool                                             `json:"serviceRestIncomingTlsEnabled,omitempty"`
	ServiceRestIncomingTLSListenPort                        int                                              `json:"serviceRestIncomingTlsListenPort,omitempty"`
	ServiceRestMode                                         string                                           `json:"serviceRestMode,omitempty"`
	ServiceRestOutgoingMaxConnectionCount                   int                                              `json:"serviceRestOutgoingMaxConnectionCount,omitempty"`
	ServiceSmfMaxConnectionCount                            int                                              `json:"serviceSmfMaxConnectionCount,omitempty"`
	ServiceSmfPlainTextEnabled                              bool                                             `json:"serviceSmfPlainTextEnabled,omitempty"`
	ServiceSmfTLSEnabled                                    bool                                             `json:"serviceSmfTlsEnabled,omitempty"`
	ServiceWebAuthenticationClientCertRequest               string                                           `json:"serviceWebAuthenticationClientCertRequest,omitempty"`
	ServiceWebMaxConnectionCount                            int                                              `json:"serviceWebMaxConnectionCount,omitempty"`
	ServiceWebPlainTextEnabled                              bool                                             `json:"serviceWebPlainTextEnabled,omitempty"`
	ServiceWebTLSEnabled                                    bool                                             `json:"serviceWebTlsEnabled,omitempty"`
	TLSAllowDowngradeToPlainTextEnabled                     bool                                             `json:"tlsAllowDowngradeToPlainTextEnabled,omitempty"`
}

func (h *APIHanlder) MsgVpnAPIRequest(method string, body io.Reader) (*APIResponse, error) {

	var response APIResponse
	response.Data = MsgVpnObject{}

	err := h.APIRequest("GET", nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
