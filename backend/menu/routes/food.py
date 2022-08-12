from fastapi import APIRouter, HTTPException
from models.food import Food
router = APIRouter()

@router.get('/')
def get_foods():
    from lib.db import connection_pool
    con = connection_pool.getconn()
    cur = con.cursor()
    cur.execute("select json_agg(t) from (select * from foods) as t")
    foods = cur.fetchall()[0][0]
    cur.close()
    connection_pool.putconn(con)
    return foods

@router.get('/{name}')
async def get_food_by_name(name):
    from lib.db import connection_pool
    con = connection_pool.getconn()
    cur = con.cursor()
    cur.execute('select json_agg(t) from (select * from foods where name = %s) as t',(name,))
    food = cur.fetchall()[0][0]
    if not food:
        cur.close()
        connection_pool.putconn(con)
        raise HTTPException(status_code=404,detail='Objcet not found')
    
    food = food[0]
    print(food)
    cur.close()
    connection_pool.putconn(con)
    return food
    

@router.post('/')
async def add_food(food: Food):
    from lib.db import connection_pool
    con = connection_pool.getconn()
    cur = con.cursor()
    cur.execute("select * from foods where name = %s",(food.name,))
    if cur.fetchall():
        cur.close()
        connection_pool.putconn(con)
        raise HTTPException(status_code=409,detail="this item exists alredy")
    cur.execute('insert into foods(name, description, type, price, is_vegan, is_vegetarian, is_spicy) values \
(%(name)s, %(description)s, %(type)s, %(price)s, %(is_vegan)s, %(is_vegetarian)s, %(is_spicy)s)', food.dict())
    con.commit()
    cur.close()
    connection_pool.putconn(con)
    return food

@router.put('/')
async def change_food(food:Food):
    from lib.db import connection_pool

    con = connection_pool.getconn()
    cur = con.cursor()
    cur.execute("select * from foods where name = %s",(food.name,))
    if not cur.fetchall():
        cur.close()
        connection_pool.putconn(con)
        raise HTTPException(status_code=404,detail="item was not found")
    cur.execute('Update foods \
        set description=%(description)s, type=%(type)s, price=%(price)s, is_vegan=%(is_vegan)s, is_vegetarian=%(is_vegetarian)s, is_spicy=%(is_spicy)s\
            where name=%(name)s', food.dict())
    con.commit()
    cur.close()
    connection_pool.putconn(con)
    return food

@router.delete('/{name}')
async def delete_food(name):
    from lib.db import connection_pool

    con = connection_pool.getconn()
    cur = con.cursor()
    cur.execute("select * from foods where name = %s",(name,))
    if not cur.fetchall():
        cur.close()
        connection_pool.putconn(con)
        raise HTTPException(status_code=404,detail="item was not found")
    cur.execute('delete from foods where name=%s',(name,))
    con.commit()
    cur.close()
    connection_pool.putconn(con)
    return {}
    