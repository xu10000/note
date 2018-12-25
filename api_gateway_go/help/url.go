package help

import (
	"regexp"
	"strings"
)

func Get_url_regexp(url string) *regexp.Regexp {
	// 把字符串中的**全部替换为\S*
	url = "^" + strings.Replace(url, "/**", "\\S*", -1)
	// 直接结尾，以/结尾， ?带参数都合规
	url = url + "($|/$|\\?\\S*)"
	url = strings.Replace(url, "/**", "\\S*", -1)
	re, _ := regexp.Compile(url)

	return re
}
