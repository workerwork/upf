package elem

type IEType uint16

const (
	//M
	IETYPE_NODE_ID                          IEType = 60
	IETYPE_F_SEID                           IEType = 57
	IETYPE_F_TEID                           IEType = 21
	IETYPE_MAC_ADDRESS                      IEType = 133
	IETYPE_C_TAG                            IEType = 134
	IETYPE_S_TAG                            IEType = 135
	IETYPE_ETHERTYPE                        IEType = 136
	IETYPE_ETHERNET_FILTER_ID               IEType = 138
	IETYPE_ETHERNET_FILTER_PROERTIER        IEType = 139
	IETYPE_SDFFILTER                        IEType = 23
	IETYPE_UE_IP_ADDRESS                    IEType = 93
	IETYPE_SOURCE_INTERFACE                 IEType = 20
	IETYPE_NETWORK_INSTANCE                 IEType = 22
	IETYPE_TRAFFIC_ENDPOINT_ID              IEType = 131
	IETYPE_APPLICATION_ID                   IEType = 24
	IETYPE_ETHERNET_PDU_SESSION_INFORMATION IEType = 142
	IETYPE_ETHERNET_PACKET_FILTER           IEType = 132
	IETYPE_QFI                              IEType = 124
	IETYPE_FRAMED_ROUTE                     IEType = 153
	IETYPE_FRAMED_ROUTING                   IEType = 154
	IETYPE_FRAMED_IPv6_ROUTE                IEType = 155
	IETYPE_3GPP_INTERFACE_TYPE              IEType = 160

	IETYPE_PDI                  IEType = 2
	IETYPE_PDR_ID               IEType = 56
	IETYPE_OUTER_HEADER_REMOVAL IEType = 95

	IETYPE_PRECEDENCE                IEType = 29
	IETYPE_OUTER_HEADER_CREATION     IEType = 84
	IETYPE_FAR_ID                    IEType = 108
	IETYPE_URR_ID                    IEType = 81
	IETYPE_QER_ID                    IEType = 109
	IETYPE_ACTIVATE_PREDEFINED_RULES IEType = 106
	IETYPE_FORWARDING_PARAMETERS     IEType = 4

	IETYPE_CREATE_PDR              IEType = 1
	IETYPE_CREATE_FAR              IEType = 3
	IETYPE_APPLY_ACTION            IEType = 44
	IETYPE_DESTINATION_INTERFACE   IEType = 42
	IETYPE_REDIRECT_INFORAMTION    IEType = 38
	IETYPE_TRANSPORT_LEVEL_MARKING IEType = 30
	IETYPE_FORWARDING_POLICY       IEType = 41
	IETYPE_FORWARDING_PROXYING     IEType = 137
	IETYPE_HEADER_ENRICHMENT       IEType = 98
	IETYPE_DUPLICATING_PARAMETERS  IEType = 5
	IETYPE_BAR_ID                  IEType = 88

	IETYPE_REMOVE_PDR IEType = 15
	IETYPE_REMOVE_FAR IEType = 16

	IETYPE_RECOVERY_TIME_STAMP               IEType = 96
	IETYPE_CAUSE                             IEType = 19
	IETYPE_OFFENDING_IE                      IEType = 40
	IETYPE_REPORT_TYPE                       IEType = 39
	IETYPE_DOWNLINK_DATA_SERVICE_INFORMATION IEType = 45
	IETYPE_DOWNLINK_DATA_REPORT              IEType = 83
	IETYPE_VOLUME_MEASUREMENT                IEType = 66
	IETYPE_DURATION_MEASUREMENT              IEType = 67
	IETYPE_APPLICATION_INSTANCE_ID           IEType = 91

	IETYPE_USAGE_REPORT         IEType = 80
	IETYPE_UR_SEQN              IEType = 104
	IETYPE_USAGE_REPORT_TRIGGER IEType = 63
	IETYPE_START_TIME           IEType = 75
	IETYPE_END_TIME             IEType = 76
	IETYPE_FLOW_INFORMATON      IEType = 92

	IETYPE_UP_Function_Features               IEType = 43
	IETYPE_CP_Function_Features               IEType = 89
	IETYPE_USER_PLANE_IP_RESOURCE_INFORMATION IEType = 116

	IETYPE_APPLICATION_IDS_PFDS IEType = 58
	IETYPE_PFD_CONTENT          IEType = 59
	IETYPE_PFD_CONTENTS         IEType = 61

	IETYPE_CREATE_URR         IEType = 6
	IETYPE_MEASUREMENT_METHOD IEType = 62
	IETYPE_REPORTING_TRIGGERS IEType = 37

	IETYPE_CREATE_QER                   IEType = 7
	IETYPE_GATE_STATUS                  IEType = 25
	IETYPE_MBR                          IEType = 26
	IETYPE_GBR                          IEType = 27
	IETYPE_QER_CORRELATION_ID           IEType = 28
	IETYPE_PACKET_RATE                  IEType = 94
	IETYPE_DL_FLOW_LEVEL_MARKING        IEType = 97
	IETYPE_RQI                          IEType = 123
	IETYPE_AVERAGING_WINDOW             IEType = 157
	IETYPE_PPI                          IEType = 158
	IETYPE_PDN_TYPE                     IEType = 113
	IETYPE_UPDATE_PDR                   IEType = 9
	IETYPE_UPDATE_FAR                   IEType = 10
	IETYPE_MEASURE_PERIOD               IEType = 64
	IETYPE_UPDATE_FORWARDING_PARAMETERS IEType = 11
	//TODO::
	CREATE_BAR IEType = 85
	FAR        IEType = 0
	//C
	URR IEType = 0
	QER IEType = 0
	BAR IEType = 0 //O

	TRAFFIC_ENDPOINT IEType = 0
	PDN_TYPE         IEType = 113
	SGW_C_FQ_CSID    IEType = 0
	MME_FQ_CSID      IEType = 0
	PGW_C_FQ_CSID    IEType = 0
	TWAN_FQ_CSID     IEType = 0
	//O
	USER_PLANE_INACTIVITY_TIMER IEType = 0
	USER_ID                     IEType = 0
	TRACE_INFOMATION            IEType = 0
	APN_OR_DNN                  IEType = 0

	PRECEDENCE                IEType = 29
	PDI                       IEType = 0
	OUTER_HEADER_REMOVAL      IEType = 95
	OUTER_HEADER_CREATION     IEType = 84
	FAR_ID                    IEType = 108
	URR_ID                    IEType = 0
	QER_ID                    IEType = 0
	ACTIVATE_PREDEFINED_RULES IEType = 106

	//PDI IE
	SOURCE_INTERFACE IEType = 0
	F_TEID           IEType = 0
	NEWWORK_INSTANCE IEType = 0
	UE_IP_ADDR       IEType = 0

	SDF_FILTER     IEType = 23
	APPLICATION_ID IEType = 0
	//....

	FORWARDING_POLICY IEType = 41

	REPORTING_TRIGGERS IEType = 37
	MEASURE_PERIOD     IEType = 64

	QER_CORRELATION_ID IEType = 28
	GATE_STATUS        IEType = 25
	MBR_TYPE           IEType = 26
	GBR_TYPE           IEType = 27

	PACKET_RATE      IEType = 94
	AVERAGING_WINDOW IEType = 156
	FQ_CSID          IEType = 65
)
