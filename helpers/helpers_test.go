package helpers

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"testing"
)

func Test_CaptureOutput(t *testing.T) {
	expected := "hello world"
	out := CaptureOutput(func() {
		fmt.Print(expected)
	})
	if out != expected {
		t.Fatalf("fmt.Print must be captured, got %s but want %s", out, expected)
	}

	out = CaptureOutput(func() {
		log.Print(expected)
	})
	matched, err := regexp.MatchString(expected, out)
	if matched != true || err != nil {
		t.Fatalf("log.Print must be captured, got %s but want %s", out, expected)
	}
}

func Test_MockHits(t *testing.T) {
	expected := map[string]map[string]int64{}
	if got := reflect.DeepEqual(expected, members); !got {
		t.Fatalf("%#v is not equal to expected (%#v)", members, expected)
	}

	myHits := NewMockHits("myhits")

	myHits.Add("myhit", 1)
	expected = map[string]map[string]int64{
		"myhits": {"myhit": 1},
	}
	if got := reflect.DeepEqual(expected, members); !got {
		t.Fatalf("%#v is not equal to expected (%#v)", members, expected)
	}

	myHits.Add("myhit2", 1)
	expected = map[string]map[string]int64{
		"myhits": {"myhit": 1, "myhit2": 1},
	}
	if got := reflect.DeepEqual(expected, members); !got {
		t.Fatalf("%#v is not equal to expected (%#v)", members, expected)
	}

	myHits.Add("myhit", 1)
	expected = map[string]map[string]int64{
		"myhits": {"myhit": 2, "myhit2": 1},
	}
	if got := reflect.DeepEqual(expected, members); !got {
		t.Fatalf("%#v is not equal to expected (%#v)", members, expected)
	}

	top, count, err := myHits.Top()
	if got, want := top, "myhit"; got != want {
		t.Fatalf("%s is not equal to expected (%s)", got, want)
	}
	if got, want := count, int64(2); got != want {
		t.Fatalf("%d is not equal to expected (%d)", got, want)
	}
	if err != nil {
		t.Fatalf("non error expected (%s)", err.Error())
	}

	myHits.Reset()
	expected = map[string]map[string]int64{"myhits": {}}
	if got := reflect.DeepEqual(expected, members); !got {
		t.Fatalf("%#v is not equal to expected (%#v)", members, expected)
	}
}
