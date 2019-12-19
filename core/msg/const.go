package msg

type PFCPMsgType byte

const (
	PFCPMsgTypeHeartbeatRequest  PFCPMsgType = 1
	PFCPMsgTypeHeartbeatResponse PFCPMsgType = 2

	PFCPMsgTypePFDManagementRequest  PFCPMsgType = 3
	PFCPMsgTypePFDManagementResponse PFCPMsgType = 4

	PFCPMsgTypeAssociationSetupRequest    PFCPMsgType = 5
	PFCPMsgTypeAssociationSetupResponse   PFCPMsgType = 6
	PFCPMsgTypeAssociationUpdateRequest   PFCPMsgType = 7
	PFCPMsgTypeAssociationUpdateResponse  PFCPMsgType = 8
	PFCPMsgTypeAssociationReleaseRequest  PFCPMsgType = 9
	PFCPMsgTypeAssociationReleaseResponse PFCPMsgType = 10

	PFCPMsgTypeVersionNotSupported PFCPMsgType = 11

	PFCPMsgTypeNodeReportRequest  PFCPMsgType = 12
	PFCPMsgTypeNodeReportResponse PFCPMsgType = 13

	PFCPMsgTypeSessionSetDeletionRequest    PFCPMsgType = 14
	PFCPMsgTypeSessionSetDeletionResponse   PFCPMsgType = 15
	PFCPMsgTypeSessionEstablishmentRequest  PFCPMsgType = 50
	PFCPMsgTypeSessionEstablishmentResponse PFCPMsgType = 51
	PFCPMsgTypeSessionModificationRequest   PFCPMsgType = 52
	PFCPMsgTypeSessionModificationResponse  PFCPMsgType = 53
	PFCPMsgTypeSessionDeletionRequest       PFCPMsgType = 54
	PFCPMsgTypeSessionDeletionResponse      PFCPMsgType = 55
	PFCPMsgTypeSessionReportRequest         PFCPMsgType = 56
	PFCPMsgTypeSessionReportResponse        PFCPMsgType = 57
)
