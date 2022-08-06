#!/bin/python3
from fastapi import FastAPI

from routes.food import router as mealsRouter

app = FastAPI()


@app.get("/")
async def root():
    return {"message": "Hello World"}

app.include_router(mealsRouter,prefix='/food')