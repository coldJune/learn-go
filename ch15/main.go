package main

import (
	"fmt"
	"io"
	"reflect"
	"encoding/json"
	"strings"
)

func main(){
	fmt.Println("reflect.Value和reflect.Type")
	i := 1
	iv := reflect.ValueOf(i)
	it := reflect.TypeOf(i)
	fmt.Println(iv,it)
	
	fmt.Println("\n获取原始类型")
	
	i1 := iv.Interface().(int)
	fmt.Println(i1)
	
	fmt.Println("\n修改对应的值")
	ipv := reflect.ValueOf(&i)
	ipv.Elem().SetInt(4)
	fmt.Println(i)
	
	p := person{Name:"cold",Age:18}
	ppv := reflect.ValueOf(&p)
	fmt.Println(ppv)
	ppv.Elem().Field(0).SetString("june")
	fmt.Println(p)
	
	fmt.Println(ppv.Kind())
	pv := reflect.ValueOf(p)
	fmt.Println(pv.Kind())
	
	fmt.Println("\nrefelect.Type")
	pt := reflect.TypeOf(p)
	// 遍历person的方法
	for i:=0;i<pt.NumField();i++{
		fmt.Println("字段：",pt.Field(i).Name)
	}
	// 遍历person的字段
	for i:=0;i<pt.NumMethod();i++{
		fmt.Println("方法：",pt.Method(i).Name)
	}
	
	fmt.Println("\n是否实现了某接口")
	stringerType := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	writerType := reflect.TypeOf((*io.Writer)(nil)).Elem()
	fmt.Println("是否实现了fmt.Stringer:",pt.Implements(stringerType))
	fmt.Println("是否实现了io.Writer:",pt.Implements(writerType))

	fmt.Println("\n字符串和结构体互转")
	// struct to json
	jsonB,err := json.Marshal(p)
	if err == nil{
		fmt.Println(string(jsonB))
	}
	
	// json to struct
	respJson :="{\"Name\":\"cold june\",\"Age\":26}"
	json.Unmarshal([]byte(respJson),&p)
	fmt.Println(p)
	
	fmt.Println("\nStruct Tag")
	respJson1 :="{\"name\":\"coldjune\",\"age\":21}"
	json.Unmarshal([]byte(respJson1),&p)
	fmt.Println(p)
	
	// 遍历person字段中key为json的tag
	for i:=0;i<pt.NumField();i++{
		sf := pt.Field(i)
		fmt.Printf("字段%s上，json tag为%s\n",sf.Name,sf.Tag.Get("json"))
		fmt.Printf("字段%s上，bson tag为%s\n",sf.Name,sf.Tag.Get("bson"))
	}
	
	
	fmt.Println("\n实现Struct转json")
	pv = reflect.ValueOf(p)
	pt = reflect.TypeOf(p)
	
	jsonBuilder := strings.Builder{}
	jsonBuilder.WriteString("{")
	num := pt.NumField()
	for i:=0;i<num;i++{
		jsonTag := pt.Field(i).Tag.Get("json")
		jsonBuilder.WriteString("\""+jsonTag+"\"")
		jsonBuilder.WriteString(":")
		jsonBuilder.WriteString(fmt.Sprintf("\"%v\"",pv.Field(i)))
		if i<num-1{
			jsonBuilder.WriteString(",")
		}
	}
	jsonBuilder.WriteString("}")
	fmt.Println(jsonBuilder.String())
}

type person struct{
	Name string `json:"name" bson:"b_name"`
	Age int	`json:"age" bson:"b_age"`
}
func (p person) String() string{
	return fmt.Sprintf("Name is %s,Age is %d",p.Name,p.Age)
}