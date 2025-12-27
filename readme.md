1. 
   basic version of project.
    1. i have made some serious design issue i.e. keeping num and servers apart, while better solution should be keeping all the available server in the array,
       so that there should not occur any issue when any of the server left the room.
    2. the basic version of this load balancer does only job is that it takes single route and transfer then our load balancer act as a client and hit the 
       server, and generated response should be transfered to the original client.
    3. there is also no fault tolerence of servers, detection of failure of server is some crucial task of manager that needs to be implemented.
    -------------------
