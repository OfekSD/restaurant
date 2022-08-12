from psycopg2 import pool

DB_NAME = "restaurant"
DB_USER = "postgres"
DB_PASS = "Aa123456"
DB_HOST = "localhost"
DB_PORT = "5432"
 
def initialize():
    global connection_pool
    connection_pool = pool.SimpleConnectionPool(1,20,database=DB_NAME, 
                                user=DB_USER,
                                password=DB_PASS, 
                                host=DB_HOST, 
                                port=DB_PORT)