package main
import (
    "fmt"

    "github.com/hefju/StockA/Dal"
)
func main() {







    fmt.Println("fff")
}

//从文件中读取code,并写入数据库.得到stock名单
func loadStockCode(){
    a:=Dal.Loadcode()
    Dal.SaveStockCode(a)
    // Dal.Set("allstock",a)
    //    for k,v:=range a{
    //        fmt.Println(k,v)
    //    }
}

