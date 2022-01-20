## ACMDiffer
```
make
./acmdiffer testdata/step.judger/user.out testdata/step.judger/data.out false false
./acmdiffer testdata/step.judger/user.out testdata/step.judger/datav5.out false false
```

## The Function
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