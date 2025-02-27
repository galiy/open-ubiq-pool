package util

import (
	"math/big"
	"regexp"
	"strconv"
	"time"

	"github.com/ubiq/go-ubiq/v7/common/hexutil"
	"github.com/ubiq/go-ubiq/v7/common/math"
)

var Ether = math.BigPow(10, 18)
var Shannon = math.BigPow(10, 9)

var pow256 = math.BigPow(2, 256)
var addressPattern = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
var zeroHash = regexp.MustCompile("^0?x?0+$")

func IsValidHexAddress(s string) bool {
	if IsZeroHash(s) || !addressPattern.MatchString(s) {
		return false
	}
	return true
}

func IsZeroHash(s string) bool {
	return zeroHash.MatchString(s)
}

func MakeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func GetTargetHex(diff int64) string {
	difficulty := big.NewInt(diff)
	diff1 := new(big.Int).Div(pow256, difficulty)
	return string(hexutil.Encode(diff1.Bytes()))
}

func TargetHexToDiff(targetHex string) (*big.Int, error) {
	targetBytes, err := hexutil.Decode(targetHex)
	if err != nil {
		return new(big.Int), err
	}
	return new(big.Int).Div(pow256, new(big.Int).SetBytes(targetBytes)), nil
}

func ToHex(n int64) string {
	return "0x0" + strconv.FormatInt(n, 16)
}

func FormatReward(reward *big.Int) string {
	return reward.String()
}

func FormatRatReward(reward *big.Rat) string {
	wei := new(big.Rat).SetInt(Ether)
	reward = reward.Quo(reward, wei)
	return reward.FloatString(8)
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func MustParseDuration(s string) time.Duration {
	value, err := time.ParseDuration(s)
	if err != nil {
		panic("util: Can't parse duration `" + s + "`: " + err.Error())
	}
	return value
}

func String2Big(num string) *big.Int {
	n := new(big.Int)
	n.SetString(num, 0)
	return n
}

func DiffFloatToInt(diffFloat float64) (diffInt int64) {
	diffInt = int64(diffFloat * float64(1<<48) / float64(0xffff)) // 48 = 256 - 26*8
	return
}

func DiffIntToFloat(diffInt int64) (diffFloat float64) {
	diffFloat = float64(diffInt*0xffff) / float64(1<<48) // 48 = 256 - 26*8
	return
}
