package main
import(
      "github.com/prathyushnallamothu/dbconnection"
    )
func main(){
    db:=dbconnection.Connect()
    defer db.Close()
   }
   
//Now db is my database and you can perform all the database operations like insert,delete,update and many mysql queries
