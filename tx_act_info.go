package gosf

import (
	"encoding/json"
	"fmt"
)

// MarketplaceWithdrawInfo represents marketplace withdrawal information
type MarketplaceWithdrawInfo struct {
	Currency                     string  `json:"currency" bson:"currency"`
	Marketplace                  string  `json:"marketplace" bson:"marketplace"`
	TreasuryAccount              string  `json:"treasury_account" bson:"treasuryAccount"`
	MarketplaceAuthority         string  `json:"marketplace_authority" bson:"marketplaceAuthority"`
	WithdrawalDestinationAccount string  `json:"withdrawal_destination_account" bson:"withdrawalDestinationAccount"`
	Amount                       float64 `json:"amount" bson:"amount"`
}

// NFTBidInfo represents NFT bid information
type NFTBidInfo struct {
	Bidder      string  `json:"bidder" bson:"bidder"`
	Currency    string  `json:"currency" bson:"currency"`
	Marketplace string  `json:"marketplace" bson:"marketplace"`
	Price       float64 `json:"price" bson:"price"`
	NFTAddress  string  `json:"nft_address" bson:"nftAddress"`
}

// NFTBurnInfo represents NFT burn information
type NFTBurnInfo struct {
	Wallet     string  `json:"wallet" bson:"wallet"`
	NFTAddress string  `json:"nft_address" bson:"nftAddress"`
	Amount     float64 `json:"amount" bson:"amount"`
}

// NFTListCancelInfo represents NFT list cancellation information
type NFTListCancelInfo struct {
	Seller      string  `json:"seller" bson:"seller"`
	Marketplace string  `json:"marketplace" bson:"marketplace"`
	NFTAddress  string  `json:"nft_address" bson:"nftAddress"`
	Price       float64 `json:"price" bson:"price"`
}

// NFTListInfo represents NFT listing information
type NFTListInfo struct {
	Seller      string  `json:"seller" bson:"seller"`
	Marketplace string  `json:"marketplace" bson:"marketplace"`
	Price       float64 `json:"price" bson:"price"`
	NFTAddress  string  `json:"nft_address" bson:"nftAddress"`
}

// NFTListUpdateInfo represents NFT list update information
type NFTListUpdateInfo struct {
	Seller      string   `json:"seller" bson:"seller"`
	Marketplace string   `json:"marketplace" bson:"marketplace"`
	NFTAddress  string   `json:"nft_address" bson:"nftAddress"`
	OldPrice    *float64 `json:"old_price" bson:"oldPrice"`
	NewPrice    float64  `json:"new_price" bson:"newPrice"`
}

// NFTMintInfo represents NFT minting information
type NFTMintInfo struct {
	NFTAddress             string  `json:"nft_address" bson:"nftAddress"`
	Amount                 float64 `json:"amount" bson:"amount"`
	AmountRaw              float64 `json:"amount_raw" bson:"amountRaw"`
	OwnerAssociatedAccount string  `json:"owner_associated_account" bson:"ownerAssociatedAccount"`
	Owner                  string  `json:"owner" bson:"owner"`
	CollectionAddress      string  `json:"collection_address" bson:"collectionAddress"`
}

// NFTSaleInfo represents NFT sale information
type NFTSaleInfo struct {
	Buyer       string  `json:"buyer" bson:"buyer"`
	Seller      string  `json:"seller" bson:"seller"`
	NFTAddress  string  `json:"nft_address" bson:"nftAddress"`
	Currency    string  `json:"currency" bson:"currency"`
	Marketplace string  `json:"marketplace" bson:"marketplace"`
	Price       float64 `json:"price" bson:"price"`
}

// CompressedNFTSaleInfo represents compressed NFT sale information
type CompressedNFTSaleInfo struct {
	Buyer             string  `json:"buyer" bson:"buyer"`
	Seller            string  `json:"seller" bson:"seller"`
	Currency          string  `json:"currency" bson:"currency"`
	ListState         string  `json:"list_state" bson:"listState"`
	Marketplace       string  `json:"marketplace" bson:"marketplace"`
	MerkleTree        string  `json:"merkle_tree" bson:"merkleTree"`
	NFTAddress        string  `json:"nft_address" bson:"nftAddress"`
	Price             float64 `json:"price" bson:"price"`
	RoyaltyPercentage float64 `json:"royalty_percentage" bson:"royaltyPercentage"`
}

// CompressedNFTListInfo represents compressed NFT listing information
type CompressedNFTListInfo struct {
	Seller         string  `json:"seller" bson:"seller"`
	MerkleTree     string  `json:"merkle_tree" bson:"merkleTree"`
	NFTAddress     string  `json:"nft_address" bson:"nftAddress"`
	Currency       string  `json:"currency" bson:"currency"`
	Price          float64 `json:"price" bson:"price"`
	ListState      string  `json:"list_state" bson:"listState"`
	Marketplace    string  `json:"marketplace" bson:"marketplace"`
	ExpireInSecond float64 `json:"expire_in_second" bson:"expireInSecond"`
	PrivateTaker   string  `json:"private_taker" bson:"privateTaker"`
}

// CompressedNFTListCancelInfo represents compressed NFT list cancellation information
type CompressedNFTListCancelInfo struct {
	Seller      string   `json:"seller" bson:"seller"`
	MerkleTree  string   `json:"merkle_tree" bson:"merkleTree"`
	ListState   string   `json:"list_state" bson:"listState"`
	NFTAddress  string   `json:"nft_address" bson:"nftAddress"`
	Marketplace string   `json:"marketplace" bson:"marketplace"`
	Currency    string   `json:"currency" bson:"currency"`
	Price       *float64 `json:"price" bson:"price"`
}

// CompressedNFTListUpdateInfo represents compressed NFT list update information
type CompressedNFTListUpdateInfo struct {
	ListState      string  `json:"list_state" bson:"listState"`
	Marketplace    string  `json:"marketplace" bson:"marketplace"`
	Seller         string  `json:"seller" bson:"seller"`
	NFTAddress     string  `json:"nft_address" bson:"nftAddress"`
	Currency       string  `json:"currency" bson:"currency"`
	NewPrice       float64 `json:"new_price" bson:"newPrice"`
	ExpireInSecond float64 `json:"expire_in_second" bson:"expireInSecond"`
	PrivateTaker   string  `json:"private_taker" bson:"privateTaker"`
}

// CompressedNFTBidInfo represents compressed NFT bid information
type CompressedNFTBidInfo struct {
	BidState       string  `json:"bid_state" bson:"bidState"`
	Bidder         string  `json:"bidder" bson:"bidder"`
	Marketplace    string  `json:"marketplace" bson:"marketplace"`
	Target         string  `json:"target" bson:"target"`
	TargetID       string  `json:"target_id" bson:"targetId"`
	Price          float64 `json:"price" bson:"price"`
	Currency       string  `json:"currency" bson:"currency"`
	Quantity       float64 `json:"quantity" bson:"quantity"`
	ExpireInSecond float64 `json:"expire_in_second" bson:"expireInSecond"`
	PrivateTaker   string  `json:"private_taker" bson:"privateTaker"`
}

// CompressedNFTBidCancelInfo represents compressed NFT bid cancellation information
type CompressedNFTBidCancelInfo struct {
	BidState    string `json:"bid_state" bson:"bidState"`
	Bidder      string `json:"bidder" bson:"bidder"`
	Marketplace string `json:"marketplace" bson:"marketplace"`
}

// CompressedNFTTakeBidInfo represents compressed NFT take bid information
type CompressedNFTTakeBidInfo struct {
	Seller            string  `json:"seller" bson:"seller"`
	MerkleTree        string  `json:"merkle_tree" bson:"merkleTree"`
	Price             float64 `json:"price" bson:"price"`
	Currency          string  `json:"currency" bson:"currency"`
	Marketplace       string  `json:"marketplace" bson:"marketplace"`
	BidState          string  `json:"bid_state" bson:"bidState"`
	Bidder            string  `json:"bidder" bson:"bidder"`
	NFTAddress        string  `json:"nft_address" bson:"nftAddress"`
	RoyaltyPercentage float64 `json:"royalty_percentage" bson:"royaltyPercentage"`
}

// CompressedNFTMintInfo represents compressed NFT mint information
type CompressedNFTMintInfo struct {
	TreeAuthority   string          `json:"tree_authority" bson:"treeAuthority"`
	MerkleTree      string          `json:"merkle_tree" bson:"merkleTree"`
	Payer           string          `json:"payer" bson:"payer"`
	NFTAddress      string          `json:"nft_address" bson:"nftAddress"`
	Owner           string          `json:"owner" bson:"owner"`
	UpdateAuthority string          `json:"update_authority" bson:"updateAuthority"`
	NFTMetadata     NFTMetadataInfo `json:"nft_metadata" bson:"nftMetadata"`
}

// CompressedNFTTransferInfo represents compressed NFT transfer information
type CompressedNFTTransferInfo struct {
	TreeAuthority   string `json:"tree_authority" bson:"treeAuthority"`
	MerkleTree      string `json:"merkle_tree" bson:"merkleTree"`
	Sender          string `json:"sender" bson:"sender"`
	Receiver        string `json:"receiver" bson:"receiver"`
	NFTAddress      string `json:"nft_address" bson:"nftAddress"`
	UpdateAuthority string `json:"update_authority" bson:"updateAuthority"`
}

// CompressedNFTBurnInfo represents compressed NFT burn information
type CompressedNFTBurnInfo struct {
	TreeAuthority   string `json:"tree_authority" bson:"treeAuthority"`
	MerkleTree      string `json:"merkle_tree" bson:"merkleTree"`
	Owner           string `json:"owner" bson:"owner"`
	NFTAddress      string `json:"nft_address" bson:"nftAddress"`
	UpdateAuthority string `json:"update_authority" bson:"updateAuthority"`
}

// NFTMetadataInfo represents NFT metadata information
type NFTMetadataInfo struct {
	Name                 string              `json:"name" bson:"name"`
	Symbol               string              `json:"symbol" bson:"symbol"`
	URI                  string              `json:"uri" bson:"uri"`
	SellerFeeBasisPoints float64             `json:"sellerFeeBasisPoints" bson:"sellerFeeBasisPoints"`
	PrimarySaleHappened  bool                `json:"primarySaleHappened" bson:"primarySaleHappened"`
	IsMutable            bool                `json:"isMutable" bson:"isMutable"`
	EditionNonce         float64             `json:"editionNonce" bson:"editionNonce"`
	TokenStandard        TokenStandard       `json:"tokenStandard" bson:"tokenStandard"`
	Collection           CollectionInfo      `json:"collection" bson:"collection"`
	Uses                 interface{}         `json:"uses" bson:"uses"`
	TokenProgramVersion  TokenProgramVersion `json:"tokenProgramVersion" bson:"tokenProgramVersion"`
	Creators             []CreatorInfo       `json:"creators" bson:"creators"`
}

// TokenStandard represents the token standard type
type TokenStandard struct {
	NonFungible        *struct{} `json:"nonFungible,omitempty" bson:"nonFungible,omitempty"`
	FungibleAsset      *struct{} `json:"fungibleAsset,omitempty" bson:"fungibleAsset,omitempty"`
	Fungible           *struct{} `json:"fungible,omitempty" bson:"fungible,omitempty"`
	NonFungibleEdition *struct{} `json:"nonFungibleEdition,omitempty" bson:"nonFungibleEdition,omitempty"`
}

// CollectionInfo represents collection information
type CollectionInfo struct {
	Verified bool   `json:"verified" bson:"verified"`
	Key      string `json:"key" bson:"key"`
}

// TokenProgramVersion represents the token program version
type TokenProgramVersion struct {
	Original  *struct{} `json:"original,omitempty" bson:"original,omitempty"`
	Token2022 *struct{} `json:"token2022,omitempty" bson:"token2022,omitempty"`
}

// CreatorInfo represents creator information
type CreatorInfo struct {
	Address  string  `json:"address" bson:"address"`
	Verified bool    `json:"verified" bson:"verified"`
	Share    float64 `json:"share" bson:"share"`
}

// NFTTransferInfo represents NFT transfer information
type NFTTransferInfo struct {
	Amount                    float64 `json:"amount" bson:"amount"`
	Receiver                  string  `json:"receiver" bson:"receiver"`
	Sender                    string  `json:"sender" bson:"sender"`
	ReceiverAssociatedAccount string  `json:"receiver_associated_account" bson:"receiverAssociatedAccount"`
	NFTAddress                string  `json:"nft_address" bson:"nftAddress"`
}

// SOLTransferInfo represents SOL transfer information
type SOLTransferInfo struct {
	Sender   string  `json:"sender" bson:"sender"`
	Receiver string  `json:"receiver" bson:"receiver"`
	Amount   float64 `json:"amount" bson:"amount"`
}

// TokenBurnInfo represents token burn information
type TokenBurnInfo struct {
	Wallet       string  `json:"wallet" bson:"wallet"`
	Amount       float64 `json:"amount" bson:"amount"`
	TokenAddress string  `json:"token_address" bson:"tokenAddress"`
}

// TokenCreateInfo represents token creation information
type TokenCreateInfo struct {
	TokenAddress string `json:"token_address" bson:"tokenAddress"`
}

// TokenMintInfo represents token minting information
type TokenMintInfo struct {
	Amount          float64 `json:"amount" bson:"amount"`
	ReceiverAddress string  `json:"receiver_address" bson:"receiverAddress"`
	TokenAddress    string  `json:"token_address" bson:"tokenAddress"`
}

// TokenTransferInfo represents token transfer information
type TokenTransferInfo struct {
	Amount                    interface{} `json:"amount" bson:"amount"` // can be float64 or string
	Receiver                  string      `json:"receiver" bson:"receiver"`
	Sender                    string      `json:"sender" bson:"sender"`
	ReceiverAssociatedAccount string      `json:"receiver_associated_account" bson:"receiverAssociatedAccount"`
	TokenAddress              string      `json:"token_address" bson:"tokenAddress"`
}

// OfferLoanInfo represents loan offer information
type OfferLoanInfo struct {
	Lender              string  `json:"lender" bson:"lender"`
	Currency            string  `json:"currency" bson:"currency"`
	Amount              float64 `json:"amount" bson:"amount"`
	Loan                string  `json:"loan" bson:"loan"`
	Escrow              string  `json:"escrow" bson:"escrow"`
	EscrowTokenAccount  string  `json:"escrow_token_account" bson:"escrowTokenAccount"`
	OrderBook           string  `json:"order_book" bson:"orderBook"`
	Borrower            string  `json:"borrower" bson:"borrower"`
	NFTAddress          string  `json:"nft_address" bson:"nftAddress"`
	APY                 float64 `json:"apy" bson:"apy"`
	LoanDurationSeconds float64 `json:"loan_duration_seconds" bson:"loanDurationSeconds"`
}

// CancelLoanInfo represents loan cancellation information
type CancelLoanInfo struct {
	Lender               string  `json:"lender" bson:"lender"`
	Borrower             string  `json:"borrower" bson:"borrower"`
	Currency             string  `json:"currency" bson:"currency"`
	Amount               float64 `json:"amount" bson:"amount"`
	Loan                 string  `json:"loan" bson:"loan"`
	Escrow               string  `json:"escrow" bson:"escrow"`
	EscrowTokenAccount   string  `json:"escrow_token_account" bson:"escrowTokenAccount"`
	ReimbursedToBorrower float64 `json:"reimbursed_to_borrower" bson:"reimbursedToBorrower"`
}

// RepayLoanInfo represents loan repayment information
type RepayLoanInfo struct {
	Lender             string  `json:"lender" bson:"lender"`
	Borrower           string  `json:"borrower" bson:"borrower"`
	Loan               string  `json:"loan" bson:"loan"`
	Currency           string  `json:"currency" bson:"currency"`
	Amount             float64 `json:"amount" bson:"amount"`
	FeePaid            float64 `json:"fee_paid" bson:"feePaid"`
	NFTAddress         string  `json:"nft_address" bson:"nftAddress"`
	Escrow             string  `json:"escrow" bson:"escrow"`
	EscrowTokenAccount string  `json:"escrow_token_account" bson:"escrowTokenAccount"`
	OrderBook          string  `json:"order_book" bson:"orderBook"`
	AdminPayment       float64 `json:"admin_payment" bson:"adminPayment"`
	PaybackToLiqOwner  float64 `json:"payback_to_liq_owner" bson:"paybackToLiqOwner"`
}

// RepayEscrowLoanInfo represents escrow loan repayment information
type RepayEscrowLoanInfo struct {
	Amount             float64 `json:"amount" bson:"amount"`
	FeePaid            float64 `json:"fee_paid" bson:"feePaid"`
	Borrower           string  `json:"borrower" bson:"borrower"`
	CollateralMint     string  `json:"collateral_mint" bson:"collateralMint"`
	Currency           string  `json:"currency" bson:"currency"`
	Escrow             string  `json:"escrow" bson:"escrow"`
	EscrowTokenAccount string  `json:"escrow_token_account" bson:"escrowTokenAccount"`
	Lender             string  `json:"lender" bson:"lender"`
	Loan               string  `json:"loan" bson:"loan"`
}

// TakeLoanInfo represents loan taking information
type TakeLoanInfo struct {
	Lender              string  `json:"lender" bson:"lender"`
	Borrower            string  `json:"borrower" bson:"borrower"`
	Currency            string  `json:"currency" bson:"currency"`
	Amount              float64 `json:"amount" bson:"amount"`
	NFTAddress          string  `json:"nft_address" bson:"nftAddress"`
	Loan                string  `json:"loan" bson:"loan"`
	Escrow              string  `json:"escrow" bson:"escrow"`
	EscrowTokenAccount  string  `json:"escrow_token_account" bson:"escrowTokenAccount"`
	OrderBook           string  `json:"order_book" bson:"orderBook"`
	APY                 float64 `json:"apy" bson:"apy"`
	LoanDurationSeconds float64 `json:"loan_duration_seconds" bson:"loanDurationSeconds"`
	Discount            float64 `json:"discount" bson:"discount"`
	TransferToBorrower  float64 `json:"transfer_to_borrower" bson:"transferToBorrower"`
}

// ExtendLoanInfo represents loan extension information
type ExtendLoanInfo struct {
	OldLender             string  `json:"old_lender" bson:"oldLender"`
	NewLender             string  `json:"new_lender" bson:"newLender"`
	OldLoan               string  `json:"old_loan" bson:"oldLoan"`
	NewLoan               string  `json:"new_loan" bson:"newLoan"`
	Borrower              string  `json:"borrower" bson:"borrower"`
	Currency              string  `json:"currency" bson:"currency"`
	Amount                float64 `json:"amount" bson:"amount"`
	NFTAddress            string  `json:"nft_address" bson:"nftAddress"`
	OldEscrow             string  `json:"old_escrow" bson:"oldEscrow"`
	NewEscrow             string  `json:"new_escrow" bson:"newEscrow"`
	NewEscrowTokenAccount string  `json:"new_escrow_token_account" bson:"newEscrowTokenAccount"`
	OrderBook             string  `json:"order_book" bson:"orderBook"`
	APY                   float64 `json:"apy" bson:"apy"`
	LoanDurationSeconds   float64 `json:"loan_duration_seconds" bson:"loanDurationSeconds"`
}

// ExtendEscrowLoanInfo represents escrow loan extension information
type ExtendEscrowLoanInfo struct {
	OldLender             string  `json:"old_lender" bson:"oldLender"`
	NewLender             string  `json:"new_lender" bson:"newLender"`
	OldLoan               string  `json:"old_loan" bson:"oldLoan"`
	NewLoan               string  `json:"new_loan" bson:"newLoan"`
	Borrower              string  `json:"borrower" bson:"borrower"`
	Currency              string  `json:"currency" bson:"currency"`
	Amount                float64 `json:"amount" bson:"amount"`
	CollateralMint        string  `json:"collateral_mint" bson:"collateralMint"`
	OldEscrow             string  `json:"old_escrow" bson:"oldEscrow"`
	NewEscrow             string  `json:"new_escrow" bson:"newEscrow"`
	OldEscrowTokenAccount string  `json:"old_escrow_token_account" bson:"oldEscrowTokenAccount"`
	NewEscrowTokenAccount string  `json:"new_escrow_token_account" bson:"newEscrowTokenAccount"`
	OrderBook             string  `json:"order_book" bson:"orderBook"`
	LoanDurationSeconds   float64 `json:"loan_duration_seconds" bson:"loanDurationSeconds"`
}

// RequestLoanInfo represents loan request information
type RequestLoanInfo struct {
	Loan                string  `json:"loan" bson:"loan"`
	Borrower            string  `json:"borrower" bson:"borrower"`
	Currency            string  `json:"currency" bson:"currency"`
	Amount              float64 `json:"amount" bson:"amount"`
	NFTAddress          string  `json:"nft_address" bson:"nftAddress"`
	APY                 float64 `json:"apy" bson:"apy"`
	LoanDurationSeconds float64 `json:"loan_duration_seconds" bson:"loanDurationSeconds"`
	LTV                 float64 `json:"ltv" bson:"ltv"`
	AdminPayment        float64 `json:"admin_payment" bson:"adminPayment"`
}

// CancelRequestLoanInfo represents loan request cancellation information
type CancelRequestLoanInfo struct {
	Loan       string `json:"loan" bson:"loan"`
	Borrower   string `json:"borrower" bson:"borrower"`
	NFTAddress string `json:"nft_address" bson:"nftAddress"`
}

// LiquidateLoanInfo represents loan liquidation information
type LiquidateLoanInfo struct {
	Lender             string  `json:"lender" bson:"lender"`
	Borrower           string  `json:"borrower" bson:"borrower"`
	Loan               string  `json:"loan" bson:"loan"`
	Currency           string  `json:"currency" bson:"currency"`
	Amount             float64 `json:"amount" bson:"amount"`
	NFTAddress         string  `json:"nft_address" bson:"nftAddress"`
	GracePeriodSeconds float64 `json:"grace_period_seconds" bson:"gracePeriodSeconds"`
}

// BuyNowPayLaterInfo represents buy now pay later information
type BuyNowPayLaterInfo struct {
	Loan                string  `json:"loan" bson:"loan"`
	Lender              string  `json:"lender" bson:"lender"`
	Borrower            string  `json:"borrower" bson:"borrower"`
	NFTAddress          string  `json:"nft_address" bson:"nftAddress"`
	Currency            string  `json:"currency" bson:"currency"`
	Amount              float64 `json:"amount" bson:"amount"`
	APY                 float64 `json:"apy" bson:"apy"`
	LoanDurationSeconds float64 `json:"loan_duration_seconds" bson:"loanDurationSeconds"`
}

// TokenSwapInfo represents token swap information for tokens_swapped field
type TokenSwapInfo struct {
	TokenAddress string  `json:"token_address" bson:"tokenAddress"`
	Name         string  `json:"name" bson:"name"`
	Symbol       string  `json:"symbol" bson:"symbol"`
	ImageURI     string  `json:"image_uri" bson:"imageUri"`
	Amount       float64 `json:"amount" bson:"amount"`
}

// SwapLiquidityPoolInfo represents liquidity pool information for swaps field
type SwapLiquidityPoolInfo struct {
	LiquidityPoolAddress string `json:"liquidity_pool_address" bson:"liquidityPoolAddress"`
	Name                 string `json:"name" bson:"name"`
	Source               string `json:"source" bson:"source"`
}

// TokensSwappedInfo represents tokens swapped information
type TokensSwappedInfo struct {
	In  TokenSwapInfo `json:"in" bson:"in"`
	Out TokenSwapInfo `json:"out" bson:"out"`
}

// SwapInfo represents swap information
type SwapInfo struct {
	Swapper       string                  `json:"swapper" bson:"swapper"`
	TokensSwapped TokensSwappedInfo       `json:"tokens_swapped" bson:"tokensSwapped"`
	Swaps         []SwapLiquidityPoolInfo `json:"swaps" bson:"swaps"`
}

// CreatePoolInfo represents pool creation information
type CreatePoolInfo struct {
	PoolCreator          string `json:"pool_creator" bson:"poolCreator"`
	LiquidityPoolAddress string `json:"liquidity_pool_address" bson:"liquidityPoolAddress"`
	TokenMintOne         string `json:"token_mint_one" bson:"tokenMintOne"`
	TokenMintTwo         string `json:"token_mint_two" bson:"tokenMintTwo"`
	TokenVaultOne        string `json:"token_vault_one" bson:"tokenVaultOne"`
	TokenVaultTwo        string `json:"token_vault_two" bson:"tokenVaultTwo"`
}

// LiquidityTokenInfo represents liquidity token information
type LiquidityTokenInfo struct {
	TokenAddress string  `json:"token_address" bson:"tokenAddress"`
	Amount       float64 `json:"amount" bson:"amount"`
	Name         string  `json:"name" bson:"name"`
	Symbol       string  `json:"symbol" bson:"symbol"`
	ImageURI     string  `json:"image_uri" bson:"imageUri"`
}

// AddLiquidityInfo represents add liquidity information
type AddLiquidityInfo struct {
	LiquidityProviderAddress string               `json:"liquidity_provider_address" bson:"liquidityProviderAddress"`
	LiquidityPoolAddress     string               `json:"liquidity_pool_address" bson:"liquidityPoolAddress"`
	NFTAddress               string               `json:"nft_address" bson:"nftAddress"`
	LiquidityAdded           []LiquidityTokenInfo `json:"liquidity_added" bson:"liquidityAdded"`
}

// RemoveLiquidityInfo represents remove liquidity information
type RemoveLiquidityInfo struct {
	LiquidityProviderAddress string               `json:"liquidity_provider_address" bson:"liquidityProviderAddress"`
	LiquidityPoolAddress     string               `json:"liquidity_pool_address" bson:"liquidityPoolAddress"`
	NFTAddress               string               `json:"nft_address" bson:"nftAddress"`
	LiquidityRemoved         []LiquidityTokenInfo `json:"liquidity_removed" bson:"liquidityRemoved"`
}

// CollectFeesInfo represents collect fees information
type CollectFeesInfo struct {
	LiquidityPoolAddress     string               `json:"liquidity_pool_address" bson:"liquidityPoolAddress"`
	LiquidityProviderAddress string               `json:"liquidity_provider_address" bson:"liquidityProviderAddress"`
	FeesTaken                []LiquidityTokenInfo `json:"fees_taken" bson:"feesTaken"`
}

// CollectRewardInfo represents collect reward information
type CollectRewardInfo struct {
	LiquidityPoolAddress     string             `json:"liquidity_pool_address" bson:"liquidityPoolAddress"`
	LiquidityProviderAddress string             `json:"liquidity_provider_address" bson:"liquidityProviderAddress"`
	Reward                   LiquidityTokenInfo `json:"reward" bson:"reward"`
}

// CreateRaffleInfo represents raffle creation information
type CreateRaffleInfo struct {
	RaffleAddress string  `json:"raffle_address" bson:"raffleAddress"`
	RaffleCreator string  `json:"raffle_creator" bson:"raffleCreator"`
	RaffleToken   string  `json:"raffle_token" bson:"raffleToken"`
	Currency      string  `json:"currency" bson:"currency"`
	TicketPrice   float64 `json:"ticket_price" bson:"ticketPrice"`
	Tickets       float64 `json:"tickets" bson:"tickets"`
	StartDate     string  `json:"start_date" bson:"startDate"`
	EndDate       string  `json:"end_date" bson:"endDate"`
}

// UpdateRaffleInfo represents raffle update information
type UpdateRaffleInfo struct {
	RaffleAddress string  `json:"raffle_address" bson:"raffleAddress"`
	RaffleCreator string  `json:"raffle_creator" bson:"raffleCreator"`
	Currency      string  `json:"currency" bson:"currency"`
	TicketPrice   float64 `json:"ticket_price" bson:"ticketPrice"`
	Tickets       float64 `json:"tickets" bson:"tickets"`
}

// BuyTicketsInfo represents ticket purchase information
type BuyTicketsInfo struct {
	RaffleAddress string  `json:"raffle_address" bson:"raffleAddress"`
	Currency      string  `json:"currency" bson:"currency"`
	TicketPrice   float64 `json:"ticket_price" bson:"ticketPrice"`
	Tickets       float64 `json:"tickets" bson:"tickets"`
	Buyer         string  `json:"buyer" bson:"buyer"`
}

// RevealWinnersInfo represents winner reveal information
type RevealWinnersInfo struct {
	RaffleAddress string `json:"raffle_address" bson:"raffleAddress"`
	RaffleWinner  string `json:"raffle_winner" bson:"raffleWinner"`
}

// ClaimPrizeInfo represents prize claim information
type ClaimPrizeInfo struct {
	RaffleAddress string `json:"raffle_address" bson:"raffleAddress"`
	RaffleWinner  string `json:"raffle_winner" bson:"raffleWinner"`
	RaffleToken   string `json:"raffle_token" bson:"raffleToken"`
}

// CloseRaffleInfo represents raffle closure information
type CloseRaffleInfo struct {
	RaffleAddress       string  `json:"raffle_address" bson:"raffleAddress"`
	RaffleCreator       string  `json:"raffle_creator" bson:"raffleCreator"`
	Currency            string  `json:"currency" bson:"currency"`
	RaffleClosureAmount float64 `json:"raffle_closure_amount" bson:"raffleClosureAmount"`
	FeeTaker            string  `json:"fee_taker" bson:"feeTaker"`
	FeeTaken            float64 `json:"fee_taken" bson:"feeTaken"`
}

// CancelRaffleInfo represents raffle cancellation information
type CancelRaffleInfo struct {
	RaffleAddress string `json:"raffle_address" bson:"raffleAddress"`
	RaffleCreator string `json:"raffle_creator" bson:"raffleCreator"`
	RaffleToken   string `json:"raffle_token" bson:"raffleToken"`
}

// CreateTreeInfo represents tree creation information
type CreateTreeInfo struct {
	TreeAuthority string  `json:"tree_authority" bson:"treeAuthority"`
	MerkleTree    string  `json:"merkle_tree" bson:"merkleTree"`
	Payer         string  `json:"payer" bson:"payer"`
	TreeCreator   string  `json:"tree_creator" bson:"treeCreator"`
	MaxDepth      float64 `json:"max_depth" bson:"maxDepth"`
	MaxBufferSize float64 `json:"max_buffer_size" bson:"maxBufferSize"`
}

// CommunityTokenConfig represents community token configuration
type CommunityTokenConfig struct {
	UseVoterWeightAddin    bool   `json:"use_voter_weight_addin" bson:"useVoterWeightAddin"`
	UseMaxVoterWeightAddin bool   `json:"use_max_voter_weight_addin" bson:"useMaxVoterWeightAddin"`
	TokenType              string `json:"token_type" bson:"tokenType"`
}

// CreateRealmInfo represents realm creation information
type CreateRealmInfo struct {
	RealmAddress                         string               `json:"realm_address" bson:"realmAddress"`
	RealmAuthority                       string               `json:"realm_authority" bson:"realmAuthority"`
	CommunityTokenMint                   string               `json:"community_token_mint" bson:"communityTokenMint"`
	Payer                                string               `json:"payer" bson:"payer"`
	RealmName                            string               `json:"realm_name" bson:"realmName"`
	UseCouncilMint                       bool                 `json:"use_council_mint" bson:"useCouncilMint"`
	MinCommunityTokensToCreateGovernance float64              `json:"min_community_tokens_to_create_governance" bson:"minCommunityTokensToCreateGovernance"`
	CommunityMintMaxVoteWeightSource     float64              `json:"community_mint_max_vote_weight_source" bson:"communityMintMaxVoteWeightSource"`
	CommunityTokenConfig                 CommunityTokenConfig `json:"community_token_config" bson:"communityTokenConfig"`
	CouncilTokenConfig                   CommunityTokenConfig `json:"council_token_config" bson:"councilTokenConfig"`
}

// CreateProgramGovernanceInfo represents program governance creation information
type CreateProgramGovernanceInfo struct {
	RealmAddress                       string  `json:"realm_address" bson:"realmAddress"`
	ProgramGovernanceAddress           string  `json:"program_governance_address" bson:"programGovernanceAddress"`
	GovernedProgram                    string  `json:"governed_program" bson:"governedProgram"`
	TokenOwnerRecord                   string  `json:"token_owner_record" bson:"tokenOwnerRecord"`
	Payer                              string  `json:"payer" bson:"payer"`
	CreateAuthority                    string  `json:"create_authority" bson:"createAuthority"`
	MinCommunityTokensToCreateProposal float64 `json:"min_community_tokens_to_create_proposal" bson:"minCommunityTokensToCreateProposal"`
	MinInstructionHoldUpTime           float64 `json:"min_instruction_hold_up_time" bson:"minInstructionHoldUpTime"`
	BaseVotingTime                     float64 `json:"base_voting_time" bson:"baseVotingTime"`
	CommunityVoteTipping               string  `json:"community_vote_tipping" bson:"communityVoteTipping"`
	MinCouncilTokensToCreateProposal   float64 `json:"min_council_tokens_to_create_proposal" bson:"minCouncilTokensToCreateProposal"`
	CouncilVoteTipping                 string  `json:"council_vote_tipping" bson:"councilVoteTipping"`
	VotingCoolOffTime                  float64 `json:"voting_cool_off_time" bson:"votingCoolOffTime"`
	DepositExemptProposalCount         float64 `json:"deposit_exempt_proposal_count" bson:"depositExemptProposalCount"`
	IsTransferUpgradeAuthority         bool    `json:"is_transfer_upgrade_authority" bson:"isTransferUpgradeAuthority"`
}

// CreateProposalInfo represents proposal creation information
type CreateProposalInfo struct {
	RealmAddress        string `json:"realm_address" bson:"realmAddress"`
	ProposalAddress     string `json:"proposal_address" bson:"proposalAddress"`
	Governance          string `json:"governance" bson:"governance"`
	ProposalOwnerRecord string `json:"proposal_owner_record" bson:"proposalOwnerRecord"`
	GoverningToken      string `json:"governing_token" bson:"governingToken"`
	GovernanceAuthority string `json:"governance_authority" bson:"governanceAuthority"`
	Payer               string `json:"payer" bson:"payer"`
	ProposalName        string `json:"proposal_name" bson:"proposalName"`
	ProposalDescription string `json:"proposal_description" bson:"proposalDescription"`
	VoteType            string `json:"vote_type" bson:"voteType"`
	Option              string `json:"option" bson:"option"`
	UseDenyOption       bool   `json:"use_deny_option" bson:"useDenyOption"`
	ProposalSeed        string `json:"proposal_seed" bson:"proposalSeed"`
}

// AddSignatoryInfo represents signatory addition information
type AddSignatoryInfo struct {
	Proposal               string `json:"proposal" bson:"proposal"`
	TokenOwnerRecord       string `json:"token_owner_record" bson:"tokenOwnerRecord"`
	GovernanceAuthority    string `json:"governance_authority" bson:"governanceAuthority"`
	SignatoryRecordAddress string `json:"signatory_record_address" bson:"signatoryRecordAddress"`
	Payer                  string `json:"payer" bson:"payer"`
	Signatory              string `json:"signatory" bson:"signatory"`
}

// RemoveSignatoryInfo represents signatory removal information
type RemoveSignatoryInfo struct {
	Proposal               string `json:"proposal" bson:"proposal"`
	TokenOwnerRecord       string `json:"token_owner_record" bson:"tokenOwnerRecord"`
	GovernanceAuthority    string `json:"governance_authority" bson:"governanceAuthority"`
	SignatoryRecordAddress string `json:"signatory_record_address" bson:"signatoryRecordAddress"`
	Beneficiary            string `json:"beneficiary" bson:"beneficiary"`
}

// InsertTransactionInfo represents transaction insertion information
type InsertTransactionInfo struct {
	Governance          string `json:"governance" bson:"governance"`
	Proposal            string `json:"proposal" bson:"proposal"`
	TokenOwnerRecord    string `json:"token_owner_record" bson:"tokenOwnerRecord"`
	GovernanceAuthority string `json:"governance_authority" bson:"governanceAuthority"`
	ProposalTransaction string `json:"proposal_transaction" bson:"proposalTransaction"`
	Payer               string `json:"payer" bson:"payer"`
}

// RemoveTransactionInfo represents transaction removal information
type RemoveTransactionInfo struct {
	Proposal            string `json:"proposal" bson:"proposal"`
	TokenOwnerRecord    string `json:"token_owner_record" bson:"tokenOwnerRecord"`
	GovernanceAuthority string `json:"governance_authority" bson:"governanceAuthority"`
	ProposalTransaction string `json:"proposal_transaction" bson:"proposalTransaction"`
	Beneficiary         string `json:"beneficiary" bson:"beneficiary"`
}

// CancelProposalInfo represents proposal cancellation information
type CancelProposalInfo struct {
	RealmAddress        string `json:"realm_address" bson:"realmAddress"`
	Governance          string `json:"governance" bson:"governance"`
	Proposal            string `json:"proposal" bson:"proposal"`
	ProposalOwnerRecord string `json:"proposal_owner_record" bson:"proposalOwnerRecord"`
	GovernanceAuthority string `json:"governance_authority" bson:"governanceAuthority"`
}

// SignOffProposalInfo represents proposal sign-off information
type SignOffProposalInfo struct {
	RealmAddress string `json:"realm_address" bson:"realmAddress"`
	Governance   string `json:"governance" bson:"governance"`
	Proposal     string `json:"proposal" bson:"proposal"`
	Signatory    string `json:"signatory" bson:"signatory"`
}

// CastVoteInfo represents vote casting information
type CastVoteInfo struct {
	RealmAddress          string   `json:"realm_address" bson:"realmAddress"`
	Governance            string   `json:"governance" bson:"governance"`
	Proposal              string   `json:"proposal" bson:"proposal"`
	ProposalOwnerRecord   string   `json:"proposal_owner_record" bson:"proposalOwnerRecord"`
	VoterTokenOwnerRecord string   `json:"voter_token_owner_record" bson:"voterTokenOwnerRecord"`
	GovernanceAuthority   string   `json:"governance_authority" bson:"governanceAuthority"`
	VoteRecordAddress     string   `json:"vote_record_address" bson:"voteRecordAddress"`
	VoteGoverningToken    string   `json:"vote_governing_token" bson:"voteGoverningToken"`
	Payer                 string   `json:"payer" bson:"payer"`
	VoteType              string   `json:"vote_type" bson:"voteType"`
	Rank                  *float64 `json:"rank,omitempty" bson:"rank,omitempty"`
	WeightPercentage      *float64 `json:"weight_percentage,omitempty" bson:"weightPercentage,omitempty"`
}

// FinalizeVoteInfo represents vote finalization information
type FinalizeVoteInfo struct {
	RealmAddress        string `json:"realm_address" bson:"realmAddress"`
	Governance          string `json:"governance" bson:"governance"`
	Proposal            string `json:"proposal" bson:"proposal"`
	ProposalOwnerRecord string `json:"proposal_owner_record" bson:"proposalOwnerRecord"`
	GoverningToken      string `json:"governing_token" bson:"governingToken"`
}

// RelinquishVoteInfo represents vote relinquishment information
type RelinquishVoteInfo struct {
	RealmAddress       string `json:"realm_address" bson:"realmAddress"`
	Governance         string `json:"governance" bson:"governance"`
	Proposal           string `json:"proposal" bson:"proposal"`
	TokenOwnerRecord   string `json:"token_owner_record" bson:"tokenOwnerRecord"`
	VoteRecordAddress  string `json:"vote_record_address" bson:"voteRecordAddress"`
	VoteGoverningToken string `json:"vote_governing_token" bson:"voteGoverningToken"`
}

// ExecuteTransactionInfo represents transaction execution information
type ExecuteTransactionInfo struct {
	Governance string `json:"governance" bson:"governance"`
	Proposal   string `json:"proposal" bson:"proposal"`
}

// CreateMintGovernanceInfo represents mint governance creation information
type CreateMintGovernanceInfo struct {
	RealmAddress                       string  `json:"realm_address" bson:"realmAddress"`
	MintGovernanceAddress              string  `json:"mint_governance_address" bson:"mintGovernanceAddress"`
	GovernedMint                       string  `json:"governed_mint" bson:"governedMint"`
	TokenOwnerRecord                   string  `json:"token_owner_record" bson:"tokenOwnerRecord"`
	Payer                              string  `json:"payer" bson:"payer"`
	CreateAuthority                    string  `json:"create_authority" bson:"createAuthority"`
	MinCommunityTokensToCreateProposal float64 `json:"min_community_tokens_to_create_proposal" bson:"minCommunityTokensToCreateProposal"`
	MinInstructionHoldUpTime           float64 `json:"min_instruction_hold_up_time" bson:"minInstructionHoldUpTime"`
	BaseVotingTime                     float64 `json:"base_voting_time" bson:"baseVotingTime"`
	CommunityVoteTipping               string  `json:"community_vote_tipping" bson:"communityVoteTipping"`
	MinCouncilTokensToCreateProposal   float64 `json:"min_council_tokens_to_create_proposal" bson:"minCouncilTokensToCreateProposal"`
	CouncilVoteTipping                 string  `json:"council_vote_tipping" bson:"councilVoteTipping"`
	VotingCoolOffTime                  float64 `json:"voting_cool_off_time" bson:"votingCoolOffTime"`
	DepositExemptProposalCount         float64 `json:"deposit_exempt_proposal_count" bson:"depositExemptProposalCount"`
	IsTransferMintAuthorities          bool    `json:"is_transfer_mint_authorities" bson:"isTransferMintAuthorities"`
}

// CreateTokenGovernanceInfo represents token governance creation information
type CreateTokenGovernanceInfo struct {
	RealmAddress                       string  `json:"realm_address" bson:"realmAddress"`
	TokenGovernanceAddress             string  `json:"token_governance_address" bson:"tokenGovernanceAddress"`
	GovernedToken                      string  `json:"governed_token" bson:"governedToken"`
	GovernedTokenOwner                 string  `json:"governed_token_owner" bson:"governedTokenOwner"`
	TokenOwnerRecord                   string  `json:"token_owner_record" bson:"tokenOwnerRecord"`
	Payer                              string  `json:"payer" bson:"payer"`
	CreateAuthority                    string  `json:"create_authority" bson:"createAuthority"`
	MinCommunityTokensToCreateProposal float64 `json:"min_community_tokens_to_create_proposal" bson:"minCommunityTokensToCreateProposal"`
	MinInstructionHoldUpTime           float64 `json:"min_instruction_hold_up_time" bson:"minInstructionHoldUpTime"`
	BaseVotingTime                     float64 `json:"base_voting_time" bson:"baseVotingTime"`
	CommunityVoteTipping               string  `json:"community_vote_tipping" bson:"communityVoteTipping"`
	MinCouncilTokensToCreateProposal   float64 `json:"min_council_tokens_to_create_proposal" bson:"minCouncilTokensToCreateProposal"`
	CouncilVoteTipping                 string  `json:"council_vote_tipping" bson:"councilVoteTipping"`
	VotingCoolOffTime                  float64 `json:"voting_cool_off_time" bson:"votingCoolOffTime"`
	DepositExemptProposalCount         float64 `json:"deposit_exempt_proposal_count" bson:"depositExemptProposalCount"`
	IsTransferTokenOwner               bool    `json:"is_transfer_token_owner" bson:"isTransferTokenOwner"`
}

// SetGovernanceConfigInfo represents governance configuration setting information
type SetGovernanceConfigInfo struct {
	GovernanceAddress                  string  `json:"governance_address" bson:"governanceAddress"`
	MinCommunityTokensToCreateProposal float64 `json:"min_community_tokens_to_create_proposal" bson:"minCommunityTokensToCreateProposal"`
	MinInstructionHoldUpTime           float64 `json:"min_instruction_hold_up_time" bson:"minInstructionHoldUpTime"`
	BaseVotingTime                     float64 `json:"base_voting_time" bson:"baseVotingTime"`
	CommunityVoteTipping               string  `json:"community_vote_tipping" bson:"communityVoteTipping"`
	MinCouncilTokensToCreateProposal   float64 `json:"min_council_tokens_to_create_proposal" bson:"minCouncilTokensToCreateProposal"`
	CouncilVoteTipping                 string  `json:"council_vote_tipping" bson:"councilVoteTipping"`
	VotingCoolOffTime                  float64 `json:"voting_cool_off_time" bson:"votingCoolOffTime"`
	DepositExemptProposalCount         float64 `json:"deposit_exempt_proposal_count" bson:"depositExemptProposalCount"`
}

// PostMessageInfo represents message posting information
type PostMessageInfo struct {
	GovernanceProgram   string `json:"governance_program" bson:"governanceProgram"`
	RealmAddress        string `json:"realm_address" bson:"realmAddress"`
	Governance          string `json:"governance" bson:"governance"`
	Proposal            string `json:"proposal" bson:"proposal"`
	TokenOwnerRecord    string `json:"token_owner_record" bson:"tokenOwnerRecord"`
	GovernanceAuthority string `json:"governance_authority" bson:"governanceAuthority"`
	ChatMessageAddress  string `json:"chat_message_address" bson:"chatMessageAddress"`
	Payer               string `json:"payer" bson:"payer"`
	ChatType            string `json:"chatType" bson:"chatType"`
	Message             string `json:"message" bson:"message"`
	IsReply             bool   `json:"isReply" bson:"isReply"`
}

// DepositGoverningTokensInfo represents governing tokens deposit information
type DepositGoverningTokensInfo struct {
	RealmAddress            string  `json:"realm_address" bson:"realmAddress"`
	GoverningTokenSource    string  `json:"governing_token_source" bson:"governingTokenSource"`
	GoverningTokenOwner     string  `json:"governing_token_owner" bson:"governingTokenOwner"`
	TokenOwnerRecordAddress string  `json:"token_owner_record_address" bson:"tokenOwnerRecordAddress"`
	GoverningToken          string  `json:"governing_token" bson:"governingToken"`
	Payer                   string  `json:"payer" bson:"payer"`
	RealmConfigAddress      string  `json:"realm_config_address" bson:"realmConfigAddress"`
	Amount                  float64 `json:"amount" bson:"amount"`
}

// WithdrawGoverningTokensInfo represents governing tokens withdrawal information
type WithdrawGoverningTokensInfo struct {
	RealmAddress              string `json:"realm_address" bson:"realmAddress"`
	GoverningTokenDestination string `json:"governing_token_destination" bson:"governingTokenDestination"`
	GoverningTokenOwner       string `json:"governing_token_owner" bson:"governingTokenOwner"`
	TokenOwnerRecordAddress   string `json:"token_owner_record_address" bson:"tokenOwnerRecordAddress"`
	GoverningToken            string `json:"governing_token" bson:"governingToken"`
	Payer                     string `json:"payer" bson:"payer"`
	RealmConfigAddress        string `json:"realm_config_address" bson:"realmConfigAddress"`
}

// SetGovernanceDelegateInfo represents governance delegate setting information
type SetGovernanceDelegateInfo struct {
	GovernanceAuthority   string `json:"governance_authority" bson:"governanceAuthority"`
	VoteRecordAddress     string `json:"vote_record_address" bson:"voteRecordAddress"`
	NewGovernanceDelegate string `json:"new_governance_delegate" bson:"newGovernanceDelegate"`
}

// CreateGovernanceInfo represents governance creation information
type CreateGovernanceInfo struct {
	RealmAddress                       string  `json:"realm_address" bson:"realmAddress"`
	GovernanceAddress                  string  `json:"governance_address" bson:"governanceAddress"`
	GovernedAccountAddress             string  `json:"governed_account_address" bson:"governedAccountAddress"`
	TokenOwnerRecord                   string  `json:"token_owner_record" bson:"tokenOwnerRecord"`
	Payer                              string  `json:"payer" bson:"payer"`
	CreateAuthority                    string  `json:"create_authority" bson:"createAuthority"`
	MinCommunityTokensToCreateProposal float64 `json:"min_community_tokens_to_create_proposal" bson:"minCommunityTokensToCreateProposal"`
	MinInstructionHoldUpTime           float64 `json:"min_instruction_hold_up_time" bson:"minInstructionHoldUpTime"`
	BaseVotingTime                     float64 `json:"base_voting_time" bson:"baseVotingTime"`
	CommunityVoteTipping               string  `json:"community_vote_tipping" bson:"communityVoteTipping"`
	MinCouncilTokensToCreateProposal   float64 `json:"min_council_tokens_to_create_proposal" bson:"minCouncilTokensToCreateProposal"`
	CouncilVoteTipping                 string  `json:"council_vote_tipping" bson:"councilVoteTipping"`
	VotingCoolOffTime                  float64 `json:"voting_cool_off_time" bson:"votingCoolOffTime"`
	DepositExemptProposalCount         float64 `json:"deposit_exempt_proposal_count" bson:"depositExemptProposalCount"`
}

// ParseInfo parses the Info field into a specific struct based on the Type
func ParseInfo(t TxType, info any) (any, error) {
	switch t {
	case TxTypeMarketplaceWithdraw:
		return parseJSON[MarketplaceWithdrawInfo](info)
	case TxTypeNFTBid:
		return parseJSON[NFTBidInfo](info)
	case TxTypeNFTBurn:
		return parseJSON[NFTBurnInfo](info)
	case TxTypeNFTListCancel:
		return parseJSON[NFTListCancelInfo](info)
	case TxTypeNFTList:
		return parseJSON[NFTListInfo](info)
	case TxTypeNFTListUpdate:
		return parseJSON[NFTListUpdateInfo](info)
	case TxTypeNFTMint:
		return parseJSON[NFTMintInfo](info)
	case TxTypeNFTSale:
		return parseJSON[NFTSaleInfo](info)
	case TxTypeCompressedNFTSale:
		return parseJSON[CompressedNFTSaleInfo](info)
	case TxTypeCompressedNFTList:
		return parseJSON[CompressedNFTListInfo](info)
	case TxTypeCompressedNFTListCancel:
		return parseJSON[CompressedNFTListCancelInfo](info)
	case TxTypeCompressedNFTListUpdate:
		return parseJSON[CompressedNFTListUpdateInfo](info)
	case TxTypeCompressedNFTBid:
		return parseJSON[CompressedNFTBidInfo](info)
	case TxTypeCompressedNFTBidCancel:
		return parseJSON[CompressedNFTBidCancelInfo](info)
	case TxTypeCompressedNFTTakeBid:
		return parseJSON[CompressedNFTTakeBidInfo](info)
	case TxTypeCompressedNFTMint:
		return parseJSON[CompressedNFTMintInfo](info)
	case TxTypeCompressedNFTTransfer:
		return parseJSON[CompressedNFTTransferInfo](info)
	case TxTypeCompressedNFTBurn:
		return parseJSON[CompressedNFTBurnInfo](info)
	case TxTypeNFTTransfer:
		return parseJSON[NFTTransferInfo](info)
	case TxTypeSolTransfer:
		return parseJSON[SOLTransferInfo](info)
	case TxTypeTokenBurn:
		return parseJSON[TokenBurnInfo](info)
	case TxTypeTokenCreate:
		return parseJSON[TokenCreateInfo](info)
	case TxTypeTokenMint:
		return parseJSON[TokenMintInfo](info)
	case TxTypeTokenTransfer:
		return parseJSON[TokenTransferInfo](info)
	case TxTypeOfferLoan:
		return parseJSON[OfferLoanInfo](info)
	case TxTypeCancelLoan:
		return parseJSON[CancelLoanInfo](info)
	case TxTypeRepayLoan:
		return parseJSON[RepayLoanInfo](info)
	case TxTypeRepayEscrowLoan:
		return parseJSON[RepayEscrowLoanInfo](info)
	case TxTypeTakeLoan:
		return parseJSON[TakeLoanInfo](info)
	case TxTypeExtendLoan:
		return parseJSON[ExtendLoanInfo](info)
	case TxTypeExtendEscrowLoan:
		return parseJSON[ExtendEscrowLoanInfo](info)
	case TxTypeRequestLoan:
		return parseJSON[RequestLoanInfo](info)
	case TxTypeCancelRequestLoan:
		return parseJSON[CancelRequestLoanInfo](info)
	case TxTypeLiquidateLoan:
		return parseJSON[LiquidateLoanInfo](info)
	case TxTypeBuyNowPayLater:
		return parseJSON[BuyNowPayLaterInfo](info)
	case TxTypeSwap:
		return parseJSON[SwapInfo](info)
	case TxTypeCreatePool:
		return parseJSON[CreatePoolInfo](info)
	case TxTypeAddLiquidity:
		return parseJSON[AddLiquidityInfo](info)
	case TxTypeRemoveLiquidity:
		return parseJSON[RemoveLiquidityInfo](info)
	case TxTypeCollectFees:
		return parseJSON[CollectFeesInfo](info)
	case TxTypeCollectReward:
		return parseJSON[CollectRewardInfo](info)
	case TxTypeCreateRaffle:
		return parseJSON[CreateRaffleInfo](info)
	case TxTypeUpdateRaffle:
		return parseJSON[UpdateRaffleInfo](info)
	case TxTypeBuyTickets:
		return parseJSON[BuyTicketsInfo](info)
	case TxTypeRevealWinners:
		return parseJSON[RevealWinnersInfo](info)
	case TxTypeClaimPrize:
		return parseJSON[ClaimPrizeInfo](info)
	case TxTypeCloseRaffle:
		return parseJSON[CloseRaffleInfo](info)
	case TxTypeCancelRaffle:
		return parseJSON[CancelRaffleInfo](info)
	case TxTypeCreateTree:
		return parseJSON[CreateTreeInfo](info)
	case TxTypeCreateRealm:
		return parseJSON[CreateRealmInfo](info)
	case TxTypeDepositGoverningTokens:
		return parseJSON[DepositGoverningTokensInfo](info)
	case TxTypeWithdrawGoverningTokens:
		return parseJSON[WithdrawGoverningTokensInfo](info)
	case TxTypeSetGovernanceDelegate:
		return parseJSON[SetGovernanceDelegateInfo](info)
	case TxTypeCreateGovernance:
		return parseJSON[CreateGovernanceInfo](info)
	case TxTypeCreateProgramGovernance:
		return parseJSON[CreateProgramGovernanceInfo](info)
	case TxTypeCreateProposal:
		return parseJSON[CreateProposalInfo](info)
	case TxTypeAddSignatory:
		return parseJSON[AddSignatoryInfo](info)
	case TxTypeRemoveSignatory:
		return parseJSON[RemoveSignatoryInfo](info)
	case TxTypeInsertTransaction:
		return parseJSON[InsertTransactionInfo](info)
	case TxTypeRemoveTransaction:
		return parseJSON[RemoveTransactionInfo](info)
	case TxTypeCancelProposal:
		return parseJSON[CancelProposalInfo](info)
	case TxTypeSignOffProposal:
		return parseJSON[SignOffProposalInfo](info)
	case TxTypeCastVote:
		return parseJSON[CastVoteInfo](info)
	case TxTypeFinalizeVote:
		return parseJSON[FinalizeVoteInfo](info)
	case TxTypeRelinquishVote:
		return parseJSON[RelinquishVoteInfo](info)
	case TxTypeExecuteTransaction:
		return parseJSON[ExecuteTransactionInfo](info)
	case TxTypeCreateMintGovernance:
		return parseJSON[CreateMintGovernanceInfo](info)
	case TxTypeCreateTokenGovernance:
		return parseJSON[CreateTokenGovernanceInfo](info)
	case TxTypeSetGovernanceConfig:
		return parseJSON[SetGovernanceConfigInfo](info)
	case TxTypePostMessage:
		return parseJSON[PostMessageInfo](info)
	default:
		return info, nil
	}
}

// parseJSON is a helper function to parse interface{} into a specific struct
func parseJSON[T any](data interface{}) (T, error) {
	var result T

	// Convert the interface{} to JSON bytes
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return result, fmt.Errorf("failed to marshal data: %w", err)
	}

	// Unmarshal JSON bytes into the target struct
	if err := json.Unmarshal(jsonBytes, &result); err != nil {
		return result, fmt.Errorf("failed to unmarshal data: %w", err)
	}

	return result, nil
}

func ParseTxActions(tx *Tx) error {
	for i, action := range tx.Actions {
		info, err := ParseInfo(action.Type, action.Info)
		if err != nil {
			return err
		}
		tx.Actions[i].Info = info
	}
	return nil
}
