package enum

type GHNInternalStatus string

type internalNewStatus struct {
	StatusDraft                  GHNInternalStatus
	StatusReadyToPick            GHNInternalStatus
	StatusPicking                GHNInternalStatus
	StatusMoneyCollectPicking    GHNInternalStatus
	StatusPicked                 GHNInternalStatus
	StatusStoring                GHNInternalStatus
	StatusSorting                GHNInternalStatus
	StatusTransporting           GHNInternalStatus
	StatusDelivering             GHNInternalStatus
	StatusDeliveryFail           GHNInternalStatus
	StatusMoneyCollectDelivering GHNInternalStatus
	StatusDelivered              GHNInternalStatus
	StatusReadyToReturn          GHNInternalStatus
	StatusReturn                 GHNInternalStatus
	StatusReturning              GHNInternalStatus
	StatusReturnFail             GHNInternalStatus
	StatusReturnTransporting     GHNInternalStatus
	StatusReturnSorting          GHNInternalStatus
	StatusReturned               GHNInternalStatus
	StatusWaitingToFinish        GHNInternalStatus
	StatusFinish                 GHNInternalStatus
	StatusCancel                 GHNInternalStatus
	StatusException              GHNInternalStatus
	StatusWaitingToReturn        GHNInternalStatus
	StatusLost                   GHNInternalStatus
	StatusDamage                 GHNInternalStatus
}

var InNewStatusEnum = &internalNewStatus{
	StatusDraft:                  "draft",
	StatusReadyToPick:            "ready_to_pick",
	StatusPicking:                "picking",
	StatusMoneyCollectPicking:    "money_collect_picking",
	StatusPicked:                 "picked",
	StatusStoring:                "storing",
	StatusSorting:                "sorting",
	StatusTransporting:           "transporting",
	StatusDelivering:             "delivering",
	StatusDeliveryFail:           "delivery_fail",
	StatusMoneyCollectDelivering: "money_collect_delivering",
	StatusDelivered:              "delivered",
	StatusReadyToReturn:          "ready_to_return",
	StatusReturn:                 "return",
	StatusReturning:              "returning",
	StatusReturnFail:             "return_fail",
	StatusReturnTransporting:     "return_transporting",
	StatusReturnSorting:          "return_sorting",
	StatusReturned:               "returned",
	StatusWaitingToFinish:        "waiting_to_finish",
	StatusFinish:                 "finish",
	StatusCancel:                 "cancel",
	StatusException:              "exception",
	StatusWaitingToReturn:        "waiting_to_return",
	StatusLost:                   "lost",
	StatusDamage:                 "damage",
}

type ghnInternalStatus struct {
	ReadyToPick           GHNInternalStatus
	Cancel                GHNInternalStatus
	Delivering            GHNInternalStatus
	Delivered             GHNInternalStatus
	Picking               GHNInternalStatus
	Picked                GHNInternalStatus
	Storing               GHNInternalStatus
	NotPicked             GHNInternalStatus
	NotReturn             GHNInternalStatus
	NotReturned           GHNInternalStatus
	LostOrder             GHNInternalStatus
	S2R                   GHNInternalStatus
	ConfirmedReturn       GHNInternalStatus
	WaitingToFinish       GHNInternalStatus
	R2DA                  GHNInternalStatus
	OnReturn              GHNInternalStatus
	Return                GHNInternalStatus
	RReturnToStoring      GHNInternalStatus
	RStoringToReadyToPick GHNInternalStatus
	RCancelToReadyToPick  GHNInternalStatus
	Finish                GHNInternalStatus
	NotDelivered          GHNInternalStatus
	StationPicked         GHNInternalStatus
	StationDelivered      GHNInternalStatus
	RPickingToReadyToPick GHNInternalStatus
	RDeliveredToStoring   GHNInternalStatus
	NewStationPicked      GHNInternalStatus
	StationCancel         GHNInternalStatus
}

var GHNInternalStatusEnum = &ghnInternalStatus{
	ReadyToPick:           "ReadyToPick",
	Cancel:                "Cancel",
	Delivering:            "Delivering",
	Delivered:             "Delivered",
	Picking:               "Picking",
	Picked:                "Picked",
	Storing:               "Storing",
	NotPicked:             "NotPicked",
	NotReturn:             "NotReturn",
	NotReturned:           "NotReturned",
	LostOrder:             "LostOrder",
	S2R:                   "S2R",
	ConfirmedReturn:       "ConfirmedReturn",
	WaitingToFinish:       "WaitingToFinish",
	R2DA:                  "R2DA",
	OnReturn:              "OnReturn",
	Return:                "Return",
	RReturnToStoring:      "RReturnToStoring",
	RStoringToReadyToPick: "RStoringToReadyToPick",
	RCancelToReadyToPick:  "RCancelToReadyToPick",
	Finish:                "Finish",
	NotDelivered:          "NotDelivered",
	StationPicked:         "StationPicked",
	StationDelivered:      "StationDelivered",
	RPickingToReadyToPick: "RPickingToReadyToPick",
	RDeliveredToStoring:   "RDeliveredToStoring",
	NewStationPicked:      "NewStationPicked",
	StationCancel:         "StationCancel",
}
