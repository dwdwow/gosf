package gosf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dwdwow/golimiter"
	"github.com/gagliardetto/solana-go/rpc"
)

const (
	RPC_BASE_URL     = "https://rpc.shyft.to"
	DEFI_BASE_URL    = "https://defi.shyft.to"
	API_BASE_URL     = "https://api.shyft.to"
	GRAPHQL_BASE_URL = "https://programs.shyft.to/v0/graphql"
)

var (
	FreeApiLimiter   = golimiter.NewReqLimiter(time.Second, 1)
	FreeRpcLimiter   = golimiter.NewReqLimiter(time.Second, 30)
	HackApiLimiter   = golimiter.NewReqLimiter(time.Second, 10)
	HackRpcLimiter   = golimiter.NewReqLimiter(time.Second, 50)
	LaunchApiLimiter = golimiter.NewReqLimiter(time.Second, 30)
	LaunchRpcLimiter = golimiter.NewReqLimiter(time.Second, 150)
	ScaleApiLimiter  = golimiter.NewReqLimiter(time.Second, 150)
	ScaleRpcLimiter  = golimiter.NewReqLimiter(time.Second, 400)
)

type Network string

const (
	NetworkMainnetBeta Network = "mainnet-beta"
	NetworkDevnet      Network = "devnet"
)

type ParsedTxType string

const (
	ParsedTxTypeNFTMint                 ParsedTxType = "NFT_MINT"
	ParsedTxTypeTokenMint               ParsedTxType = "TOKEN_MINT"
	ParsedTxTypeTokenCreate             ParsedTxType = "TOKEN_CREATE"
	ParsedTxTypeSolTransfer             ParsedTxType = "SOL_TRANSFER"
	ParsedTxTypeTokenTransfer           ParsedTxType = "TOKEN_TRANSFER"
	ParsedTxTypeNFTTransfer             ParsedTxType = "NFT_TRANSFER"
	ParsedTxTypeNFTBurn                 ParsedTxType = "NFT_BURN"
	ParsedTxTypeTokenBurn               ParsedTxType = "TOKEN_BURN"
	ParsedTxTypeNFTSale                 ParsedTxType = "NFT_SALE"
	ParsedTxTypeNFTBid                  ParsedTxType = "NFT_BID"
	ParsedTxTypeNFTBidCancel            ParsedTxType = "NFT_BID_CANCEL"
	ParsedTxTypeNFTList                 ParsedTxType = "NFT_LIST"
	ParsedTxTypeNFTListUpdate           ParsedTxType = "NFT_LIST_UPDATE"
	ParsedTxTypeNFTListCancel           ParsedTxType = "NFT_LIST_CANCEL"
	ParsedTxTypeMarketplaceWithdraw     ParsedTxType = "MARKETPLACE_WITHDRAW"
	ParsedTxTypeCompressedNFTSale       ParsedTxType = "COMPRESSED_NFT_SALE"
	ParsedTxTypeCompressedNFTList       ParsedTxType = "COMPRESSED_NFT_LIST"
	ParsedTxTypeCompressedNFTListCancel ParsedTxType = "COMPRESSED_NFT_LIST_CANCEL"
	ParsedTxTypeCompressedNFTListUpdate ParsedTxType = "COMPRESSED_NFT_LIST_UPDATE"
	ParsedTxTypeCompressedNFTBid        ParsedTxType = "COMPRESSED_NFT_BID"
	ParsedTxTypeCompressedNFTBidCancel  ParsedTxType = "COMPRESSED_NFT_BID_CANCEL"
	ParsedTxTypeCompressedNFTTakeBid    ParsedTxType = "COMPRESSED_NFT_TAKE_BID"
	ParsedTxTypeOfferLoan               ParsedTxType = "OFFER_LOAN"
	ParsedTxTypeCancelLoan              ParsedTxType = "CANCEL_LOAN"
	ParsedTxTypeTakeLoan                ParsedTxType = "TAKE_LOAN"
	ParsedTxTypeRepayLoan               ParsedTxType = "REPAY_LOAN"
	ParsedTxTypeForecloseLoan           ParsedTxType = "FORECLOSE_LOAN"
	ParsedTxTypeRepayEscrowLoan         ParsedTxType = "REPAY_ESCROW_LOAN"
	ParsedTxTypeExtendLoan              ParsedTxType = "EXTEND_LOAN"
	ParsedTxTypeExtendEscrowLoan        ParsedTxType = "EXTEND_ESCROW_LOAN"
	ParsedTxTypeRequestLoan             ParsedTxType = "REQUEST_LOAN"
	ParsedTxTypeCancelRequestLoan       ParsedTxType = "CANCEL_REQUEST_LOAN"
	ParsedTxTypeLiquidateLoan           ParsedTxType = "LIQUIDATE_LOAN"
	ParsedTxTypeBuyNowPayLater          ParsedTxType = "BUY_NOW_PAY_LATER"
	ParsedTxTypeMemo                    ParsedTxType = "MEMO"
	ParsedTxTypeSwap                    ParsedTxType = "SWAP"
	ParsedTxTypeCreatePool              ParsedTxType = "CREATE_POOL"
	ParsedTxTypeAddLiquidity            ParsedTxType = "ADD_LIQUIDITY"
	ParsedTxTypeRemoveLiquidity         ParsedTxType = "REMOVE_LIQUIDITY"
	ParsedTxTypeCollectFees             ParsedTxType = "COLLECT_FEES"
	ParsedTxTypeCollectReward           ParsedTxType = "COLLECT_REWARD"
	ParsedTxTypeCreateRaffle            ParsedTxType = "CREATE_RAFFLE"
	ParsedTxTypeBuyTickets              ParsedTxType = "BUY_TICKETS"
	ParsedTxTypeRevealWinners           ParsedTxType = "REVEAL_WINNERS"
	ParsedTxTypeClaimPrize              ParsedTxType = "CLAIM_PRIZE"
	ParsedTxTypeCloseRaffle             ParsedTxType = "CLOSE_RAFFLE"
	ParsedTxTypeCancelRaffle            ParsedTxType = "CANCEL_RAFFLE"
	ParsedTxTypeCreateTree              ParsedTxType = "CREATE_TREE"
	ParsedTxTypeCompressedNFTMint       ParsedTxType = "COMPRESSED_NFT_MINT"
	ParsedTxTypeCompressedNFTTransfer   ParsedTxType = "COMPRESSED_NFT_TRANSFER"
	ParsedTxTypeCompressedNFTBurn       ParsedTxType = "COMPRESSED_NFT_BURN"
	ParsedTxTypeCreateRealm             ParsedTxType = "CREATE_REALM"
	ParsedTxTypeDepositGoverningTokens  ParsedTxType = "DEPOSIT_GOVERNING_TOKENS"
	ParsedTxTypeWithdrawGoverningTokens ParsedTxType = "WITHDRAW_GOVERNING_TOKENS"
	ParsedTxTypeSetGovernanceDelegate   ParsedTxType = "SET_GOVERNANCE_DELEGATE"
	ParsedTxTypeCreateGovernance        ParsedTxType = "CREATE_GOVERNANCE"
	ParsedTxTypeCreateProgramGovernance ParsedTxType = "CREATE_PROGRAM_GOVERNANCE"
	ParsedTxTypeCreateProposal          ParsedTxType = "CREATE_PROPOSAL"
	ParsedTxTypeAddSignatory            ParsedTxType = "ADD_SIGNATORY"
	ParsedTxTypeRemoveSignatory         ParsedTxType = "REMOVE_SIGNATORY"
	ParsedTxTypeInsertTransaction       ParsedTxType = "INSERT_TRANSACTION"
	ParsedTxTypeRemoveTransaction       ParsedTxType = "REMOVE_TRANSACTION"
	ParsedTxTypeCancelProposal          ParsedTxType = "CANCEL_PROPOSAL"
	ParsedTxTypeSignOffProposal         ParsedTxType = "SIGN_OFF_PROPOSAL"
	ParsedTxTypeCastVote                ParsedTxType = "CAST_VOTE"
	ParsedTxTypeFinalizeVote            ParsedTxType = "FINALIZE_VOTE"
	ParsedTxTypeRelinquishVote          ParsedTxType = "RELINQUISH_VOTE"
	ParsedTxTypeExecuteTransaction      ParsedTxType = "EXECUTE_TRANSACTION"
	ParsedTxTypeCreateMintGovernance    ParsedTxType = "CREATE_MINT_GOVERNANCE"
	ParsedTxTypeCreateTokenGovernance   ParsedTxType = "CREATE_TOKEN_GOVERNANCE"
	ParsedTxTypeSetGovernanceConfig     ParsedTxType = "SET_GOVERNANCE_CONFIG"
	ParsedTxTypeSetRealmAuthority       ParsedTxType = "SET_REALM_AUTHORITY"
	ParsedTxTypeSetRealmConfig          ParsedTxType = "SET_REALM_CONFIG"
	ParsedTxTypeCreateTokenOwnerRecord  ParsedTxType = "CREATE_TOKEN_OWNER_RECORD"
	ParsedTxTypePostMessage             ParsedTxType = "POST_MESSAGE"
)

type CallbackType string

const (
	CallbackTypeDiscord  CallbackType = "DISCORD"
	CallbackTypeCallback CallbackType = "CALLBACK"
	CallbackTypeAccount  CallbackType = "ACCOUNT"
)

type Encoding string

const (
	EncodingRaw    Encoding = "RAW"
	EncodingParsed Encoding = "PARSED"
)

type RespData[D any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  D      `json:"result"`
}

type CallbackInfo struct {
	ID          string   `json:"id"`
	ID_         string   `json:"_id"`
	Events      []string `json:"events"`
	CallbackURL string   `json:"callback_url"`
	Network     string   `json:"network"`
	Addresses   []string `json:"addresses"`
	APIKey      string   `json:"api_key"`
}

type Client struct {
	apiLimiter *golimiter.ReqLimiter
	rpcLimiter *golimiter.ReqLimiter

	apiKey string

	rpcClient *rpc.Client
}

func NewClient(apiKey string, apiLimiter *golimiter.ReqLimiter, rpcLimiter *golimiter.ReqLimiter) *Client {
	return &Client{
		apiLimiter: apiLimiter,
		rpcLimiter: rpcLimiter,
		apiKey:     apiKey,
		rpcClient:  rpc.New(RPC_BASE_URL + "?api_key=" + apiKey),
	}
}

func (c *Client) newApiHeader() http.Header {
	header := http.Header{}
	header.Set("X-API-KEY", c.apiKey)
	header.Set("Content-Type", "application/json")
	return header
}

func req[D any](url, method string, header http.Header, body any) (D, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return *new(D), err
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return *new(D), err
	}
	req.Header = header

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return *new(D), err
	}
	defer resp.Body.Close()

	var d RespData[D]
	err = json.NewDecoder(resp.Body).Decode(&d)
	if err != nil {
		return *new(D), err
	}

	if !d.Success {
		return *new(D), fmt.Errorf("shyft: http status code %d, message: %s", resp.StatusCode, d.Message)
	}

	return d.Result, nil
}

func (c *Client) CallbackList() ([]CallbackInfo, error) {
	return req[[]CallbackInfo](API_BASE_URL+"/sol/v1/callback/list", "GET", c.newApiHeader(), nil)
}

type CallbackCreateBody struct {
	Network      Network      `json:"network"`
	Addresses    []string     `json:"addresses"`
	CallbackUrl  string       `json:"callback_url"`
	Events       []string     `json:"events,omitempty"`
	EnableRaw    bool         `json:"enable_raw,omitempty"`
	EnableEvents bool         `json:"enable_events,omitempty"`
	Type         CallbackType `json:"type,omitempty"`
	Encoding     Encoding     `json:"encoding,omitempty"`
}

func (c *Client) CreateCallback(body CallbackCreateBody) (CallbackInfo, error) {
	return req[CallbackInfo](API_BASE_URL+"/sol/v1/callback/create", "POST", c.newApiHeader(), body)
}

func (c *Client) RemoveCallback(id string) error {
	_, err := req[map[string]string](API_BASE_URL+"/sol/v1/callback/remove", "DELETE", c.newApiHeader(), map[string]string{"id": id})
	return err
}

func (c *Client) PauseCallback(id string) error {
	_, err := req[map[string]string](API_BASE_URL+"/sol/v1/callback/pause", "POST", c.newApiHeader(), map[string]string{"id": id})
	return err
}

func (c *Client) ResumeCallback(id string) error {
	_, err := req[map[string]string](API_BASE_URL+"/sol/v1/callback/resume", "POST", c.newApiHeader(), map[string]string{"id": id})
	return err
}

type CallbackUpdateBody struct {
	ID           string   `json:"id"`
	Addresses    []string `json:"addresses,omitempty"`
	Events       []string `json:"events,omitempty"`
	CallbackUrl  string   `json:"callback_url,omitempty"`
	EnableRaw    bool     `json:"enable_raw,omitempty"`
	EnableEvents bool     `json:"enable_events,omitempty"`
	Encoding     Encoding `json:"encoding,omitempty"`
}

func (c *Client) UpdateCallback(body CallbackUpdateBody) (CallbackInfo, error) {
	return req[CallbackInfo](API_BASE_URL+"/sol/v1/callback/update", "POST", c.newApiHeader(), body)
}

func (c *Client) AddCallbackAddresses(id string, addresses []string) (CallbackInfo, error) {
	return req[CallbackInfo](API_BASE_URL+"/sol/v1/callback/add-addresses", "POST", c.newApiHeader(), map[string]any{
		"id":        id,
		"addresses": addresses,
	})
}

func (c *Client) RemoveCallbackAddresses(id string, addresses []string) (CallbackInfo, error) {
	return req[CallbackInfo](API_BASE_URL+"/sol/v1/callback/remove-addresses", "POST", c.newApiHeader(), map[string]any{
		"id":        id,
		"addresses": addresses,
	})
}
