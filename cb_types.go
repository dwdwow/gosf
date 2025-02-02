package gosf

type CBTxProtocol struct {
	Address string `json:"address" bson:"address"`
	Name    string `json:"name" bson:"name"`
}

type CBTxActionInfo struct {
	Sender   string  `json:"sender" bson:"sender"`
	Receiver string  `json:"receiver" bson:"receiver"`
	Amount   float64 `json:"amount" bson:"amount"`
}

type CBTxAction struct {
	Info           CBTxActionInfo `json:"info" bson:"info"`
	SourceProtocol CBTxProtocol   `json:"source_protocol" bson:"source_protocol"`
	Type           string         `json:"type" bson:"type"`
}

type CBTxAccountData struct {
	Address  string `json:"address" bson:"address"`
	Owner    string `json:"owner" bson:"owner"`
	Lamports int64  `json:"lamports" bson:"lamports"`
	Data     any    `json:"data" bson:"data"`
}

type CBTxRawMeta struct {
	ComputeUnitsConsumed int64          `json:"computeUnitsConsumed" bson:"computeUnitsConsumed"`
	Err                  any            `json:"err" bson:"err"`
	Fee                  int64          `json:"fee" bson:"fee"`
	InnerInstructions    []any          `json:"innerInstructions" bson:"innerInstructions"`
	LogMessages          []string       `json:"logMessages" bson:"logMessages"`
	PostBalances         []int64        `json:"postBalances" bson:"postBalances"`
	PostTokenBalances    []any          `json:"postTokenBalances" bson:"postTokenBalances"`
	PreBalances          []int64        `json:"preBalances" bson:"preBalances"`
	PreTokenBalances     []any          `json:"preTokenBalances" bson:"preTokenBalances"`
	Rewards              []any          `json:"rewards" bson:"rewards"`
	Status               map[string]any `json:"status" bson:"status"`
}

type CBTxAccountKey struct {
	Pubkey   string `json:"pubkey" bson:"pubkey"`
	Signer   bool   `json:"signer" bson:"signer"`
	Source   string `json:"source" bson:"source"`
	Writable bool   `json:"writable" bson:"writable"`
}

type CBTxInstructionInfo struct {
	Destination string `json:"destination" bson:"destination"`
	Lamports    int64  `json:"lamports" bson:"lamports"`
	Source      string `json:"source" bson:"source"`
}

type CBTxInstruction struct {
	Parsed struct {
		Info CBTxInstructionInfo `json:"info" bson:"info"`
		Type string              `json:"type" bson:"type"`
	} `json:"parsed" bson:"parsed"`
	Program   string `json:"program" bson:"program"`
	ProgramId string `json:"programId" bson:"programId"`
}

type CBTxMessage struct {
	AccountKeys         []CBTxAccountKey  `json:"accountKeys" bson:"accountKeys"`
	AddressTableLookups any               `json:"addressTableLookups" bson:"addressTableLookups"`
	Instructions        []CBTxInstruction `json:"instructions" bson:"instructions"`
	RecentBlockhash     string            `json:"recentBlockhash" bson:"recentBlockhash"`
}

type CBTxTransaction struct {
	Message    CBTxMessage `json:"message" bson:"message"`
	Signatures []string    `json:"signatures" bson:"signatures"`
}

type CBTxParsed struct {
	Timestamp  string       `json:"timestamp" bson:"timestamp"`
	Fee        float64      `json:"fee" bson:"fee"`
	FeePayer   string       `json:"fee_payer" bson:"fee_payer"`
	Signers    []string     `json:"signers" bson:"signers"`
	Signatures []string     `json:"signatures" bson:"signatures"`
	Protocol   CBTxProtocol `json:"protocol" bson:"protocol"`
	Type       string       `json:"type" bson:"type"`
	Status     string       `json:"status" bson:"status"`
	Actions    []CBTxAction `json:"actions" bson:"actions"`
}

type CBTxRawTransaction struct {
	BlockTime   int64           `json:"blockTime" bson:"blockTime"`
	Meta        CBTxRawMeta     `json:"meta" bson:"meta"`
	Slot        int64           `json:"slot" bson:"slot"`
	Transaction CBTxTransaction `json:"transaction" bson:"transaction"`
	Version     string          `json:"version" bson:"version"`
	Parsed      CBTxParsed      `json:"parsed" bson:"parsed"`
}

type CBTx struct {
	Timestamp  string             `json:"timestamp" bson:"timestamp"`
	Fee        float64            `json:"fee" bson:"fee"`
	FeePayer   string             `json:"fee_payer" bson:"fee_payer"`
	Signers    []string           `json:"signers" bson:"signers"`
	Signatures []string           `json:"signatures" bson:"signatures"`
	Protocol   CBTxProtocol       `json:"protocol" bson:"protocol"`
	Type       string             `json:"type" bson:"type"`
	Status     string             `json:"status" bson:"status"`
	Actions    []CBTxAction       `json:"actions" bson:"actions"`
	Raw        CBTxRawTransaction `json:"raw" bson:"raw"`
	Accounts   []CBTxAccountData  `json:"accounts" bson:"accounts"`
}

type CBAcct struct {
	Account     string     `json:"account" bson:"account"`
	AccountInfo CBAcctInfo `json:"accountInfo" bson:"accountInfo"`
}

type CBAcctInfo struct {
	Slot     int64      `json:"slot" bson:"slot"`
	Pubkey   string     `json:"pubkey" bson:"pubkey"`
	Lamports int64      `json:"lamports" bson:"lamports"`
	Owner    string     `json:"owner" bson:"owner"`
	Data     CBAcctData `json:"data" bson:"data"`
}

type CBAcctData struct {
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
