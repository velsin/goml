package bases

import "testing"

func TestBinaryToDecimal(t *testing.T) {
	testCases := []struct {
		in  string
		exp string
	}{
		{"11010", "26"},
		{"100", "4"},
		{"0", "0"},
		{"10011", "19"},
		{"1111111111111111111111111111111", "2147483647"},
	}

	for _, testCase := range testCases {
		res := BinaryToDecimal(testCase.in)
		if testCase.exp != res {
			t.Errorf("Received %v for %v, expected %v", res, testCase.in, testCase.exp)
		}
	}
}

func TestDecimalToBinary(t *testing.T) {
	testCases := []struct {
		in  string
		exp string
	}{
		{"26", "11010"},
		{"4", "100"},
		{"0", "0"},
		{"19", "10011"},
		{"2147483647", "1111111111111111111111111111111"},
	}

	for _, testCase := range testCases {
		res := DecimalToBinary(testCase.in)
		if testCase.exp != res {
			t.Errorf("Received %v for %v, expected %v", res, testCase.in, testCase.exp)
		}
	}
}
