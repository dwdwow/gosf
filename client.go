package gosf

import (
	"bytes"
	"context"
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
	Type        string   `json:"type"`
	Encoding    Encoding `json:"encoding"`
	Active      bool     `json:"active"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
	V           int64    `json:"__v"`
}

type Client struct {
	apiLimiter *golimiter.ReqLimiter
	rpcLimiter *golimiter.ReqLimiter

	apiKey string

	rpcClient *rpc.Client
}

func NewClient(apiKey string, apiLimiter *golimiter.ReqLimiter, rpcLimiter *golimiter.ReqLimiter) *Client {
	if apiKey == "" {
		panic("shyft: apiKey is required")
	}
	return &Client{
		apiLimiter: apiLimiter,
		rpcLimiter: rpcLimiter,
		apiKey:     apiKey,
		rpcClient:  rpc.New(RPC_BASE_URL + "?api_key=" + apiKey),
	}
}

func (c *Client) newGetApiHeader() http.Header {
	header := http.Header{}
	header.Set("x-api-key", c.apiKey)
	return header
}

func (c *Client) newPostApiHeader() http.Header {
	header := http.Header{}
	header.Set("x-api-key", c.apiKey)
	header.Set("Content-Type", "application/json")
	return header
}

func req[D any](limiter *golimiter.ReqLimiter, url, method string, header http.Header, body any) (D, error) {
	if limiter != nil {
		limiter.Wait(context.Background())
	}
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

	// bs, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return *new(D), err
	// }
	// fmt.Println(string(bs))

	// return *new(D), nil

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
	return req[[]CallbackInfo](c.apiLimiter, API_BASE_URL+"/sol/v1/callback/list", "GET", c.newGetApiHeader(), nil)
}

type CallbackCreateBody struct {
	Network      Network      `json:"network"`
	Addresses    []string     `json:"addresses"`
	CallbackUrl  string       `json:"callback_url"`
	Events       []TxType     `json:"events,omitempty"`
	EnableRaw    bool         `json:"enable_raw,omitempty"`
	EnableEvents bool         `json:"enable_events,omitempty"`
	Type         CallbackType `json:"type,omitempty"`
	Encoding     Encoding     `json:"encoding,omitempty"`
}

func (c *Client) CreateCallback(body CallbackCreateBody) (CallbackInfo, error) {
	return req[CallbackInfo](c.apiLimiter, API_BASE_URL+"/sol/v1/callback/create", "POST", c.newPostApiHeader(), body)
}

func (c *Client) RemoveCallback(id string) error {
	_, err := req[map[string]string](c.apiLimiter, API_BASE_URL+"/sol/v1/callback/remove", "DELETE", c.newPostApiHeader(), map[string]string{"id": id})
	return err
}

func (c *Client) PauseCallback(id string) error {
	_, err := req[map[string]string](c.apiLimiter, API_BASE_URL+"/sol/v1/callback/pause", "POST", c.newPostApiHeader(), map[string]string{"id": id})
	return err
}

func (c *Client) ResumeCallback(id string) error {
	_, err := req[map[string]string](c.apiLimiter, API_BASE_URL+"/sol/v1/callback/resume", "POST", c.newPostApiHeader(), map[string]string{"id": id})
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
	return req[CallbackInfo](c.apiLimiter, API_BASE_URL+"/sol/v1/callback/update", "POST", c.newPostApiHeader(), body)
}

func (c *Client) AddCallbackAddresses(id string, addresses []string) (CallbackInfo, error) {
	return req[CallbackInfo](c.apiLimiter, API_BASE_URL+"/sol/v1/callback/add-addresses", "POST", c.newPostApiHeader(), map[string]any{
		"id":        id,
		"addresses": addresses,
	})
}

func (c *Client) RemoveCallbackAddresses(id string, addresses []string) (CallbackInfo, error) {
	return req[CallbackInfo](c.apiLimiter, API_BASE_URL+"/sol/v1/callback/remove-addresses", "POST", c.newPostApiHeader(), map[string]any{
		"id":        id,
		"addresses": addresses,
	})
}
