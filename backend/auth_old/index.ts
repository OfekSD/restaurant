
const dotenv = require('dotenv');
import express, {json} from 'express';
import authRouter from './routes/auth'

dotenv.config();

const app: express.Application = express();
const port: number = 8080;
app.use(json())

app.use('/',authRouter)
// Server setup
app.listen(port, () => {
    console.log(`TypeScript with Express
         http://localhost:${port}/`);
});