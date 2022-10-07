package enum

type ItemType int

var (
	ItemTypeMoney    ItemType = 1
	ItemTypeTicket   ItemType = 2
	ItemTypeArtifact ItemType = 3
)

type ItemStatus int

var (
	ItemStatusActive  ItemStatus = 1
	ItemStatusDisable ItemStatus = 2
	ItemStatusExpire  ItemStatus = 3
)

type InventoryStatus int

var (
	InventoryStatusDefault   InventoryStatus = 1
	InventoryStatusReady     InventoryStatus = 2
	InventoryStatusRequested InventoryStatus = 3
	InventoryStatusConfirmed InventoryStatus = 4
)

type ItemSourceType int

var (
	ItemSourceTypeSystem    ItemSourceType = 8888
	ItemSourceTypeAdminGift ItemSourceType = 1000
	ItemSourceTypeCampaign  ItemSourceType = 2000
	ItemSourceTypeGame      ItemSourceType = 3000
)

type ActionId int32
type ActionName string

const (
	// moneys
	ActionClaimReward         ActionId   = 1
	ActionNameClaimReward     ActionName = "Thêm vào thùng đồ"
	ActionMoneyUsedReward     ActionId   = 2
	ActionNameMoneyUsedReward ActionName = "Chi tiền từ thùng đồ"

	// tickets
	ActionTicketClaimReward      ActionId   = 3
	ActionTicketUsedReward       ActionId   = 4
	ActionNameTicketUsedReward   ActionName = "Sử dụng vật phẩm"
	ActionTicketCancelReward     ActionId   = 5
	ActionNameTicketCancelReward ActionName = "Hủy vé"

	// artifacts
	ActionArtifactClaimReward         ActionId   = 6
	ActionArtifactReadyReward         ActionId   = 7
	ActionNameArtifactReadyReward     ActionName = "Mở khóa quà"
	ActionArtifactRequestedReward     ActionId   = 8
	ActionNameArtifactRequestedReward ActionName = "Đòi quà"
	ActionArtifactConfirmedReward     ActionId   = 9
	ActionNameArtifactConfirmedReward ActionName = "Phát quà"
	ActionArtifactReceivedReward      ActionId   = 10
	ActionNameArtifactReceivedReward  ActionName = "Đã nhận quà"
	ActionArtifactExpiredReward       ActionId   = 11
	ActionNameArtifactExpiredReward   ActionName = "Hết hạn"
)

func (action ActionId) String() string {
	switch action {
	case ActionClaimReward:
		return string(ActionNameClaimReward)

	case ActionMoneyUsedReward:
		return string(ActionNameMoneyUsedReward)

	case ActionTicketClaimReward:
		return string(ActionNameClaimReward)

	case ActionTicketUsedReward:
		return string(ActionNameTicketUsedReward)

	case ActionTicketCancelReward:
		return string(ActionNameTicketCancelReward)

	case ActionArtifactClaimReward:
		return string(ActionNameClaimReward)

	case ActionArtifactReadyReward:
		return string(ActionNameArtifactReadyReward)

	case ActionArtifactRequestedReward:
		return string(ActionNameArtifactRequestedReward)

	case ActionArtifactConfirmedReward:
		return string(ActionNameArtifactConfirmedReward)

	case ActionArtifactReceivedReward:
		return string(ActionNameArtifactReceivedReward)

	case ActionArtifactExpiredReward:
		return string(ActionNameArtifactExpiredReward)
	}

	return ""
}
