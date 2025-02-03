package gosf

type TxStatus string

const (
	TxStatusSuccess TxStatus = "Success"
)

type TxProtocol struct {
	Address string `json:"address" bson:"address"`
	Name    string `json:"name" bson:"name"`
}

type TxAction struct {
	Info           any        `json:"info" bson:"info"`
	SourceProtocol TxProtocol `json:"source_protocol" bson:"sourceProtocol"`
	Type           TxType     `json:"type" bson:"type"`
	IxIndex        int64      `json:"ix_index" bson:"ixIndex"`
}

type TxTokenBalanceChange struct {
	Address      string  `json:"address" bson:"address"`
	Decimals     int64   `json:"decimals" bson:"decimals"`
	ChangeAmount float64 `json:"change_amount" bson:"changeAmount"`
	// PostBalance is int, but may be big, so use float64
	PostBalance float64 `json:"post_balance" bson:"postBalance"`
	// PreBalance is int, but may be big, so use float64
	PreBalance float64 `json:"pre_balance" bson:"preBalance"`
	Mint       string  `json:"mint" bson:"mint"`
	Owner      string  `json:"owner" bson:"owner"`
}

type TxAccountData struct {
	Address  string `json:"address" bson:"address"`
	Owner    string `json:"owner" bson:"owner"`
	Lamports int64  `json:"lamports" bson:"lamports"`
	Data     any    `json:"data" bson:"data"`
}

type TxRawMeta struct {
	ComputeUnitsConsumed int64    `json:"computeUnitsConsumed" bson:"computeUnitsConsumed"`
	Err                  any      `json:"err" bson:"err"`
	Fee                  int64    `json:"fee" bson:"fee"`
	InnerInstructions    []any    `json:"innerInstructions" bson:"innerInstructions"`
	LogMessages          []string `json:"logMessages" bson:"logMessages"`
	// PostBalances is int, but may be big, so use float64
	PostBalances      []float64 `json:"postBalances" bson:"postBalances"`
	PostTokenBalances []any     `json:"postTokenBalances" bson:"postTokenBalances"`
	// PreBalances is int, but may be big, so use float64
	PreBalances      []float64      `json:"preBalances" bson:"preBalances"`
	PreTokenBalances []any          `json:"preTokenBalances" bson:"preTokenBalances"`
	Rewards          []any          `json:"rewards" bson:"rewards"`
	Status           map[string]any `json:"status" bson:"status"`
}

type TxAccountKey struct {
	Pubkey   string `json:"pubkey" bson:"pubkey"`
	Signer   bool   `json:"signer" bson:"signer"`
	Source   string `json:"source" bson:"source"`
	Writable bool   `json:"writable" bson:"writable"`
}

type TxInstructionInfo struct {
	Destination string `json:"destination" bson:"destination"`
	Lamports    int64  `json:"lamports" bson:"lamports"`
	Source      string `json:"source" bson:"source"`
}

type TxInstruction struct {
	Parsed struct {
		Info TxInstructionInfo `json:"info" bson:"info"`
		Type string            `json:"type" bson:"type"`
	} `json:"parsed" bson:"parsed"`
	Program   string `json:"program" bson:"program"`
	ProgramId string `json:"programId" bson:"programId"`
}

type TxMessage struct {
	AccountKeys         []TxAccountKey  `json:"accountKeys" bson:"accountKeys"`
	AddressTableLookups any             `json:"addressTableLookups" bson:"addressTableLookups"`
	Instructions        []TxInstruction `json:"instructions" bson:"instructions"`
	RecentBlockhash     string          `json:"recentBlockhash" bson:"recentBlockhash"`
}

type TxTransaction struct {
	Message    TxMessage `json:"message" bson:"message"`
	Signatures []string  `json:"signatures" bson:"signatures"`
}

type TxParsed struct {
	Timestamp  string     `json:"timestamp" bson:"timestamp"`
	Fee        float64    `json:"fee" bson:"fee"`
	FeePayer   string     `json:"fee_payer" bson:"feePayer"`
	Signers    []string   `json:"signers" bson:"signers"`
	Signatures []string   `json:"signatures" bson:"signatures"`
	Protocol   TxProtocol `json:"protocol" bson:"protocol"`
	Type       string     `json:"type" bson:"type"`
	Status     string     `json:"status" bson:"status"`
	Actions    []TxAction `json:"actions" bson:"actions"`
}

type TxRawTransaction struct {
	BlockTime   int64         `json:"blockTime" bson:"blockTime"`
	Meta        TxRawMeta     `json:"meta" bson:"meta"`
	Slot        int64         `json:"slot" bson:"slot"`
	Transaction TxTransaction `json:"transaction" bson:"transaction"`
	Version     string        `json:"version" bson:"version"`
	Parsed      TxParsed      `json:"parsed" bson:"parsed"`
}

type Tx struct {
	TsMilli             int64                  `json:"ts_milli" bson:"tsMilli"`
	Timestamp           string                 `json:"timestamp" bson:"timestamp"`
	Fee                 float64                `json:"fee" bson:"fee"`
	FeePayer            string                 `json:"fee_payer" bson:"feePayer"`
	Signers             []string               `json:"signers" bson:"signers"`
	Signatures          []string               `json:"signatures" bson:"signatures"`
	Protocol            TxProtocol             `json:"protocol" bson:"protocol"`
	Type                TxType                 `json:"type" bson:"type"`
	Status              TxStatus               `json:"status" bson:"status"`
	Actions             []TxAction             `json:"actions" bson:"actions"`
	Raw                 *TxRawTransaction      `json:"raw,omitempty" bson:"raw,omitempty"`
	Accounts            []TxAccountData        `json:"accounts" bson:"accounts"`
	TokenBalanceChanges []TxTokenBalanceChange `json:"token_balance_changes" bson:"tokenBalanceChanges"`
	TriggeredFor        string                 `json:"triggered_for" bson:"triggeredFor"`
}
