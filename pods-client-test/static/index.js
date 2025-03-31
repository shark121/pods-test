import express, { json } from "express"
import grapher from "./grapher.js";
import cors from "cors"




const app = express();



const port = 3000;


app.use(cors({
    origin: "*", 
    methods: ['GET', 'POST'], 
    allowedHeaders: ['Content-Type', 'Authorization'] 
}))


app.use(json())

let data = []

// app.get('/', (req, res) => {
//   res.send(data)
// });

// app.get("/", (req, res)=>{
//     res.sendFile("C:/Users/Hashira/Desktop/nuclear-launch-codes/pods-client-test/index.html")
//     res.sendFile("C:/Users/Hashira/Desktop/nuclear-launch-codes/pods-client-test/index.js")
// })

app.post("/", (req, res)=>{
    res.send("i heard you the first time")
    console.log(req.body)
    data.push(req.body)
    // grapher(req.body)
   
})

app.listen(port, () => {
  console.log(`Server listening at http://localhost:${port}`);
});