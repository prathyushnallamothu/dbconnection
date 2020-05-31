# dbconnection
connect to a database in MySql using Golang
You can easily connect to a MySql database in Golang by importing this repository .
simply use go get github.com/prathyushnallamothu/dbconnection in your golang working directory.
and you can use the package to connect to database as follows
example:

package main
import(
      "github.com/prathyushnallamothu/dbconnection"
    )
func main(){
    db:=dbconnection.Connect()
    defer db.Close()
   }
   
Now db is my database and you can perform all the database operations like insert,delete,update and many mysql queries
