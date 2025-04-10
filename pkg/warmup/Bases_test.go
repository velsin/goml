package warmup

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

func TestHexToDecimal(t *testing.T) {
	testCases := []struct {
		in  string
		exp string
	}{
		{"15", "21"},
		{"F1", "241"},
		{"7FFFFFFF", "2147483647"},
		{"0", "0"},
	}

	for _, testCase := range testCases {
		res := HexToDecimal(testCase.in)
		if testCase.exp != res {
			t.Errorf("Received %v for %v, expected %v", res, testCase.in, testCase.exp)
		}
	}
}

func TestDecimalToHex(t *testing.T) {
	testCases := []struct {
		in  string
		exp string
	}{
		{"21", "15"},
		{"241", "F1"},
		{"2147483647", "7FFFFFFF"},
		{"0", "0"},
	}

	for _, testCase := range testCases {
		res := DecimalToHex(testCase.in)
		if testCase.exp != res {
			t.Errorf("Received %v for %v, expected %v", res, testCase.in, testCase.exp)
		}
	}
}

func TestBinaryToHex(t *testing.T) {
	testCases := []struct {
		in  string
		exp string
	}{
		{"11010", "1A"},
		{"100", "4"},
		{"0", "0"},
		{"11110001", "F1"},
		{"1111111111111111111111111111111", "7FFFFFFF"},
	}

	for _, testCase := range testCases {
		res := BinaryToHex(testCase.in)
		if testCase.exp != res {
			t.Errorf("Received %v for %v, expected %v", res, testCase.in, testCase.exp)
		}
	}

}

func TestHexToBinary(t *testing.T) {
	testCases := []struct {
		in  string
		exp string
	}{
		{"1A", "11010"},
		{"4", "100"},
		{"0", "0"},
		{"F1", "11110001"},
		{"7FFFFFFF", "1111111111111111111111111111111"},
	}

	for _, testCase := range testCases {
		res := HexToBinary(testCase.in)
		if testCase.exp != res {
			t.Errorf("Received %v for %v, expected %v", res, testCase.in, testCase.exp)
		}
	}

}
