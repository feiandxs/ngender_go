package ngender_go_test

import (
	"fmt"
	"testing"

	"github.com/feiandxs/ngender_go"
)

func TestGuess(t *testing.T) {
	guesser, err := ngender_go.NewGuesser()
	if err != nil {
		t.Fatalf("创建猜测器失败：%v", err)
	}

	testCases := []struct {
		Name           string
		ExpectedGender string
	}{
		{"段彦宇", "female"},
		{"尹茜", "female"},
		{"黄姣", "female"},
		{"刘丰鑫", "male"},
		{"马化腾", "male"},
		{"孙思邈", "male"},
	}

	for _, tc := range testCases {
		gender, rate := guesser.Guess(tc.Name)
		rateStr := fmt.Sprintf("%.2f%%", rate*100)
		fmt.Printf("姓名：%s，性别：%s， 概率： %s\n", tc.Name, gender, rateStr)

		if gender != tc.ExpectedGender {
			t.Errorf("姓名：%s，预期性别：%s，实际性别：%s, 概率： %s", tc.Name, tc.ExpectedGender, gender, rateStr)
		}
	}
}
