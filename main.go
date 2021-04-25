package main

import (
	"database/sql"
	"fmt"
	"strconv"

	"log"

	"github.com/pkg/errors"
)

func main() {
	/*数据库连接*/
	db, err := sql.Open("mysql",
		"user:XXX@tcp(127.0.0.1:3306)/XXX")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//1.如果是单个查询
	if   err:= SingleQueryTest(db);err!=nil{
		switch errors.Unwrap(err) {
		//比如查询没这个就直接创建
		case sql.ErrNoRows:
			 /*创建这条数据*/
			 case errors.New("其他的error"):
		case errors.New("致命错误的话"):
			//然后如果是例如gin网络层应该价格recover的middleware
			//例如   router.Use(gin.Recovery()) 保证还能恢复
			panic(err)
			
		}
	}else{
		//这里可能是更新这个数据
	}

   var userIds = []uint{1,2,3,4,5}

	//2.如果是循环差很多的
	errArray:= ArrayQueryTest(userIds,db)
	for _,err:=range errArray{
		switch errors.Unwrap(err) {
		case sql.ErrNoRows:
			log.Println("创建新的或者其他处理")
			//其他错误
			case errors.New("fatal"):
		case nil:
			log.Println("可能是更新或者")
		 if  userID,err:=	  strconv.Atoi(err.Error()) ;err!=nil{
		 	log.Println(err)
		 }else{
			 log.Printf("userID=%d",userID)
		 }

		}
		
	}


}


func SingleQueryTest(db *sql.DB) error {
	var name string
	err := db.QueryRow("select name from users where id = ?", 1).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
		  return 	errors.Wrap(err,"packageName functionName query")
		} else {
			return err
		}
	}
	return nil
}

func ArrayQueryTest(userIds []uint,db*sql.DB) []error{
	var errorArray []error
for k,userId:=range userIds{
	var name string
	err := db.QueryRow("select name from users where id = ?", 1).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			errorArray=append(errorArray,errors.Wrap(err, fmt.Sprintf("packageName functionName query %d and userID=%d",k,userId)))
		} else {
			errorArray=append(errorArray,errors.Wrap(err, "other error"))
		}
	}else{
		errorArray=append(errorArray,errors.Wrap(nil,fmt.Sprintf("%d",userId)))
	}
}
	return errorArray
}
