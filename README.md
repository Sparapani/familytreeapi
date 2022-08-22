# familytreeAPI

- To start the API, execute "go build" and "./familytreeapi" or use a docker-compose file
- To make a request, use "Postman" or another software which supports REST calls

routes:
GET:
- /familytree/ - This endpoint return all family tree members with relatives 
- /familytree/relative/ - This endpoint return all relatives of a member 
- /familytree/baconsnumber/ - This endpoint return the baconsnumber of a member from and member to
- /familytree/member/ - This endpoint return the baconsnumber of a member from and member to
POST, PUT (UPDATE) AND DELETE:
- /familytree/member - This endpoint is gonna make a crud of a member (depends what you put on request)

initial data:
by default is create some members and relatives

The response statuses can be:
200 - OK
400 - Bad request (if some required field is empty)