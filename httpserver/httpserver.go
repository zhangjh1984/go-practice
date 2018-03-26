package main

import(
    "log"
    "net/http"  
    "httpserver/util" 
    "io"
    "os"
)

var configcache map[string]string

// var session=make(map[string]string)
func upload(w http.ResponseWriter, req *http.Request){
    file, handler, err :=req.FormFile("file") 
    if err==nil{
        f, err := os.OpenFile(configcache["filePath"]+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
        if err==nil{
            log.Println("copy file")
            defer file.Close()
            defer f.Close()
            io.Copy(f, file)
        }
    }
}
func loginhandler(w http.ResponseWriter, req *http.Request){    
    // user:=configcache["user"];
    // password:=configcache["password"];
    // dburl:=configcache["db.url"]
    // dbuser:=configcache["db.user"]
    // password:=configcache["db.password"]
    // dbname:=configcache["db.name"]    
    requser:=req.FormValue("user")
    reqpass:=req.FormValue("password")
    var u=util.GetUser(requser)
    log.Println(u)
    if u.Username==requser && u.Password==reqpass{
        var cookie=new(http.Cookie)
        cookie.Name="haslogin"
        cookie.Value="true"
        req.AddCookie(cookie);
        log.Println("login success")
    }
}

func main(){
    configcache=util.LoadConfig()     
    log.Println(configcache)
    log.Println(configcache["filePath"])
    fs:=http.FileServer(http.Dir(configcache["filePath"]))
    http.Handle("/",fs)
    http.HandleFunc("/login",loginhandler)
    http.HandleFunc("/upload",upload)
    log.Println("start server")
    http.ListenAndServe("0.0.0.0:11111",nil)
}