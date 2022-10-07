package enum

import "gitlab.ghn.vn/nhanh/backend/work-shift-service/pkg/constants/message"

type UserActionHistory int

var (
	LayoutHistoryCreate       UserActionHistory = 1
	LayoutHistoryUpdate       UserActionHistory = 2
	LayoutHistoryDelete       UserActionHistory = 3
	LayoutHistoryActive       UserActionHistory = 4
	LayoutHistoryDeactive     UserActionHistory = 5
	AreaHistoryCreate         UserActionHistory = 6
	AreaHistoryUpdate         UserActionHistory = 7
	AreaHistoryDelete         UserActionHistory = 8
	ShiftHistoryImport        UserActionHistory = 9
	ShiftHistoryUpdate        UserActionHistory = 10
	ShiftHistoryDelete        UserActionHistory = 11
	ScheduleHistoryDelete     UserActionHistory = 12
	AreaHistorySwitchUserArea UserActionHistory = 13
	AreaHistoryCheckoutUser   UserActionHistory = 14
	SoundConfigHistoryImport  UserActionHistory = 15
	ControlInOutHistory       UserActionHistory = 16
	ScanMemberHistory         UserActionHistory = 17
)

func ParaseUserActionHistory(data UserActionHistory) string {
	switch data {
	case LayoutHistoryCreate:
		return message.CreateLayout
	case LayoutHistoryUpdate:
		return message.UpdateLayout
	case LayoutHistoryDelete:
		return message.DeleteLayout
	case LayoutHistoryActive:
		return message.ActivateLayout
	case LayoutHistoryDeactive:
		return message.DeactivateLayout
	case AreaHistoryCreate:
		return message.CreateArea
	case AreaHistoryUpdate:
		return message.UpdateArea
	case AreaHistoryDelete:
		return message.DeleteArea
	case ShiftHistoryImport:
		return message.ImportShift
	case ShiftHistoryUpdate:
		return message.UpdateShift
	case ShiftHistoryDelete:
		return message.DeleteShift
	case ScheduleHistoryDelete:
		return message.DeleteSchedule
	case AreaHistorySwitchUserArea:
		return message.SwitchUserArea
	case AreaHistoryCheckoutUser:
		return message.CheckoutUserArea
	case SoundConfigHistoryImport:
		return message.SoundConfigImport
	case ControlInOutHistory:
		return message.ControlInOut
	case ScanMemberHistory:
		return message.ScanMember
	default:
		return ""
	}
}
