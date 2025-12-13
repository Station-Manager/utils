package utils

//func TestIsValidFrequencyMHz(t *testing.T) {
//	valid := []string{
//		"7",
//		"7.0",
//		"7.050",
//		"14.074",
//		"144.390000",
//		"014.074", // leading zeros should be acceptable as numeric string
//	}
//	for _, v := range valid {
//		if !IsValidFrequencyMHz(v) {
//			t.Fatalf("expected valid ADIF MHz frequency: %q", v)
//		}
//	}
//
//	invalid := []string{
//		"",
//		"7.",
//		"7.1234567", // too many decimals
//		"7,050",     // comma not allowed
//		"7..050",
//		"abc",
//		"-7.0",
//		"0",
//		"0.000000",
//	}
//	for _, v := range invalid {
//		if IsValidFrequencyMHz(v) {
//			t.Fatalf("expected invalid ADIF MHz frequency: %q", v)
//		}
//	}
//}
