#!/bin/python3
from fastapi import FastAPI,Request
from lib.db import initialize


from routes.food import router as mealsRouter
initialize()
app = FastAPI()





app.include_router(mealsRouter,prefix='/')