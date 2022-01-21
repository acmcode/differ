package tool

import (
	"path/filepath"
	"testing"
)

func TestDiffOutFunc(t *testing.T) {
	path, err := filepath.Abs(".")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		desc       string
		userOut    string
		dataOut    string
		ignoreHead bool
		strictMode bool
		shouldSame bool
	}{
		{
			desc:       "(L)There is an enter key in each line",
			userOut:    "user.out",
			dataOut:    "data.out",
			ignoreHead: false,
			strictMode: false,
			shouldSame: true,
		},
		{
			desc:       "(S)There is an enter key in each line",
			userOut:    "user.out",
			dataOut:    "data.out",
			ignoreHead: false,
			strictMode: true,
			shouldSame: false,
		},
		{
			desc:       "(S)Both are the same contents",
			userOut:    "userv1.out",
			dataOut:    "datav1.out",
			ignoreHead: false,
			strictMode: true,
			shouldSame: true,
		},
		{
			desc:       "(L)Enter VS Space",
			userOut:    "userv2.out",
			dataOut:    "datav2.out",
			ignoreHead: false,
			strictMode: false,
			shouldSame: true,
		},
		{
			desc:       "(S)Enter VS Space",
			userOut:    "userv2.out",
			dataOut:    "datav2.out",
			ignoreHead: false,
			strictMode: true,
			shouldSame: false,
		},
		{
			desc:       "(L)New lines in the last",
			userOut:    "userv3.out",
			dataOut:    "datav3.out",
			ignoreHead: false,
			strictMode: false,
			shouldSame: true,
		},
		{
			desc:       "(S)New lines in the last",
			userOut:    "userv3.out",
			dataOut:    "datav3.out",
			ignoreHead: false,
			strictMode: true,
			shouldSame: true,
		},
		{
			desc:       "(L)Ignore the first line",
			userOut:    "userv4.out",
			dataOut:    "datav4.out",
			ignoreHead: true,
			strictMode: false,
			shouldSame: true,
		},
		{
			desc:       "(L)So many white characters",
			userOut:    "userv5.out",
			dataOut:    "datav5.out",
			ignoreHead: false,
			strictMode: false,
			shouldSame: true,
		},
		{
			desc:       "(S)So many white characters",
			userOut:    "userv5.out",
			dataOut:    "datav5.out",
			ignoreHead: false,
			strictMode: true,
			shouldSame: false,
		},
	}
	for idx, tc := range tests {
		t.Logf("%d. %s:\n", idx+1, tc.desc)
		nodiff, err := DiffOut(filepath.Join(path, "../testdata/step.judger/", tc.userOut), filepath.Join(path, "../testdata/step.judger/", tc.dataOut), tc.ignoreHead, tc.strictMode)
		if err != nil {
			t.Fatal(err)
		}
		if nodiff == tc.shouldSame {
			t.Logf("PASS!")
		} else {
			t.Fatalf("Should %v, got %v: %s\n", tc.shouldSame, nodiff, tc.desc)
		}
	}
}
