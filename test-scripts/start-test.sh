#! /bin/bash
ip=$(ip -4 -o addr show enp0s3 | awk '{print $4}' | rev | cut -c 4- | rev)
gfsh << EOF
start locator --name=testLocator
start server --name=testServer --start-rest-api=true --http-service-port=8080 --http-service-bind-address=$ip
create region --name=testRegion --type=REPLICATE
EOF

echo "Adding testKey data."

curl -v -H 'Content-Type: application/json' -X PUT http://$ip:8080/gemfire-api/v1/testRegion/testKey -d @- << EOF
{
  "message": "hello, world.",
  "value": 1
}
EOF

echo "Adding testNestedKey data."

curl -v -H 'Content-Type: application/json' -X PUT http://$ip:8080/gemfire-api/v1/testRegion/testNestedKey -d @- << EOF
{
  "message": "hello, world.",
  "object": {
    "message": "hello, it's me.",
    "value": 2
  }
}
EOF

echo "Adding testTypeData data."

curl -v -H 'Content-Type: application/json' -X PUT http://$ip:8080/gemfire-api/v1/testRegion/testTypeDataKey -d @- << EOF
{
  "message": "hello, world.",
  "value": 3,
  "@type": "testKey"
}
EOF