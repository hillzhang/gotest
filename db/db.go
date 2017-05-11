package db


import (
	_"nas_monitor/utils/mysql"
	"database/sql"
	"log"
	"fmt"
	"time"
)

var(
	DB *sql.DB
)

func InitDatabase(){

	var err error

	DB,err = sql.Open("mysql", "hillzhang:-sAQ=t88@tcp(10.100.6.76:3306)/nas_history")
	if err != nil {
		log.Fatalln("open mysql rrd error",err)
	}
	DB.SetMaxIdleConns(5)
	err = DB.Ping()
	if err != nil {
		log.Fatalln("ping rrd database error",err)
	}
}

func Do() {
	InitDatabase()
	sql_real := "select value,metric from 102102223_store5mins where ts>=? && ts<=?"
	real_data := make(map[string][]int)
	rows,err := DB.Query(sql_real,time.Now().Add(-time.Minute*30).Unix(),time.Now().Unix())
	if err != nil {
		fmt.Println("error:",err)
	}
	defer rows.Close()
	for rows.Next(){
		var value int
		var metric string
		err := rows.Scan(&value,&metric)
		if err != nil {
			log.Println(err)
		}
		if _,ok := real_data[metric];ok{
			real_data[metric] = append(real_data[metric],value)
		}else {
			real_data[metric] = []int{value}
		}
	}
	for key,val := range real_data{
		count  := 0
		for _,num := range val{
			count += num
		}
		value := count/len(val)
		fmt.Println(key,value)
		if err != nil {
			log.Println(err)
		}
	}
}
