from email import message
from fastapi import APIRouter, HTTPException
from models.food import Food
from lib.db import cur
router = APIRouter()


@router.get('/')
def get_foods():
    cur.execute("select * from foods")
    foods = cur.fetchall()
    return foods

@router.post('/')
async def add_food(food: Food):
    cur.execute("select * from foods where name = %s",(food.name,))
    if cur.fetchall():
        raise HTTPException(status_code=409,detail="this item exists alredy")
    cur.execute('insert into foods(name, description, type, price, is_vegan, is_vegetarian, is_spicy) values \
(%(name)s, %(description)s, %(type)s, %(price)s, %(is_vegan)s, %(is_vegetarian)s, %(is_spicy)s)', food.dict())
    return food
        
    