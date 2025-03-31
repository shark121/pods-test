import express from "express";
import cors from "cors";
import path from 'path';
import { fileURLToPath } from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const app = express();
const port = 3000;

app.use(cors());

// Serve static files from the 'static' directory (relative to the current file)
app.use(express.static(path.join(__dirname, 'static')));

// Optional: Serve index.html from the root URL
// app.get('/', (req, res) => {
//     res.sendFile(path.join(__dirname, 'static', 'index.html'));
// });

app.listen(port, () => {
    console.log(`Server listening at http://localhost:${port}`);
});