package web

import (
	"container/list"
	"net/url"
	"regexp"

	"github.com/hereyou-go/logs"
	"github.com/hereyou-go/web/http"
)

type RouteData struct {
	entry      *RouteEntry
	matchedUrl string
	values     map[string]string
}

type RouteEntry struct {
	method      http.HttpMethod
	pattern     *regexp.Regexp
	handler     HandlerFunc
	middlewares []Middleware
	paramNames  []string
}

type RouteTable struct {
	routes *list.List
}

func (rt *RouteTable) Register(method http.HttpMethod, pattern *regexp.Regexp, paramNames []string, handler HandlerFunc, segments int, endsWildcard bool, middlewares []Middleware) {
	route := &RouteEntry{
		method:      method,
		pattern:     pattern,
		handler:     handler,
		middlewares: middlewares,
		paramNames:  paramNames,
	}
	logs.Debug("map url %v to %v %v", pattern.String(), handler, method)
	//TODO:路由的优先级
	rt.routes.PushBack(route)
}

//http://stackoverflow.com/questions/30483652/how-to-get-capturing-group-functionality-in-golang-regular-expressions

// Match 匹配
func (rt *RouteTable) Match(method http.HttpMethod, url *url.URL) (*RouteData, bool) {
	//url := ""
	// logs.Debug("%v in %v = %v", method, url.Path, "route.pattern")
	var route *RouteEntry
	elem := rt.routes.Front()
	for elem != nil {
		route, _ = elem.Value.(*RouteEntry)
		elem = elem.Next()
		logs.Debug("%s %v in %v = %v", url.Path, method, route.method, route.method.In(method))
		if route == nil || !route.method.In(method) {
			continue
		}

		if route.pattern.MatchString(url.Path) {
			// logs.Debug("%v in %v = %v", method, route.method, method.In(route.method))
			// if !method.In(route.method) {
			// 	logs.Debug("not ok")
			// 	continue
			// }
			return &RouteData{
				entry:      route,
				matchedUrl: url.Path,
			}, true
		}
	}
	return nil, false
}
//
//type Entry struct {
//	path        string
//	patternType string
//	entries     []Entry
//}
//
//type RouteTable2 struct {
//	entries []Entry
//}
//
//func (*RouteTable2) Register(path string) {
//	bytes := []byte(path)
//	var arr []byte
//	for i:=0;i<len(bytes);i++ {
//		c := bytes[i]
//		if c == '\\'{
//			i++
//			arr = append(arr, bytes[i])
//			continue
//		} else if c == '/'{
//
//		}
//		arr = append(arr, c)
//	}
//
//}