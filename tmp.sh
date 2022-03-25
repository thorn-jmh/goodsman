
# /ping

curl -X GET "http://localhost:1926/ping"

# /api/users/records/goodsid

curl -X GET "http://localhost:1926/api/users/records/goods_id?employee_id=123456&goods_id=<goods_id>"

# /api/goods/new

curl -X POST "http://localhost:1926/api/goods/new" -H "Content-Type: application/json" --data '{
    "employee_id": "123456",
    "number": 1,
    "goods_authority": 1,
    "state": 1,
    "owner": "1919810",
    "goods_msg": {
        "Name": "namaiwa",
        "Type": "AOE-1",
        "Photo": "1234567",
        "cost": 123.5,
        "Consumables": 0
    }
}'

# /api/goods/states

curl -X POST "http://localhost:1926/api/goods/state" -H "Content-Type: application/json" --data '{
    "employee_id": "123456",
    "goods_id" :"<goods_id>",
    "goods_state": {
        "State": 0,
        "Number": 1
    }
}'

# /api/goods/borrow

curl -X POST "http://localhost:1926/api/goods/borrow" -H "Content-Type: application/json" --data '{
    "employee_id": "123456",
    "goods_id" :"<goods_id>",
    "goods_num" : 1
}'

# /api/goods/msg

curl -X GET "http://localhost:1926/api/goods/msg?employee_id=123456&goods_id=<goods_id>"