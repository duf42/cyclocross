
package router

import(
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

var routes = []route{	
   newRoute("GET",  "/startlist", getStartList),
}

func newRoute(method, pattern string, handler http.HandlerFunc) route {
   return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

type route struct {
   method  string
   regex   *regexp.Regexp
   handler http.HandlerFunc
}


func Serve(w http.ResponseWriter, r *http.Request) {
   var allow []string
   for _, route := range routes {
       matches := route.regex.FindStringSubmatch(r.URL.Path)
       if len(matches) > 0 {
           if r.Method != route.method {
               allow = append(allow, route.method)
               continue
           }
           ctx := context.WithValue(r.Context(), ctxKey{}, matches[1:])
           route.handler(w, r.WithContext(ctx))
           return
       }
   }
   if len(allow) > 0 {
       w.Header().Set("Allow", strings.Join(allow, ", "))
       http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
       return
   }
   http.NotFound(w, r)
}

/*
* Utilities
*/
type ctxKey struct{}

func getField(r *http.Request, index int) string {
   fields := r.Context().Value(ctxKey{}).([]string)
   return fields[index]
}

func getStartList(w http.ResponseWriter, r *http.Request) {
	// name := getField(r, 0)
	fmt.Fprintf(w, "Get start list\n")
}


/*
* Entry-point
*/
func Router() http.HandlerFunc {

   router := http.HandlerFunc(Serve)
   return router

}