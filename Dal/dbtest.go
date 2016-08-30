package Dal
//import (
//    "github.com/mediocregopher/radix.v2/redis"
//    "log"
//    "fmt"
//)
////http://www.alexedwards.net/blog/working-with-redis
//func Connect(){
//    conn, err := redis.Dial("tcp", "localhost:6379")
//    if err != nil {
//        log.Fatal(err)
//    }
//    // Importantly, use defer to ensure the connection is always properly
//    // closed before exiting the main() function.
//    defer conn.Close()
//
//    // Send our command across the connection. The first parameter to Cmd()
//    // is always the name of the Redis command (in this example HMSET),
//    // optionally followed by any necessary arguments (in this example the
//    // key, followed by the various hash fields and values).
//    resp :=conn.Cmd("set", "AUTH", "abc","EX","3600") //conn.Cmd("set", "AUTH", "123456")
//
//    // Check the Err field of the *Resp object for any errors.
//
//    if resp.Err != nil {
//        log.Fatal("nimei",resp.Err)
//    }
//
//    title, err := conn.Cmd("GET", "EX").Str()
//    if err != nil {
//        log.Fatal(err)
//    }
//    fmt.Printf("%T\n",title)
//    fmt.Println("title:",title)
//
// //   fmt.Println("Electric Ladyland added!")
//}
//func TestA(){
//   // Set("AUTH",123456)
//    v,err:=Get("AUTH2")
//    if err!=nil{
//
//    }
//    fmt.Println(v)
//}
//func Set(Key string,Value interface{}){
//    client,_:=getClient()
//    resp :=client.Cmd("set", Key, Value)//client.Cmd("set", "AUTH", "123456")
//    if resp.Err != nil {
//        log.Fatal("nimei",resp.Err)
//    }
//
//}
//func Get(Key string)(interface{},error){
//    client,_:=getClient()
//    return client.Cmd("GET",Key).Str()
////    foo, err := client.Cmd("GET",Key).Str()
////    if err != nil {
////        return  nil,error.Error("找不到")
////    // handle err
////        log.Fatal("Get error ",err)
////    }
////    return foo,nil
//}
//func getClient()(*redis.Client, error){
//    conn, err := redis.Dial("tcp", "localhost:6379")
//    if err != nil {
//        log.Fatal(err)
//    }
//    return conn,nil
//}