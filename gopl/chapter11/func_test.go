package chapter11

import (
	"math/rand"
	"testing"
	"time"
	"unicode"
)

func TestPalindrome(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Error(`IsPalindrome("detartrated") = false`)
	}
	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") = false`)
	}
	if !IsPalindrome("}å") {
		t.Error(`IsPalindrome("kayak") = false`)
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Error(`IsPalindrome("palindrome") = true`)
	}
}

func TestFrenchPalindrome(t *testing.T) {
	if !IsPalindrome("été") {
		t.Error(`IsPalindrome("été") = false`)
	}
}

func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		t.Errorf(`IsPalindrome(%q) = false`, input) // 避免重复
	}
}

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // non-palindrome
		{"desserts", false},   // semi-palindrome
	}

	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q)==%v", test.input, got)
		}
	}
}

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < n; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
		runes[n-i-1] = r
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		input := randomPalindrome(rng)
		if got := IsPalindrome(input); !got {
			t.Errorf("IsPalindrome(%q) == %v, want true", input, got)
		}
	}
}

func randNonPalindromes(rng *rand.Rand) string {
	n := rng.Intn(25)
	if n < 2 {
		n = 2
	}
	runes := make([]rune, n)
	for i := 0; i < n; i++ {
		v := rng.Intn(0x1000)
		for !unicode.IsLetter(rune(v)){
			v = rng.Intn(0x1000)
		}
		runes[i] = rune(v)
		v2 := rng.Intn(0x1000)
		for !unicode.IsLetter(rune(v2)) || v2 == v{
			v2 = rng.Intn(0x1000)
		}
		runes[n-i-1] = rune(v2)
	}
	return string(runes)
}

func TestRandNonPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		input := randNonPalindromes(rng)
		if got := IsPalindrome(input); got {
			t.Errorf("IsPalindrome(%q) = %v, want false", input, got)
		}
	}
}