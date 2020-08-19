#### To build client - launch `make client-build`

#####To put the graph into the server:
`./client -method=put -body='{"Edges":[{"A":1,"B":2},{"A":2,"B":3}]}'`

#####To get last graph shortest distance from the server: 
`./client -method=get -body='{"Src":1,"Dst":3}'`

#####To delete graph from server (ID here is the ID returned by "put" method):  
`./client -method=delete -body={"ID":"89508b8c-f9f2-0b0c-6ccd-cc8b2b9c0aa1"}`

