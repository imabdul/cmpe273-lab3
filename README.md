# cmpe273-lab3
A simple RESTful key-value cache data store. Idea is to understand consistent hashing.
<pre>
1. Server Side
PUT http://localhost:3000/keys/{key_id}/{value}
E.g. http://localhost:3000/keys/1/foobar
Response: 200
GET http://localhost:3000/keys/{key_id}
E.g. http://localhost:3000/keys/1
Response: {
                     “key” : 1,
                     “value” : “foobar”
                   }
        
GET http://localhost:3000/keys
E.g. http://localhost:3000/keys
Response: [
          {
                     “key” : 1,
                     “value” : “foobar”
           },
                       {
                                 “key” : 2,
                                “value” : “b”
                       }
            ]

2. Consistent Hashing on Client Side
Command to execute: go run /path/to/your/folder/consHashing-client.go
</pre>
