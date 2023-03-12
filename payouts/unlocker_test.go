package payouts

import (
	"math/big"
	"os"
	"testing"

	"github.com/galiy/open-ubiq-pool/rpc"
	"github.com/galiy/open-ubiq-pool/storage"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestCalculateRewards(t *testing.T) {
	blockReward, _ := new(big.Rat).SetString("5000000000000000000")
	shares := map[string]int64{"0x0": 1000000, "0x1": 20000, "0x2": 5000, "0x3": 10, "0x4": 1}
	expectedRewards := map[string]int64{"0x0": 4877996431, "0x1": 97559929, "0x2": 24389982, "0x3": 48780, "0x4": 4878}
	totalShares := int64(1025011)

	rewards := calculateRewardsForShares(shares, totalShares, blockReward)
	expectedTotalAmount := int64(5000000000)

	totalAmount := int64(0)
	for login, amount := range rewards {
		totalAmount += amount

		if expectedRewards[login] != amount {
			t.Errorf("Amount for %v must be equal to %v vs %v", login, expectedRewards[login], amount)
		}
	}
	if totalAmount != expectedTotalAmount {
		t.Errorf("Total reward must be equal to block reward in Shannon: %v vs %v", expectedTotalAmount, totalAmount)
	}
}

func TestChargeFee(t *testing.T) {
	orig, _ := new(big.Rat).SetString("5000000000000000000")
	value, _ := new(big.Rat).SetString("5000000000000000000")
	expectedNewValue, _ := new(big.Rat).SetString("3750000000000000000")
	expectedFee, _ := new(big.Rat).SetString("1250000000000000000")
	newValue, fee := chargeFee(orig, 25.0)

	if orig.Cmp(value) != 0 {
		t.Error("Must not change original value")
	}
	if newValue.Cmp(expectedNewValue) != 0 {
		t.Error("Must charge and deduct correct fee")
	}
	if fee.Cmp(expectedFee) != 0 {
		t.Error("Must charge fee")
	}
}

func TestWeiToShannonInt64(t *testing.T) {
	wei, _ := new(big.Rat).SetString("1000000000000000000")
	origWei, _ := new(big.Rat).SetString("1000000000000000000")
	shannon := int64(1000000000)

	if weiToShannonInt64(wei) != shannon {
		t.Error("Must convert to Shannon")
	}
	if wei.Cmp(origWei) != 0 {
		t.Error("Must charge original value")
	}
}

func TestGetUncleReward(t *testing.T) {
	// Year 4 - Pre-Orion
	rewardsPreOrion := make(map[int64]string)
	expectedRewardsPreOrion := map[int64]string{
		1774000: "2000000000000000000",
		1774001: "0", //previous blocks not rewarded
		1774002: "0",
		1774003: "0",
		1774004: "0",
		1774005: "0",
	}
	for i := int64(1774000); i < 1774000+6; i++ {
		rewardsPreOrion[i] = getUncleReward(1774000, i+1).String()
	}
	for i, reward := range rewardsPreOrion {
		if expectedRewardsPreOrion[i] != rewardsPreOrion[i] {
			t.Errorf("Incorrect uncle reward for %v, expected %v vs %v", i, expectedRewardsPreOrion[i], reward)
		}
	}

	// Orion
	rewardsOrion := make(map[int64]string)
	expectedRewardsOrion := map[int64]string{
		1791793: "750000000000000000",
		1791794: "0", //previous blocks not rewarded
		1791795: "0",
		1791796: "0",
		1791797: "0",
		1791798: "0",
	}
	for i := int64(1791793); i < 1791793+6; i++ {
		rewardsOrion[i] = getUncleReward(1791793, i+1).String()
	}
	for i, reward := range rewardsOrion {
		if expectedRewardsOrion[i] != rewardsOrion[i] {
			t.Errorf("Incorrect uncle reward for %v, expected %v vs %v", i, expectedRewardsOrion[i], reward)
		}
	}

	// Year 4 - Pre-Orion - Simple
	expectedRewardsPreOrionSimple := "2000000000000000000"
	rewardsPreOrionSimple := getUncleReward(1433454, 1433454+1).String()
	if expectedRewardsPreOrionSimple != rewardsPreOrionSimple {
		t.Errorf("Incorrect uncle reward, expected %v vs %v", expectedRewardsPreOrionSimple, rewardsPreOrionSimple)
	}

	// Year 4 - Orion - Simple
	expectedRewardsOrionSimple := "750000000000000000"
	rewardsOrionSimple := getUncleReward(2000000, 2000000+1).String()
	if expectedRewardsOrionSimple != rewardsOrionSimple {
		t.Errorf("Incorrect uncle reward, expected %v vs %v", expectedRewardsOrionSimple, rewardsOrionSimple)
	}
}

func TestMatchCandidate(t *testing.T) {
	gethBlock := &rpc.GetBlockReply{Hash: "0x12345A", Nonce: "0x1A"}
	parityBlock := &rpc.GetBlockReply{Hash: "0x12345A", SealFields: []string{"0x0A", "0x1A"}}
	candidate := &storage.BlockData{Nonce: "0x1a"}
	orphan := &storage.BlockData{Nonce: "0x1abc"}

	if !matchCandidate(gethBlock, candidate) {
		t.Error("Must match with nonce")
	}
	if !matchCandidate(parityBlock, candidate) {
		t.Error("Must match with seal fields")
	}
	if matchCandidate(gethBlock, orphan) {
		t.Error("Must not match with orphan with nonce")
	}
	if matchCandidate(parityBlock, orphan) {
		t.Error("Must not match orphan with seal fields")
	}

	block := &rpc.GetBlockReply{Hash: "0x12345A"}
	immature := &storage.BlockData{Hash: "0x12345a", Nonce: "0x0"}
	if !matchCandidate(block, immature) {
		t.Error("Must match with hash")
	}
}
