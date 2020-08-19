StackServer instructions:

Tests 

- [x] test_pops_to_empty_stack
- [x] test_one_slow_client_gets_killed_for_fast_clients (response code contradicting with test_server_resource_limit)
- [x] The test suite will cleanly close all of its connections when the teardown # runs after each test. Your code will need to account for, and cleanup closed# connections. To do this in a POSIX world, you must attempt to read from the # socket; if your read call returns 0, then the connection has been closed. # See the man page on recv: - implemented server reload for this.
- [x] test_full_stack_ignore
- [x] test_slowest_client_gets_killed 
- [x] test_slow_client_gets_killed_for_fast_client 
- [x] test_server_resource_limit (0xFF)
- [x] test_long_polling_push
- [x] test_server_queues_slow_message_correctly 
- [x] test_server_survives_half_message 
- [x] test_long_polling_get
- [x] test_serialized_requests
- [x] test_interleaved_requests
- [x] test_single_request
- [x] test_slow_clients_are_not_disconnected_for_no_reason

Technologies used: GoLang, Docker

To start server just go the root and type “make takeoff”, if it doesn't work, ensure make tool is installed.
To restart it - push on port 8081 “rel” command to the socket, in stack-test.rb it is “push_reload” action. Each test prepended with push_reload call to reset it’s state (as mentioned in the spec) except for test_single_request because it is used in different tests and may reset the connection in the middle. 

IF something is not working after socket server restart command - just CTRL+C in the container frontend and “make takeoff” once again but it’s unlikely - test “test_multiple_reload” was added to the test suite to ensure works normally.

The app is built with enabled race detector, this is not an option for the prod of course since it slows down performance, but in our case is to ensure there are no race conditions. The profiler on port 8090 commented and left just in case. 

If there is too much logging - you may reduce logging level change LOG_LEVEL: debug in docker-compose file to “info” or “error”. The stack-related logs are on the info level.

Also some unit tests are available, they are in *_test.go files. 

Both stack-test.rb and stack-test-original.rb are included in the root. Original version was not changed,
stack-test.rb contains some debugging outpus and server relaunch commands.

Known issues:
note#1 The tests test_pops_to_empty_stack and test_server_resource_limit seems to be contradicting each other. If 
1. comment the code which begins from note#1 and ends EOF note#1 - the test_server_resource_limit will be working
2. uncomment the which begins from note#1 and ends EOF note#1 in test_pops_to_empty_stack works test_server_resource_limit doesn’t, although the code 0xFF is sent to the client.

