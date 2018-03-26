package util

import(
    "log"    
    "os"
    "bufio"
    "strings"
)
var configcache map[string]string
// LoadConfig 加载配置
func LoadConfig()(map[string]string){       
    configcache=make(map[string]string)
    var f *os.File
    var err error   
    f,err=os.OpenFile("./config.txt",os.O_RDONLY,0666)
    if err==nil{        
        defer f.Close()
        r:=bufio.NewReader(f) 
        loop:=true
        for loop{       
            b,_,err:=r.ReadLine()
            if err==nil && b!=nil{
                log.Println(b)
                config:=strings.Split(string(b),"=")            
                configcache[config[0]]=config[1]
            }else{
                loop=false
            }
        }
    }
   return configcache
}
