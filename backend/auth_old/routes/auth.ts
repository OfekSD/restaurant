import jwt from 'jsonwebtoken'
import express, {Response,Request} from 'express' 
const router = express.Router()
import {env} from 'process'
import pool from '../lib/database'


const generateAccessToken = (username:String) => {
    return jwt.sign({username}, env.TOKEN_SECRET || "" , { expiresIn: '60d' });
  }


router.post('/login', (req: Request,res: Response) =>{
    console.log(req.body);
    
    const username: String = req.body["username"]
    const password: String = req.body["password"]
    console.log(username,password);
    
    res.json(generateAccessToken(username))

})


export default router