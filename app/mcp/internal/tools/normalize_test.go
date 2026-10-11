package tools

import (
	"testing"
)

func TestAsInt(t *testing.T) {
	cases := []struct {
		in    any
		want  int
		isErr bool
	}{
		{float64(5), 5, false},
		{int(7), 7, false},
		{"12", 12, false},
		{" 12 ", 12, false},
		{"abc", 0, true},
		{true, 0, true},
	}
	for _, c := range cases {
		got, err := asInt(c.in)
		if c.isErr {
			if err == nil {
				t.Errorf("asInt(%v) 应报错", c.in)
			}
			continue
		}
		if err != nil || got != c.want {
			t.Errorf("asInt(%v) = %d, %v；期望 %d", c.in, got, err, c.want)
		}
	}
}

func TestAsIntList(t *testing.T) {
	got, err := asIntList(3)
	if err != nil || len(got) != 1 || got[0] != 3 {
		t.Errorf("asIntList(3) = %v, %v", got, err)
	}
	got, err = asIntList([]any{float64(1), "2", 3})
	if err != nil || len(got) != 3 || got[0] != 1 || got[1] != 2 || got[2] != 3 {
		t.Errorf("asIntList 混合列表 = %v, %v", got, err)
	}
	if _, err := asIntList("x"); err == nil {
		t.Error("asIntList(\"x\") 应报错")
	}
}

func TestAsStrList(t *testing.T) {
	got, _ := asStrList("net1")
	if len(got) != 1 || got[0] != "net1" {
		t.Errorf("asStrList 单值 = %v", got)
	}
	got, _ = asStrList([]any{"a", "b"})
	if len(got) != 2 || got[0] != "a" || got[1] != "b" {
		t.Errorf("asStrList 列表 = %v", got)
	}
}
