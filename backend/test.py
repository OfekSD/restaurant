#!/bin/python3
from psycopg2 import connect

DB_NAME = "restaurant"
DB_USER = "postgres"
DB_PASS = "Aa123456"
DB_HOST = "localhost"
DB_PORT = "5432"
 

conn = connect(database=DB_NAME,
                            user=DB_USER,
                            password=DB_PASS,
                            host=DB_HOST,
                            port=DB_PORT)
cur = conn.cursor()
cur.execute("select * from foods where type = %s",('ראשונות',))
print(cur.fetchall())