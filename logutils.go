package logutils

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"strings"
)

// 打印更漂亮的对象
func PrintObj(obj interface{}) {

	log.Println(FormatJSON(obj))

}

func PrintObjWithMsg(obj interface{}, msg string) {

	log.Println(msg+":", "\n", FormatJSON(obj))

}

func PrintError(err error, extraMsg ...string) {

	PrintErrorMsg(err.Error(), extraMsg...)

}

func PrintErrorMsg(errMsg string, extraMsg ...string) {

	msg := GetCallerInfo() + "\n"

	if len(extraMsg) > 0 {
		msg += strings.Join(extraMsg, "") + "\n"
	}

	msg += errMsg

	log.Println(msg)

}

func PrintCaller() {
	log.Println(GetCallerInfo())

}

func GetCallerInfo() string {
	_, file, line, _ := runtime.Caller(2)
	return fmt.Sprintf("%s:%d", file, line)

}

func FormatJSON(obj interface{}) string {
	jsIndent, _ := json.MarshalIndent(&obj, "", "\t")

	return string(jsIndent)
}
