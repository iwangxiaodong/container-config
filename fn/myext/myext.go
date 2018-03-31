    package myext
    import (
        "fmt"
	"net/http"
	"time"
        "github.com/fnproject/fn/api/server"
        "github.com/fnproject/fn/fnext"
    )
    func init() {
        server.RegisterExtension(&myExt{})
    }

type myExt struct {
}

func (e *myExt) Name() string {
	return "github.com/iwangxiaodong/container-config/fn/myext"
}

func (e *myExt) Setup(s fnext.ExtServer) error {
        s.AddAPIMiddlewareFunc(func(next http.Handler) http.Handler {
            return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                fmt.Println("My ext - ", time.Now())
            })
        })
        return nil
    }
