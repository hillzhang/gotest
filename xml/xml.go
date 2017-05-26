package xml

import (
	"encoding/xml"
	"os"
	"fmt"
	"io/ioutil"
	"io"
	"github.com/qiniu/errors"
)
type Recurlyservers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	//Description string   `xml:",innerxml"`
}
type server struct {
	XMLName xml.Name `xml:"server"`
	ServerName string `xml:"serverName"`
	ServerIP string `xml:"serverIP"`
	Des string `xml:",innerxml"`
}

func Xml(){
	file,err := os.Open("test.xml")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	data,err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	v := Recurlyservers{}
	err = xml.Unmarshal(data,&v)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
}

func Xml1(){
	file,err := os.Open("test.xml")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	de := xml.NewDecoder(file)

	for{
		token,err := de.Token()
		if err != nil {
			if err == io.EOF{
				break
			}
		}
		switch  tt := token.(type) {
		case xml.SyntaxError:
			fmt.Println(errors.New(tt.Error()))
		case xml.StartElement:
			fmt.Println(tt.Attr)
		}
	}
}

type StringResources struct {
	XMLName        xml.Name         `xml:"resources"`
	ResourceString []ResourceString `xml:"string"`
}

type ResourceString struct {
	XMLName    xml.Name `xml:"string"`
	StringName string   `xml:"name,attr"`
	InnerText  string   `xml:",innerxml"`
}
func XML2(){
	

}