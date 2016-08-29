package Dal

import (
    _ "github.com/mattn/go-sqlite3"
    "github.com/go-xorm/xorm"
    "github.com/go-xorm/core"
    "log"
    "github.com/hefju/StockA/model"
)

var engine *xorm.Engine

func init(){
    var err error
    engine, err = xorm.NewEngine("sqlite3", "./test.db")
    if err!=nil{
        log.Fatal("sqlite init falie.",err)
    }
    engine.SetMapper(core.SameMapper{})
    engine.ShowSQL(true)
    err = engine.Sync2(new(model.Stock))
    if err!=nil{
        log.Fatal("xorm sync falie.",err)
    }
}

func GetAllStock()[]model.Stock{
    list := make([]model.Stock, 0)
    err := engine.Find(&list)
    if err!=nil{
        log.Println("xorm GetAllStock falie.",err)
    }
    return list
}

func SaveStockCode(list []model.Stock){
    session := engine.NewSession()
    defer session.Close()
    // add Begin() before any action
    err := session.Begin()
    for _,st:=range list{
        _, err = session.Insert(&st)
    }

    if err != nil {
        session.Rollback()
        return
    }
    err = session.Commit()
    if err != nil {
        return
    }
}