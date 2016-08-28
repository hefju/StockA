package Dal
import (
    "io/ioutil"
    "fmt"
    "bytes"
    "github.com/hefju/StockA/model"
)

func Loadcode()[]model.Stock{
    list:=make([]model.Stock,0)
    filename:="c:/AMD/stock2.txt"//文件stock2为测试,只有10行数据,正式请使用stock
    b,err:=ioutil.ReadFile(filename)
    if err!=nil{
        fmt.Println(err)
        return nil
    }
    var sep []byte=[]byte(")")
    a := bytes.Split(b, sep)
    sep=[]byte("(")
    for _,v:=range a{
        namecode:=bytes.Split(v,sep)
        if len(namecode)==2{
            st:=model.Stock{Name:string(namecode[0]),Code:string(namecode[1])}
            list=append(list,st)
        }
    }
    return list
}
