import express from  "express"
import path from "path"
import cors from "cors"


const app = express();
const port = 3000;


app.use(cors())

// Serve static files from the 'public' directory
app.use(express.static("C:/Users/Hashira/Desktop/nuclear-launch-codes/pods-client-test/static"));

// // Optional: Serve index.html from the root URL
// app.get('/', (req, res) => {
//     res.sendFile(path.join(__dirname, 'public', 'index.html'));
// });

app.listen(port, () => {
    console.log(`Server listening at http://localhost:${port}`);
});