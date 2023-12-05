### This is an example update file you must create your own update.sh in order to work.

steps:
  name: Update Database Schema
    script:
      # shellcheck disable=SC2164
      cd ../schema
      go install github.com/pressly/goose/v3/cmd/goose@latest
      goose mysql "username:password@tcp(127.0.0.1:3306)/dbname?parseTime=true" up
