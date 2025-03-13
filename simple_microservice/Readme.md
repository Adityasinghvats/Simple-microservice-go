## A simple microservice project
- This project uses jwt_creater and jwt_parser running on ports 8080 and 9001 respectively.
- Both srvices run on different port mimicing a microsevices architecture where a service produces jwt token and other service is reponsible for user authentication.

### Use postman to test
- Make a GET request on http://localhost:8080 to get the token and copy it.
- Make a GET request on http://localhost:9001 with a field `Token` and value of the copied token to get the secret information.