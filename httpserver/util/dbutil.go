package util

import(   
	"log"
    "database/sql"
    _"github.com/Go-SQL-Driver/MySQL"
)

//User
type User struct{
	Username string
	Password string
}

var dbconn *sql.DB
// InitDbConn 初始化数据库连接
func InitDbConn(){
	log.Println("init db conn")
	db, err := sql.Open("mysql", configcache["db.user"]+":"+configcache["db.password"]+"@tcp("+configcache["db.url"]+")/"+configcache["db.dbname"]+"?charset=utf8") 
	if err==nil{
		dbconn=db
	}	
 }

 // GetUser 获取用户
 func GetUser(username string)(*User){
	//var u=user{username:"",password:""}
	 var u=new(User)
	 if dbconn==nil{
		InitDbConn()
	 }
	 result,err:=dbconn.Query("SELECT name,password FROM res_user where name=?",username)
	 defer result.Close()
	 var name,password string
	 if err==nil && result.Next(){
		 cols,err:=result.Columns()
		 if err==nil{
			 log.Println(cols)
		 }		 
		 log.Println("query result")
		 result.Scan(&name,&password)
		 u.Username=name
		 u.Password=password
	 }
	 return u
 }