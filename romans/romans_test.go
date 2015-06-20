package romans

import "testing"

func TestFromArabicString(t *testing.T) {
	nums := []string{"1", "2", "5", "9", "10", "19", "20", "1999", "1000000", "10000000"}

	expected := []string{"I", "II", "V", "IX", "X", "XIX", "XX", "MCMXCIX", "M̅", "M̅M̅M̅M̅M̅M̅M̅M̅M̅M̅"}

	for i, num := range nums {
		res, err := FromArabicString(num)
		if err != nil {
			t.Error(err)
		}
		if res != expected[i] {
			t.Errorf("Expected %s, got %s", expected[i], res)
		}

	}
}

func TestToArabicString(t *testing.T) {
	romans := []string{"I", "MCMXCIX", "XXV", "VL", "M̅"}
	expected := []string{"1", "1999", "25", "45", "1000000"}

	for i, roman := range romans {
		res, err := ToArabicString(roman)
		if err != nil {
			t.Error(err)
		}
		if res != expected[i] {
			t.Errorf("Expected %s, got %s", expected[i], res)
		}

	}
}
