package main  

  
import (  
    "fmt"
    "io/ioutil"
    "sort"
    "encoding/json"
    "net/http"
    "hash/crc32"
)  
   
type circ []uint32

type keyVals struct{
    Key         int         `json:"key,omitempty"`
    Value       string      `json:"value,omitempty"`
}



func (h circ) Len() int {
    return len(h)
}

func (h circ) Less(i, j int) bool {
    return h[i] < h[j]
}

func (h circ) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

type node struct {
    id       int
    ip       string
}  
  
func newNode(id int, ip string) *node {
    return &node{
        id:       id,
        ip:       ip,
    }  
}  
  
type consHash struct {
    nodes       map[uint32]node
    doesExists  map[int]bool
    cir circ

}  
  
func newConsHash() *consHash {
    return &consHash{
        nodes:     make(map[uint32]node),
        doesExists: make(map[int]bool),
        cir:      circ{},
    }  
}

func (h *consHash) addNode(node *node) bool {
 
    if _, ok := h.doesExists[node.id]; ok {
        return false  
    }  
    str := h.returnNode(node)
    h.nodes[h.getHashVal(str)] = *(node)
    h.doesExists[node.id] = true
    h.sortHashCirc()
    return true  
}  
  
func (hr *consHash) sortHashCirc() {
    hr.cir = circ{}
    for k := range hr.nodes {
        hr.cir = append(hr.cir, k)
    }  
    sort.Sort(hr.cir)
}  
  
func (h *consHash) returnNode(node *node) string {
    return node.ip
}  
  
func (h *consHash) getHashVal(key string) uint32 {
    return crc32.ChecksumIEEE([]byte(key))  
}  
  
func (h *consHash) get(key string) node {
    hash := h.getHashVal(key)
    i := h. findNode(hash)
    return h.nodes[h.cir[i]]
}  

func (h *consHash) findNode(hash uint32) int {
    i := sort.Search(len(h.cir), func(i int) bool {return h.cir[i] >= hash })
    if i < len(h.cir) {
        if i == len(h.cir)-1 {
            return 0  
        } else {  
            return i  
        }  
    } else {  
        return len(h.cir) - 1
    }  
}

//Puts the key value pair
func keyValPut(circle *consHash, str string, input string){
        ipAddress := circle.get(str)
        address := "http://"+ipAddress.ip +"/keys/"+str+"/"+input
		fmt.Println(address)
        req,err := http.NewRequest("PUT",address,nil)
        client := &http.Client{}
        resp, err := client.Do(req)
        if err!=nil{
            fmt.Println("Error:",err)
        }else{
            defer resp.Body.Close()
            fmt.Println("PUT Request succeeded")
        }  
}

//Gets you key value pair on a given key
func keyValGet(key string,circle *consHash){
    var out keyVals
    ipAddress:= circle.get(key)
	address := "http://"+ipAddress.ip +"/keys/"+key
	fmt.Println(address)
    response,err:= http.Get(address)
    if err!=nil{
        fmt.Println("Error==>",err)
    }else{
        defer response.Body.Close()
        contents,err:= ioutil.ReadAll(response.Body)
        if(err!=nil){
            fmt.Println(err)
        }
        json.Unmarshal(contents,&out)
        result,_:= json.Marshal(out)
        fmt.Println(string(result))
    }
}

//Gets you all the key value pairs
func keyValsGet(address string){
     
    var out []keyVals
    response,err:= http.Get(address)
    if err!=nil{
        fmt.Println("Error:",err)
    }else{
        defer response.Body.Close()
        contents,err:= ioutil.ReadAll(response.Body)
        if(err!=nil){
            fmt.Println(err)
        }
        json.Unmarshal(contents,&out)
        result,_:= json.Marshal(out)
        fmt.Println(string(result))
    }
}

//main function here
func main() {   
    consHash := newConsHash()
    consHash.addNode(newNode(0, "127.0.0.1:3000"))
	consHash.addNode(newNode(1, "127.0.0.1:3001"))
	consHash.addNode(newNode(2, "127.0.0.1:3002"))

    //PUT Requests for key value pair
    fmt.Println(" ")
    fmt.Println("PUT Requests")
    keyValPut(consHash,"1","a")
	keyValPut(consHash,"2","b")
    keyValPut(consHash,"3","c")
	keyValPut(consHash,"4","d")
	keyValPut(consHash,"5","e")
	keyValPut(consHash,"6","f")
	keyValPut(consHash,"7","g")
	keyValPut(consHash,"8","h")
	keyValPut(consHash,"9","i")
	keyValPut(consHash,"10","j")

    //GET Requests for key value pair based on key
    fmt.Println(" ")
    fmt.Println("GET Requests")
    keyValGet("1", consHash)
	keyValGet("2", consHash)
	keyValGet("3", consHash)
	keyValGet("4", consHash)
	keyValGet("5", consHash)
	keyValGet("6", consHash)
	keyValGet("7", consHash)
	keyValGet("8", consHash)
	keyValGet("9", consHash)
	keyValGet("10", consHash)

    //GET all key value pairs requests
    fmt.Println(" ")
    fmt.Println("key value pairs stored on Instance 1 running on port 3000")
    keyValsGet("http://127.0.0.1:3000/keys")
    fmt.Println(" ")
	fmt.Println("key value pairs stored on Instance 2 running on port 3001")
    keyValsGet("http://127.0.0.1:3001/keys")
    fmt.Println(" ")
    fmt.Println("key value pairs stored on Instance 3 running on port 3002")
    keyValsGet("http://127.0.0.1:3002/keys")

}  
