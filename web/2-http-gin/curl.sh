curl --location 'localhost:8089/private/v1/users' \
--header 'username: admin' \
--header 'password: admin123' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"jiten",
    "email":"jiten@outlook.com"
}'