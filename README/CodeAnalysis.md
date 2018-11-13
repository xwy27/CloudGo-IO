# Code Analysis

1. main.go
    ```go
    func main() {
      port := os.Getenv("PORT")
      if len(port) == 0 {
        port = PORT
      }

      customPort := flag.StringP("port", "p", PORT, "server listening port")
      flag.Parse()
      if len(*customPort) != 0 {
        port = *customPort
      }

      server := xservice.NewServer()
      server.Run(":" + port)
    }
    ```
    我们首先在环境变量中查找是否有自定义端口环境变量，如果没有，则获取 terminal 参数，如果依旧没有，则让 server 监听默认端口 9000
1. ./service/server.go
    ```go
    func NewServer() *negroni.Negroni {
      formatter := render.New(render.Options{
        Directory:  "templates",
        Extensions: []string{".html"},
        IndentJSON: true,
      })

      router := mux.NewRouter()
      initRouter(router, formatter)

      n := negroni.Classic()
      n.UseHandler(router)
      return n
    }
    ```
    formatter 是对 [Render](https://github.com/unrolled/render/) 中间件的[参数设置](https://github.com/unrolled/render#available-options)，以提供自定义的数据渲染方式。
    router 为新建的路由，我们利用设置的 render 来初始化。
    n 为新 server，使用 [Negroni](https://github.com/urfave/negroni) 中间件，其中底层依旧是 go 的 net/http server，满足request 的并发请求。其中 Classic 为默认 server，可以调用 New 方法来实现自定义。

    ```go
    // initial router with a given render
    func initRouter(router *mux.Router, render *render.Render) {
      staticRoot, err := os.Getwd()
      if err != nil {
        panic("Could not retrive static file directory")
      }

      // path router
      router.HandleFunc("/", xrouter.HomeHandler(render))
      router.HandleFunc("/home", xrouter.HomeHandler(render))
      router.HandleFunc("/login", xrouter.LoginHandler(render))
      router.HandleFunc("/index", xrouter.LoginHandler(render))

      // Api router
      router.HandleFunc("/api/info", xrouter.InfoHandler(render))

      // Static file router
      router.PathPrefix("/templates").Handler(
        http.StripPrefix("/templates/", http.FileServer(http.Dir(staticRoot+"/templates/"))))
      router.PathPrefix("/static").Handler(
        http.StripPrefix("/static/", http.FileServer(http.Dir(staticRoot+"/static"))))

      // Not implement router
      router.NotFoundHandler = xrouter.DevelopHandler()
    }
    ```
    首先使用 `os.Getwd()` 获取 server 的运行目录，以支持静态文件的访问支持。
    下面实现 url 的 handler 分配，这里没有必须的格式要求，但个人根据 url 的作用，人为分隔代码块，可以帮助快速定位，这个看个人风格。
1. ./router/home.go
    ```go
    // HomeHandler renders the default home page to user
    func HomeHandler(formatter *render.Render) http.HandlerFunc {
      return func(w http.ResponseWriter, req *http.Request) {
        formatter.HTML(w, http.StatusOK, "home", struct{}{})
      }
    }
    ```
    handler 都返回一个 http.HandlerFunc，即，默认格式：
    ```go
    func Handler(formatter *render.Render) http.HandlerFunc {
      return func(w http.ResponseWriter, req *http.Request) {
        // Code depends on the return data
        // Refer to render data
        // formatter.XXX(w, statusCode, ...)
      }
    }
    ```
    这里，我们返回 home.html(其查找文件目录在 server.go 中指定)，后面的 struct 是返回的数据类型，这里没有返回数据。
1. ./router/info.go
    ```go
    type info struct {
      Author  string `json:"author"`
      Contact string `json:"contact"`
    }

    // InfoHandler returns the author info for info request
    func InfoHandler(formatter *render.Render) http.HandlerFunc {
      return func(w http.ResponseWriter, req *http.Request) {
        formatter.JSON(w, http.StatusOK, info{
          Contact: "xuwy27@mail2.sysu.edu.cn",
          Author:  "xwy27"})
      }
    }
    ```
    这里，我们返回的是一个 json，它的格式可以预先定义(在上面的 info)，然后构造一个 info 返回。
1. ./router/login.go
    ```go
    // Index info struct
    type indexInfo struct {
      Email    string `json:"Email"`
      Password string `json:"Password"`
    }

    // LoginHandler returns the login page with GET method and
    // deals with a login form and returns the result with POST method
    func LoginHandler(formatter *render.Render) http.HandlerFunc {
      return func(w http.ResponseWriter, req *http.Request) {
        if req.Method == "GET" {
          formatter.HTML(w, http.StatusOK, "login", struct{}{})
        } else if req.Method == "POST" {
          var email, password []string
          req.ParseForm()
          for k, v := range req.Form {
            switch k {
            case "email":
              email = v
            case "password":
              password = v
            }
          }
          formatter.HTML(w, http.StatusOK, "index", indexInfo{
            Email:    email[0],
            Password: password[0],
          })
        }
      }
    }
    ```
    这里处理登陆的表单和登陆页面，如果是 GET 方法，那么返回登陆页面，如果是 POST，说明是提交表单，接着处理表单，利用 http.Request 的 `ParseForm()` 函数，解析表单(**默认不解析**)，然后循环读出表单内容，后面可以添加自己的处理函数，这里直接返回个人页面，并附带一个 json 数据。这个 json 数据会在发送回客户端前渲染进入页面(后端渲染)。
    其绑定方法简单的介绍一下：
    如果返回页面时，同时返回了数据：
    ```json
    {
      Content: "hello"
    }
    ```
    页面内容如下：
    ```html
    <div>{{.Content}}</div>
    ```
    页面中的 .Content 会被 json 中的 Content 字段内容替代。