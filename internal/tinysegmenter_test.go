package internal

import (
	_ "embed"
	"reflect"
	"testing"
)

//go:embed testdata/timemachineu8j.txt
var sampletext string

func TestSegment(t *testing.T) {
	ary := Segment("私の名前は中野です")
	expect := []string{"私", "の", "名前", "は", "中野", "です"}
	if !reflect.DeepEqual(ary, expect) {
		t.Errorf("got %+v, expected %v", ary, expect)
	}
}

func BenchmarkSegment(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Segment("私の名前は中野です")
	}
}

func BenchmarkSegmentLargeText(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Segment(sampletext)
	}
}

func TestSegmentYear(t *testing.T) {
	ary := Segment("2025年")
	expect := []string{"2025", "年"}
	if !reflect.DeepEqual(ary, expect) {
		t.Errorf("got %+v, expected %v", ary, expect)
	}
}

func TestSegmentNumbers(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect []string
	}{
		{"year", "2025年", []string{"2025", "年"}},
		{"date", "2025年12月31日", []string{"2025", "年", "12月", "31", "日"}},
		{"numbers in sentence", "私は2025年に生まれました", []string{"私", "は", "2025", "年", "に", "生まれ", "まし", "た"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ary := Segment(tt.input)
			if !reflect.DeepEqual(ary, tt.expect) {
				t.Errorf("got %+v, expected %v", ary, tt.expect)
			}
		})
	}
}
