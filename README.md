# 🦕 greenlight

playing around with go, nothing fancy

## Migrations

For migrations I used golang-migrate

installation on mac

`brew install golang-migrate`

### To create migrations with golang-migrate

To generate your migration file with sequential numbering:

`migrate create -seq -ext=.sql -dir=./migrations migration_name`

To run your migration file:

`migrate -database "dsn" -path ./migrations up`
