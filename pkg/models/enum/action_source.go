package enum

type ActionSourceType string
type actionSourceType struct {
	GHN_ACCOUNT ActionSourceType
	API_PARTNER ActionSourceType
}

var ActionSourceTypeEnum = &actionSourceType{
	GHN_ACCOUNT: "GHN_ACCOUNT",
	API_PARTNER: "API_PARTNER",
}
