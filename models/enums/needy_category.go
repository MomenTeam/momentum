package enums

type NeedyCategoryType int

const (
	None         NeedyCategoryType = 0
	WithChildren NeedyCategoryType = 1
	Elder        NeedyCategoryType = 2
)

func GenerateNeedyCategoryFromInt(category int) NeedyCategoryType {
	if category == 1 {
		return WithChildren
	} else if category == 2 {
		return Elder
	} else {
		return None
	}
}
