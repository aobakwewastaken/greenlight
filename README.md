# 🦕 greenlight

playing around with go, nothing fancy

## Migrations

For migrations I used golang-migrate

installation on mac

`brew install golang-migrate`

### To create migrations with golang-migrate

In your terminal run:

`migrate create -seq -ext=.sql -dir=./migrations migration_name`
