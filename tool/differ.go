package tool

import (
	"bufio"
	"os"
	"strings"
)

// DiffOut is to diff the output between user and testcase
// TODO: we need to implement the "strictMode" mode in the future.
func DiffOut(userOut, dataOut string, diffIgnoreHead bool, strictMode bool) (bool, error) {
	src, err := os.Open(userOut)
	if err != nil {
		return false, err
	}
	defer src.Close()
	dst, err := os.Open(dataOut)
	if err != nil {
		return false, err
	}
	defer dst.Close()

	var strSrc, strDest strings.Builder

	rd := bufio.NewScanner(src)
	srcIgnoreHead := diffIgnoreHead
	for rd.Scan() {
		str := rd.Text()
		trimStr := strings.Trim(str, " ")
		if srcIgnoreHead {
			srcIgnoreHead = false
		} else {
			if len(trimStr) > 0 {
				strSrc.WriteString(trimStr + " ")
			}
		}
	}

	rd = bufio.NewScanner(dst)
	dstIgnoreHead := diffIgnoreHead
	for rd.Scan() {
		str := rd.Text()
		trimStr := strings.Trim(str, " ")
		if dstIgnoreHead {
			dstIgnoreHead = false
		} else {
			if len(trimStr) > 0 {
				strDest.WriteString(trimStr + " ")
			}
		}
	}

	return strings.Trim(strSrc.String(), " ") == strings.Trim(strDest.String(), " "), nil
}
