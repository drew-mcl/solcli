package models

type Bridge struct {
	Data  []BridgeData  `json:"data,omitempty"`
	Links []BridgeLinks `json:"links,omitempty"`
	Meta  Meta          `json:"meta,omitempty"`
}
type BridgeData struct {
	BridgeName                              string `json:"bridgeName,omitempty"`
	BridgeVirtualRouter                     string `json:"bridgeVirtualRouter,omitempty"`
	Enabled                                 bool   `json:"enabled,omitempty"`
	MaxTTL                                  int    `json:"maxTtl,omitempty"`
	MsgVpnName                              string `json:"msgVpnName,omitempty"`
	RemoteAuthenticationBasicClientUsername string `json:"remoteAuthenticationBasicClientUsername,omitempty"`
	RemoteAuthenticationBasicPassword       string `json:"remoteAuthenticationBasicPassword,omitempty"`
	RemoteAuthenticationClientCertContent   string `json:"remoteAuthenticationClientCertContent,omitempty"`
	RemoteAuthenticationClientCertPassword  string `json:"remoteAuthenticationClientCertPassword,omitempty"`
	RemoteAuthenticationScheme              string `json:"remoteAuthenticationScheme,omitempty"`
	RemoteConnectionRetryCount              int    `json:"remoteConnectionRetryCount,omitempty"`
	RemoteConnectionRetryDelay              int    `json:"remoteConnectionRetryDelay,omitempty"`
	RemoteDeliverToOnePriority              string `json:"remoteDeliverToOnePriority,omitempty"`
	TLSCipherSuiteList                      string `json:"tlsCipherSuiteList,omitempty"`
}

type BridgeLinks struct {
	RemoteMsgVpnsURI         string `json:"remoteMsgVpnsUri,omitempty"`
	RemoteSubscriptionsURI   string `json:"remoteSubscriptionsUri,omitempty"`
	TLSTrustedCommonNamesURI string `json:"tlsTrustedCommonNamesUri,omitempty"`
	URI                      string `json:"uri,omitempty"`
}
