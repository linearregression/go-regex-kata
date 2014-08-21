package regexkata

import "testing"
import "regexp"
import "unicode/utf8"

//Match pattern against byte array
func TestMatch(t *testing.T) {
	pattern := "^(B|b)rian$"
	if match, err := regexp.Match(pattern, []byte("Brian")); match != true {
		t.Errorf("Brian did not match %q %v", pattern, err)
	}

	if match, err := regexp.Match(pattern, []byte("brian")); match != true {
		t.Errorf("brian did not match %q %v", pattern, err)
	}
}

//MatchReader against a rune reader
type TestRuneReader struct{}
type RuneError struct{}

func (e *RuneError) Error() string {
	return "Rune error!!!"
}
func (reader *TestRuneReader) ReadRune() (r rune, size int, e error) {
	r = '\a'
	size = utf8.RuneLen(r)
	if r == '\v' {
		return 0, 0, &RuneError{}
	}
	return
}
func TestMatchReader(t *testing.T) {
	pattern, reader := "\a", &TestRuneReader{}
	run := '\a'
	reader.ReadRune()

	if match, err := regexp.MatchReader(pattern, reader); match != true {
		t.Errorf("MatchReader did not match %v %v", run, err)
	}
}

//MatchString against a string literal
func TestMatchString(t *testing.T) {
	pattern, upper, lower := "^(B|b)rian$", "Brian", "brian"

	if match, err := regexp.MatchString(pattern, upper); match != true {
		t.Errorf("MatchString did not match %q %v", upper, err)
	}

	if match, err := regexp.MatchString(pattern, lower); match != true {
		t.Errorf("MatchString did not match %q %v", lower, err)
	}
}

func TestMatchStringWithQuoteMeta(t *testing.T) {
	pattern, str := regexp.QuoteMeta("[foo]"), "[foo]"

	if match, err := regexp.MatchString(pattern, str); match != true {
		t.Errorf("MatchString did not match %q %v", str, err)
	}
}
