package gcfg

import (
	"fmt"
	"math/big"
	"strings"
	"testing"
)

type Config1 struct {
	Section struct {
		Int    int
		BigInt big.Int
	}
}

var testsGoogleCodeIssue1 = []struct {
	cfg      string
	typename string
}{
	{"[section]\nint=X", "int"},
	{"[section]\nint=", "int"},
	{"[section]\nint=1A", "int"},
	{"[section]\nbigint=X", "big.Int"},
	{"[section]\nbigint=", "big.Int"},
	{"[section]\nbigint=1A", "big.Int"},
}

// Value parse error should:
//  - include plain type name
//  - not include reflect internals
func TestGoogleCodeIssue1(t *testing.T) {
	for i, tt := range testsGoogleCodeIssue1 {
		var c Config1
		err := ReadStringInto(&c, tt.cfg)
		switch {
		case err == nil:
			t.Errorf("%d fail: got ok; wanted error", i)
		case !strings.Contains(err.Error(), tt.typename):
			t.Errorf("%d fail: error message doesn't contain type name %q: %v",
				i, tt.typename, err)
		case strings.Contains(err.Error(), "reflect"):
			t.Errorf("%d fail: error message includes reflect internals: %v",
				i, err)
		default:
			t.Logf("%d pass: %v", i, err)
		}
	}
}

type confGoogleCodeIssue2 struct{ Main struct{ Foo string } }

var testsGoogleCodeIssue2 = []readtest{
	{"[main]\n;\nfoo = bar\n", &confGoogleCodeIssue2{struct{ Foo string }{"bar"}}, true},
	{"[main]\r\n;\r\nfoo = bar\r\n", &confGoogleCodeIssue2{struct{ Foo string }{"bar"}}, true},
}

func TestGoogleCodeIssue2(t *testing.T) {
	for i, tt := range testsGoogleCodeIssue2 {
		id := fmt.Sprintf("issue2:%d", i)
		testRead(t, id, tt)
	}
}

type ConfigIssue11 struct {
	Sect struct {
		Var bool
	}
}

// Escaped double quote should be supported in "raw" string literals
func TestIssue12(t *testing.T) {
	var c struct {
		Section struct {
			Name string
		}
	}
	err := ReadFileInto(&c, "testdata/issue12.gcfg")
	if err != nil {
		t.Fatalf("fail: want ok, got error %v", err)
	}
	if c.Section.Name != `"value"` {
		t.Errorf("fail: want `\"value\"`, got %q", c.Section.Name)
	}
}
