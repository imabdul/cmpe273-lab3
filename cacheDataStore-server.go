package main
import  (
		"fmt"
		"strings"
		"strconv"
		"github.com/julienschmidt/httprouter"
		"encoding/json"
		"net/http"
		"sort")



var key1, key2, key3 int

type keyVals struct{
	Key int	`json:"key,omitempty"`
	Value string	`json:"value,omitempty"`
}


var instance1, instance2, instance3 [] keyVals

type keyBased []keyVals

func (a keyBased) Len() int {
	return len(a)
}

func (a keyBased) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a keyBased) Less(i, j int) bool {
	return a[i].Key < a[j].Key
}


//Puts the key value pair
func keyValPUT(rw http.ResponseWriter, request *http.Request,p httprouter.Params){
	port := strings.Split(request.Host,":")
	key,_ := strconv.Atoi(p.ByName("key_id"))
	if(port[1]=="3000"){
		instance1 = append(instance1, keyVals{key,p.ByName("value")})
		key1++
	}else if(port[1]=="3001"){
		instance2 = append(instance2, keyVals{key,p.ByName("value")})
		key2++
	}else{
		instance3 = append(instance3, keyVals{key,p.ByName("value")})
		key3++
	}	
}

//Gets you key value pair on a given key
func keyValGET(rw http.ResponseWriter, request *http.Request,p httprouter.Params){
	out := instance1
	ind := key1
	port := strings.Split(request.Host,":")
	if(port[1]=="3001"){
		out = instance2
		ind = key2
	}else if(port[1]=="3002"){
		out = instance3
		ind = key3
	}	
	key,_ := strconv.Atoi(p.ByName("key_id"))
	for i:=0 ; i< ind ;i++{
		if(out[i].Key==key){
			result,_:= json.Marshal(out[i])
			fmt.Fprintln(rw,string(result))
		}
	}
}


//Gets you all the key value pairs
func keyValsGET(rw http.ResponseWriter, request *http.Request,p httprouter.Params){
	port := strings.Split(request.Host,":")
	if(port[1]=="3000"){
		sort.Sort(keyBased(instance1))
		result,_:= json.Marshal(instance1)
		fmt.Fprintln(rw,string(result))
	}else if(port[1]=="3001"){
		sort.Sort(keyBased(instance2))
		result,_:= json.Marshal(instance2)
		fmt.Fprintln(rw,string(result))
	}else{
		sort.Sort(keyBased(instance3))
		result,_:= json.Marshal(instance3)
		fmt.Fprintln(rw,string(result))
	}
}


//main function here
func main(){

	key1 = 0
	key2 = 0
	key3 = 0

	router := httprouter.New()
    router.GET("/keys", keyValsGET) // endpoint that gives all key value pairs
    router.GET("/keys/:key_id", keyValGET) // endpoint that gives the key value pair on a given id
    router.PUT("/keys/:key_id/:value", keyValPUT) // endpoint that puts the key value pair

    go http.ListenAndServe(":3000", router) //instance 1
    go http.ListenAndServe(":3001", router) //instance 2
    go http.ListenAndServe(":3002", router) //instance 3

    select {}
}