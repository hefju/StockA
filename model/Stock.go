
package model
type Stock struct{
    ID int64
    Name,Code string
	Prefix string //163指定1是泸,0是深
	DownFlag int//下载万要更新状态
	Statu int
}