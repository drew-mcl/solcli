package models

type QueueList struct {
	Data  []QueueData  `json:"data,omitempty"`
	Links []QueueLinks `json:"links,omitempty"`
	Meta  Meta         `json:"meta,omitempty"`
}

type QueueSingle struct {
	Data  QueueData  `json:"data,omitempty"`
	Links QueueLinks `json:"links,omitempty"`
	Meta  Meta       `json:"meta,omitempty"`
}

type EventBindCountThreshold struct {
	ClearPercent int `json:"clearPercent,omitempty"`
	ClearValue   int `json:"clearValue,omitempty"`
	SetPercent   int `json:"setPercent,omitempty"`
	SetValue     int `json:"setValue,omitempty"`
}

type EventRejectLowPriorityMsgLimitThreshold struct {
	ClearPercent int `json:"clearPercent,omitempty"`
	ClearValue   int `json:"clearValue,omitempty"`
	SetPercent   int `json:"setPercent,omitempty"`
	SetValue     int `json:"setValue,omitempty"`
}
type QueueData struct {
	AccessType                              string                                  `json:"accessType,omitempty"`
	ConsumerAckPropagationEnabled           bool                                    `json:"consumerAckPropagationEnabled,omitempty"`
	DeadMsgQueue                            string                                  `json:"deadMsgQueue,omitempty"`
	DeliveryCountEnabled                    bool                                    `json:"deliveryCountEnabled,omitempty"`
	DeliveryDelay                           int                                     `json:"deliveryDelay,omitempty"`
	EgressEnabled                           bool                                    `json:"egressEnabled,omitempty"`
	EventBindCountThreshold                 EventBindCountThreshold                 `json:"eventBindCountThreshold,omitempty"`
	EventMsgSpoolUsageThreshold             EventMsgSpoolUsageThreshold             `json:"eventMsgSpoolUsageThreshold,omitempty"`
	EventRejectLowPriorityMsgLimitThreshold EventRejectLowPriorityMsgLimitThreshold `json:"eventRejectLowPriorityMsgLimitThreshold,omitempty"`
	IngressEnabled                          bool                                    `json:"ingressEnabled,omitempty"`
	MaxBindCount                            int                                     `json:"maxBindCount,omitempty"`
	MaxDeliveredUnackedMsgsPerFlow          int                                     `json:"maxDeliveredUnackedMsgsPerFlow,omitempty"`
	MaxMsgSize                              int                                     `json:"maxMsgSize,omitempty"`
	MaxMsgSpoolUsage                        int                                     `json:"maxMsgSpoolUsage,omitempty"`
	MaxRedeliveryCount                      int                                     `json:"maxRedeliveryCount,omitempty"`
	MaxTTL                                  int                                     `json:"maxTtl,omitempty"`
	MsgVpnName                              string                                  `json:"msgVpnName,omitempty"`
	Owner                                   string                                  `json:"owner,omitempty"`
	PartitionCount                          int                                     `json:"partitionCount,omitempty"`
	PartitionRebalanceDelay                 int                                     `json:"partitionRebalanceDelay,omitempty"`
	PartitionRebalanceMaxHandoffTime        int                                     `json:"partitionRebalanceMaxHandoffTime,omitempty"`
	Permission                              string                                  `json:"permission,omitempty"`
	QueueName                               string                                  `json:"queueName,omitempty"`
	RedeliveryDelayEnabled                  bool                                    `json:"redeliveryDelayEnabled,omitempty"`
	RedeliveryDelayInitialInterval          int                                     `json:"redeliveryDelayInitialInterval,omitempty"`
	RedeliveryDelayMaxInterval              int                                     `json:"redeliveryDelayMaxInterval,omitempty"`
	RedeliveryDelayMultiplier               int                                     `json:"redeliveryDelayMultiplier,omitempty"`
	RedeliveryEnabled                       bool                                    `json:"redeliveryEnabled,omitempty"`
	RejectLowPriorityMsgEnabled             bool                                    `json:"rejectLowPriorityMsgEnabled,omitempty"`
	RejectLowPriorityMsgLimit               int                                     `json:"rejectLowPriorityMsgLimit,omitempty"`
	RejectMsgToSenderOnDiscardBehavior      string                                  `json:"rejectMsgToSenderOnDiscardBehavior,omitempty"`
	RespectMsgPriorityEnabled               bool                                    `json:"respectMsgPriorityEnabled,omitempty"`
	RespectTTLEnabled                       bool                                    `json:"respectTtlEnabled,omitempty"`
}

type QueueLinks struct {
	SubscriptionsURI string `json:"subscriptionsUri,omitempty"`
	URI              string `json:"uri,omitempty"`
}
