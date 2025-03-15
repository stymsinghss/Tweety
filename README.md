## Tweety

A simple social network

### Tech Stack

Backend -> Golang

Database -> Postgres ( CockroachDb )

Frontend -> Vanilla JS

### Dependencies

| Package | Description | Link | Download |
|----------|----------|----------|----------|
| Branca   | Secure alternative to JWT   | https://github.com/hako/branca   | go get -u github.com/hako/branca |
| Pgx   | Postgres driver   | https://github.com/jackc/pgx/wiki/Getting-started-with-pgx   | go get github.com/jackc/pgx/v5 |
| Way   | HTTP router for Go   | https://github.com/matryer/way   | go get github.com/matryer/way |


# Database

- Install -> `brew install cockroach`

- Start a Single Node (Insecure Mode) ->
`cockroach start-single-node --insecure --host 127.0.0.1`

- Apply the schema.sql File -> `cockroach sql --insecure --database=mydb --file=schema.sql`

- Connect with a postgres client ->

| **Field**    | **Value**                             |
|--------------|---------------------------------------|
| **Host**     | localhost                             |
| **Port**     | 26257                                 |
| **Database** | mydb (or the name of your database)    |
| **Username** | root (or your configured user)         |
| **Password** | *(leave empty if using insecure mode)* |
