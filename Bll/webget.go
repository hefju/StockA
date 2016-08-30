package Bll
import(
	"fmt"
	"net/http"
	"io/ioutil"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"os"

	"bytes"
)

//http://quotes.money.163.com/service/chddata.html?code=1002566&start=20160829&end=20160829
func DownOne(code string){
	url:="http://quotes.money.163.com/service/chddata.html?code="+code+"&start=20160829&end=20160829"
	httpGet(code,url)
}
func httpGet(code,url string){
	resp, err := http.Get(url)
		if err != nil {
			// handle error
			fmt.Println("httpGet",err)
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("httpGet",err)
		}
		writetofile(code,body)
//	bs,_:=GbkToUtf8(body)//使用转码, 耗时将会翻倍
//		fmt.Println(string(bs))

	fmt.Println(string(body))
}
func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func GetUtf8(url string){
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept-Encoding", "utf-8")
	client := http.DefaultClient
	resp, err := client.Do(req)
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(result), err)
}
func writetofile(code string,body []byte){
	userFile := "data/"+code+".csv"
	fout,err := os.Create(userFile)
	defer fout.Close()
	if err != nil {
		fmt.Println(userFile,err)
		return
	}
	fout.Write(body)

}
func DoNothing(){

}