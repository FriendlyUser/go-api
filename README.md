# restapi-go-vue
Go REST Api + postresql db + Vue.js frontend

## Getting started

Create a postgresql database 

Create a products Table
``` sql
Create Table products(
id int primary key,
name varchar(30),
price float
);
```

``` bash
# install client dependencies
cd client
yarn

# build the client
yarn build

# build the backend
cd ..
go build

# setup the environment variables
# db username
APP_DB_USERNAME=...
# db password
APP_DB_PASSWORD=...
# db name
APP_DB_NAME=...
# db host
APP_DB_HOST=...
# application port
PORT=...

##### run the server
./restapi-go-vue
```

Browse http://localhost:{PORT}


# run the server
```
./go-api
```

### Running tests

Pass in all the parameters.
```sh
export TEST_DB_USERNAME=testuser TEST_DB_PASSWORD=testing TEST_DB_NAME=rgmp TEST_DB_HOST=localhost TEST_DB_PORT=5432; go test -v
```
