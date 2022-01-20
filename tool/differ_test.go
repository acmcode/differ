package tool

import (
	"path/filepath"
	"testing"
)

func TestDiffOut(t *testing.T) {
	path, err := filepath.Abs(".")
	if err != nil {
		t.Fatal(err)
	}
	nodiff, err := DiffOut(filepath.Join(path, "../testdata/step.judger/user.out"), filepath.Join(path, "../testdata/step.judger/data.out"), false, false)
	if err != nil {
		t.Fatal(err)
	}
	if nodiff {
		t.Logf("No diff")
	} else {
		t.Fatal("What's the diff?")
	}
}

func TestDiffOutV1(t *testing.T) {
	path, err := filepath.Abs(".")
	if err != nil {
		t.Fatal(err)
	}
	nodiff, err := DiffOut(filepath.Join(path, "../testdata/step.judger/userv1.out"), filepath.Join(path, "../testdata/step.judger/datav1.out"), false, false)
	if err != nil {
		t.Fatal(err)
	}
	if nodiff {
		t.Logf("No diff")
	} else {
		t.Fatal("What's the diff?")
	}
}

func TestDiffOutV2(t *testing.T) {
	path, err := filepath.Abs(".")
	if err != nil {
		t.Fatal(err)
	}
	nodiff, err := DiffOut(filepath.Join(path, "../testdata/step.judger/userv2.out"), filepath.Join(path, "../testdata/step.judger/datav2.out"), false, false)
	if err != nil {
		t.Fatal(err)
	}
	if nodiff {
		t.Logf("No diff")
	} else {
		t.Fatal("What's the diff?")
	}
}

func TestDiffOutV3(t *testing.T) {
	path, err := filepath.Abs(".")
	if err != nil {
		t.Fatal(err)
	}
	nodiff, err := DiffOut(filepath.Join(path, "../testdata/step.judger/userv3.out"), filepath.Join(path, "../testdata/step.judger/datav3.out"), false, false)
	if err != nil {
		t.Fatal(err)
	}
	if nodiff {
		t.Logf("No diff")
	} else {
		t.Fatal("What's the diff?")
	}
}

func TestDiffOutV4(t *testing.T) {
	path, err := filepath.Abs(".")
	if err != nil {
		t.Fatal(err)
	}
	nodiff, err := DiffOut(filepath.Join(path, "../testdata/step.judger/userv4.out"), filepath.Join(path, "../testdata/step.judger/datav4.out"), true, false)
	if err != nil {
		t.Fatal(err)
	}
	if nodiff {
		t.Logf("No diff")
	} else {
		t.Fatal("What's the diff?")
	}
}

func TestDiffOutV5(t *testing.T) {
	path, err := filepath.Abs(".")
	if err != nil {
		t.Fatal(err)
	}
	nodiff, err := DiffOut(filepath.Join(path, "../testdata/step.judger/userv5.out"), filepath.Join(path, "../testdata/step.judger/datav5.out"), false, false)
	if err != nil {
		t.Fatal(err)
	}
	if nodiff {
		t.Logf("No diff")
	} else {
		t.Fatal("What's the diff?")
	}
}
