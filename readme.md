1. 
   basic version of project.
    1. i have made some serious design issue i.e. keeping num and servers apart, while better solution should be keeping all the available server in the array,
       so that there should not occur any issue when any of the server left the room.
    2. the basic version of this load balancer does only job is that it takes single route and transfer then our load balancer act as a client and hit the 
       server, and generated response should be transfered to the original client.
    3. there is also no fault tolerence of servers, detection of failure of server is some crucial task of manager that needs to be implemented.
    -------------------

2. 
   version 2 of project. 
   1. i have added resolved issue, all the servers details in array to avoid any of the error, also corrected the logic so that didn't get divide by 0 error
      when entire room got entry.
   2. maintained the basic health check of all the servers i.e. servers who are either getting late or not responding after some time, are monitored and directly 
      removed from the load balancer counting them as effective. -> upgradation would be if some of the server are slow then we should mark the server as unhealthy but also keep taking use of it according to its capacity to increase efficiency.
   3. i haved used algorithm of simple round robin or say token passing which will only work in the ideal scenarios in real life, every servers capability depend
      on different thing algorithm should mo realistic.
   4. added the feature of whatever route is calling load balancer will take it and then just transfer to the deserving server initially there was only 1
      hardcoded route that any client can hit request right now every route that is being defined on server is managable.
      it also includes every routes with their headers, body, tokens,  even response  everything that is required in http coommunication.
   5. as written active montioring of health of each server is done basically may be there will be something that can i learn more, i will implement that but 
      right now i didn't implemented passive checking of the servers.
         example :- 

               client hits -> http.xyz.com/api/ping
               load balancer -> http.server1.com/api/ping.
               at this point if load balancer couldn't complete it request or say server 1 is down,
              that doesn't mean user should not get service, immediately we should remove that server
              from the active healthy list so that another user should not recieve this same error,
              while new req should be transfer to the another server and serve the response to the client.

               basically, 
               load balancer <- status 500 <- http.server1.com/api/ping.
               then 
               load balancer -> http.server2.com/ping.
   -------------------------------------

