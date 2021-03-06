//log.SetFlags(log.Llongfile)
//使log输出加上行数

import(
  "github.com/gocraft/dbr"
	"log"
	"os"
	"io"
)

var Warning,Error * log.Logger

func init() {
	errFile,err:=os.OpenFile("errors.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err!=nil{
		log.Fatalln("打开日志文件失败：",err)
	}
  Info = log.New(os.Stdout,"Info:",log.Ldate | log.Ltime | log.Lshortfile)
  Error = log.New(io.MultiWriter(errFile),"Error:", log.Ltime | log.Lshortfile)
  输出到控制台和文件
  //Error = log.New(io.MultiWriter(os.Stderr,errFile),"Error:",log.Ldate | log.Ltime | log.Lshortfile)
}

func main(){
  //按格式输出
  Info.Println("。。。。。")
  Error.Println("。。。。。。")
}


const (
	Ldate         = 1 << iota     //日期示例： 2009/01/23
	Ltime                         //时间示例: 01:23:23
	Lmicroseconds                 //毫秒示例: 01:23:23.123123.
	Llongfile                     //绝对路径和行号: /a/b/c/d.go:23
	Lshortfile                    //文件和行号: d.go:23.
	LUTC                          //日期时间转为0时区的
	LstdFlags     = Ldate | Ltime //Go提供的标准抬头信息
)
