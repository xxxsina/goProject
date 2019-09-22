package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
)

//结构体1
type response1 struct {
	Page int
	Recorde []string
}
//结构体2 使用tag改变键值
type response2 struct {
	Page int `json:"xpage"`
	Recorde []string `json:"xrecord"`
}
//json
func funcJson() {
	//单字符
	j1, _ :=json.Marshal(true)
	println(string(j1))
	j2, _ := json.Marshal(1)
	println(string(j2))
	j3, _ := json.Marshal(2.34)
	println(string(j3))
	//切片与字典
	a := []string{"abc", "def", "hij"}
	j4, _ := json.Marshal(a)
	println(string(j4))
	b := map[string]int{"a" : 1, "b" : 2, "c" : 3}
	j5, _ := json.Marshal(b)
	println(string(j5))
	//结构体1
	c := &response1{
		Page:    1,
		Recorde: []string{"b", "c", "d"},
	}
	j6, _ := json.Marshal(c)
	println(string(j6))
	//结构体2
	d := &response2{
		Page:    2,
		Recorde: []string{"e", "f", "g"},
	}
	j7, _ := json.Marshal(d)
	println(string(j7))

	//解码json
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	// 保存解码后的数据，Value可以为任意数据类型
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	num := dat["num"].(float64)	//这里对字典里面的值强制转化类型
	fmt.Println(num)
	//这里对strs里面的值转化，这里interface为任意数据类型
	strs := dat["strs"].([]interface{})
	fmt.Println(strs)
	fmt.Println(strs[0].(string))
	//结构体解码
	str := `{"xpage":2,"xrecord":["e","f","g"]}`
	res := &response2{}
	if err := json.Unmarshal([]byte(str), &res); err != nil {
		panic(err)
	}
	fmt.Println(res)
	fmt.Println(res.Recorde[2])
	//std
	enc := json.NewEncoder(os.Stdout)
	dd := map[string]int{"apple" : 1, "lettuce" : 2}
	enc.Encode(dd)
}
//url
func funcUrl() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme)	//postgres
	fmt.Println(u.User)	//user:pass
	fmt.Println(u.User.Username())	//user
	fmt.Println(u.User.Password())	//pass true
	fmt.Println(u.Host)	//host.com:5432
	fmt.Println(u.Path)	// /path
	fmt.Println(u.Fragment)	//f
	fmt.Println(u.RawQuery)	//k=v
	m, _:= url.ParseQuery(u.RawQuery)
	fmt.Println(m["k"][0])
}
func main() {
	//funcJson()
	//funcUrl()
}
