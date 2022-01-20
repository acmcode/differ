## ACMDiffer
```
This is a core module used by ACMCoder OJ.
It's used to compare the user's out and the test case.
```

## How to make and test
```
go test github.com/acmcode/differ/tool
make
./acmdiffer testdata/step.judger/user.out testdata/step.judger/data.out false false
./acmdiffer testdata/step.judger/user.out testdata/step.judger/datav5.out false false
```

## The Function Definition
```
func DiffOut(userOut, dataOut string, diffIgnoreHead bool, strictMode bool) (bool, error)
```

## Params
```
userOut: user's output file
dataOut: the test case file
diffIgnoreHead: ignore the first line when comparing the contents
strictMode: whether ignore the space and enter key or not when comparing the contents
```

## Special test cases
```
TestDiffOutV2: ignore the space when comparing
TestDiffOutV5: ignore the enter key when comparing

Welcome some other special test cases.
```

## TODO
```
implement the strict mode
```

Welcome send issues or PRs.