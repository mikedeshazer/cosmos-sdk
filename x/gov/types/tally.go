package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ValidatorGovInfo used for tallying
type ValidatorGovInfo struct {
	Address             sdk.ValAddress // address of the validator operator
	BondedTokens        sdk.Int        // Power of a Validator
	DelegatorShares     sdk.Dec        // Total outstanding delegator shares
	DelegatorDeductions sdk.Dec        // Delegator deductions from validator's delegators voting independently
	Vote                VoteOption     // Vote of the validator
}

// NewValidatorGovInfo creates a ValidatorGovInfo instance
func NewValidatorGovInfo(address sdk.ValAddress, bondedTokens sdk.Int, delegatorShares,
	delegatorDeductions sdk.Dec, vote VoteOption) ValidatorGovInfo {

	return ValidatorGovInfo{
		Address:             address,
		BondedTokens:        bondedTokens,
		DelegatorShares:     delegatorShares,
		DelegatorDeductions: delegatorDeductions,
		Vote:                vote,
	}
}

type TallyResult struct {
	Yes        sdk.Int `json:"yes" yaml:"yes"`
	Abstain    sdk.Int `json:"abstain" yaml:"abstain"`
	No         sdk.Int `json:"no" yaml:"no"`
	NoWithVeto sdk.Int `json:"no_with_veto" yaml:"no_with_veto"`
}

func NewTallyResult(yes, abstain, no, noWithVeto sdk.Int) TallyResult {
	return TallyResult{
		Yes:        yes,
		Abstain:    abstain,
		No:         no,
		NoWithVeto: noWithVeto,
	}
}

func NewTallyResultFromMap(results map[VoteOption]sdk.Dec) TallyResult {
	return TallyResult{
		Yes:        results[OptionYes].TruncateInt(),
		Abstain:    results[OptionAbstain].TruncateInt(),
		No:         results[OptionNo].TruncateInt(),
		NoWithVeto: results[OptionNoWithVeto].TruncateInt(),
	}
}

// EmptyTallyResult returns an empty TallyResult.
func EmptyTallyResult() TallyResult {
	return TallyResult{
		Yes:        sdk.ZeroInt(),
		Abstain:    sdk.ZeroInt(),
		No:         sdk.ZeroInt(),
		NoWithVeto: sdk.ZeroInt(),
	}
}

// Equals returns if two proposals are equal.
func (tr TallyResult) Equals(comp TallyResult) bool {
	return tr.Yes.Equal(comp.Yes) &&
		tr.Abstain.Equal(comp.Abstain) &&
		tr.No.Equal(comp.No) &&
		tr.NoWithVeto.Equal(comp.NoWithVeto)
}

func (tr TallyResult) String() string {
	return fmt.Sprintf(`Tally Result:
  Yes:        %s
  Abstain:    %s
  No:         %s
  NoWithVeto: %s`, tr.Yes, tr.Abstain, tr.No, tr.NoWithVeto)
}
