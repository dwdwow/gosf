package gosf

type CBTxProtocol struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type CBTxActionInfo struct {
	Sender   string  `json:"sender"`
	Receiver string  `json:"receiver"`
	Amount   float64 `json:"amount"`
}

type CBTxAction struct {
	Info           CBTxActionInfo `json:"info"`
	SourceProtocol CBTxProtocol   `json:"source_protocol"`
	Type           string         `json:"type"`
}

type CBTxAccountData struct {
	Address  string `json:"address"`
	Owner    string `json:"owner"`
	Lamports int64  `json:"lamports"`
	Data     any    `json:"data"`
}

type CBTxRawMeta struct {
	ComputeUnitsConsumed int64          `json:"computeUnitsConsumed"`
	Err                  any            `json:"err"`
	Fee                  int64          `json:"fee"`
	InnerInstructions    []any          `json:"innerInstructions"`
	LogMessages          []string       `json:"logMessages"`
	PostBalances         []int64        `json:"postBalances"`
	PostTokenBalances    []any          `json:"postTokenBalances"`
	PreBalances          []int64        `json:"preBalances"`
	PreTokenBalances     []any          `json:"preTokenBalances"`
	Rewards              []any          `json:"rewards"`
	Status               map[string]any `json:"status"`
}

type CBTxAccountKey struct {
	Pubkey   string `json:"pubkey"`
	Signer   bool   `json:"signer"`
	Source   string `json:"source"`
	Writable bool   `json:"writable"`
}

type CBTxInstructionInfo struct {
	Destination string `json:"destination"`
	Lamports    int64  `json:"lamports"`
	Source      string `json:"source"`
}

type CBTxInstruction struct {
	Parsed struct {
		Info CBTxInstructionInfo `json:"info"`
		Type string              `json:"type"`
	} `json:"parsed"`
	Program   string `json:"program"`
	ProgramId string `json:"programId"`
}

type CBTxMessage struct {
	AccountKeys         []CBTxAccountKey  `json:"accountKeys"`
	AddressTableLookups any               `json:"addressTableLookups"`
	Instructions        []CBTxInstruction `json:"instructions"`
	RecentBlockhash     string            `json:"recentBlockhash"`
}

type CBTxTransaction struct {
	Message    CBTxMessage `json:"message"`
	Signatures []string    `json:"signatures"`
}

type CBTxParsed struct {
	Timestamp  string       `json:"timestamp"`
	Fee        float64      `json:"fee"`
	FeePayer   string       `json:"fee_payer"`
	Signers    []string     `json:"signers"`
	Signatures []string     `json:"signatures"`
	Protocol   CBTxProtocol `json:"protocol"`
	Type       string       `json:"type"`
	Status     string       `json:"status"`
	Actions    []CBTxAction `json:"actions"`
}

type CBTxRawTransaction struct {
	BlockTime   int64           `json:"blockTime"`
	Meta        CBTxRawMeta     `json:"meta"`
	Slot        int64           `json:"slot"`
	Transaction CBTxTransaction `json:"transaction"`
	Version     string          `json:"version"`
	Parsed      CBTxParsed      `json:"parsed"`
}

type CBTx struct {
	Timestamp  string             `json:"timestamp"`
	Fee        float64            `json:"fee"`
	FeePayer   string             `json:"fee_payer"`
	Signers    []string           `json:"signers"`
	Signatures []string           `json:"signatures"`
	Protocol   CBTxProtocol       `json:"protocol"`
	Type       string             `json:"type"`
	Status     string             `json:"status"`
	Actions    []CBTxAction       `json:"actions"`
	Raw        CBTxRawTransaction `json:"raw"`
	Accounts   []CBTxAccountData  `json:"accounts"`
}

type CBAcct struct {
	Account     string     `json:"account"`
	AccountInfo CBAcctInfo `json:"accountInfo"`
}

type CBAcctInfo struct {
	Slot     int64      `json:"slot"`
	Pubkey   string     `json:"pubkey"`
	Lamports int64      `json:"lamports"`
	Owner    string     `json:"owner"`
	Data     CBAcctData `json:"data"`
}

type CBAcctData struct {
	Version      int64   `json:"version"`
	Bump         []int64 `json:"bump"`
	Owner        string  `json:"owner"`
	AssetId      string  `json:"assetId"`
	Amount       string  `json:"amount"`
	Currency     *string `json:"currency"`
	Expiry       string  `json:"expiry"`
	PrivateTaker *string `json:"privateTaker"`
	MakerBroker  *string `json:"makerBroker"`
	Raw          string  `json:"raw"`
}
