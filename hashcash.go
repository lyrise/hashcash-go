package hashcash

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"math/bits"
	"time"

	"github.com/vpxyz/xorshift/xorshift4096star"
)

func SimpleSha2_256(target []byte) {
	var xs = xorshift4096star.NewSource(time.Now().Unix())

	buffer := make([]byte, 32+len(target))
	copy(buffer[32:], target)

	var lastCost int

	lastWriteTime := time.Now()
	var loopCount int64

	for {
		loopCount++

		binary.LittleEndian.PutUint64(buffer[0:8], xs.Uint64())
		binary.LittleEndian.PutUint64(buffer[8:16], xs.Uint64())
		binary.LittleEndian.PutUint64(buffer[16:24], xs.Uint64())
		binary.LittleEndian.PutUint64(buffer[24:32], xs.Uint64())

		var hash = sha256.Sum256(buffer)

		var currentCost int

		for index := 0; index < 32; index++ {
			var zeroBits = bits.LeadingZeros8(hash[index])
			currentCost += zeroBits

			if zeroBits < 8 {
				break
			}
		}

		if lastCost < currentCost {
			fmt.Printf("%v %v\n", currentCost, base64.StdEncoding.EncodeToString(buffer[:32]))
			lastCost = currentCost
		}

		if loopCount%10000000 == 0 {
			fmt.Println(time.Since(lastWriteTime))
			lastWriteTime = time.Now()
		}
	}
}
