package context

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type MsContext struct {
	Request  *http.Request
	W        http.ResponseWriter
	routers  map[string]func(ctx *MsContext)
	pathArgs map[string]map[string]string
}

func NewContext() *MsContext {
	ctx := &MsContext{}
	ctx.routers = make(map[string]func(ctx2 *MsContext))
	ctx.pathArgs = make(map[string]map[string]string)
	return ctx
}

var UrlTree = NewTrie()

// 前缀树结构 用于路径参数匹配
type Trie struct {
	next   map[string]*Trie
	isWord bool
}

func NewTrie() Trie {
	root := new(Trie)
	root.next = make(map[string]*Trie)
	root.isWord = false
	return *root
}

// 插入数据， 路由根据 "/" 进行拆分
func (t *Trie) Insert(word string) {
	for _, v := range strings.Split(word, "/") {
		if t.next[v] == nil {
			node := new(Trie)
			node.next = make(map[string]*Trie)
			node.isWord = false
			t.next[v] = node
		}
		// * 匹配所有
		// {X}  匹配路由参数 X
		if v == "*" || strings.Index(v, "{") != -1 {
			t.isWord = true
		}
		t = t.next[v]
	}
	t.isWord = true
}

// 匹配路由
func (t *Trie) Search(word string) (isHave bool, arg map[string]string) {
	arg = make(map[string]string)
	isHave = false
	for _, v := range strings.Split(word, "/") {
		if t.isWord {
			for k := range t.next {
				if strings.Index(k, "{") != -1 {
					key := strings.Replace(k, "{", "", -1)
					key = strings.Replace(key, "}", "", -1)
					arg[key] = v
				}
				v = k
			}
			//log.Println("v = ", v)
		}
		if t.next[v] == nil {
			log.Println("找不到了, 匹配不上")
			return
		}
		t = t.next[v]
	}
	//log.Println(t.next, len(t.next))
	// 必须匹配全  比如: /v1/{b}/{a}  /v1/123匹配不到， /v1/123/456才可匹配
	if len(t.next) == 0 {
		isHave = t.isWord
		return
	}
	return
}
func (ctx *MsContext) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx.W = w
	ctx.Request = r
	path := r.URL.Path
	// /c/1  /c/{id}
	f := ctx.routers[path]
	if f == nil {
		for key, value := range ctx.routers {
			//判断是否为携带路径参数的
			reg, _ := regexp.Compile("(/\\w+)*(/{\\w+})+(/\\w+)*")
			match := reg.MatchString(key)
			if !match {
				continue
			}
			isHav, args := UrlTree.Search(path)
			if isHav {
				//匹配上 存储路径对应参数 /c/1 /c/{id}  id=1
				ctx.pathArgs[path] = args
				//pre handler
				value(ctx)
				//post handler
			}
		}
	} else {
		f(ctx)
	}
}

func (ctx *MsContext) Handler(url string, f func(context *MsContext)) {
	//前缀树的放入  /c/{id}  /c/1
	UrlTree.Insert(url)
	ctx.routers[url] = f
}

func (ctx *MsContext) GetPathVariable(key string) string {
	return ctx.pathArgs[ctx.Request.URL.Path][key]
}

func (ctx *MsContext) GetForm(key string) (string, error) {
	if err := ctx.Request.ParseForm(); err != nil {
		log.Println("表单获取失败：", err)
		return "", err
	}
	return ctx.Request.Form.Get(key), nil
}
func (ctx *MsContext) GetJson(key string) (interface{}, error) {
	// 读取请求体
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return nil, err
	}

	// 重置请求体，避免后续读取问题
	err = ctx.Request.Body.Close()
	if err != nil {
		return nil, err
	}
	ctx.Request.Body = http.NoBody

	// 解析 JSON
	var params map[string]interface{}
	err = json.Unmarshal(body, &params)
	if err != nil {
		return nil, err
	}

	// 返回指定键的值
	value, ok := params[key]
	if !ok {
		return nil, fmt.Errorf("key '%s' not found in JSON payload", key)
	}

	return value, nil
}
