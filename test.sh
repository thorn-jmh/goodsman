# /ping

curl -X GET "http://{url}:1926/ping"

# /api/users/records/goodsid

curl -X GET "http://{url}:1926/api/users/records/goods_id?employee_id=123456&goods_id=<goods_id>"

# /api/goods/msg

curl -X GET "http://{url}:1926/api/goods/msg?employee_id=123456&goods_id=<goods_id>"

# /users/records/goods_id

curl -X GET "http://{url}:1926/api/users/records/goods_id?employee_id=123456&goods_id=<goods_id>"

# /api/goods/new

curl -X POST "http://{url}:1926/api/goods/new" -H "Content-Type: application/json" --data '{
    "employee_id": "123456",
    "number": 1,
    "goods_authority": 1,
    "state": 1,
    "owner": "1919810",
    "goods_msg": {
        "name": "namaiwa",
        "type": "AOE-1",
        "photo": "1234567",
        "cost": 123.5,
        "consumables": 0
    }
}'

# /api/goods/states

curl -X POST "http://{url}:1926/api/goods/state" -H "Content-Type: application/json" --data '{
    "employee_id": "123456",
    "goods_id" :"<goods_id>",
    "goods_state": {
        "state": 0,
        "number": 1
    }
}'

# /api/goods/borrow

curl -X POST "http://{url}:1926/api/goods/borrow" -H "Content-Type: application/json" --data '{
    "employee_id" : "123456",
    "goods_id" : "<goods_id>",
    "goods_num" : 1
}'

# /users/return_goods

curl -X POST "http://{url}:1926/api/users/return_goods" -H "Content-Type: application/json" --data '{
    "employee_id" : "123456",
    "goods_id" : "<goods_id>",
    "number" : 1
}'

# /users/return_goods/all

curl -X POST "http://{url}:1926/api/users/return_goods/all" -H "Content-Type: application/json" --data '{
    "employee_id" : "123456",
    "goods_id" : "<goods_id>"
}'

#需和前端联调部分
#/api/users/id

curl -X GET "http://127.0.0.1:1926/api/users/id" --data '{"code": <code> }'
 
#/api/users/authority
 
curl -X GET "http://127.0.0.1:1926/api/users/authority?employee_id=12dafadsf"

#机器人部分可直接在飞书上调试