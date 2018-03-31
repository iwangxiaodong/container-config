    package myext
    import (
        "fmt"
	    "time"
        "github.com/fnproject/fn/api/server"
        "github.com/fnproject/fn/fnext"
    )
    func init() {
        server.RegisterExtension(&fnext.Extension{
            Name:  "github.com/iwangxiaodong/container-config/fn/my-ext", // Should be the import name
            Setup: setup, // Fn will call this during startup
        })
    }
    func Setup(s fnext.ExtServer) error {
        s.AddAPIMiddlewareFunc(func(next http.Handler) http.Handler {
            return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                fmt.Println("My ext - ", time.Now())
            })
        })
        return nil
    }
