package enum

type CampaignType string
type CampaignTypeName string
type campaignType struct {
	NewClient        CampaignType
	NewClientRevenue CampaignType
}

var CampaignTypeEnum = &campaignType{
	NewClientRevenue: "new_client_revenue",
	NewClient:        "num_new_client",
}

const (
	CampaignTypeDefaultName          CampaignTypeName = "Điểm"
	CampaignTypeNewClientName        CampaignTypeName = "SL KH"
	CampaignTypeNewClientRevenueName CampaignTypeName = "Doanh thu"
)

func (action CampaignType) String() string {
	switch action {
	case CampaignTypeEnum.NewClient:
		return string(CampaignTypeEnum.NewClient)
	case CampaignTypeEnum.NewClientRevenue:
		return string(CampaignTypeEnum.NewClientRevenue)
	}
	return ""
}

type CampaignUserType string
type campaignUserType struct {
	Employee CampaignUserType
	Client   CampaignUserType
}

var CampaignUserTypeEnum = &campaignUserType{
	Employee: "employee",
	Client:   "client",
}

type CampaignStatus string
type campaignStatus struct {
	Prepare      CampaignStatus
	Started      CampaignStatus
	NotQualified CampaignStatus
	CalcResult   CampaignStatus
	Finished     CampaignStatus
	Cancel       CampaignStatus
}

var CampaignStatusEnum = &campaignStatus{
	Prepare:      "prepare",
	Started:      "started",
	NotQualified: "not_qualified",
	CalcResult:   "calc_result",
	Finished:     "finished",
	Cancel:       "cancel",
}

var StatusDelete = []string{
	string(CampaignStatusEnum.Prepare),
	string(CampaignStatusEnum.NotQualified),
}
