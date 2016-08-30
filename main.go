package main
import (
    "fmt"
    "github.com/hefju/StockA/Dal"
	"github.com/hefju/StockA/Bll"
	"time"
	"bufio"
	"os"
)
func main() {
//	Dal.UpdateTable()//更新数据库,什么操作也不做

	Bll.TeatA()
//	a:=Dal.GetAllStock()
//	fmt.Println(a)
//	for i:=0;i<2;i++{
//		//fmt.Println(a[i].Prefix+a[i].Code)
//		code:=a[i].Prefix+a[i].Code
//		start := time.Now()//记录开始时间
//		Bll.DownOne(code)
//		end := time.Now()	//记录结束时间
//
//		fmt.Println(" 耗时:",end.Sub(start).Nanoseconds() / 1000000,"ms")//输出执行时间，单位为毫秒。
//	}

    fmt.Println("fff")
//	fmt.Print("Press 'Enter' to continue...")
//	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func donothing(){
	println (time.Now().Day())
	Bll.DoNothing()

	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
func showdbcode(){
	a:=Dal.GetAllStock()
	for _,v:=range a {
		fmt.Println(v)
	}
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

