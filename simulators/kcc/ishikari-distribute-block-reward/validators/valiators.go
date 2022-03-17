// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package validators

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IValidatorsVotingData is an auto generated low-level Go binding around an user-defined struct.
type IValidatorsVotingData struct {
	Validator            common.Address
	ValidatorBallot      *big.Int
	FeeShares            *big.Int
	Ballot               *big.Int
	PendingReward        *big.Int
	RevokingBallot       *big.Int
	RevokeLockingEndTime *big.Int
}

// ValidatorsMetaData contains all meta data concerning the Validators contract.
var ValidatorsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeShares\",\"type\":\"uint256\"}],\"name\":\"AddValidatorPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeReward\",\"type\":\"uint256\"}],\"name\":\"ClaimFeeReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pendingReward\",\"type\":\"uint256\"}],\"name\":\"ClaimReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositMargin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward2\",\"type\":\"uint256\"}],\"name\":\"NotifyReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inputLength\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"okLength\",\"type\":\"uint256\"}],\"name\":\"NotifyRewardSummary\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RedeemMargin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockingEndTime\",\"type\":\"uint256\"}],\"name\":\"Revoke\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"SetFeeSetLockingDuration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"SetFeeShares\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"SetMaxPunishmentBallots\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"}],\"name\":\"SetMinSelfBallots\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"SetPoolStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"SetRevokeLockingDuration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pendingReward\",\"type\":\"uint256\"}],\"name\":\"ValidatorClaimReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Vote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE_TOTAL_SHARES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FEE_SHARES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_VALIDATORS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROPOSAL_CONTRACT\",\"outputs\":[{\"internalType\":\"contractIProposal\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PUBLISH_CONTRACT\",\"outputs\":[{\"internalType\":\"contractIPunish\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESERVEPOOL_CONTRACT\",\"outputs\":[{\"internalType\":\"contractIReservePool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_CONTRACT\",\"outputs\":[{\"internalType\":\"contractIValidators\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VOTE_UNIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"activeValidators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_proposalID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_feeShares\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"}],\"name\":\"addValidator\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"candidateInfos\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_val\",\"type\":\"address\"}],\"name\":\"claimFeeReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_val\",\"type\":\"address\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_val\",\"type\":\"address\"}],\"name\":\"claimSelfBallotsReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_val\",\"type\":\"address\"}],\"name\":\"depositMargin\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributeBlockReward\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeSetLockingDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"getPoolSelfBallots\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"getPoolSelfBallotsRewardsDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"getPoolaccRewardPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"getPoolelectedNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"getPoolenabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"getPoolfeeDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"getPoolfeeSettLockingEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"getPoolfeeShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"getPoollastRewardBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"getPoolpendingFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"getPoolsuppliedBallot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"getPoolvoterNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTopValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserVotingSummary\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validatorBallot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeShares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ballot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revokingBallot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revokeLockingEndTime\",\"type\":\"uint256\"}],\"internalType\":\"structIValidators.VotingData[]\",\"name\":\"votingDataList\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorsOfManager\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_managers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_feeShares\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validatorsContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_punishContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proposalContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reservePool\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"isActiveValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"isPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_val\",\"type\":\"address\"}],\"name\":\"isWithdrawable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marginLockingDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPunishmentAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minSelfBallots\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"enumIValidators.Operation\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"operationsDone\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_val\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"remove\",\"type\":\"bool\"}],\"name\":\"punish\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_val\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"redeemMargin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokeLockingDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_val\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"revokeVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"revokingInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockingEndTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsLeft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lockingDuration\",\"type\":\"uint256\"}],\"name\":\"setFeeSetLockingDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"setFeeSharesOfValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lockingDuration\",\"type\":\"uint256\"}],\"name\":\"setMarginLockingDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_max\",\"type\":\"uint256\"}],\"name\":\"setMaxPunishmentAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_min\",\"type\":\"uint256\"}],\"name\":\"setMinSelfBallots\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_val\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"setPoolStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lockingDuration\",\"type\":\"uint256\"}],\"name\":\"setRevokeLockingDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBallot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"newSet\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"updateActiveValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"}],\"name\":\"updateCandidateInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"usedProposals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_val\",\"type\":\"address\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_val\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_val\",\"type\":\"address\"}],\"name\":\"withdrawMargin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// ValidatorsABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorsMetaData.ABI instead.
var ValidatorsABI = ValidatorsMetaData.ABI

// Validators is an auto generated Go binding around an Ethereum contract.
type Validators struct {
	ValidatorsCaller     // Read-only binding to the contract
	ValidatorsTransactor // Write-only binding to the contract
	ValidatorsFilterer   // Log filterer for contract events
}

// ValidatorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorsSession struct {
	Contract     *Validators       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorsCallerSession struct {
	Contract *ValidatorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ValidatorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorsTransactorSession struct {
	Contract     *ValidatorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ValidatorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorsRaw struct {
	Contract *Validators // Generic contract binding to access the raw methods on
}

// ValidatorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorsCallerRaw struct {
	Contract *ValidatorsCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorsTransactorRaw struct {
	Contract *ValidatorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidators creates a new instance of Validators, bound to a specific deployed contract.
func NewValidators(address common.Address, backend bind.ContractBackend) (*Validators, error) {
	contract, err := bindValidators(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Validators{ValidatorsCaller: ValidatorsCaller{contract: contract}, ValidatorsTransactor: ValidatorsTransactor{contract: contract}, ValidatorsFilterer: ValidatorsFilterer{contract: contract}}, nil
}

// NewValidatorsCaller creates a new read-only instance of Validators, bound to a specific deployed contract.
func NewValidatorsCaller(address common.Address, caller bind.ContractCaller) (*ValidatorsCaller, error) {
	contract, err := bindValidators(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorsCaller{contract: contract}, nil
}

// NewValidatorsTransactor creates a new write-only instance of Validators, bound to a specific deployed contract.
func NewValidatorsTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorsTransactor, error) {
	contract, err := bindValidators(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorsTransactor{contract: contract}, nil
}

// NewValidatorsFilterer creates a new log filterer instance of Validators, bound to a specific deployed contract.
func NewValidatorsFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorsFilterer, error) {
	contract, err := bindValidators(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorsFilterer{contract: contract}, nil
}

// bindValidators binds a generic wrapper to an already deployed contract.
func bindValidators(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validators *ValidatorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validators.Contract.ValidatorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validators *ValidatorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validators.Contract.ValidatorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validators *ValidatorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validators.Contract.ValidatorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validators *ValidatorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validators.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validators *ValidatorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validators.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validators *ValidatorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validators.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Validators *ValidatorsCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Validators *ValidatorsSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Validators.Contract.DEFAULTADMINROLE(&_Validators.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Validators *ValidatorsCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Validators.Contract.DEFAULTADMINROLE(&_Validators.CallOpts)
}

// FEETOTALSHARES is a free data retrieval call binding the contract method 0xdbff2dfd.
//
// Solidity: function FEE_TOTAL_SHARES() view returns(uint256)
func (_Validators *ValidatorsCaller) FEETOTALSHARES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "FEE_TOTAL_SHARES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEETOTALSHARES is a free data retrieval call binding the contract method 0xdbff2dfd.
//
// Solidity: function FEE_TOTAL_SHARES() view returns(uint256)
func (_Validators *ValidatorsSession) FEETOTALSHARES() (*big.Int, error) {
	return _Validators.Contract.FEETOTALSHARES(&_Validators.CallOpts)
}

// FEETOTALSHARES is a free data retrieval call binding the contract method 0xdbff2dfd.
//
// Solidity: function FEE_TOTAL_SHARES() view returns(uint256)
func (_Validators *ValidatorsCallerSession) FEETOTALSHARES() (*big.Int, error) {
	return _Validators.Contract.FEETOTALSHARES(&_Validators.CallOpts)
}

// MAXFEESHARES is a free data retrieval call binding the contract method 0x496188e2.
//
// Solidity: function MAX_FEE_SHARES() view returns(uint256)
func (_Validators *ValidatorsCaller) MAXFEESHARES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "MAX_FEE_SHARES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFEESHARES is a free data retrieval call binding the contract method 0x496188e2.
//
// Solidity: function MAX_FEE_SHARES() view returns(uint256)
func (_Validators *ValidatorsSession) MAXFEESHARES() (*big.Int, error) {
	return _Validators.Contract.MAXFEESHARES(&_Validators.CallOpts)
}

// MAXFEESHARES is a free data retrieval call binding the contract method 0x496188e2.
//
// Solidity: function MAX_FEE_SHARES() view returns(uint256)
func (_Validators *ValidatorsCallerSession) MAXFEESHARES() (*big.Int, error) {
	return _Validators.Contract.MAXFEESHARES(&_Validators.CallOpts)
}

// MAXVALIDATORS is a free data retrieval call binding the contract method 0x714897df.
//
// Solidity: function MAX_VALIDATORS() view returns(uint16)
func (_Validators *ValidatorsCaller) MAXVALIDATORS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "MAX_VALIDATORS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXVALIDATORS is a free data retrieval call binding the contract method 0x714897df.
//
// Solidity: function MAX_VALIDATORS() view returns(uint16)
func (_Validators *ValidatorsSession) MAXVALIDATORS() (uint16, error) {
	return _Validators.Contract.MAXVALIDATORS(&_Validators.CallOpts)
}

// MAXVALIDATORS is a free data retrieval call binding the contract method 0x714897df.
//
// Solidity: function MAX_VALIDATORS() view returns(uint16)
func (_Validators *ValidatorsCallerSession) MAXVALIDATORS() (uint16, error) {
	return _Validators.Contract.MAXVALIDATORS(&_Validators.CallOpts)
}

// PROPOSALCONTRACT is a free data retrieval call binding the contract method 0x863a20b7.
//
// Solidity: function PROPOSAL_CONTRACT() view returns(address)
func (_Validators *ValidatorsCaller) PROPOSALCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "PROPOSAL_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PROPOSALCONTRACT is a free data retrieval call binding the contract method 0x863a20b7.
//
// Solidity: function PROPOSAL_CONTRACT() view returns(address)
func (_Validators *ValidatorsSession) PROPOSALCONTRACT() (common.Address, error) {
	return _Validators.Contract.PROPOSALCONTRACT(&_Validators.CallOpts)
}

// PROPOSALCONTRACT is a free data retrieval call binding the contract method 0x863a20b7.
//
// Solidity: function PROPOSAL_CONTRACT() view returns(address)
func (_Validators *ValidatorsCallerSession) PROPOSALCONTRACT() (common.Address, error) {
	return _Validators.Contract.PROPOSALCONTRACT(&_Validators.CallOpts)
}

// PUBLISHCONTRACT is a free data retrieval call binding the contract method 0xe5a99f4f.
//
// Solidity: function PUBLISH_CONTRACT() view returns(address)
func (_Validators *ValidatorsCaller) PUBLISHCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "PUBLISH_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PUBLISHCONTRACT is a free data retrieval call binding the contract method 0xe5a99f4f.
//
// Solidity: function PUBLISH_CONTRACT() view returns(address)
func (_Validators *ValidatorsSession) PUBLISHCONTRACT() (common.Address, error) {
	return _Validators.Contract.PUBLISHCONTRACT(&_Validators.CallOpts)
}

// PUBLISHCONTRACT is a free data retrieval call binding the contract method 0xe5a99f4f.
//
// Solidity: function PUBLISH_CONTRACT() view returns(address)
func (_Validators *ValidatorsCallerSession) PUBLISHCONTRACT() (common.Address, error) {
	return _Validators.Contract.PUBLISHCONTRACT(&_Validators.CallOpts)
}

// RESERVEPOOLCONTRACT is a free data retrieval call binding the contract method 0x60c80cbf.
//
// Solidity: function RESERVEPOOL_CONTRACT() view returns(address)
func (_Validators *ValidatorsCaller) RESERVEPOOLCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "RESERVEPOOL_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RESERVEPOOLCONTRACT is a free data retrieval call binding the contract method 0x60c80cbf.
//
// Solidity: function RESERVEPOOL_CONTRACT() view returns(address)
func (_Validators *ValidatorsSession) RESERVEPOOLCONTRACT() (common.Address, error) {
	return _Validators.Contract.RESERVEPOOLCONTRACT(&_Validators.CallOpts)
}

// RESERVEPOOLCONTRACT is a free data retrieval call binding the contract method 0x60c80cbf.
//
// Solidity: function RESERVEPOOL_CONTRACT() view returns(address)
func (_Validators *ValidatorsCallerSession) RESERVEPOOLCONTRACT() (common.Address, error) {
	return _Validators.Contract.RESERVEPOOLCONTRACT(&_Validators.CallOpts)
}

// VALIDATORCONTRACT is a free data retrieval call binding the contract method 0x46f75138.
//
// Solidity: function VALIDATOR_CONTRACT() view returns(address)
func (_Validators *ValidatorsCaller) VALIDATORCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "VALIDATOR_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VALIDATORCONTRACT is a free data retrieval call binding the contract method 0x46f75138.
//
// Solidity: function VALIDATOR_CONTRACT() view returns(address)
func (_Validators *ValidatorsSession) VALIDATORCONTRACT() (common.Address, error) {
	return _Validators.Contract.VALIDATORCONTRACT(&_Validators.CallOpts)
}

// VALIDATORCONTRACT is a free data retrieval call binding the contract method 0x46f75138.
//
// Solidity: function VALIDATOR_CONTRACT() view returns(address)
func (_Validators *ValidatorsCallerSession) VALIDATORCONTRACT() (common.Address, error) {
	return _Validators.Contract.VALIDATORCONTRACT(&_Validators.CallOpts)
}

// VOTEUNIT is a free data retrieval call binding the contract method 0xc531869d.
//
// Solidity: function VOTE_UNIT() view returns(uint256)
func (_Validators *ValidatorsCaller) VOTEUNIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "VOTE_UNIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VOTEUNIT is a free data retrieval call binding the contract method 0xc531869d.
//
// Solidity: function VOTE_UNIT() view returns(uint256)
func (_Validators *ValidatorsSession) VOTEUNIT() (*big.Int, error) {
	return _Validators.Contract.VOTEUNIT(&_Validators.CallOpts)
}

// VOTEUNIT is a free data retrieval call binding the contract method 0xc531869d.
//
// Solidity: function VOTE_UNIT() view returns(uint256)
func (_Validators *ValidatorsCallerSession) VOTEUNIT() (*big.Int, error) {
	return _Validators.Contract.VOTEUNIT(&_Validators.CallOpts)
}

// ActiveValidators is a free data retrieval call binding the contract method 0x14f64c78.
//
// Solidity: function activeValidators(uint256 ) view returns(address)
func (_Validators *ValidatorsCaller) ActiveValidators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "activeValidators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ActiveValidators is a free data retrieval call binding the contract method 0x14f64c78.
//
// Solidity: function activeValidators(uint256 ) view returns(address)
func (_Validators *ValidatorsSession) ActiveValidators(arg0 *big.Int) (common.Address, error) {
	return _Validators.Contract.ActiveValidators(&_Validators.CallOpts, arg0)
}

// ActiveValidators is a free data retrieval call binding the contract method 0x14f64c78.
//
// Solidity: function activeValidators(uint256 ) view returns(address)
func (_Validators *ValidatorsCallerSession) ActiveValidators(arg0 *big.Int) (common.Address, error) {
	return _Validators.Contract.ActiveValidators(&_Validators.CallOpts, arg0)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Validators *ValidatorsCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Validators *ValidatorsSession) Admin() (common.Address, error) {
	return _Validators.Contract.Admin(&_Validators.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Validators *ValidatorsCallerSession) Admin() (common.Address, error) {
	return _Validators.Contract.Admin(&_Validators.CallOpts)
}

// CandidateInfos is a free data retrieval call binding the contract method 0x2e011c1e.
//
// Solidity: function candidateInfos(address ) view returns(string website, string email, string details)
func (_Validators *ValidatorsCaller) CandidateInfos(opts *bind.CallOpts, arg0 common.Address) (struct {
	Website string
	Email   string
	Details string
}, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "candidateInfos", arg0)

	outstruct := new(struct {
		Website string
		Email   string
		Details string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Website = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Email = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Details = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// CandidateInfos is a free data retrieval call binding the contract method 0x2e011c1e.
//
// Solidity: function candidateInfos(address ) view returns(string website, string email, string details)
func (_Validators *ValidatorsSession) CandidateInfos(arg0 common.Address) (struct {
	Website string
	Email   string
	Details string
}, error) {
	return _Validators.Contract.CandidateInfos(&_Validators.CallOpts, arg0)
}

// CandidateInfos is a free data retrieval call binding the contract method 0x2e011c1e.
//
// Solidity: function candidateInfos(address ) view returns(string website, string email, string details)
func (_Validators *ValidatorsCallerSession) CandidateInfos(arg0 common.Address) (struct {
	Website string
	Email   string
	Details string
}, error) {
	return _Validators.Contract.CandidateInfos(&_Validators.CallOpts, arg0)
}

// FeeSetLockingDuration is a free data retrieval call binding the contract method 0x9e83513d.
//
// Solidity: function feeSetLockingDuration() view returns(uint256)
func (_Validators *ValidatorsCaller) FeeSetLockingDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "feeSetLockingDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeSetLockingDuration is a free data retrieval call binding the contract method 0x9e83513d.
//
// Solidity: function feeSetLockingDuration() view returns(uint256)
func (_Validators *ValidatorsSession) FeeSetLockingDuration() (*big.Int, error) {
	return _Validators.Contract.FeeSetLockingDuration(&_Validators.CallOpts)
}

// FeeSetLockingDuration is a free data retrieval call binding the contract method 0x9e83513d.
//
// Solidity: function feeSetLockingDuration() view returns(uint256)
func (_Validators *ValidatorsCallerSession) FeeSetLockingDuration() (*big.Int, error) {
	return _Validators.Contract.FeeSetLockingDuration(&_Validators.CallOpts)
}

// GetActiveValidators is a free data retrieval call binding the contract method 0x9de70258.
//
// Solidity: function getActiveValidators() view returns(address[])
func (_Validators *ValidatorsCaller) GetActiveValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "getActiveValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetActiveValidators is a free data retrieval call binding the contract method 0x9de70258.
//
// Solidity: function getActiveValidators() view returns(address[])
func (_Validators *ValidatorsSession) GetActiveValidators() ([]common.Address, error) {
	return _Validators.Contract.GetActiveValidators(&_Validators.CallOpts)
}

// GetActiveValidators is a free data retrieval call binding the contract method 0x9de70258.
//
// Solidity: function getActiveValidators() view returns(address[])
func (_Validators *ValidatorsCallerSession) GetActiveValidators() ([]common.Address, error) {
	return _Validators.Contract.GetActiveValidators(&_Validators.CallOpts)
}

// GetPoolSelfBallots is a free data retrieval call binding the contract method 0xc2657c4d.
//
// Solidity: function getPoolSelfBallots(address val) view returns(uint256)
func (_Validators *ValidatorsCaller) GetPoolSelfBallots(opts *bind.CallOpts, val common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "getPoolSelfBallots", val)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolSelfBallots is a free data retrieval call binding the contract method 0xc2657c4d.
//
// Solidity: function getPoolSelfBallots(address val) view returns(uint256)
func (_Validators *ValidatorsSession) GetPoolSelfBallots(val common.Address) (*big.Int, error) {
	return _Validators.Contract.GetPoolSelfBallots(&_Validators.CallOpts, val)
}

// GetPoolSelfBallots is a free data retrieval call binding the contract method 0xc2657c4d.
//
// Solidity: function getPoolSelfBallots(address val) view returns(uint256)
func (_Validators *ValidatorsCallerSession) GetPoolSelfBallots(val common.Address) (*big.Int, error) {
	return _Validators.Contract.GetPoolSelfBallots(&_Validators.CallOpts, val)
}

// GetPoolSelfBallotsRewardsDebt is a free data retrieval call binding the contract method 0x61167fe7.
//
// Solidity: function getPoolSelfBallotsRewardsDebt(address val) view returns(uint256)
func (_Validators *ValidatorsCaller) GetPoolSelfBallotsRewardsDebt(opts *bind.CallOpts, val common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "getPoolSelfBallotsRewardsDebt", val)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolSelfBallotsRewardsDebt is a free data retrieval call binding the contract method 0x61167fe7.
//
// Solidity: function getPoolSelfBallotsRewardsDebt(address val) view returns(uint256)
func (_Validators *ValidatorsSession) GetPoolSelfBallotsRewardsDebt(val common.Address) (*big.Int, error) {
	return _Validators.Contract.GetPoolSelfBallotsRewardsDebt(&_Validators.CallOpts, val)
}

// GetPoolSelfBallotsRewardsDebt is a free data retrieval call binding the contract method 0x61167fe7.
//
// Solidity: function getPoolSelfBallotsRewardsDebt(address val) view returns(uint256)
func (_Validators *ValidatorsCallerSession) GetPoolSelfBallotsRewardsDebt(val common.Address) (*big.Int, error) {
	return _Validators.Contract.GetPoolSelfBallotsRewardsDebt(&_Validators.CallOpts, val)
}

// GetPoolaccRewardPerShare is a free data retrieval call binding the contract method 0xce3d1daf.
//
// Solidity: function getPoolaccRewardPerShare(address val) view returns(uint256)
func (_Validators *ValidatorsCaller) GetPoolaccRewardPerShare(opts *bind.CallOpts, val common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "getPoolaccRewardPerShare", val)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolaccRewardPerShare is a free data retrieval call binding the contract method 0xce3d1daf.
//
// Solidity: function getPoolaccRewardPerShare(address val) view returns(uint256)
func (_Validators *ValidatorsSession) GetPoolaccRewardPerShare(val common.Address) (*big.Int, error) {
	return _Validators.Contract.GetPoolaccRewardPerShare(&_Validators.CallOpts, val)
}

// GetPoolaccRewardPerShare is a free data retrieval call binding the contract method 0xce3d1daf.
//
// Solidity: function getPoolaccRewardPerShare(address val) view returns(uint256)
func (_Validators *ValidatorsCallerSession) GetPoolaccRewardPerShare(val common.Address) (*big.Int, error) {
	return _Validators.Contract.GetPoolaccRewardPerShare(&_Validators.CallOpts, val)
}

// GetPoolelectedNumber is a free data retrieval call binding the contract method 0xf1656296.
//
// Solidity: function getPoolelectedNumber(address val) view returns(uint256)
func (_Validators *ValidatorsCaller) GetPoolelectedNumber(opts *bind.CallOpts, val common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "getPoolelectedNumber", val)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolelectedNumber is a free data retrieval call binding the contract method 0xf1656296.
//
// Solidity: function getPoolelectedNumber(address val) view returns(uint256)
func (_Validators *ValidatorsSession) GetPoolelectedNumber(val common.Address) (*big.Int, error) {
	return _Validators.Contract.GetPoolelectedNumber(&_Validators.CallOpts, val)
}

// GetPoolelectedNumber is a free data retrieval call binding the contract method 0xf1656296.
//
// Solidity: function getPoolelectedNumber(address val) view returns(uint256)
func (_Validators *ValidatorsCallerSession) GetPoolelectedNumber(val common.Address) (*big.Int, error) {
	return _Validators.Contract.GetPoolelectedNumber(&_Validators.CallOpts, val)
}

// GetPoolenabled is a free data retrieval call binding the contract method 0x56713343.
//
// Solidity: function getPoolenabled(address val) view returns(bool)
func (_Validators *ValidatorsCaller) GetPoolenabled(opts *bind.CallOpts, val common.Address) (bool, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "getPoolenabled", val)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetPoolenabled is a free data retrieval call binding the contract method 0x56713343.
//
// Solidity: function getPoolenabled(address val) view returns(bool)
func (_Validators *ValidatorsSession) GetPoolenabled(val common.Address) (bool, error) {
	return _Validators.Contract.GetPoolenabled(&_Validators.CallOpts, val)
}

// GetPoolenabled is a free data retrieval call binding the contract method 0x56713343.
//
// Solidity: function getPoolenabled(address val) view returns(bool)
func (_Validators *ValidatorsCallerSession) GetPoolenabled(val common.Address) (bool, error) {
	return _Validators.Contract.GetPoolenabled(&_Validators.CallOpts, val)
}

// GetPoolfeeDebt is a free data retrieval call binding the contract method 0x18e018dc.
//
// Solidity: function getPoolfeeDebt(address val) view returns(uint256)
func (_Validators *ValidatorsCaller) GetPoolfeeDebt(opts *bind.CallOpts, val common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "getPoolfeeDebt", val)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolfeeDebt is a free data retrieval call binding the contract method 0x18e018dc.
//
// Solidity: function getPoolfeeDebt(address val) view returns(uint256)
func (_Validators *ValidatorsSession) GetPoolfeeDebt(val common.Address) (*big.Int, error) {
	return _Validators.Contract.GetPoolfeeDebt(&_Validators.CallOpts, val)
}

// GetPoolfeeDebt is a free data retrieval call binding the contract method 0x18e018dc.
//
// Solidity: function getPoolfeeDebt(address val) view returns(uint256)
func (_Validators *ValidatorsCallerSession) GetPoolfeeDebt(val common.Address) (*big.Int, error) {
	return _Validators.Contract.GetPoolfeeDebt(&_Validators.CallOpts, val)
}

// GetPoolfeeSettLockingEndTime is a free data retrieval call binding the contract method 0x967945c7.
//
// Solidity: function getPoolfeeSettLockingEndTime(address val) view returns(uint256)
func (_Validators *ValidatorsCaller) GetPoolfeeSettLockingEndTime(opts *bind.CallOpts, val common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "getPoolfeeSettLockingEndTime", val)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolfeeSettLockingEndTime is a free data retrieval call binding the contract method 0x967945c7.
//
// Solidity: function getPoolfeeSettLockingEndTime(address val) view returns(uint256)
func (_Validators *ValidatorsSession) GetPoolfeeSettLockingEndTime(val common.Address) (*big.Int, error) {
	return _Validators.Contract.GetPoolfeeSettLockingEndTime(&_Validators.CallOpts, val)
}

// GetPoolfeeSettLockingEndTime is a free data retrieval call binding the contract method 0x967945c7.
//
// Solidity: function getPoolfeeSettLockingEndTime(address val) view returns(uint256)
func (_Validators *ValidatorsCallerSession) GetPoolfeeSettLockingEndTime(val common.Address) (*big.Int, error) {
	return _Validators.Contract.GetPoolfeeSettLockingEndTime(&_Validators.CallOpts, val)
}

// GetPoolfeeShares is a free data retrieval call binding the contract method 0xa2340d37.
//
// Solidity: function getPoolfeeShares(address val) view returns(uint256)
func (_Validators *ValidatorsCaller) GetPoolfeeShares(opts *bind.CallOpts, val common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "getPoolfeeShares", val)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolfeeShares is a free data retrieval call binding the contract method 0xa2340d37.
//
// Solidity: function getPoolfeeShares(address val) view returns(uint256)
func (_Validators *ValidatorsSession) GetPoolfeeShares(val common.Address) (*big.Int, error) {
	return _Validators.Contract.GetPoolfeeShares(&_Validators.CallOpts, val)
}

// GetPoolfeeShares is a free data retrieval call binding the contract method 0xa2340d37.
//
// Solidity: function getPoolfeeShares(address val) view returns(uint256)
func (_Validators *ValidatorsCallerSession) GetPoolfeeShares(val common.Address) (*big.Int, error) {
	return _Validators.Contract.GetPoolfeeShares(&_Validators.CallOpts, val)
}

// GetPoollastRewardBlock is a free data retrieval call binding the contract method 0x3619e9c3.
//
// Solidity: function getPoollastRewardBlock(address val) view returns(uint256)
func (_Validators *ValidatorsCaller) GetPoollastRewardBlock(opts *bind.CallOpts, val common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "getPoollastRewardBlock", val)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoollastRewardBlock is a free data retrieval call binding the contract method 0x3619e9c3.
//
// Solidity: function getPoollastRewardBlock(address val) view returns(uint256)
func (_Validators *ValidatorsSession) GetPoollastRewardBlock(val common.Address) (*big.Int, error) {
	return _Validators.Contract.GetPoollastRewardBlock(&_Validators.CallOpts, val)
}

// GetPoollastRewardBlock is a free data retrieval call binding the contract method 0x3619e9c3.
//
// Solidity: function getPoollastRewardBlock(address val) view returns(uint256)
func (_Validators *ValidatorsCallerSession) GetPoollastRewardBlock(val common.Address) (*big.Int, error) {
	return _Validators.Contract.GetPoollastRewardBlock(&_Validators.CallOpts, val)
}

// GetPoolpendingFee is a free data retrieval call binding the contract method 0x2687e03a.
//
// Solidity: function getPoolpendingFee(address val) view returns(uint256)
func (_Validators *ValidatorsCaller) GetPoolpendingFee(opts *bind.CallOpts, val common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "getPoolpendingFee", val)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolpendingFee is a free data retrieval call binding the contract method 0x2687e03a.
//
// Solidity: function getPoolpendingFee(address val) view returns(uint256)
func (_Validators *ValidatorsSession) GetPoolpendingFee(val common.Address) (*big.Int, error) {
	return _Validators.Contract.GetPoolpendingFee(&_Validators.CallOpts, val)
}

// GetPoolpendingFee is a free data retrieval call binding the contract method 0x2687e03a.
//
// Solidity: function getPoolpendingFee(address val) view returns(uint256)
func (_Validators *ValidatorsCallerSession) GetPoolpendingFee(val common.Address) (*big.Int, error) {
	return _Validators.Contract.GetPoolpendingFee(&_Validators.CallOpts, val)
}

// GetPoolsuppliedBallot is a free data retrieval call binding the contract method 0xd0840298.
//
// Solidity: function getPoolsuppliedBallot(address val) view returns(uint256)
func (_Validators *ValidatorsCaller) GetPoolsuppliedBallot(opts *bind.CallOpts, val common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "getPoolsuppliedBallot", val)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolsuppliedBallot is a free data retrieval call binding the contract method 0xd0840298.
//
// Solidity: function getPoolsuppliedBallot(address val) view returns(uint256)
func (_Validators *ValidatorsSession) GetPoolsuppliedBallot(val common.Address) (*big.Int, error) {
	return _Validators.Contract.GetPoolsuppliedBallot(&_Validators.CallOpts, val)
}

// GetPoolsuppliedBallot is a free data retrieval call binding the contract method 0xd0840298.
//
// Solidity: function getPoolsuppliedBallot(address val) view returns(uint256)
func (_Validators *ValidatorsCallerSession) GetPoolsuppliedBallot(val common.Address) (*big.Int, error) {
	return _Validators.Contract.GetPoolsuppliedBallot(&_Validators.CallOpts, val)
}

// GetPoolvoterNumber is a free data retrieval call binding the contract method 0x69ee53fa.
//
// Solidity: function getPoolvoterNumber(address val) view returns(uint256)
func (_Validators *ValidatorsCaller) GetPoolvoterNumber(opts *bind.CallOpts, val common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "getPoolvoterNumber", val)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolvoterNumber is a free data retrieval call binding the contract method 0x69ee53fa.
//
// Solidity: function getPoolvoterNumber(address val) view returns(uint256)
func (_Validators *ValidatorsSession) GetPoolvoterNumber(val common.Address) (*big.Int, error) {
	return _Validators.Contract.GetPoolvoterNumber(&_Validators.CallOpts, val)
}

// GetPoolvoterNumber is a free data retrieval call binding the contract method 0x69ee53fa.
//
// Solidity: function getPoolvoterNumber(address val) view returns(uint256)
func (_Validators *ValidatorsCallerSession) GetPoolvoterNumber(val common.Address) (*big.Int, error) {
	return _Validators.Contract.GetPoolvoterNumber(&_Validators.CallOpts, val)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Validators *ValidatorsCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Validators *ValidatorsSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Validators.Contract.GetRoleAdmin(&_Validators.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Validators *ValidatorsCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Validators.Contract.GetRoleAdmin(&_Validators.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Validators *ValidatorsCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Validators *ValidatorsSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Validators.Contract.GetRoleMember(&_Validators.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Validators *ValidatorsCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Validators.Contract.GetRoleMember(&_Validators.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Validators *ValidatorsCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Validators *ValidatorsSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Validators.Contract.GetRoleMemberCount(&_Validators.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Validators *ValidatorsCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Validators.Contract.GetRoleMemberCount(&_Validators.CallOpts, role)
}

// GetTopValidators is a free data retrieval call binding the contract method 0xafeea115.
//
// Solidity: function getTopValidators() view returns(address[])
func (_Validators *ValidatorsCaller) GetTopValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "getTopValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTopValidators is a free data retrieval call binding the contract method 0xafeea115.
//
// Solidity: function getTopValidators() view returns(address[])
func (_Validators *ValidatorsSession) GetTopValidators() ([]common.Address, error) {
	return _Validators.Contract.GetTopValidators(&_Validators.CallOpts)
}

// GetTopValidators is a free data retrieval call binding the contract method 0xafeea115.
//
// Solidity: function getTopValidators() view returns(address[])
func (_Validators *ValidatorsCallerSession) GetTopValidators() ([]common.Address, error) {
	return _Validators.Contract.GetTopValidators(&_Validators.CallOpts)
}

// GetUserVotingSummary is a free data retrieval call binding the contract method 0xbeb2291a.
//
// Solidity: function getUserVotingSummary(address _user) view returns((address,uint256,uint256,uint256,uint256,uint256,uint256)[] votingDataList)
func (_Validators *ValidatorsCaller) GetUserVotingSummary(opts *bind.CallOpts, _user common.Address) ([]IValidatorsVotingData, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "getUserVotingSummary", _user)

	if err != nil {
		return *new([]IValidatorsVotingData), err
	}

	out0 := *abi.ConvertType(out[0], new([]IValidatorsVotingData)).(*[]IValidatorsVotingData)

	return out0, err

}

// GetUserVotingSummary is a free data retrieval call binding the contract method 0xbeb2291a.
//
// Solidity: function getUserVotingSummary(address _user) view returns((address,uint256,uint256,uint256,uint256,uint256,uint256)[] votingDataList)
func (_Validators *ValidatorsSession) GetUserVotingSummary(_user common.Address) ([]IValidatorsVotingData, error) {
	return _Validators.Contract.GetUserVotingSummary(&_Validators.CallOpts, _user)
}

// GetUserVotingSummary is a free data retrieval call binding the contract method 0xbeb2291a.
//
// Solidity: function getUserVotingSummary(address _user) view returns((address,uint256,uint256,uint256,uint256,uint256,uint256)[] votingDataList)
func (_Validators *ValidatorsCallerSession) GetUserVotingSummary(_user common.Address) ([]IValidatorsVotingData, error) {
	return _Validators.Contract.GetUserVotingSummary(&_Validators.CallOpts, _user)
}

// GetValidatorsOfManager is a free data retrieval call binding the contract method 0x347255bc.
//
// Solidity: function getValidatorsOfManager() view returns(address[])
func (_Validators *ValidatorsCaller) GetValidatorsOfManager(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "getValidatorsOfManager")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidatorsOfManager is a free data retrieval call binding the contract method 0x347255bc.
//
// Solidity: function getValidatorsOfManager() view returns(address[])
func (_Validators *ValidatorsSession) GetValidatorsOfManager() ([]common.Address, error) {
	return _Validators.Contract.GetValidatorsOfManager(&_Validators.CallOpts)
}

// GetValidatorsOfManager is a free data retrieval call binding the contract method 0x347255bc.
//
// Solidity: function getValidatorsOfManager() view returns(address[])
func (_Validators *ValidatorsCallerSession) GetValidatorsOfManager() ([]common.Address, error) {
	return _Validators.Contract.GetValidatorsOfManager(&_Validators.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Validators *ValidatorsCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Validators *ValidatorsSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Validators.Contract.HasRole(&_Validators.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Validators *ValidatorsCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Validators.Contract.HasRole(&_Validators.CallOpts, role, account)
}

// IsActiveValidator is a free data retrieval call binding the contract method 0x40550a1c.
//
// Solidity: function isActiveValidator(address val) view returns(bool)
func (_Validators *ValidatorsCaller) IsActiveValidator(opts *bind.CallOpts, val common.Address) (bool, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "isActiveValidator", val)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveValidator is a free data retrieval call binding the contract method 0x40550a1c.
//
// Solidity: function isActiveValidator(address val) view returns(bool)
func (_Validators *ValidatorsSession) IsActiveValidator(val common.Address) (bool, error) {
	return _Validators.Contract.IsActiveValidator(&_Validators.CallOpts, val)
}

// IsActiveValidator is a free data retrieval call binding the contract method 0x40550a1c.
//
// Solidity: function isActiveValidator(address val) view returns(bool)
func (_Validators *ValidatorsCallerSession) IsActiveValidator(val common.Address) (bool, error) {
	return _Validators.Contract.IsActiveValidator(&_Validators.CallOpts, val)
}

// IsPool is a free data retrieval call binding the contract method 0x5b16ebb7.
//
// Solidity: function isPool(address _validator) view returns(bool)
func (_Validators *ValidatorsCaller) IsPool(opts *bind.CallOpts, _validator common.Address) (bool, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "isPool", _validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPool is a free data retrieval call binding the contract method 0x5b16ebb7.
//
// Solidity: function isPool(address _validator) view returns(bool)
func (_Validators *ValidatorsSession) IsPool(_validator common.Address) (bool, error) {
	return _Validators.Contract.IsPool(&_Validators.CallOpts, _validator)
}

// IsPool is a free data retrieval call binding the contract method 0x5b16ebb7.
//
// Solidity: function isPool(address _validator) view returns(bool)
func (_Validators *ValidatorsCallerSession) IsPool(_validator common.Address) (bool, error) {
	return _Validators.Contract.IsPool(&_Validators.CallOpts, _validator)
}

// IsWithdrawable is a free data retrieval call binding the contract method 0xbadca819.
//
// Solidity: function isWithdrawable(address _user, address _val) view returns(bool)
func (_Validators *ValidatorsCaller) IsWithdrawable(opts *bind.CallOpts, _user common.Address, _val common.Address) (bool, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "isWithdrawable", _user, _val)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWithdrawable is a free data retrieval call binding the contract method 0xbadca819.
//
// Solidity: function isWithdrawable(address _user, address _val) view returns(bool)
func (_Validators *ValidatorsSession) IsWithdrawable(_user common.Address, _val common.Address) (bool, error) {
	return _Validators.Contract.IsWithdrawable(&_Validators.CallOpts, _user, _val)
}

// IsWithdrawable is a free data retrieval call binding the contract method 0xbadca819.
//
// Solidity: function isWithdrawable(address _user, address _val) view returns(bool)
func (_Validators *ValidatorsCallerSession) IsWithdrawable(_user common.Address, _val common.Address) (bool, error) {
	return _Validators.Contract.IsWithdrawable(&_Validators.CallOpts, _user, _val)
}

// MarginLockingDuration is a free data retrieval call binding the contract method 0x60b1d728.
//
// Solidity: function marginLockingDuration() view returns(uint256)
func (_Validators *ValidatorsCaller) MarginLockingDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "marginLockingDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarginLockingDuration is a free data retrieval call binding the contract method 0x60b1d728.
//
// Solidity: function marginLockingDuration() view returns(uint256)
func (_Validators *ValidatorsSession) MarginLockingDuration() (*big.Int, error) {
	return _Validators.Contract.MarginLockingDuration(&_Validators.CallOpts)
}

// MarginLockingDuration is a free data retrieval call binding the contract method 0x60b1d728.
//
// Solidity: function marginLockingDuration() view returns(uint256)
func (_Validators *ValidatorsCallerSession) MarginLockingDuration() (*big.Int, error) {
	return _Validators.Contract.MarginLockingDuration(&_Validators.CallOpts)
}

// MaxPunishmentAmount is a free data retrieval call binding the contract method 0x3c3000ba.
//
// Solidity: function maxPunishmentAmount() view returns(uint256)
func (_Validators *ValidatorsCaller) MaxPunishmentAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "maxPunishmentAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPunishmentAmount is a free data retrieval call binding the contract method 0x3c3000ba.
//
// Solidity: function maxPunishmentAmount() view returns(uint256)
func (_Validators *ValidatorsSession) MaxPunishmentAmount() (*big.Int, error) {
	return _Validators.Contract.MaxPunishmentAmount(&_Validators.CallOpts)
}

// MaxPunishmentAmount is a free data retrieval call binding the contract method 0x3c3000ba.
//
// Solidity: function maxPunishmentAmount() view returns(uint256)
func (_Validators *ValidatorsCallerSession) MaxPunishmentAmount() (*big.Int, error) {
	return _Validators.Contract.MaxPunishmentAmount(&_Validators.CallOpts)
}

// MinSelfBallots is a free data retrieval call binding the contract method 0xbbb9e4d1.
//
// Solidity: function minSelfBallots() view returns(uint256)
func (_Validators *ValidatorsCaller) MinSelfBallots(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "minSelfBallots")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinSelfBallots is a free data retrieval call binding the contract method 0xbbb9e4d1.
//
// Solidity: function minSelfBallots() view returns(uint256)
func (_Validators *ValidatorsSession) MinSelfBallots() (*big.Int, error) {
	return _Validators.Contract.MinSelfBallots(&_Validators.CallOpts)
}

// MinSelfBallots is a free data retrieval call binding the contract method 0xbbb9e4d1.
//
// Solidity: function minSelfBallots() view returns(uint256)
func (_Validators *ValidatorsCallerSession) MinSelfBallots() (*big.Int, error) {
	return _Validators.Contract.MinSelfBallots(&_Validators.CallOpts)
}

// OperationsDone is a free data retrieval call binding the contract method 0x4a9b826a.
//
// Solidity: function operationsDone(uint256 , uint8 ) view returns(bool)
func (_Validators *ValidatorsCaller) OperationsDone(opts *bind.CallOpts, arg0 *big.Int, arg1 uint8) (bool, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "operationsDone", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperationsDone is a free data retrieval call binding the contract method 0x4a9b826a.
//
// Solidity: function operationsDone(uint256 , uint8 ) view returns(bool)
func (_Validators *ValidatorsSession) OperationsDone(arg0 *big.Int, arg1 uint8) (bool, error) {
	return _Validators.Contract.OperationsDone(&_Validators.CallOpts, arg0, arg1)
}

// OperationsDone is a free data retrieval call binding the contract method 0x4a9b826a.
//
// Solidity: function operationsDone(uint256 , uint8 ) view returns(bool)
func (_Validators *ValidatorsCallerSession) OperationsDone(arg0 *big.Int, arg1 uint8) (bool, error) {
	return _Validators.Contract.OperationsDone(&_Validators.CallOpts, arg0, arg1)
}

// PendingReward is a free data retrieval call binding the contract method 0x9ced7e76.
//
// Solidity: function pendingReward(address _val, address _user) view returns(uint256)
func (_Validators *ValidatorsCaller) PendingReward(opts *bind.CallOpts, _val common.Address, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "pendingReward", _val, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingReward is a free data retrieval call binding the contract method 0x9ced7e76.
//
// Solidity: function pendingReward(address _val, address _user) view returns(uint256)
func (_Validators *ValidatorsSession) PendingReward(_val common.Address, _user common.Address) (*big.Int, error) {
	return _Validators.Contract.PendingReward(&_Validators.CallOpts, _val, _user)
}

// PendingReward is a free data retrieval call binding the contract method 0x9ced7e76.
//
// Solidity: function pendingReward(address _val, address _user) view returns(uint256)
func (_Validators *ValidatorsCallerSession) PendingReward(_val common.Address, _user common.Address) (*big.Int, error) {
	return _Validators.Contract.PendingReward(&_Validators.CallOpts, _val, _user)
}

// RevokeLockingDuration is a free data retrieval call binding the contract method 0x205f5180.
//
// Solidity: function revokeLockingDuration() view returns(uint256)
func (_Validators *ValidatorsCaller) RevokeLockingDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "revokeLockingDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RevokeLockingDuration is a free data retrieval call binding the contract method 0x205f5180.
//
// Solidity: function revokeLockingDuration() view returns(uint256)
func (_Validators *ValidatorsSession) RevokeLockingDuration() (*big.Int, error) {
	return _Validators.Contract.RevokeLockingDuration(&_Validators.CallOpts)
}

// RevokeLockingDuration is a free data retrieval call binding the contract method 0x205f5180.
//
// Solidity: function revokeLockingDuration() view returns(uint256)
func (_Validators *ValidatorsCallerSession) RevokeLockingDuration() (*big.Int, error) {
	return _Validators.Contract.RevokeLockingDuration(&_Validators.CallOpts)
}

// RevokingInfo is a free data retrieval call binding the contract method 0xc76f8d40.
//
// Solidity: function revokingInfo(address , address ) view returns(uint256 amount, uint256 lockingEndTime)
func (_Validators *ValidatorsCaller) RevokingInfo(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
	Amount         *big.Int
	LockingEndTime *big.Int
}, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "revokingInfo", arg0, arg1)

	outstruct := new(struct {
		Amount         *big.Int
		LockingEndTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LockingEndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RevokingInfo is a free data retrieval call binding the contract method 0xc76f8d40.
//
// Solidity: function revokingInfo(address , address ) view returns(uint256 amount, uint256 lockingEndTime)
func (_Validators *ValidatorsSession) RevokingInfo(arg0 common.Address, arg1 common.Address) (struct {
	Amount         *big.Int
	LockingEndTime *big.Int
}, error) {
	return _Validators.Contract.RevokingInfo(&_Validators.CallOpts, arg0, arg1)
}

// RevokingInfo is a free data retrieval call binding the contract method 0xc76f8d40.
//
// Solidity: function revokingInfo(address , address ) view returns(uint256 amount, uint256 lockingEndTime)
func (_Validators *ValidatorsCallerSession) RevokingInfo(arg0 common.Address, arg1 common.Address) (struct {
	Amount         *big.Int
	LockingEndTime *big.Int
}, error) {
	return _Validators.Contract.RevokingInfo(&_Validators.CallOpts, arg0, arg1)
}

// RewardsLeft is a free data retrieval call binding the contract method 0xa8bc58f2.
//
// Solidity: function rewardsLeft() view returns(uint256)
func (_Validators *ValidatorsCaller) RewardsLeft(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "rewardsLeft")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardsLeft is a free data retrieval call binding the contract method 0xa8bc58f2.
//
// Solidity: function rewardsLeft() view returns(uint256)
func (_Validators *ValidatorsSession) RewardsLeft() (*big.Int, error) {
	return _Validators.Contract.RewardsLeft(&_Validators.CallOpts)
}

// RewardsLeft is a free data retrieval call binding the contract method 0xa8bc58f2.
//
// Solidity: function rewardsLeft() view returns(uint256)
func (_Validators *ValidatorsCallerSession) RewardsLeft() (*big.Int, error) {
	return _Validators.Contract.RewardsLeft(&_Validators.CallOpts)
}

// TotalBallot is a free data retrieval call binding the contract method 0xedf4a3c3.
//
// Solidity: function totalBallot() view returns(uint256)
func (_Validators *ValidatorsCaller) TotalBallot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "totalBallot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBallot is a free data retrieval call binding the contract method 0xedf4a3c3.
//
// Solidity: function totalBallot() view returns(uint256)
func (_Validators *ValidatorsSession) TotalBallot() (*big.Int, error) {
	return _Validators.Contract.TotalBallot(&_Validators.CallOpts)
}

// TotalBallot is a free data retrieval call binding the contract method 0xedf4a3c3.
//
// Solidity: function totalBallot() view returns(uint256)
func (_Validators *ValidatorsCallerSession) TotalBallot() (*big.Int, error) {
	return _Validators.Contract.TotalBallot(&_Validators.CallOpts)
}

// UsedProposals is a free data retrieval call binding the contract method 0x5785a619.
//
// Solidity: function usedProposals(bytes32 ) view returns(bool)
func (_Validators *ValidatorsCaller) UsedProposals(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "usedProposals", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsedProposals is a free data retrieval call binding the contract method 0x5785a619.
//
// Solidity: function usedProposals(bytes32 ) view returns(bool)
func (_Validators *ValidatorsSession) UsedProposals(arg0 [32]byte) (bool, error) {
	return _Validators.Contract.UsedProposals(&_Validators.CallOpts, arg0)
}

// UsedProposals is a free data retrieval call binding the contract method 0x5785a619.
//
// Solidity: function usedProposals(bytes32 ) view returns(bool)
func (_Validators *ValidatorsCallerSession) UsedProposals(arg0 [32]byte) (bool, error) {
	return _Validators.Contract.UsedProposals(&_Validators.CallOpts, arg0)
}

// UserInfo is a free data retrieval call binding the contract method 0x0f208beb.
//
// Solidity: function userInfo(address , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_Validators *ValidatorsCaller) UserInfo(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "userInfo", arg0, arg1)

	outstruct := new(struct {
		Amount     *big.Int
		RewardDebt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x0f208beb.
//
// Solidity: function userInfo(address , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_Validators *ValidatorsSession) UserInfo(arg0 common.Address, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _Validators.Contract.UserInfo(&_Validators.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x0f208beb.
//
// Solidity: function userInfo(address , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_Validators *ValidatorsCallerSession) UserInfo(arg0 common.Address, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _Validators.Contract.UserInfo(&_Validators.CallOpts, arg0, arg1)
}

// AddValidator is a paid mutator transaction binding the contract method 0x79c72850.
//
// Solidity: function addValidator(address _validator, address _manager, bytes32 _proposalID, uint256 _feeShares, string description, string website, string email) payable returns()
func (_Validators *ValidatorsTransactor) AddValidator(opts *bind.TransactOpts, _validator common.Address, _manager common.Address, _proposalID [32]byte, _feeShares *big.Int, description string, website string, email string) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "addValidator", _validator, _manager, _proposalID, _feeShares, description, website, email)
}

// AddValidator is a paid mutator transaction binding the contract method 0x79c72850.
//
// Solidity: function addValidator(address _validator, address _manager, bytes32 _proposalID, uint256 _feeShares, string description, string website, string email) payable returns()
func (_Validators *ValidatorsSession) AddValidator(_validator common.Address, _manager common.Address, _proposalID [32]byte, _feeShares *big.Int, description string, website string, email string) (*types.Transaction, error) {
	return _Validators.Contract.AddValidator(&_Validators.TransactOpts, _validator, _manager, _proposalID, _feeShares, description, website, email)
}

// AddValidator is a paid mutator transaction binding the contract method 0x79c72850.
//
// Solidity: function addValidator(address _validator, address _manager, bytes32 _proposalID, uint256 _feeShares, string description, string website, string email) payable returns()
func (_Validators *ValidatorsTransactorSession) AddValidator(_validator common.Address, _manager common.Address, _proposalID [32]byte, _feeShares *big.Int, description string, website string, email string) (*types.Transaction, error) {
	return _Validators.Contract.AddValidator(&_Validators.TransactOpts, _validator, _manager, _proposalID, _feeShares, description, website, email)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _admin) returns()
func (_Validators *ValidatorsTransactor) ChangeAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "changeAdmin", _admin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _admin) returns()
func (_Validators *ValidatorsSession) ChangeAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Validators.Contract.ChangeAdmin(&_Validators.TransactOpts, _admin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _admin) returns()
func (_Validators *ValidatorsTransactorSession) ChangeAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Validators.Contract.ChangeAdmin(&_Validators.TransactOpts, _admin)
}

// ClaimFeeReward is a paid mutator transaction binding the contract method 0xc6a84c03.
//
// Solidity: function claimFeeReward(address _val) returns()
func (_Validators *ValidatorsTransactor) ClaimFeeReward(opts *bind.TransactOpts, _val common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "claimFeeReward", _val)
}

// ClaimFeeReward is a paid mutator transaction binding the contract method 0xc6a84c03.
//
// Solidity: function claimFeeReward(address _val) returns()
func (_Validators *ValidatorsSession) ClaimFeeReward(_val common.Address) (*types.Transaction, error) {
	return _Validators.Contract.ClaimFeeReward(&_Validators.TransactOpts, _val)
}

// ClaimFeeReward is a paid mutator transaction binding the contract method 0xc6a84c03.
//
// Solidity: function claimFeeReward(address _val) returns()
func (_Validators *ValidatorsTransactorSession) ClaimFeeReward(_val common.Address) (*types.Transaction, error) {
	return _Validators.Contract.ClaimFeeReward(&_Validators.TransactOpts, _val)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address _val) returns()
func (_Validators *ValidatorsTransactor) ClaimReward(opts *bind.TransactOpts, _val common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "claimReward", _val)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address _val) returns()
func (_Validators *ValidatorsSession) ClaimReward(_val common.Address) (*types.Transaction, error) {
	return _Validators.Contract.ClaimReward(&_Validators.TransactOpts, _val)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address _val) returns()
func (_Validators *ValidatorsTransactorSession) ClaimReward(_val common.Address) (*types.Transaction, error) {
	return _Validators.Contract.ClaimReward(&_Validators.TransactOpts, _val)
}

// ClaimSelfBallotsReward is a paid mutator transaction binding the contract method 0x9d771a04.
//
// Solidity: function claimSelfBallotsReward(address _val) returns()
func (_Validators *ValidatorsTransactor) ClaimSelfBallotsReward(opts *bind.TransactOpts, _val common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "claimSelfBallotsReward", _val)
}

// ClaimSelfBallotsReward is a paid mutator transaction binding the contract method 0x9d771a04.
//
// Solidity: function claimSelfBallotsReward(address _val) returns()
func (_Validators *ValidatorsSession) ClaimSelfBallotsReward(_val common.Address) (*types.Transaction, error) {
	return _Validators.Contract.ClaimSelfBallotsReward(&_Validators.TransactOpts, _val)
}

// ClaimSelfBallotsReward is a paid mutator transaction binding the contract method 0x9d771a04.
//
// Solidity: function claimSelfBallotsReward(address _val) returns()
func (_Validators *ValidatorsTransactorSession) ClaimSelfBallotsReward(_val common.Address) (*types.Transaction, error) {
	return _Validators.Contract.ClaimSelfBallotsReward(&_Validators.TransactOpts, _val)
}

// DepositMargin is a paid mutator transaction binding the contract method 0x0c65f52c.
//
// Solidity: function depositMargin(address _val) payable returns()
func (_Validators *ValidatorsTransactor) DepositMargin(opts *bind.TransactOpts, _val common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "depositMargin", _val)
}

// DepositMargin is a paid mutator transaction binding the contract method 0x0c65f52c.
//
// Solidity: function depositMargin(address _val) payable returns()
func (_Validators *ValidatorsSession) DepositMargin(_val common.Address) (*types.Transaction, error) {
	return _Validators.Contract.DepositMargin(&_Validators.TransactOpts, _val)
}

// DepositMargin is a paid mutator transaction binding the contract method 0x0c65f52c.
//
// Solidity: function depositMargin(address _val) payable returns()
func (_Validators *ValidatorsTransactorSession) DepositMargin(_val common.Address) (*types.Transaction, error) {
	return _Validators.Contract.DepositMargin(&_Validators.TransactOpts, _val)
}

// DistributeBlockReward is a paid mutator transaction binding the contract method 0xd6c0edad.
//
// Solidity: function distributeBlockReward() payable returns()
func (_Validators *ValidatorsTransactor) DistributeBlockReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "distributeBlockReward")
}

// DistributeBlockReward is a paid mutator transaction binding the contract method 0xd6c0edad.
//
// Solidity: function distributeBlockReward() payable returns()
func (_Validators *ValidatorsSession) DistributeBlockReward() (*types.Transaction, error) {
	return _Validators.Contract.DistributeBlockReward(&_Validators.TransactOpts)
}

// DistributeBlockReward is a paid mutator transaction binding the contract method 0xd6c0edad.
//
// Solidity: function distributeBlockReward() payable returns()
func (_Validators *ValidatorsTransactorSession) DistributeBlockReward() (*types.Transaction, error) {
	return _Validators.Contract.DistributeBlockReward(&_Validators.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Validators *ValidatorsTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Validators *ValidatorsSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Validators.Contract.GrantRole(&_Validators.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Validators *ValidatorsTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Validators.Contract.GrantRole(&_Validators.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x65534cc0.
//
// Solidity: function initialize(address[] _validators, address[] _managers, uint256[] _feeShares, address _admin, address _validatorsContract, address _punishContract, address _proposalContract, address _reservePool) returns()
func (_Validators *ValidatorsTransactor) Initialize(opts *bind.TransactOpts, _validators []common.Address, _managers []common.Address, _feeShares []*big.Int, _admin common.Address, _validatorsContract common.Address, _punishContract common.Address, _proposalContract common.Address, _reservePool common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "initialize", _validators, _managers, _feeShares, _admin, _validatorsContract, _punishContract, _proposalContract, _reservePool)
}

// Initialize is a paid mutator transaction binding the contract method 0x65534cc0.
//
// Solidity: function initialize(address[] _validators, address[] _managers, uint256[] _feeShares, address _admin, address _validatorsContract, address _punishContract, address _proposalContract, address _reservePool) returns()
func (_Validators *ValidatorsSession) Initialize(_validators []common.Address, _managers []common.Address, _feeShares []*big.Int, _admin common.Address, _validatorsContract common.Address, _punishContract common.Address, _proposalContract common.Address, _reservePool common.Address) (*types.Transaction, error) {
	return _Validators.Contract.Initialize(&_Validators.TransactOpts, _validators, _managers, _feeShares, _admin, _validatorsContract, _punishContract, _proposalContract, _reservePool)
}

// Initialize is a paid mutator transaction binding the contract method 0x65534cc0.
//
// Solidity: function initialize(address[] _validators, address[] _managers, uint256[] _feeShares, address _admin, address _validatorsContract, address _punishContract, address _proposalContract, address _reservePool) returns()
func (_Validators *ValidatorsTransactorSession) Initialize(_validators []common.Address, _managers []common.Address, _feeShares []*big.Int, _admin common.Address, _validatorsContract common.Address, _punishContract common.Address, _proposalContract common.Address, _reservePool common.Address) (*types.Transaction, error) {
	return _Validators.Contract.Initialize(&_Validators.TransactOpts, _validators, _managers, _feeShares, _admin, _validatorsContract, _punishContract, _proposalContract, _reservePool)
}

// Punish is a paid mutator transaction binding the contract method 0x7c01f053.
//
// Solidity: function punish(address validator, bool remove) returns()
func (_Validators *ValidatorsTransactor) Punish(opts *bind.TransactOpts, validator common.Address, remove bool) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "punish", validator, remove)
}

// Punish is a paid mutator transaction binding the contract method 0x7c01f053.
//
// Solidity: function punish(address validator, bool remove) returns()
func (_Validators *ValidatorsSession) Punish(validator common.Address, remove bool) (*types.Transaction, error) {
	return _Validators.Contract.Punish(&_Validators.TransactOpts, validator, remove)
}

// Punish is a paid mutator transaction binding the contract method 0x7c01f053.
//
// Solidity: function punish(address validator, bool remove) returns()
func (_Validators *ValidatorsTransactorSession) Punish(validator common.Address, remove bool) (*types.Transaction, error) {
	return _Validators.Contract.Punish(&_Validators.TransactOpts, validator, remove)
}

// RedeemMargin is a paid mutator transaction binding the contract method 0xc4578658.
//
// Solidity: function redeemMargin(address _val, uint256 _amount) returns()
func (_Validators *ValidatorsTransactor) RedeemMargin(opts *bind.TransactOpts, _val common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "redeemMargin", _val, _amount)
}

// RedeemMargin is a paid mutator transaction binding the contract method 0xc4578658.
//
// Solidity: function redeemMargin(address _val, uint256 _amount) returns()
func (_Validators *ValidatorsSession) RedeemMargin(_val common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.RedeemMargin(&_Validators.TransactOpts, _val, _amount)
}

// RedeemMargin is a paid mutator transaction binding the contract method 0xc4578658.
//
// Solidity: function redeemMargin(address _val, uint256 _amount) returns()
func (_Validators *ValidatorsTransactorSession) RedeemMargin(_val common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.RedeemMargin(&_Validators.TransactOpts, _val, _amount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Validators *ValidatorsTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Validators *ValidatorsSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Validators.Contract.RenounceRole(&_Validators.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Validators *ValidatorsTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Validators.Contract.RenounceRole(&_Validators.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Validators *ValidatorsTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Validators *ValidatorsSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Validators.Contract.RevokeRole(&_Validators.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Validators *ValidatorsTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Validators.Contract.RevokeRole(&_Validators.TransactOpts, role, account)
}

// RevokeVote is a paid mutator transaction binding the contract method 0xee16442e.
//
// Solidity: function revokeVote(address _val, uint256 _amount) returns()
func (_Validators *ValidatorsTransactor) RevokeVote(opts *bind.TransactOpts, _val common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "revokeVote", _val, _amount)
}

// RevokeVote is a paid mutator transaction binding the contract method 0xee16442e.
//
// Solidity: function revokeVote(address _val, uint256 _amount) returns()
func (_Validators *ValidatorsSession) RevokeVote(_val common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.RevokeVote(&_Validators.TransactOpts, _val, _amount)
}

// RevokeVote is a paid mutator transaction binding the contract method 0xee16442e.
//
// Solidity: function revokeVote(address _val, uint256 _amount) returns()
func (_Validators *ValidatorsTransactorSession) RevokeVote(_val common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.RevokeVote(&_Validators.TransactOpts, _val, _amount)
}

// SetFeeSetLockingDuration is a paid mutator transaction binding the contract method 0xfd2d4b6a.
//
// Solidity: function setFeeSetLockingDuration(uint256 _lockingDuration) returns()
func (_Validators *ValidatorsTransactor) SetFeeSetLockingDuration(opts *bind.TransactOpts, _lockingDuration *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "setFeeSetLockingDuration", _lockingDuration)
}

// SetFeeSetLockingDuration is a paid mutator transaction binding the contract method 0xfd2d4b6a.
//
// Solidity: function setFeeSetLockingDuration(uint256 _lockingDuration) returns()
func (_Validators *ValidatorsSession) SetFeeSetLockingDuration(_lockingDuration *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetFeeSetLockingDuration(&_Validators.TransactOpts, _lockingDuration)
}

// SetFeeSetLockingDuration is a paid mutator transaction binding the contract method 0xfd2d4b6a.
//
// Solidity: function setFeeSetLockingDuration(uint256 _lockingDuration) returns()
func (_Validators *ValidatorsTransactorSession) SetFeeSetLockingDuration(_lockingDuration *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetFeeSetLockingDuration(&_Validators.TransactOpts, _lockingDuration)
}

// SetFeeSharesOfValidator is a paid mutator transaction binding the contract method 0x7b415b12.
//
// Solidity: function setFeeSharesOfValidator(uint256 _shares) returns()
func (_Validators *ValidatorsTransactor) SetFeeSharesOfValidator(opts *bind.TransactOpts, _shares *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "setFeeSharesOfValidator", _shares)
}

// SetFeeSharesOfValidator is a paid mutator transaction binding the contract method 0x7b415b12.
//
// Solidity: function setFeeSharesOfValidator(uint256 _shares) returns()
func (_Validators *ValidatorsSession) SetFeeSharesOfValidator(_shares *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetFeeSharesOfValidator(&_Validators.TransactOpts, _shares)
}

// SetFeeSharesOfValidator is a paid mutator transaction binding the contract method 0x7b415b12.
//
// Solidity: function setFeeSharesOfValidator(uint256 _shares) returns()
func (_Validators *ValidatorsTransactorSession) SetFeeSharesOfValidator(_shares *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetFeeSharesOfValidator(&_Validators.TransactOpts, _shares)
}

// SetMarginLockingDuration is a paid mutator transaction binding the contract method 0xf1143a46.
//
// Solidity: function setMarginLockingDuration(uint256 _lockingDuration) returns()
func (_Validators *ValidatorsTransactor) SetMarginLockingDuration(opts *bind.TransactOpts, _lockingDuration *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "setMarginLockingDuration", _lockingDuration)
}

// SetMarginLockingDuration is a paid mutator transaction binding the contract method 0xf1143a46.
//
// Solidity: function setMarginLockingDuration(uint256 _lockingDuration) returns()
func (_Validators *ValidatorsSession) SetMarginLockingDuration(_lockingDuration *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetMarginLockingDuration(&_Validators.TransactOpts, _lockingDuration)
}

// SetMarginLockingDuration is a paid mutator transaction binding the contract method 0xf1143a46.
//
// Solidity: function setMarginLockingDuration(uint256 _lockingDuration) returns()
func (_Validators *ValidatorsTransactorSession) SetMarginLockingDuration(_lockingDuration *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetMarginLockingDuration(&_Validators.TransactOpts, _lockingDuration)
}

// SetMaxPunishmentAmount is a paid mutator transaction binding the contract method 0x5463080f.
//
// Solidity: function setMaxPunishmentAmount(uint256 _max) returns()
func (_Validators *ValidatorsTransactor) SetMaxPunishmentAmount(opts *bind.TransactOpts, _max *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "setMaxPunishmentAmount", _max)
}

// SetMaxPunishmentAmount is a paid mutator transaction binding the contract method 0x5463080f.
//
// Solidity: function setMaxPunishmentAmount(uint256 _max) returns()
func (_Validators *ValidatorsSession) SetMaxPunishmentAmount(_max *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetMaxPunishmentAmount(&_Validators.TransactOpts, _max)
}

// SetMaxPunishmentAmount is a paid mutator transaction binding the contract method 0x5463080f.
//
// Solidity: function setMaxPunishmentAmount(uint256 _max) returns()
func (_Validators *ValidatorsTransactorSession) SetMaxPunishmentAmount(_max *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetMaxPunishmentAmount(&_Validators.TransactOpts, _max)
}

// SetMinSelfBallots is a paid mutator transaction binding the contract method 0xc3d117bf.
//
// Solidity: function setMinSelfBallots(uint256 _min) returns()
func (_Validators *ValidatorsTransactor) SetMinSelfBallots(opts *bind.TransactOpts, _min *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "setMinSelfBallots", _min)
}

// SetMinSelfBallots is a paid mutator transaction binding the contract method 0xc3d117bf.
//
// Solidity: function setMinSelfBallots(uint256 _min) returns()
func (_Validators *ValidatorsSession) SetMinSelfBallots(_min *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetMinSelfBallots(&_Validators.TransactOpts, _min)
}

// SetMinSelfBallots is a paid mutator transaction binding the contract method 0xc3d117bf.
//
// Solidity: function setMinSelfBallots(uint256 _min) returns()
func (_Validators *ValidatorsTransactorSession) SetMinSelfBallots(_min *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetMinSelfBallots(&_Validators.TransactOpts, _min)
}

// SetPoolStatus is a paid mutator transaction binding the contract method 0x33ad1e01.
//
// Solidity: function setPoolStatus(address _val, bool _enabled) returns()
func (_Validators *ValidatorsTransactor) SetPoolStatus(opts *bind.TransactOpts, _val common.Address, _enabled bool) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "setPoolStatus", _val, _enabled)
}

// SetPoolStatus is a paid mutator transaction binding the contract method 0x33ad1e01.
//
// Solidity: function setPoolStatus(address _val, bool _enabled) returns()
func (_Validators *ValidatorsSession) SetPoolStatus(_val common.Address, _enabled bool) (*types.Transaction, error) {
	return _Validators.Contract.SetPoolStatus(&_Validators.TransactOpts, _val, _enabled)
}

// SetPoolStatus is a paid mutator transaction binding the contract method 0x33ad1e01.
//
// Solidity: function setPoolStatus(address _val, bool _enabled) returns()
func (_Validators *ValidatorsTransactorSession) SetPoolStatus(_val common.Address, _enabled bool) (*types.Transaction, error) {
	return _Validators.Contract.SetPoolStatus(&_Validators.TransactOpts, _val, _enabled)
}

// SetRevokeLockingDuration is a paid mutator transaction binding the contract method 0x5b337c2c.
//
// Solidity: function setRevokeLockingDuration(uint256 _lockingDuration) returns()
func (_Validators *ValidatorsTransactor) SetRevokeLockingDuration(opts *bind.TransactOpts, _lockingDuration *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "setRevokeLockingDuration", _lockingDuration)
}

// SetRevokeLockingDuration is a paid mutator transaction binding the contract method 0x5b337c2c.
//
// Solidity: function setRevokeLockingDuration(uint256 _lockingDuration) returns()
func (_Validators *ValidatorsSession) SetRevokeLockingDuration(_lockingDuration *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetRevokeLockingDuration(&_Validators.TransactOpts, _lockingDuration)
}

// SetRevokeLockingDuration is a paid mutator transaction binding the contract method 0x5b337c2c.
//
// Solidity: function setRevokeLockingDuration(uint256 _lockingDuration) returns()
func (_Validators *ValidatorsTransactorSession) SetRevokeLockingDuration(_lockingDuration *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetRevokeLockingDuration(&_Validators.TransactOpts, _lockingDuration)
}

// UpdateActiveValidatorSet is a paid mutator transaction binding the contract method 0x6846992a.
//
// Solidity: function updateActiveValidatorSet(address[] newSet, uint256 epoch) returns()
func (_Validators *ValidatorsTransactor) UpdateActiveValidatorSet(opts *bind.TransactOpts, newSet []common.Address, epoch *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "updateActiveValidatorSet", newSet, epoch)
}

// UpdateActiveValidatorSet is a paid mutator transaction binding the contract method 0x6846992a.
//
// Solidity: function updateActiveValidatorSet(address[] newSet, uint256 epoch) returns()
func (_Validators *ValidatorsSession) UpdateActiveValidatorSet(newSet []common.Address, epoch *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.UpdateActiveValidatorSet(&_Validators.TransactOpts, newSet, epoch)
}

// UpdateActiveValidatorSet is a paid mutator transaction binding the contract method 0x6846992a.
//
// Solidity: function updateActiveValidatorSet(address[] newSet, uint256 epoch) returns()
func (_Validators *ValidatorsTransactorSession) UpdateActiveValidatorSet(newSet []common.Address, epoch *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.UpdateActiveValidatorSet(&_Validators.TransactOpts, newSet, epoch)
}

// UpdateCandidateInfo is a paid mutator transaction binding the contract method 0x862fab42.
//
// Solidity: function updateCandidateInfo(address _validator, string details, string website, string email) returns()
func (_Validators *ValidatorsTransactor) UpdateCandidateInfo(opts *bind.TransactOpts, _validator common.Address, details string, website string, email string) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "updateCandidateInfo", _validator, details, website, email)
}

// UpdateCandidateInfo is a paid mutator transaction binding the contract method 0x862fab42.
//
// Solidity: function updateCandidateInfo(address _validator, string details, string website, string email) returns()
func (_Validators *ValidatorsSession) UpdateCandidateInfo(_validator common.Address, details string, website string, email string) (*types.Transaction, error) {
	return _Validators.Contract.UpdateCandidateInfo(&_Validators.TransactOpts, _validator, details, website, email)
}

// UpdateCandidateInfo is a paid mutator transaction binding the contract method 0x862fab42.
//
// Solidity: function updateCandidateInfo(address _validator, string details, string website, string email) returns()
func (_Validators *ValidatorsTransactorSession) UpdateCandidateInfo(_validator common.Address, details string, website string, email string) (*types.Transaction, error) {
	return _Validators.Contract.UpdateCandidateInfo(&_Validators.TransactOpts, _validator, details, website, email)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(address _val) payable returns()
func (_Validators *ValidatorsTransactor) Vote(opts *bind.TransactOpts, _val common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "vote", _val)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(address _val) payable returns()
func (_Validators *ValidatorsSession) Vote(_val common.Address) (*types.Transaction, error) {
	return _Validators.Contract.Vote(&_Validators.TransactOpts, _val)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(address _val) payable returns()
func (_Validators *ValidatorsTransactorSession) Vote(_val common.Address) (*types.Transaction, error) {
	return _Validators.Contract.Vote(&_Validators.TransactOpts, _val)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _val) returns()
func (_Validators *ValidatorsTransactor) Withdraw(opts *bind.TransactOpts, _val common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "withdraw", _val)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _val) returns()
func (_Validators *ValidatorsSession) Withdraw(_val common.Address) (*types.Transaction, error) {
	return _Validators.Contract.Withdraw(&_Validators.TransactOpts, _val)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _val) returns()
func (_Validators *ValidatorsTransactorSession) Withdraw(_val common.Address) (*types.Transaction, error) {
	return _Validators.Contract.Withdraw(&_Validators.TransactOpts, _val)
}

// WithdrawMargin is a paid mutator transaction binding the contract method 0xc46102cc.
//
// Solidity: function withdrawMargin(address _val) returns()
func (_Validators *ValidatorsTransactor) WithdrawMargin(opts *bind.TransactOpts, _val common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "withdrawMargin", _val)
}

// WithdrawMargin is a paid mutator transaction binding the contract method 0xc46102cc.
//
// Solidity: function withdrawMargin(address _val) returns()
func (_Validators *ValidatorsSession) WithdrawMargin(_val common.Address) (*types.Transaction, error) {
	return _Validators.Contract.WithdrawMargin(&_Validators.TransactOpts, _val)
}

// WithdrawMargin is a paid mutator transaction binding the contract method 0xc46102cc.
//
// Solidity: function withdrawMargin(address _val) returns()
func (_Validators *ValidatorsTransactorSession) WithdrawMargin(_val common.Address) (*types.Transaction, error) {
	return _Validators.Contract.WithdrawMargin(&_Validators.TransactOpts, _val)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Validators *ValidatorsTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validators.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Validators *ValidatorsSession) Receive() (*types.Transaction, error) {
	return _Validators.Contract.Receive(&_Validators.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Validators *ValidatorsTransactorSession) Receive() (*types.Transaction, error) {
	return _Validators.Contract.Receive(&_Validators.TransactOpts)
}

// ValidatorsAddValidatorPoolIterator is returned from FilterAddValidatorPool and is used to iterate over the raw logs and unpacked data for AddValidatorPool events raised by the Validators contract.
type ValidatorsAddValidatorPoolIterator struct {
	Event *ValidatorsAddValidatorPool // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsAddValidatorPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsAddValidatorPool)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsAddValidatorPool)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsAddValidatorPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsAddValidatorPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsAddValidatorPool represents a AddValidatorPool event raised by the Validators contract.
type ValidatorsAddValidatorPool struct {
	User      common.Address
	Pid       *big.Int
	FeeShares *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddValidatorPool is a free log retrieval operation binding the contract event 0xfb11d94006946e73fd14303e895332be5303b4847147999d8850a73a2f208322.
//
// Solidity: event AddValidatorPool(address indexed user, uint256 indexed pid, uint256 feeShares)
func (_Validators *ValidatorsFilterer) FilterAddValidatorPool(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*ValidatorsAddValidatorPoolIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "AddValidatorPool", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsAddValidatorPoolIterator{contract: _Validators.contract, event: "AddValidatorPool", logs: logs, sub: sub}, nil
}

// WatchAddValidatorPool is a free log subscription operation binding the contract event 0xfb11d94006946e73fd14303e895332be5303b4847147999d8850a73a2f208322.
//
// Solidity: event AddValidatorPool(address indexed user, uint256 indexed pid, uint256 feeShares)
func (_Validators *ValidatorsFilterer) WatchAddValidatorPool(opts *bind.WatchOpts, sink chan<- *ValidatorsAddValidatorPool, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "AddValidatorPool", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsAddValidatorPool)
				if err := _Validators.contract.UnpackLog(event, "AddValidatorPool", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddValidatorPool is a log parse operation binding the contract event 0xfb11d94006946e73fd14303e895332be5303b4847147999d8850a73a2f208322.
//
// Solidity: event AddValidatorPool(address indexed user, uint256 indexed pid, uint256 feeShares)
func (_Validators *ValidatorsFilterer) ParseAddValidatorPool(log types.Log) (*ValidatorsAddValidatorPool, error) {
	event := new(ValidatorsAddValidatorPool)
	if err := _Validators.contract.UnpackLog(event, "AddValidatorPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsClaimFeeRewardIterator is returned from FilterClaimFeeReward and is used to iterate over the raw logs and unpacked data for ClaimFeeReward events raised by the Validators contract.
type ValidatorsClaimFeeRewardIterator struct {
	Event *ValidatorsClaimFeeReward // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsClaimFeeRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsClaimFeeReward)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsClaimFeeReward)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsClaimFeeRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsClaimFeeRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsClaimFeeReward represents a ClaimFeeReward event raised by the Validators contract.
type ValidatorsClaimFeeReward struct {
	Validator common.Address
	FeeReward *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimFeeReward is a free log retrieval operation binding the contract event 0xa847ccfc1c24e1b6493e19c018745f8e9f60288140e209713efc58517bfc0ce3.
//
// Solidity: event ClaimFeeReward(address indexed validator, uint256 feeReward)
func (_Validators *ValidatorsFilterer) FilterClaimFeeReward(opts *bind.FilterOpts, validator []common.Address) (*ValidatorsClaimFeeRewardIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ClaimFeeReward", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsClaimFeeRewardIterator{contract: _Validators.contract, event: "ClaimFeeReward", logs: logs, sub: sub}, nil
}

// WatchClaimFeeReward is a free log subscription operation binding the contract event 0xa847ccfc1c24e1b6493e19c018745f8e9f60288140e209713efc58517bfc0ce3.
//
// Solidity: event ClaimFeeReward(address indexed validator, uint256 feeReward)
func (_Validators *ValidatorsFilterer) WatchClaimFeeReward(opts *bind.WatchOpts, sink chan<- *ValidatorsClaimFeeReward, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ClaimFeeReward", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsClaimFeeReward)
				if err := _Validators.contract.UnpackLog(event, "ClaimFeeReward", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimFeeReward is a log parse operation binding the contract event 0xa847ccfc1c24e1b6493e19c018745f8e9f60288140e209713efc58517bfc0ce3.
//
// Solidity: event ClaimFeeReward(address indexed validator, uint256 feeReward)
func (_Validators *ValidatorsFilterer) ParseClaimFeeReward(log types.Log) (*ValidatorsClaimFeeReward, error) {
	event := new(ValidatorsClaimFeeReward)
	if err := _Validators.contract.UnpackLog(event, "ClaimFeeReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsClaimRewardIterator is returned from FilterClaimReward and is used to iterate over the raw logs and unpacked data for ClaimReward events raised by the Validators contract.
type ValidatorsClaimRewardIterator struct {
	Event *ValidatorsClaimReward // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsClaimRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsClaimReward)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsClaimReward)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsClaimRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsClaimRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsClaimReward represents a ClaimReward event raised by the Validators contract.
type ValidatorsClaimReward struct {
	User          common.Address
	Validator     common.Address
	PendingReward *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterClaimReward is a free log retrieval operation binding the contract event 0x7e77f685b38c861064cb08f2776eb5dfd3c82f652ed9f21221b8c53b75628e51.
//
// Solidity: event ClaimReward(address indexed user, address indexed validator, uint256 pendingReward)
func (_Validators *ValidatorsFilterer) FilterClaimReward(opts *bind.FilterOpts, user []common.Address, validator []common.Address) (*ValidatorsClaimRewardIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ClaimReward", userRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsClaimRewardIterator{contract: _Validators.contract, event: "ClaimReward", logs: logs, sub: sub}, nil
}

// WatchClaimReward is a free log subscription operation binding the contract event 0x7e77f685b38c861064cb08f2776eb5dfd3c82f652ed9f21221b8c53b75628e51.
//
// Solidity: event ClaimReward(address indexed user, address indexed validator, uint256 pendingReward)
func (_Validators *ValidatorsFilterer) WatchClaimReward(opts *bind.WatchOpts, sink chan<- *ValidatorsClaimReward, user []common.Address, validator []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ClaimReward", userRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsClaimReward)
				if err := _Validators.contract.UnpackLog(event, "ClaimReward", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimReward is a log parse operation binding the contract event 0x7e77f685b38c861064cb08f2776eb5dfd3c82f652ed9f21221b8c53b75628e51.
//
// Solidity: event ClaimReward(address indexed user, address indexed validator, uint256 pendingReward)
func (_Validators *ValidatorsFilterer) ParseClaimReward(log types.Log) (*ValidatorsClaimReward, error) {
	event := new(ValidatorsClaimReward)
	if err := _Validators.contract.UnpackLog(event, "ClaimReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsDepositMarginIterator is returned from FilterDepositMargin and is used to iterate over the raw logs and unpacked data for DepositMargin events raised by the Validators contract.
type ValidatorsDepositMarginIterator struct {
	Event *ValidatorsDepositMargin // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsDepositMarginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsDepositMargin)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsDepositMargin)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsDepositMarginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsDepositMarginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsDepositMargin represents a DepositMargin event raised by the Validators contract.
type ValidatorsDepositMargin struct {
	From      common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDepositMargin is a free log retrieval operation binding the contract event 0xf297865ac57d5e90b6602e368c1ee6df8104228d78fce8ed8923b48a96947964.
//
// Solidity: event DepositMargin(address indexed from, address indexed validator, uint256 amount)
func (_Validators *ValidatorsFilterer) FilterDepositMargin(opts *bind.FilterOpts, from []common.Address, validator []common.Address) (*ValidatorsDepositMarginIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "DepositMargin", fromRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsDepositMarginIterator{contract: _Validators.contract, event: "DepositMargin", logs: logs, sub: sub}, nil
}

// WatchDepositMargin is a free log subscription operation binding the contract event 0xf297865ac57d5e90b6602e368c1ee6df8104228d78fce8ed8923b48a96947964.
//
// Solidity: event DepositMargin(address indexed from, address indexed validator, uint256 amount)
func (_Validators *ValidatorsFilterer) WatchDepositMargin(opts *bind.WatchOpts, sink chan<- *ValidatorsDepositMargin, from []common.Address, validator []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "DepositMargin", fromRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsDepositMargin)
				if err := _Validators.contract.UnpackLog(event, "DepositMargin", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositMargin is a log parse operation binding the contract event 0xf297865ac57d5e90b6602e368c1ee6df8104228d78fce8ed8923b48a96947964.
//
// Solidity: event DepositMargin(address indexed from, address indexed validator, uint256 amount)
func (_Validators *ValidatorsFilterer) ParseDepositMargin(log types.Log) (*ValidatorsDepositMargin, error) {
	event := new(ValidatorsDepositMargin)
	if err := _Validators.contract.UnpackLog(event, "DepositMargin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsNotifyRewardIterator is returned from FilterNotifyReward and is used to iterate over the raw logs and unpacked data for NotifyReward events raised by the Validators contract.
type ValidatorsNotifyRewardIterator struct {
	Event *ValidatorsNotifyReward // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsNotifyRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsNotifyReward)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsNotifyReward)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsNotifyRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsNotifyRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsNotifyReward represents a NotifyReward event raised by the Validators contract.
type ValidatorsNotifyReward struct {
	User    common.Address
	Reward1 *big.Int
	Reward2 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNotifyReward is a free log retrieval operation binding the contract event 0x3c0f5c48b0ffa2c570c1a0f4fbf7b0f8982213afff9eb42cd258ead865cf3c9d.
//
// Solidity: event NotifyReward(address indexed user, uint256 reward1, uint256 reward2)
func (_Validators *ValidatorsFilterer) FilterNotifyReward(opts *bind.FilterOpts, user []common.Address) (*ValidatorsNotifyRewardIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "NotifyReward", userRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsNotifyRewardIterator{contract: _Validators.contract, event: "NotifyReward", logs: logs, sub: sub}, nil
}

// WatchNotifyReward is a free log subscription operation binding the contract event 0x3c0f5c48b0ffa2c570c1a0f4fbf7b0f8982213afff9eb42cd258ead865cf3c9d.
//
// Solidity: event NotifyReward(address indexed user, uint256 reward1, uint256 reward2)
func (_Validators *ValidatorsFilterer) WatchNotifyReward(opts *bind.WatchOpts, sink chan<- *ValidatorsNotifyReward, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "NotifyReward", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsNotifyReward)
				if err := _Validators.contract.UnpackLog(event, "NotifyReward", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNotifyReward is a log parse operation binding the contract event 0x3c0f5c48b0ffa2c570c1a0f4fbf7b0f8982213afff9eb42cd258ead865cf3c9d.
//
// Solidity: event NotifyReward(address indexed user, uint256 reward1, uint256 reward2)
func (_Validators *ValidatorsFilterer) ParseNotifyReward(log types.Log) (*ValidatorsNotifyReward, error) {
	event := new(ValidatorsNotifyReward)
	if err := _Validators.contract.UnpackLog(event, "NotifyReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsNotifyRewardSummaryIterator is returned from FilterNotifyRewardSummary and is used to iterate over the raw logs and unpacked data for NotifyRewardSummary events raised by the Validators contract.
type ValidatorsNotifyRewardSummaryIterator struct {
	Event *ValidatorsNotifyRewardSummary // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsNotifyRewardSummaryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsNotifyRewardSummary)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsNotifyRewardSummary)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsNotifyRewardSummaryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsNotifyRewardSummaryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsNotifyRewardSummary represents a NotifyRewardSummary event raised by the Validators contract.
type ValidatorsNotifyRewardSummary struct {
	InputLength *big.Int
	OkLength    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNotifyRewardSummary is a free log retrieval operation binding the contract event 0xc5bd0052fa3f4e939e9fe3624e6eb68e9370d71ed37f4a24c31dae6dcec33bbf.
//
// Solidity: event NotifyRewardSummary(uint256 inputLength, uint256 okLength)
func (_Validators *ValidatorsFilterer) FilterNotifyRewardSummary(opts *bind.FilterOpts) (*ValidatorsNotifyRewardSummaryIterator, error) {

	logs, sub, err := _Validators.contract.FilterLogs(opts, "NotifyRewardSummary")
	if err != nil {
		return nil, err
	}
	return &ValidatorsNotifyRewardSummaryIterator{contract: _Validators.contract, event: "NotifyRewardSummary", logs: logs, sub: sub}, nil
}

// WatchNotifyRewardSummary is a free log subscription operation binding the contract event 0xc5bd0052fa3f4e939e9fe3624e6eb68e9370d71ed37f4a24c31dae6dcec33bbf.
//
// Solidity: event NotifyRewardSummary(uint256 inputLength, uint256 okLength)
func (_Validators *ValidatorsFilterer) WatchNotifyRewardSummary(opts *bind.WatchOpts, sink chan<- *ValidatorsNotifyRewardSummary) (event.Subscription, error) {

	logs, sub, err := _Validators.contract.WatchLogs(opts, "NotifyRewardSummary")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsNotifyRewardSummary)
				if err := _Validators.contract.UnpackLog(event, "NotifyRewardSummary", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNotifyRewardSummary is a log parse operation binding the contract event 0xc5bd0052fa3f4e939e9fe3624e6eb68e9370d71ed37f4a24c31dae6dcec33bbf.
//
// Solidity: event NotifyRewardSummary(uint256 inputLength, uint256 okLength)
func (_Validators *ValidatorsFilterer) ParseNotifyRewardSummary(log types.Log) (*ValidatorsNotifyRewardSummary, error) {
	event := new(ValidatorsNotifyRewardSummary)
	if err := _Validators.contract.UnpackLog(event, "NotifyRewardSummary", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsRedeemMarginIterator is returned from FilterRedeemMargin and is used to iterate over the raw logs and unpacked data for RedeemMargin events raised by the Validators contract.
type ValidatorsRedeemMarginIterator struct {
	Event *ValidatorsRedeemMargin // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsRedeemMarginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsRedeemMargin)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsRedeemMargin)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsRedeemMarginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsRedeemMarginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsRedeemMargin represents a RedeemMargin event raised by the Validators contract.
type ValidatorsRedeemMargin struct {
	From      common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRedeemMargin is a free log retrieval operation binding the contract event 0x25f15cdd571a1fdd95de770368cf68023dc79c8d2d4a0f35208a3b0b07bf23bc.
//
// Solidity: event RedeemMargin(address indexed from, address indexed validator, uint256 amount)
func (_Validators *ValidatorsFilterer) FilterRedeemMargin(opts *bind.FilterOpts, from []common.Address, validator []common.Address) (*ValidatorsRedeemMarginIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "RedeemMargin", fromRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsRedeemMarginIterator{contract: _Validators.contract, event: "RedeemMargin", logs: logs, sub: sub}, nil
}

// WatchRedeemMargin is a free log subscription operation binding the contract event 0x25f15cdd571a1fdd95de770368cf68023dc79c8d2d4a0f35208a3b0b07bf23bc.
//
// Solidity: event RedeemMargin(address indexed from, address indexed validator, uint256 amount)
func (_Validators *ValidatorsFilterer) WatchRedeemMargin(opts *bind.WatchOpts, sink chan<- *ValidatorsRedeemMargin, from []common.Address, validator []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "RedeemMargin", fromRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsRedeemMargin)
				if err := _Validators.contract.UnpackLog(event, "RedeemMargin", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRedeemMargin is a log parse operation binding the contract event 0x25f15cdd571a1fdd95de770368cf68023dc79c8d2d4a0f35208a3b0b07bf23bc.
//
// Solidity: event RedeemMargin(address indexed from, address indexed validator, uint256 amount)
func (_Validators *ValidatorsFilterer) ParseRedeemMargin(log types.Log) (*ValidatorsRedeemMargin, error) {
	event := new(ValidatorsRedeemMargin)
	if err := _Validators.contract.UnpackLog(event, "RedeemMargin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsRevokeIterator is returned from FilterRevoke and is used to iterate over the raw logs and unpacked data for Revoke events raised by the Validators contract.
type ValidatorsRevokeIterator struct {
	Event *ValidatorsRevoke // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsRevokeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsRevoke)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsRevoke)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsRevokeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsRevokeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsRevoke represents a Revoke event raised by the Validators contract.
type ValidatorsRevoke struct {
	User           common.Address
	Validator      common.Address
	Amount         *big.Int
	LockingEndTime *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRevoke is a free log retrieval operation binding the contract event 0xd6f80c7d68e3e62bd7a51c3d37e575c1cfbc311c07487b69ef4eb570bc21cb68.
//
// Solidity: event Revoke(address indexed user, address indexed validator, uint256 amount, uint256 lockingEndTime)
func (_Validators *ValidatorsFilterer) FilterRevoke(opts *bind.FilterOpts, user []common.Address, validator []common.Address) (*ValidatorsRevokeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "Revoke", userRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsRevokeIterator{contract: _Validators.contract, event: "Revoke", logs: logs, sub: sub}, nil
}

// WatchRevoke is a free log subscription operation binding the contract event 0xd6f80c7d68e3e62bd7a51c3d37e575c1cfbc311c07487b69ef4eb570bc21cb68.
//
// Solidity: event Revoke(address indexed user, address indexed validator, uint256 amount, uint256 lockingEndTime)
func (_Validators *ValidatorsFilterer) WatchRevoke(opts *bind.WatchOpts, sink chan<- *ValidatorsRevoke, user []common.Address, validator []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "Revoke", userRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsRevoke)
				if err := _Validators.contract.UnpackLog(event, "Revoke", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRevoke is a log parse operation binding the contract event 0xd6f80c7d68e3e62bd7a51c3d37e575c1cfbc311c07487b69ef4eb570bc21cb68.
//
// Solidity: event Revoke(address indexed user, address indexed validator, uint256 amount, uint256 lockingEndTime)
func (_Validators *ValidatorsFilterer) ParseRevoke(log types.Log) (*ValidatorsRevoke, error) {
	event := new(ValidatorsRevoke)
	if err := _Validators.contract.UnpackLog(event, "Revoke", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsRewardTransferIterator is returned from FilterRewardTransfer and is used to iterate over the raw logs and unpacked data for RewardTransfer events raised by the Validators contract.
type ValidatorsRewardTransferIterator struct {
	Event *ValidatorsRewardTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsRewardTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsRewardTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsRewardTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsRewardTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsRewardTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsRewardTransfer represents a RewardTransfer event raised by the Validators contract.
type ValidatorsRewardTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardTransfer is a free log retrieval operation binding the contract event 0xa282a314ffba449688a7c5eb4426bd8e1781944ce824ce34dfaea4691d08f146.
//
// Solidity: event RewardTransfer(address indexed from, address indexed to, uint256 amount)
func (_Validators *ValidatorsFilterer) FilterRewardTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ValidatorsRewardTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "RewardTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsRewardTransferIterator{contract: _Validators.contract, event: "RewardTransfer", logs: logs, sub: sub}, nil
}

// WatchRewardTransfer is a free log subscription operation binding the contract event 0xa282a314ffba449688a7c5eb4426bd8e1781944ce824ce34dfaea4691d08f146.
//
// Solidity: event RewardTransfer(address indexed from, address indexed to, uint256 amount)
func (_Validators *ValidatorsFilterer) WatchRewardTransfer(opts *bind.WatchOpts, sink chan<- *ValidatorsRewardTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "RewardTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsRewardTransfer)
				if err := _Validators.contract.UnpackLog(event, "RewardTransfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardTransfer is a log parse operation binding the contract event 0xa282a314ffba449688a7c5eb4426bd8e1781944ce824ce34dfaea4691d08f146.
//
// Solidity: event RewardTransfer(address indexed from, address indexed to, uint256 amount)
func (_Validators *ValidatorsFilterer) ParseRewardTransfer(log types.Log) (*ValidatorsRewardTransfer, error) {
	event := new(ValidatorsRewardTransfer)
	if err := _Validators.contract.UnpackLog(event, "RewardTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Validators contract.
type ValidatorsRoleGrantedIterator struct {
	Event *ValidatorsRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsRoleGranted represents a RoleGranted event raised by the Validators contract.
type ValidatorsRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Validators *ValidatorsFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ValidatorsRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsRoleGrantedIterator{contract: _Validators.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Validators *ValidatorsFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ValidatorsRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsRoleGranted)
				if err := _Validators.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Validators *ValidatorsFilterer) ParseRoleGranted(log types.Log) (*ValidatorsRoleGranted, error) {
	event := new(ValidatorsRoleGranted)
	if err := _Validators.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Validators contract.
type ValidatorsRoleRevokedIterator struct {
	Event *ValidatorsRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsRoleRevoked represents a RoleRevoked event raised by the Validators contract.
type ValidatorsRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Validators *ValidatorsFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ValidatorsRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsRoleRevokedIterator{contract: _Validators.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Validators *ValidatorsFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ValidatorsRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsRoleRevoked)
				if err := _Validators.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Validators *ValidatorsFilterer) ParseRoleRevoked(log types.Log) (*ValidatorsRoleRevoked, error) {
	event := new(ValidatorsRoleRevoked)
	if err := _Validators.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsSetFeeSetLockingDurationIterator is returned from FilterSetFeeSetLockingDuration and is used to iterate over the raw logs and unpacked data for SetFeeSetLockingDuration events raised by the Validators contract.
type ValidatorsSetFeeSetLockingDurationIterator struct {
	Event *ValidatorsSetFeeSetLockingDuration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsSetFeeSetLockingDurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsSetFeeSetLockingDuration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsSetFeeSetLockingDuration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsSetFeeSetLockingDurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsSetFeeSetLockingDurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsSetFeeSetLockingDuration represents a SetFeeSetLockingDuration event raised by the Validators contract.
type ValidatorsSetFeeSetLockingDuration struct {
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetFeeSetLockingDuration is a free log retrieval operation binding the contract event 0x836d4870c480e504f7e59c68ff0ab6582d14af4475c2a66c16a8f2a0b6c3fd9b.
//
// Solidity: event SetFeeSetLockingDuration(uint256 duration)
func (_Validators *ValidatorsFilterer) FilterSetFeeSetLockingDuration(opts *bind.FilterOpts) (*ValidatorsSetFeeSetLockingDurationIterator, error) {

	logs, sub, err := _Validators.contract.FilterLogs(opts, "SetFeeSetLockingDuration")
	if err != nil {
		return nil, err
	}
	return &ValidatorsSetFeeSetLockingDurationIterator{contract: _Validators.contract, event: "SetFeeSetLockingDuration", logs: logs, sub: sub}, nil
}

// WatchSetFeeSetLockingDuration is a free log subscription operation binding the contract event 0x836d4870c480e504f7e59c68ff0ab6582d14af4475c2a66c16a8f2a0b6c3fd9b.
//
// Solidity: event SetFeeSetLockingDuration(uint256 duration)
func (_Validators *ValidatorsFilterer) WatchSetFeeSetLockingDuration(opts *bind.WatchOpts, sink chan<- *ValidatorsSetFeeSetLockingDuration) (event.Subscription, error) {

	logs, sub, err := _Validators.contract.WatchLogs(opts, "SetFeeSetLockingDuration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsSetFeeSetLockingDuration)
				if err := _Validators.contract.UnpackLog(event, "SetFeeSetLockingDuration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetFeeSetLockingDuration is a log parse operation binding the contract event 0x836d4870c480e504f7e59c68ff0ab6582d14af4475c2a66c16a8f2a0b6c3fd9b.
//
// Solidity: event SetFeeSetLockingDuration(uint256 duration)
func (_Validators *ValidatorsFilterer) ParseSetFeeSetLockingDuration(log types.Log) (*ValidatorsSetFeeSetLockingDuration, error) {
	event := new(ValidatorsSetFeeSetLockingDuration)
	if err := _Validators.contract.UnpackLog(event, "SetFeeSetLockingDuration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsSetFeeSharesIterator is returned from FilterSetFeeShares and is used to iterate over the raw logs and unpacked data for SetFeeShares events raised by the Validators contract.
type ValidatorsSetFeeSharesIterator struct {
	Event *ValidatorsSetFeeShares // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsSetFeeSharesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsSetFeeShares)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsSetFeeShares)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsSetFeeSharesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsSetFeeSharesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsSetFeeShares represents a SetFeeShares event raised by the Validators contract.
type ValidatorsSetFeeShares struct {
	User   common.Address
	Pid    *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetFeeShares is a free log retrieval operation binding the contract event 0xe296be2204f5c4eb666245e667a1cc867025f5ad6517722077a0a7662fdc35ae.
//
// Solidity: event SetFeeShares(address indexed user, uint256 indexed pid, uint256 shares)
func (_Validators *ValidatorsFilterer) FilterSetFeeShares(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*ValidatorsSetFeeSharesIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "SetFeeShares", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsSetFeeSharesIterator{contract: _Validators.contract, event: "SetFeeShares", logs: logs, sub: sub}, nil
}

// WatchSetFeeShares is a free log subscription operation binding the contract event 0xe296be2204f5c4eb666245e667a1cc867025f5ad6517722077a0a7662fdc35ae.
//
// Solidity: event SetFeeShares(address indexed user, uint256 indexed pid, uint256 shares)
func (_Validators *ValidatorsFilterer) WatchSetFeeShares(opts *bind.WatchOpts, sink chan<- *ValidatorsSetFeeShares, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "SetFeeShares", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsSetFeeShares)
				if err := _Validators.contract.UnpackLog(event, "SetFeeShares", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetFeeShares is a log parse operation binding the contract event 0xe296be2204f5c4eb666245e667a1cc867025f5ad6517722077a0a7662fdc35ae.
//
// Solidity: event SetFeeShares(address indexed user, uint256 indexed pid, uint256 shares)
func (_Validators *ValidatorsFilterer) ParseSetFeeShares(log types.Log) (*ValidatorsSetFeeShares, error) {
	event := new(ValidatorsSetFeeShares)
	if err := _Validators.contract.UnpackLog(event, "SetFeeShares", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsSetMaxPunishmentBallotsIterator is returned from FilterSetMaxPunishmentBallots and is used to iterate over the raw logs and unpacked data for SetMaxPunishmentBallots events raised by the Validators contract.
type ValidatorsSetMaxPunishmentBallotsIterator struct {
	Event *ValidatorsSetMaxPunishmentBallots // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsSetMaxPunishmentBallotsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsSetMaxPunishmentBallots)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsSetMaxPunishmentBallots)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsSetMaxPunishmentBallotsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsSetMaxPunishmentBallotsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsSetMaxPunishmentBallots represents a SetMaxPunishmentBallots event raised by the Validators contract.
type ValidatorsSetMaxPunishmentBallots struct {
	Max *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetMaxPunishmentBallots is a free log retrieval operation binding the contract event 0xd045295f12ef4f160f8803b90edaa025c769d65531eab370dc395fc297ac5ac7.
//
// Solidity: event SetMaxPunishmentBallots(uint256 max)
func (_Validators *ValidatorsFilterer) FilterSetMaxPunishmentBallots(opts *bind.FilterOpts) (*ValidatorsSetMaxPunishmentBallotsIterator, error) {

	logs, sub, err := _Validators.contract.FilterLogs(opts, "SetMaxPunishmentBallots")
	if err != nil {
		return nil, err
	}
	return &ValidatorsSetMaxPunishmentBallotsIterator{contract: _Validators.contract, event: "SetMaxPunishmentBallots", logs: logs, sub: sub}, nil
}

// WatchSetMaxPunishmentBallots is a free log subscription operation binding the contract event 0xd045295f12ef4f160f8803b90edaa025c769d65531eab370dc395fc297ac5ac7.
//
// Solidity: event SetMaxPunishmentBallots(uint256 max)
func (_Validators *ValidatorsFilterer) WatchSetMaxPunishmentBallots(opts *bind.WatchOpts, sink chan<- *ValidatorsSetMaxPunishmentBallots) (event.Subscription, error) {

	logs, sub, err := _Validators.contract.WatchLogs(opts, "SetMaxPunishmentBallots")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsSetMaxPunishmentBallots)
				if err := _Validators.contract.UnpackLog(event, "SetMaxPunishmentBallots", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetMaxPunishmentBallots is a log parse operation binding the contract event 0xd045295f12ef4f160f8803b90edaa025c769d65531eab370dc395fc297ac5ac7.
//
// Solidity: event SetMaxPunishmentBallots(uint256 max)
func (_Validators *ValidatorsFilterer) ParseSetMaxPunishmentBallots(log types.Log) (*ValidatorsSetMaxPunishmentBallots, error) {
	event := new(ValidatorsSetMaxPunishmentBallots)
	if err := _Validators.contract.UnpackLog(event, "SetMaxPunishmentBallots", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsSetMinSelfBallotsIterator is returned from FilterSetMinSelfBallots and is used to iterate over the raw logs and unpacked data for SetMinSelfBallots events raised by the Validators contract.
type ValidatorsSetMinSelfBallotsIterator struct {
	Event *ValidatorsSetMinSelfBallots // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsSetMinSelfBallotsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsSetMinSelfBallots)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsSetMinSelfBallots)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsSetMinSelfBallotsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsSetMinSelfBallotsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsSetMinSelfBallots represents a SetMinSelfBallots event raised by the Validators contract.
type ValidatorsSetMinSelfBallots struct {
	Min *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetMinSelfBallots is a free log retrieval operation binding the contract event 0xaa013f7784cb6a7e4b908ae680fb75d44d826bcec005582a91d91cc8eba9c5f8.
//
// Solidity: event SetMinSelfBallots(uint256 min)
func (_Validators *ValidatorsFilterer) FilterSetMinSelfBallots(opts *bind.FilterOpts) (*ValidatorsSetMinSelfBallotsIterator, error) {

	logs, sub, err := _Validators.contract.FilterLogs(opts, "SetMinSelfBallots")
	if err != nil {
		return nil, err
	}
	return &ValidatorsSetMinSelfBallotsIterator{contract: _Validators.contract, event: "SetMinSelfBallots", logs: logs, sub: sub}, nil
}

// WatchSetMinSelfBallots is a free log subscription operation binding the contract event 0xaa013f7784cb6a7e4b908ae680fb75d44d826bcec005582a91d91cc8eba9c5f8.
//
// Solidity: event SetMinSelfBallots(uint256 min)
func (_Validators *ValidatorsFilterer) WatchSetMinSelfBallots(opts *bind.WatchOpts, sink chan<- *ValidatorsSetMinSelfBallots) (event.Subscription, error) {

	logs, sub, err := _Validators.contract.WatchLogs(opts, "SetMinSelfBallots")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsSetMinSelfBallots)
				if err := _Validators.contract.UnpackLog(event, "SetMinSelfBallots", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetMinSelfBallots is a log parse operation binding the contract event 0xaa013f7784cb6a7e4b908ae680fb75d44d826bcec005582a91d91cc8eba9c5f8.
//
// Solidity: event SetMinSelfBallots(uint256 min)
func (_Validators *ValidatorsFilterer) ParseSetMinSelfBallots(log types.Log) (*ValidatorsSetMinSelfBallots, error) {
	event := new(ValidatorsSetMinSelfBallots)
	if err := _Validators.contract.UnpackLog(event, "SetMinSelfBallots", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsSetPoolStatusIterator is returned from FilterSetPoolStatus and is used to iterate over the raw logs and unpacked data for SetPoolStatus events raised by the Validators contract.
type ValidatorsSetPoolStatusIterator struct {
	Event *ValidatorsSetPoolStatus // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsSetPoolStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsSetPoolStatus)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsSetPoolStatus)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsSetPoolStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsSetPoolStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsSetPoolStatus represents a SetPoolStatus event raised by the Validators contract.
type ValidatorsSetPoolStatus struct {
	Validator common.Address
	Enabled   bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetPoolStatus is a free log retrieval operation binding the contract event 0x2f62a1a1d6859de456dd8107c09e738da088ceb9f87d90c2e31f91fb3e89b727.
//
// Solidity: event SetPoolStatus(address indexed validator, bool enabled)
func (_Validators *ValidatorsFilterer) FilterSetPoolStatus(opts *bind.FilterOpts, validator []common.Address) (*ValidatorsSetPoolStatusIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "SetPoolStatus", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsSetPoolStatusIterator{contract: _Validators.contract, event: "SetPoolStatus", logs: logs, sub: sub}, nil
}

// WatchSetPoolStatus is a free log subscription operation binding the contract event 0x2f62a1a1d6859de456dd8107c09e738da088ceb9f87d90c2e31f91fb3e89b727.
//
// Solidity: event SetPoolStatus(address indexed validator, bool enabled)
func (_Validators *ValidatorsFilterer) WatchSetPoolStatus(opts *bind.WatchOpts, sink chan<- *ValidatorsSetPoolStatus, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "SetPoolStatus", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsSetPoolStatus)
				if err := _Validators.contract.UnpackLog(event, "SetPoolStatus", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetPoolStatus is a log parse operation binding the contract event 0x2f62a1a1d6859de456dd8107c09e738da088ceb9f87d90c2e31f91fb3e89b727.
//
// Solidity: event SetPoolStatus(address indexed validator, bool enabled)
func (_Validators *ValidatorsFilterer) ParseSetPoolStatus(log types.Log) (*ValidatorsSetPoolStatus, error) {
	event := new(ValidatorsSetPoolStatus)
	if err := _Validators.contract.UnpackLog(event, "SetPoolStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsSetRevokeLockingDurationIterator is returned from FilterSetRevokeLockingDuration and is used to iterate over the raw logs and unpacked data for SetRevokeLockingDuration events raised by the Validators contract.
type ValidatorsSetRevokeLockingDurationIterator struct {
	Event *ValidatorsSetRevokeLockingDuration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsSetRevokeLockingDurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsSetRevokeLockingDuration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsSetRevokeLockingDuration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsSetRevokeLockingDurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsSetRevokeLockingDurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsSetRevokeLockingDuration represents a SetRevokeLockingDuration event raised by the Validators contract.
type ValidatorsSetRevokeLockingDuration struct {
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetRevokeLockingDuration is a free log retrieval operation binding the contract event 0xb1af0e731cd364da26a491d6c9e37100decbf030b0f444b8eb3601dfec2ade3a.
//
// Solidity: event SetRevokeLockingDuration(uint256 duration)
func (_Validators *ValidatorsFilterer) FilterSetRevokeLockingDuration(opts *bind.FilterOpts) (*ValidatorsSetRevokeLockingDurationIterator, error) {

	logs, sub, err := _Validators.contract.FilterLogs(opts, "SetRevokeLockingDuration")
	if err != nil {
		return nil, err
	}
	return &ValidatorsSetRevokeLockingDurationIterator{contract: _Validators.contract, event: "SetRevokeLockingDuration", logs: logs, sub: sub}, nil
}

// WatchSetRevokeLockingDuration is a free log subscription operation binding the contract event 0xb1af0e731cd364da26a491d6c9e37100decbf030b0f444b8eb3601dfec2ade3a.
//
// Solidity: event SetRevokeLockingDuration(uint256 duration)
func (_Validators *ValidatorsFilterer) WatchSetRevokeLockingDuration(opts *bind.WatchOpts, sink chan<- *ValidatorsSetRevokeLockingDuration) (event.Subscription, error) {

	logs, sub, err := _Validators.contract.WatchLogs(opts, "SetRevokeLockingDuration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsSetRevokeLockingDuration)
				if err := _Validators.contract.UnpackLog(event, "SetRevokeLockingDuration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetRevokeLockingDuration is a log parse operation binding the contract event 0xb1af0e731cd364da26a491d6c9e37100decbf030b0f444b8eb3601dfec2ade3a.
//
// Solidity: event SetRevokeLockingDuration(uint256 duration)
func (_Validators *ValidatorsFilterer) ParseSetRevokeLockingDuration(log types.Log) (*ValidatorsSetRevokeLockingDuration, error) {
	event := new(ValidatorsSetRevokeLockingDuration)
	if err := _Validators.contract.UnpackLog(event, "SetRevokeLockingDuration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsValidatorClaimRewardIterator is returned from FilterValidatorClaimReward and is used to iterate over the raw logs and unpacked data for ValidatorClaimReward events raised by the Validators contract.
type ValidatorsValidatorClaimRewardIterator struct {
	Event *ValidatorsValidatorClaimReward // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsValidatorClaimRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorClaimReward)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsValidatorClaimReward)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsValidatorClaimRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorClaimRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorClaimReward represents a ValidatorClaimReward event raised by the Validators contract.
type ValidatorsValidatorClaimReward struct {
	Validator     common.Address
	PendingReward *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterValidatorClaimReward is a free log retrieval operation binding the contract event 0xef89e1a1006e92f5c35aeca19e2099abded5e99c36b1f1dcdd7131398794c325.
//
// Solidity: event ValidatorClaimReward(address indexed validator, uint256 pendingReward)
func (_Validators *ValidatorsFilterer) FilterValidatorClaimReward(opts *bind.FilterOpts, validator []common.Address) (*ValidatorsValidatorClaimRewardIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorClaimReward", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorClaimRewardIterator{contract: _Validators.contract, event: "ValidatorClaimReward", logs: logs, sub: sub}, nil
}

// WatchValidatorClaimReward is a free log subscription operation binding the contract event 0xef89e1a1006e92f5c35aeca19e2099abded5e99c36b1f1dcdd7131398794c325.
//
// Solidity: event ValidatorClaimReward(address indexed validator, uint256 pendingReward)
func (_Validators *ValidatorsFilterer) WatchValidatorClaimReward(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorClaimReward, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorClaimReward", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorClaimReward)
				if err := _Validators.contract.UnpackLog(event, "ValidatorClaimReward", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorClaimReward is a log parse operation binding the contract event 0xef89e1a1006e92f5c35aeca19e2099abded5e99c36b1f1dcdd7131398794c325.
//
// Solidity: event ValidatorClaimReward(address indexed validator, uint256 pendingReward)
func (_Validators *ValidatorsFilterer) ParseValidatorClaimReward(log types.Log) (*ValidatorsValidatorClaimReward, error) {
	event := new(ValidatorsValidatorClaimReward)
	if err := _Validators.contract.UnpackLog(event, "ValidatorClaimReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsVoteIterator is returned from FilterVote and is used to iterate over the raw logs and unpacked data for Vote events raised by the Validators contract.
type ValidatorsVoteIterator struct {
	Event *ValidatorsVote // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsVote)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsVote)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsVote represents a Vote event raised by the Validators contract.
type ValidatorsVote struct {
	User      common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVote is a free log retrieval operation binding the contract event 0x66a9138482c99e9baf08860110ef332cc0c23b4a199a53593d8db0fc8f96fbfc.
//
// Solidity: event Vote(address indexed user, address indexed validator, uint256 amount)
func (_Validators *ValidatorsFilterer) FilterVote(opts *bind.FilterOpts, user []common.Address, validator []common.Address) (*ValidatorsVoteIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "Vote", userRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsVoteIterator{contract: _Validators.contract, event: "Vote", logs: logs, sub: sub}, nil
}

// WatchVote is a free log subscription operation binding the contract event 0x66a9138482c99e9baf08860110ef332cc0c23b4a199a53593d8db0fc8f96fbfc.
//
// Solidity: event Vote(address indexed user, address indexed validator, uint256 amount)
func (_Validators *ValidatorsFilterer) WatchVote(opts *bind.WatchOpts, sink chan<- *ValidatorsVote, user []common.Address, validator []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "Vote", userRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsVote)
				if err := _Validators.contract.UnpackLog(event, "Vote", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVote is a log parse operation binding the contract event 0x66a9138482c99e9baf08860110ef332cc0c23b4a199a53593d8db0fc8f96fbfc.
//
// Solidity: event Vote(address indexed user, address indexed validator, uint256 amount)
func (_Validators *ValidatorsFilterer) ParseVote(log types.Log) (*ValidatorsVote, error) {
	event := new(ValidatorsVote)
	if err := _Validators.contract.UnpackLog(event, "Vote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Validators contract.
type ValidatorsWithdrawIterator struct {
	Event *ValidatorsWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsWithdraw represents a Withdraw event raised by the Validators contract.
type ValidatorsWithdraw struct {
	User      common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed user, address indexed validator, uint256 amount)
func (_Validators *ValidatorsFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, validator []common.Address) (*ValidatorsWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "Withdraw", userRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsWithdrawIterator{contract: _Validators.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed user, address indexed validator, uint256 amount)
func (_Validators *ValidatorsFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ValidatorsWithdraw, user []common.Address, validator []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "Withdraw", userRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsWithdraw)
				if err := _Validators.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed user, address indexed validator, uint256 amount)
func (_Validators *ValidatorsFilterer) ParseWithdraw(log types.Log) (*ValidatorsWithdraw, error) {
	event := new(ValidatorsWithdraw)
	if err := _Validators.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

