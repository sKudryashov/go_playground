### Graph GRPC server

#### Launch:
##### To launch the app, just go to the root and perform `make start` command, it will build and launch the application under docker container.
##### Pls note that app will be built with race detector enabled to ensure there is no race conditions. If you want to disable it, just uncomment #race detector disabled instruction in Dockerfile.

#### Client / actions
##### The app has three actions, you may perform them by cd ./cmd/client, make build-client and take a look at README.MD inside The client dir. 

#### Graph operations
##### The graph supports find minimum path operation, it returns all nodes in between including source and destination nodes. I didn’t use any libraries since it was quicker and simpler to implement it by myself. 
##### The nodes are supposed to be numeric beginning from 1 and src should be smaller than dst in minimum path search

#### Application layer explanation
##### The application consist of cmd / internal and pkg directories, where in cmd are all the launchers, pkg are shared libraries for common usage which may be reused by other components, and inside "internal" directory there is all application specific logic which is not supposed to be shared/reused by anyone else: handlers for various transport layers, service layer handling all business logic, repository layer handling all storage operations, model layer performing data model operations/converting from one format to another.  For such a simple application it wasn’t necessary, but taking into account `as best practices and code structure in a golang program` requirement I think it was worth to structure app this way.

#### Tests
##### Unit: three unit tests, all in pkg directory, with bunch of test cases for graph and a few for stack.
##### Functional: internal/handler/grpc/handler_test.go covers all the actions put/get last - calculate distance/delete with bunch of test cases.
##### Performance: this test is in cmd/perf directory, with appropriate instructions in cmd/perf/README.md file
##### To launch all the tests (except of performance) - perform `make test` from the app root

#### Storage layer
##### Is in internal/repository. Basically it’s an abstract layer over any kind of storage, under the hood there may be mongo/badger etc.. 
##### For the sake of simplicity I picked up a combination of hash map (to store graphs) and stack (to track the last added).

#### Concurrency
##### The app supports concurrent requests and provides concurrency safety by putting mutex locks on shared resources. I assume that would be enough for satisfying requirement `The server should deal with concurrent operations or restrict parallelism in some way`

###### In case of any questions: Sergey Kudryashov s.a.kudryashov@gmail.com