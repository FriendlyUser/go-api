[![Build Status](https://travis-ci.org/FriendlyUser/go-api.svg?branch=master)](https://travis-ci.org/FriendlyUser/go-api)
[![Documentation](https://godoc.org/github.com/FriendlyUser/go-api?status.svg)](https://godoc.org/github.com/FriendlyUser/go-api)
# go-api 
Simple Go REST API with postresql db + Vue.js frontend

## Getting started

Create a postgresql database 

Create a jobinfo Table
``` sql
CREATE TABLE IF NOT EXISTS jobinfo
(
id SERIAL,
numjobs INT,
avgkeywords NUMERIC(5,2) NOT NULL DEFAULT 0.00,
avgskills NUMERIC(5,2) NOT NULL DEFAULT 0.00,
city TEXT NOT NULL,
searchterm TEXT NOT NULL,
searchtime TEXT NOT NULL,
CONSTRAINT jobinfo_pkey PRIMARY KEY (id)
);
```

### Runnin Application

This go application reads reads the webpacked produced minified HTML file and corresponding js/css files inside the client folder at the path `client/dist/index.html`.

``` bash
# install client dependencies
cd client
yarn

# build the client
yarn build

# build the backend
cd ..
go build

# setup the connection url



Browse http://localhost:{PORT}


# run the server
```
./scrapper-api
```

### Running tests

Start the database
pg_ctl.exe restart -D "C:\Program Files\PostgreSQL\9.6\data"

Pass in all the parameters.
```sh
export TEST_DB_USERNAME=testuser TEST_DB_PASSWORD=testing TEST_DB_NAME=rgmp TEST_DB_HOST=localhost TEST_DB_PORT=5432; go test -v
```

#### Todo List

- [x] Finish writing unit tests 
- [x] Deploy to heroku (via travis is ideal)
- [ ] Build documentation using godoc
- [ ] Build client side application 
- [ ] Configure scripts to send data to api.
