package gosf

type Acct struct {
	Account     string   `json:"account" bson:"account"`
	AccountInfo AcctInfo `json:"accountInfo" bson:"accountInfo"`
}

type AcctInfo struct {
	Slot     int64    `json:"slot" bson:"slot"`
	Pubkey   string   `json:"pubkey" bson:"pubkey"`
	Lamports int64    `json:"lamports" bson:"lamports"`
	Owner    string   `json:"owner" bson:"owner"`
	Data     AcctData `json:"data" bson:"data"`
}

type AcctData struct {
	Version      int64   `json:"version" bson:"version"`
	Bump         []int64 `json:"bump" bson:"bump"`
	Owner        string  `json:"owner" bson:"owner"`
	AssetId      string  `json:"assetId" bson:"assetId"`
	Amount       string  `json:"amount" bson:"amount"`
	Currency     *string `json:"currency" bson:"currency"`
	Expiry       string  `json:"expiry" bson:"expiry"`
	PrivateTaker *string `json:"privateTaker" bson:"privateTaker"`
	MakerBroker  *string `json:"makerBroker" bson:"makerBroker"`
	Raw          string  `json:"raw" bson:"raw"`
}
