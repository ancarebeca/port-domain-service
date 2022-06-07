# port-domain-services

## Purpose
This service is responsible to handles HTTP requests to retrieve, update and add new port information.
When the services start it also loads in  the DB the port information store in [ports.json](fixtures/ports.json) 

## HTTP API
 - Get all ports store in the system `GET /ports` 
 - Get information for a specific port given an ID `GET /ports/{id}`
 - Add/Update a port `POST /ports` 
   Example of payload: `{"id":"AAAAAA","name":"Knarrevik","city":"Knarrevik","country":"Norway","alias":[],"regions":[],"coordinates":[5.15,60.37],"province":"Hordaland","timezone":"Europe/Oslo","unlocs":["NOKRV"],"code":""`
    Note: If the user add a new Port it will be stored in the system but if the port already exists it will be updated as the exercises required.
    Personally, I would prefer to create two different endpoints for it (one for update and another one for create a port)

## How to run it?
 ``` 
    docker-compose up 
 ```

You can find in this [postman](docs/postman/port-domain-services.postman_collection.json) file all the calls need it to test the application

## Future improvements

The list below contains the TODO that can be done to improve the services that due to time constraints they weren't done

- [ ] Add tests to cover the unhappy path
- [ ] Add the requestID to the logs to be able to trace the flow execution
- [ ] Uses a real database, maybe a non-sql, currently this is using an inmemory database