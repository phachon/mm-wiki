package utils

import (
	"fmt"
	"math"
	"math/rand"
	"net"
	"strings"
	"time"
)

var Misc = NewMisc()

type misc struct{}

func NewMisc() *misc {
	return &misc{}
}

// get map default
func (m *misc) GetMapDefault(mapValue map[string]interface{}, key string, def interface{}) interface{} {
	value, ok := mapValue[key]
	if ok {
		return value
	} else {
		return def
	}
}

// rand string
func (m *misc) RandString(strlen int) string {
	codes := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	codeLen := len(codes)
	data := make([]byte, strlen)
	rand.Seed(time.Now().UnixNano() + rand.Int63() + rand.Int63() + rand.Int63() + rand.Int63())
	for i := 0; i < strlen; i++ {
		idx := rand.Intn(codeLen)
		data[i] = byte(codes[idx])
	}
	return string(data)
}

// rand int
func (m *misc) RandInt(strLen int) string {
	codes := "0123456789"
	codeLen := len(codes)
	data := make([]byte, strLen)
	rand.Seed(time.Now().UnixNano() + rand.Int63() + rand.Int63() + rand.Int63() + rand.Int63())
	for i := 0; i < strLen; i++ {
		idx := rand.Intn(codeLen)
		data[i] = byte(codes[idx])
	}
	return string(data)
}

// get local ip
func (m *misc) GetLocalIp() string {
	addrSlice, err := net.InterfaceAddrs()
	if nil != err {
		return "localhost"
	}
	for _, addr := range addrSlice {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if nil != ipnet.IP.To4() {
				return ipnet.IP.String()
			}
		}
	}
	return "localhost"
}

// 分页方法
// total: 一共多少记录
// page:  当前是第几页
// pagesize: 每页多少
// url: url里面的{page}会被替换成页码
// args:
//  order 分页条的组成，是一个数组，可以按着1-6的序号，选择分页条组成部分和每个部分的顺序
//  a_count   分页条中a页码链接的总数量,不包含当前页的a标签，默认10个。
//
func (m *misc) Page(total, page, pagesize int, url string, args ...interface{}) string {
	order := []int{1, 2, 3, 4, 5, 6}
	a_count := 10
	if len(args) >= 1 {
		order = args[0].([]int)
	}
	if len(args) >= 2 {
		a_count = args[1].(int)
	}
	a_num := a_count
	first := "首页"
	last := "尾页"
	pre := "上页"
	next := "下页"
	if a_num%2 == 0 {
		a_num++
	}
	pages := int(math.Ceil(float64(total) / float64(pagesize)))
	curpage := page
	if curpage > pages || curpage <= 0 {
		curpage = 1
	}
	body := `<span class="page_body">`
	prefix := ""
	subfix := ""
	start := curpage - ((a_num - 1) / 2)
	end := curpage + ((a_num - 1) / 2)
	if start <= 0 {
		start = 1
	}
	if end > pages {
		end = pages
	}
	if pages >= a_num {
		if curpage <= (a_num-1)/2 {
			end = a_num
		}
		if end-curpage <= (a_num-1)/2 {
			start -= int(math.Floor(float64(a_num)/float64(2))) - (end - curpage)
		}
	}
	for i := start; i <= end; i++ {
		if i == curpage {
			body += fmt.Sprintf(`<a class="page_cur_page" href="javascript:void(0);"><b>%d</b></a>`, i)
		} else {
			body += fmt.Sprintf(`<a href="%s">%d</a>`, strings.Replace(url, "{page}", fmt.Sprintf("%d", i), 1), i)

		}
	}
	body += "</span>"
	if curpage > 1 {
		prefix = fmt.Sprintf(`<span class="page_bar_prefix"><a href="%s">%s</a><a href="%s">%s</a></span>`, strings.Replace(url, "{page}", fmt.Sprintf("%d", 1), 1), first, strings.Replace(url, "{page}", fmt.Sprintf("%d", curpage-1), 1), pre)
	}
	if curpage != pages {
		subfix = fmt.Sprintf(`<span class="page_bar_subfix"><a href="%s">%s</a><a href="%s">%s</a></span>`, strings.Replace(url, "{page}", fmt.Sprintf("%d", curpage+1), 1), next, strings.Replace(url, "{page}", fmt.Sprintf("%d", pages), 1), last)
	}
	info := fmt.Sprintf(`<span class="page_cur">第%d/%d页</span>`, curpage, pages)
	id := fmt.Sprintf("gsd09fhas9d%d%d%d", rand.Intn(1000), rand.Intn(1000), rand.Intn(1000))
	gostr := fmt.Sprintf(`<script>function ekup(){if(event.keyCode==13){clkyup();}}function clkyup(){var num=document.getElementById('%s').value;if(!/^\d+$/.test(num)||num<=0||num>%d){alert('请输入正确页码!');return;};location='%s'.replace(/\{page\}/,document.getElementById('%s').value);}</script><span class="page_input_num"><input onkeyup="ekup()" type="text" id="%s" style="width:40px;vertical-align:text-baseline;padding:0 2px;font-size:10px;border:1px solid gray;"/></span><span class="page_btn_go" onclick="clkyup();" style="cursor:pointer;">转到</span>`, id, pages, url, id, id)
	totalstr := fmt.Sprintf(`<span class="page_total">共%d条</span>`, total)
	pagenation := []string{totalstr, info, prefix, body, subfix, gostr}
	output := []string{}
	for _, v := range order {
		if v-1 < len(pagenation) && v-1 >= 0 {
			output = append(output, pagenation[v-1])
		}
	}
	if pages > 1 {
		return strings.Join(output, "")
	}
	return ""
}

// 获取字符串子串的位置
func (m *misc) GetStrUnicodeIndex(str string, substr string) int {
	// 子串在字符串的字节位置
	result := strings.Index(str, substr)
	if result >= 0 {
		return m.GetStrUnicodeIndexByByteIndex(str, result)
	}
	return -1
}

// 根据字符串字节位置获取字符串位置
func (m *misc) GetStrUnicodeIndexByByteIndex(str string, subStrByteIndex int) int {
	if subStrByteIndex > len(str)-1 {
		return -1
	}
	// 获得子串之前的字符串并转换成[]byte
	prefix := []byte(str)[0:subStrByteIndex]
	// 将子串之前的字符串转换成[]rune
	rs := []rune(string(prefix))
	// 获得子串之前的字符串的长度，便是子串在字符串的字符位置
	result := len(rs)
	return result
}

// 截取包含子字符串的一段字符串，前后截取
func (m *misc) SubStrUnicode(str string, subStr string, preLen int, sufLen int) string {
	subStrRune := []rune(subStr)
	strRune := []rune(str)
	count := len(strRune)
	subStrUnicodeIndex := m.GetStrUnicodeIndex(str, subStr)
	startIndex := 0
	endIndex := count - 1
	if subStrUnicodeIndex-preLen > 0 {
		startIndex = subStrUnicodeIndex - preLen
	}
	if subStrUnicodeIndex+len(subStrRune)+sufLen < count-1 {
		endIndex = subStrUnicodeIndex + len(subStrRune) + sufLen
	}
	return string(strRune[startIndex:endIndex])
}

// 截取包含子字符串的一段字符串，前后截取
// subStrIndex 已只子串的字节位置
func (m *misc) SubStrUnicodeBySubStrIndex(str string, subStr string, subStrIndex int, preLen int, sufLen int) string {
	subStrRune := []rune(subStr)
	strRune := []rune(str)
	count := len(strRune)
	subStrUnicodeIndex := m.GetStrUnicodeIndexByByteIndex(str, subStrIndex)
	startIndex := 0
	endIndex := count - 1
	if subStrUnicodeIndex-preLen > 0 {
		startIndex = subStrUnicodeIndex - preLen
	}
	if subStrUnicodeIndex+len(subStrRune)+sufLen < count-1 {
		endIndex = subStrUnicodeIndex + len(subStrRune) + sufLen
	}
	return string(strRune[startIndex:endIndex])
}
