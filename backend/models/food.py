from pydantic import BaseModel


from pydantic import BaseModel

class Food(BaseModel):
    name: str
    description: str
    type: str
    price: float
    is_vegan = False
    is_vegetarian = False
    is_spicy = False
    