package tool

import (
	"bufio"
	"os"
	"strings"
)

// clearString replace all white chars
func clearString(str string) string {
	trimStr := strings.ReplaceAll(str, " ", "")
	trimStr = strings.ReplaceAll(trimStr, "\n", "")
	trimStr = strings.ReplaceAll(trimStr, "\r", "")
	return trimStr
}

// DiffOut is to diff the output between user and testcase
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
		if srcIgnoreHead {
			srcIgnoreHead = false
		} else {
			trimStr := str
			if !strictMode {
				trimStr = clearString(trimStr)
			}
			if len(trimStr) > 0 {
				strSrc.WriteString(trimStr)
			}
		}
	}

	rd = bufio.NewScanner(dst)
	dstIgnoreHead := diffIgnoreHead
	for rd.Scan() {
		str := rd.Text()
		if dstIgnoreHead {
			dstIgnoreHead = false
		} else {
			trimStr := str
			if !strictMode {
				trimStr = clearString(trimStr)
			}
			if len(trimStr) > 0 {
				strDest.WriteString(trimStr)
			}
		}
	}

	return strings.Trim(strSrc.String(), " ") == strings.Trim(strDest.String(), " "), nil
}
