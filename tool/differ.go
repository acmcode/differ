package tool

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// clearString replace all white chars
func clearString(str string) string {
	trimStr := strings.ReplaceAll(str, " ", "")
	trimStr = strings.ReplaceAll(trimStr, "\t", "")
	trimStr = strings.ReplaceAll(trimStr, "\n", "")
	trimStr = strings.ReplaceAll(trimStr, "\r", "")
	return trimStr
}

// clearEnter replace all '\r' chars
func clearEnter(str string) string {
	trimStr := strings.ReplaceAll(str, "\r", "")
	return trimStr
}

func getFileContent(fileName string, diffIgnoreHead bool, strictMode bool) (string, error) {
	src, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer src.Close()

	if stat, err := src.Stat(); err == nil {
		if stat.Size() == 0 {
			return "", err
		}
	} else {
		return "", err
	}

	var strBuilder strings.Builder
	rd := bufio.NewReader(src)
	srcIgnoreHead := diffIgnoreHead
	for {
		str, err := rd.ReadString('\n')
		if err == nil || err == io.EOF {
			if srcIgnoreHead {
				srcIgnoreHead = false
			} else {
				trimStr := str
				if !strictMode {
					trimStr = clearString(trimStr)
				} else {
					trimStr = clearEnter(trimStr)
				}
				if len(trimStr) > 0 {
					strBuilder.WriteString(trimStr)
				}
			}
		}
		if err == io.EOF {
			break
		}
	}
	return strBuilder.String(), nil
}

// DiffOut is to diff the output between user and testcase
func DiffOut(userOut, dataOut string, diffIgnoreHead bool, strictMode bool) (bool, error) {
	strSrc, err := getFileContent(userOut, diffIgnoreHead, strictMode)
	if err != nil {
		return false, err
	}

	strDst, err := getFileContent(dataOut, diffIgnoreHead, strictMode)
	if err != nil {
		return false, err
	}

	return strings.Trim(strSrc, " ") == strings.Trim(strDst, " "), nil
}
