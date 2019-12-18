package elem

type CauseType byte

const (
	CauseReserved                          CauseType = 0
	CauseSuccess                           CauseType = 1 //2-63,spare
	CauseUnspecifiedReason                 CauseType = 64
	CauseSessionContextNoFound             CauseType = 65 //fseid pfcp session modification/deletion is unknown
	CauseMandatoryIEMissing                CauseType = 66
	CauseConditionalIEMissing              CauseType = 67
	CauseInvalidLength                     CauseType = 68
	CauseMandatoryIEIncorrect              CauseType = 69 //eg:is malformed or it carries an invalid oe unexpected value
	CauseInvalidForwardingPolicy           CauseType = 70
	CauseInvalidFTEIDAllocation            CauseType = 71 //same TEID
	CauseNoEstablishedPFCPAssociation      CauseType = 72
	CauseRuleCreationOrModificationFailure CauseType = 73 //failed to be stored
	CausePFCPEntityInCongestion            CauseType = 74
	CauseNoResourcesAvailable              CauseType = 75
	CauseServiceNotSupported               CauseType = 76
	CauseSystemFailure                     CauseType = 77 //78-255 ,SPARE
)

type Cause struct {
	EType   IEType
	ELength uint16
	Cause   CauseType
}

func NewCause(cause CauseType) *Cause {
	return &Cause{
		EType:   IETypeCause,
		ELength: uint16(1),
		Cause:   cause,
	}
}

func DecodeCause(data []byte, len uint16) *Cause {
	return &Cause{
		EType:   IETypeCause,
		ELength: len,
		Cause:   CauseType(getValue(data, len-1)[0]),
	}
}

func EncodeCause(cause Cause) []byte {
	return setBuffer(cause.EType, cause.ELength, []byte{byte(cause.Cause)})
}
