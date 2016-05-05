package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"fmt"
	"net/http"
	"io/ioutil"
	"bytes"
	"encoding/json"
)

type ServiceController struct {
        beego.Controller
}

type Message struct {
        ID string
        Name string
        Port int64
        Address string
}

type DeRegMessage struct {
        Node string
        ServiceID string
}

type mystruct struct {
         FieldOne string `json:"field_one"`
}

func (c *ServiceController) GetNodes() {

	
        req := httplib.Get("http://52.23.208.10:8500/v1/catalog/nodes")
        req.Debug(true)
	str, err := req.String()
	
	if err != nil {
        	fmt.Printf("%s", err)
  	    	return
    	}
        fmt.Println(str)
	
	c.Data["json"] = &str
	c.ServeJSON()

}

func (c *ServiceController) GetService() {


	req := httplib.Get("http://52.23.208.10:8500/v1/catalog/services")
        req.Debug(true)
        str, err := req.String()

        if err != nil {
                fmt.Printf("%s", err)
                return
        }
        fmt.Println(str)

        c.Data["json"] = &str
        c.ServeJSON()

}

func (c *ServiceController) Service() {


        service := c.Ctx.Input.Param(":id")
        fmt.Println(service)

	s:= "http://52.23.208.10:8500/v1/catalog/service/"
        s += service
	fmt.Println(s)
	req := httplib.Get(s)
        
	req.Debug(true)
        str, err := req.String()

        if err != nil {
                fmt.Printf("%s", err)
                return
        }
        fmt.Println(str)

        if str == "[]" {
		c.Data["json"] = "Service not present"
	} else{
		c.Data["json"] = &str
        }
	c.ServeJSON()

}

func (c *ServiceController) Register() {

	var response Message
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &response)
	if err != nil {
        	fmt.Println(err)
        	return
    	}
	fmt.Println(response)

        fmt.Println(response.ID)	
	url := "http://52.23.208.10:8500/v1/agent/service/register"
    	fmt.Println("URL:>", url)


    	m := Message{response.ID, response.Name, response.Port, response.Address}
	b, err := json.Marshal(m)

    	req, err := http.NewRequest("PUT", url, bytes.NewReader(b))
    	req.Header.Set("X-Custom-Header", "myvalue")
    	req.Header.Set("Content-Type", "application/json")

    	client := &http.Client{}
    	resp, err := client.Do(req)
    	if err != nil {
		fmt.Printf("%s", err)
		return
    	}
    	defer resp.Body.Close()

    	fmt.Println("response Status:", resp.Status)
    	fmt.Println("response Headers:", resp.Header)
    	body, _ := ioutil.ReadAll(resp.Body)
    	fmt.Println("response Body:", string(body))

	if resp.Status != "200 OK" {
		c.Data["json"] = "Post failed"
	} else {
		c.Data["json"] = "Post Successful"
	}
	c.ServeJSON()

}

func (c *ServiceController) Deregister() {

        service_id := c.Ctx.Input.Param(":id")
        fmt.Println(service_id)

        url := "http://52.23.208.10:8500/v1/agent/service/deregister/"
        url += service_id
        fmt.Println("URL:>", url)
        
        req := httplib.Put(url)
        req.Debug(true)
        str, err := req.String()
        if err != nil {
                fmt.Printf("%s", err)
                return
        }
        fmt.Println(str)
        c.Data["json"] = "Deregistered"
        c.ServeJSON()


}


