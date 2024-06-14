package handler

import (
	"fmt"
	"math/rand"
	"time"
)

// 随机种子
var seeded bool = false

// 生成跟踪id 模拟
func CreateTrackingId(salt string) string {
	if !seeded {
		rand.New(rand.NewSource(time.Now().UnixNano()))
		seeded = true
	}
	return fmt.Sprintf("%c%c-%d%s-%d%s",
		getRandomLetterCode(),
		getRandomLetterCode(),
		len(salt),
		getRandomNumber(3),
		len(salt)/2,
		getRandomNumber(7),
	)
}

func getRandomLetterCode() uint32 {
	return 65 + uint32(rand.Intn(25))
}

func getRandomNumber(digits int) string {
	str := ""
	for i := 0; i < digits; i++ {
		str = fmt.Sprintf("%s%d", str, rand.Intn(10))
	}

	return str
}
