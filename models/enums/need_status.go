package enums

type NeedStatus int

const (
	NeedCreated   NeedStatus = 0
	NeedPaid      NeedStatus = 1
	NeedFulfilled NeedStatus = 2
	NeedCancelled NeedStatus = -1
)
