# to create user
curl --location 'localhost:8089/v1/private/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"karthik",
    "email":"karthik@outlook.com"
}'
