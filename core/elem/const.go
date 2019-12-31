package elem

type IEType uint16

const (
	//Grouped IE, extendable
	IETypeCreatePDR                  IEType = 1
	IETypePDI                        IEType = 2
	IETypeCreateFAR                  IEType = 3
	IETypeForwardingParameters       IEType = 4
	IETypeDuplicatingParameters      IEType = 5
	IETypeCreateURR                  IEType = 6
	IETypeCreateQER                  IEType = 7
	IETypeCreatedPDR                 IEType = 8
	IETypeUpdatePDR                  IEType = 9
	IETypeUpdateFAR                  IEType = 10
	IETypeUpdateForwardingParameters IEType = 11
	IETypeUpdateBAR                  IEType = 12 //PFCP Session Report Response
	IETypeUpdateURR                  IEType = 13
	IETypeUpdateQER                  IEType = 14
	IETypeRemovePDR                  IEType = 15
	IETypeRemoveFAR                  IEType = 16
	IETypeRemoveURR                  IEType = 17
	IETypeRemoveQER                  IEType = 18
	//
	IETypeCause                                           IEType = 19 //decimal
	IETypeSourceInterface                                 IEType = 20
	IETypeFTEID                                           IEType = 21
	IETypeNetworkInstance                                 IEType = 22
	IETypeSDFFilter                                       IEType = 23
	IETypeApplicationID                                   IEType = 24
	IETypeGateStatus                                      IEType = 25
	IETypeMBR                                             IEType = 26
	IETypeGBR                                             IEType = 27
	IETypeQERCorrelationID                                IEType = 28
	IETypePrecedence                                      IEType = 29
	IETypeTransportLevelMarking                           IEType = 30
	IETypeVolumeThreshold                                 IEType = 31
	IETypeTimeThreshold                                   IEType = 32
	IETypeMonitoringTime                                  IEType = 33
	IETypeSubsequentVolumeThreshold                       IEType = 34
	IETypeSubsequentTimeThreshold                         IEType = 35
	IETypeInactivityDetectionTime                         IEType = 36
	IETypeReportingTriggers                               IEType = 37
	IETypeRedirectInformation                             IEType = 38
	IETypeReportType                                      IEType = 39
	IETypeOffendingIE                                     IEType = 40
	IETypeForwardingPolicy                                IEType = 41
	IETypeDestinationInterface                            IEType = 42
	IETypeUPFunctionFeatures                              IEType = 43
	IETypeApplyAction                                     IEType = 44
	IETypeDownLinkDataServiceInformation                  IEType = 45
	IETypeDownLinkDataNotificationDelay                   IEType = 46
	IETypeDLBufferingDuration                             IEType = 47
	IETypeDLBufferingSuggestedPacketCount                 IEType = 48
	IETypePFCPSMReqFlags                                  IEType = 49
	IETypePFCPSRRspFlags                                  IEType = 50
	IETypeSequenceNumber                                  IEType = 52
	IETypeMetric                                          IEType = 53
	IETypeTimer                                           IEType = 55
	IETypePDRID                                           IEType = 56 //PacketDetectionRuleID
	IETypeFSEID                                           IEType = 57
	IETypeNodeID                                          IEType = 60
	IETypePFDContents                                     IEType = 61
	IETypeMeasurementMethod                               IEType = 62
	IETypeUsageReportTrigger                              IEType = 63
	IETypeMeasurementPeriod                               IEType = 64
	IETypeFullyQualifiedPDNConnectionSetIdentifier        IEType = 65 //FQ-CSID
	IETypeVolumeMeasurement                               IEType = 66
	IETypeDurationMeasurement                             IEType = 67
	IETypeTimeOfFirstPacket                               IEType = 69
	IETypeTimeOfLastPacket                                IEType = 70
	IETypeQuotaHoldingTime                                IEType = 71
	IETypeDroppedDLTrafficThreshold                       IEType = 72
	IETypeVolumeQuota                                     IEType = 73
	IETypeTimeQuota                                       IEType = 74
	IETypeStartTime                                       IEType = 75
	IETypeEndTime                                         IEType = 76
	IETypeURRID                                           IEType = 81
	IETypeLinkedURRID                                     IEType = 82
	IETypeOuterHeaderCreation                             IEType = 84
	IETypeBARID                                           IEType = 88
	IETypeCPFunctionFeatures                              IEType = 89
	IETypeUsageInformation                                IEType = 90
	IETypeApplicationInstanceID                           IEType = 91
	IETypeFlowInformation                                 IEType = 92
	IETypeUEIPAddress                                     IEType = 93
	IETypePacketRate                                      IEType = 94
	IETypeOuterHeaderRemoval                              IEType = 95
	IETypeRecoveryTimeStamp                               IEType = 96
	IETypeDLFlowLevelMarking                              IEType = 97
	IETypeHeaderEnrichment                                IEType = 98
	IETypeMeasurementInformation                          IEType = 100
	IETypeNodeReportType                                  IEType = 101
	IETypeRemoteGTPUPeer                                  IEType = 103
	IETypeURSEQN                                          IEType = 104
	IETypeActivatePredefinedRules                         IEType = 106
	IETypeDeactivatePredefinedRules                       IEType = 107
	IETypeFARID                                           IEType = 108
	IETypeQERID                                           IEType = 109
	IETypeOCIFlags                                        IEType = 100
	IETypePFCPAssociationReleaseRequest                   IEType = 111
	IETypeGracefulReleasePeriod                           IEType = 112
	IETypePDNType                                         IEType = 113
	IETypeFailedRuleID                                    IEType = 114
	IETypeTimeQuotaMechanism                              IEType = 115
	IETypeUserPlaneIPResourceInformation                  IEType = 116
	IETypeUserPlaneInactivityTimer                        IEType = 117
	IETypeMultiplier                                      IEType = 119
	IETypeAggregatedURRID                                 IEType = 120
	IETypeSubsequentVolumeQuota                           IEType = 121
	IETypeSubsequentTimeQuota                             IEType = 122
	IETypeRQI                                             IEType = 123
	IETypeQFI                                             IEType = 124
	IETypeQueryURRReference                               IEType = 125
	IETypeAdditionalUsageReportsInformation               IEType = 126
	IETypeTrafficEndpointID                               IEType = 131
	IETypeMACAddress                                      IEType = 133
	IETypeCTAG                                            IEType = 134 //Customer-VLAN tag
	IETypeSTAG                                            IEType = 135 //Service-VLAN tag
	IETypeEtherType                                       IEType = 136
	IETypeProxying                                        IEType = 137
	IETypeEthernetFilterID                                IEType = 138
	IETypeEthernetFilterProperties                        IEType = 139
	IETypeSuggestedBufferingPacketsCount                  IEType = 140
	IETypeUserID                                          IEType = 141
	IETypeEthernetPDUSessionInformation                   IEType = 142
	IETypeMACAddressesDetected                            IEType = 144
	IETypeMACAddressesRemoved                             IEType = 145
	IETypeEthernetInactivityTimer                         IEType = 146
	IETypeEventQuota                                      IEType = 148
	IETypeEventThreshold                                  IEType = 149
	IETypeSubsequentEventQuota                            IEType = 150
	IETypeSubsequentEventThreshold                        IEType = 151
	IETypeTraceInformation                                IEType = 152
	IETypeFramedRoute                                     IEType = 153
	IETypeFramedRouting                                   IEType = 154
	IETypeFramedIPv6Route                                 IEType = 155
	IETypeEventTimeStamp                                  IEType = 156
	IETypeAveragingWindow                                 IEType = 157
	IETypePagingPolicyIndicator                           IEType = 158 //PPI
	IETypeAPNDNN                                          IEType = 159
	IEType3GPPInterfaceType                               IEType = 160
	IETypePFCPSRReqFlags                                  IEType = 161
	IETypePFCPAUReqFlags                                  IEType = 162
	IETypeActivationTime                                  IEType = 163
	IETypeDeactivationTime                                IEType = 164
	IETypeMARID                                           IEType = 170
	IETypeSteeringFunctionality                           IEType = 171
	IETypeSteeringMode                                    IEType = 172
	IETypeWeight                                          IEType = 173
	IETypePriority                                        IEType = 174
	IETypeUEIPAddressPoolIdentity                         IEType = 177
	IETypeAlternativeSMFIPAddress                         IEType = 178
	IETypePacketReplicationAndDetectionCarryOnInformation IEType = 179
	IETypeSMFSetID                                        IEType = 180
	IETypeQuotaValidityTime                               IEType = 181
)

type InterfaceType byte

const (
	InterfaceTypeAccess       InterfaceType = iota //0
	InterfaceTypeCore                              //1
	InterfaceTypeSGiLAN                            //2
	InterfaceTypeCPFunction                        //3
	InterfaceType5GVNInternal                      //4
)
