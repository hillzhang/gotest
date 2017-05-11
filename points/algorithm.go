package points

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"sort"
	"strconv"
	"sys_monitor/toolkits/concurrent/semaphore"
	"time"
	"sync"
	"sync/atomic"
	"github.com/goeval"
)

var _ = bufio.ErrAdvanceTooFar
//(1)     (a o1 b) o2 c o3 d
//(2)     (a o1 b o2 c) o3 d
//(3)     a o1 (b o2 c) o3 d
//(4)     a o1 (b o2 c o3 d)
//(5)     a o1 b o2 (c o3 d)

var (
	opreation = map[int]string{
		1:"+",
		2:"-",
		3:"*",
		4:"/",
	}
)

var file *os.File
var lists []string
var points []string
var points1 []string
var lock sync.RWMutex

func do(i,j,r int){

	lists = append(lists,fmt.Sprintf("(a%sb)%s(c%sd)",opreation[i],opreation[j],opreation[r]))
	lists = append(lists,fmt.Sprintf("(a%sb)%sc%sd",opreation[i],opreation[j],opreation[r]))
	lists = append(lists,fmt.Sprintf("(a%sb%sc)%sd",opreation[i],opreation[j],opreation[r]))
	lists = append(lists,fmt.Sprintf("a%s(b%sc)%sd",opreation[i],opreation[j],opreation[r]))
	lists = append(lists,fmt.Sprintf("a%s(b%sc%sd)",opreation[i],opreation[j],opreation[r]))
	lists = append(lists,fmt.Sprintf("a%sb%s(c%sd)",opreation[i],opreation[j],opreation[r]))
}

var error_count int32
var dail_count int32

type list []string

func(s list) Len()int{
	return len(s)
}
func (s list) Less(i,j int) bool{
	s1 := strings.Fields(s[i])[:4]
	s2 := strings.Fields(s[j])[:4]
	for i := 0;i < 4; i ++{
		a1, _ := strconv.Atoi(s1[i])
		a2, _ := strconv.Atoi(s2[i])
		if a1 != a2{
			return a1 < a2
		}
	}
	return false
}
func (s list) Swap(i,j int){
	s[i],s[j] = s[j],s[i]

}

func format(str string)string{
	str = strings.Replace(str,".0","",-1)
	str = strings.Replace(str,"("," ",-1)
	str = strings.Replace(str,")"," ",-1)
	str = strings.Replace(str,"+"," ",-1)
	str = strings.Replace(str,"-"," ",-1)
	str = strings.Replace(str,"*"," ",-1)
	str = strings.Replace(str,"/"," ",-1)
	strs := strings.Fields(strings.TrimSpace(str))

	var int_V = make([]int,4)
	for i := 0; i < 4; i ++ {
		val,_ := strconv.Atoi(strs[i])
		int_V[i] = val
	}
	sort.Ints(int_V)

	return fmt.Sprintf("%d %d %d %d",int_V[0],int_V[1],int_V[2],int_V[3])
}

func Points() {
	now := time.Now()
	var err error
	file,err = os.OpenFile("points_final.txt",os.O_APPEND|os.O_WRONLY|os.O_CREATE,0777)
	if err != nil {
		fmt.Println(err)
	}

	for i := 1;i <= 4;i ++{
		for j := 1; j <= 4; j ++{
			for r := 1; r <= 4; r ++{
				do(i,j,r)
			}
		}
	}
	fmt.Println(len(lists))

	//writer := bufio.NewWriter(file)
	//for _,val := range lists{
	//	writer.WriteString(val)
	//	writer.WriteByte('\n')
	//}
	//writer.Flush()

	s := goeval.NewScope()

	sema := semaphore.NewSemaphore(5)

	for i := 1;i <= 13;i ++{
		for j := 1; j <= 13; j ++{
			for r := 1; r <= 13; r ++{
				for k := 1; k <= 13; k ++{
					for _,val := range lists{
						sema.Acquire()
						go func(val string,i,j,r,k int) {
							atomic.AddInt32(&dail_count,1)
							defer sema.Release()
							defer func() {
								if r := recover();r !=nil{
									atomic.AddInt32(&error_count,1)
								}
							}()
							val = strings.Replace(val,"a",strconv.Itoa(i)+".0",-1)
							val = strings.Replace(val,"b",strconv.Itoa(j)+".0",-1)
							val = strings.Replace(val,"c",strconv.Itoa(r)+".0",-1)
							val = strings.Replace(val,"d",strconv.Itoa(k)+".0",-1)

							value,err := s.Eval(val)
							if err != nil {
								fmt.Println(err)
							}
							if value.(float64) == 24 || value.(float64)+ 0.00000000000001==24{
								lock.Lock()
								points = append(points,val)
								lock.Unlock()
							}
						}(val,i,j,r,k)
					}
				}
			}
		}
	}


	now1:=time.Now()
	fmt.Println(now1.Sub(now))
	writer := bufio.NewWriter(file)

	time.Sleep(time.Second * 10)
	fmt.Println(len(points))

	for _,val := range points{
		if Final(val){
			points1 = append(points1,val)
		}
	}
	fmt.Println(len(points1))

	map_exist := make(map[string]string)
	var result []string
	for _,val := range points1{
		if _,ok := map_exist[format(val)];!ok{
			result = append(result,format(val) + " --> " + strings.Replace(val,".0","",-1))
			map_exist[format(val)] = val
		}
	}
	sort.Sort(list(result))

	for _,l := range list(result){
		writer.WriteString(l)
		writer.WriteByte('\n')
	}
	writer.Flush()


	fmt.Println(len(map_exist))

}