const express = require("express");
const path = require("path");
const port = 80; // Port for instance (HTTP)
const app = express();

app.use(express.static(__dirname));

app.get("*", (req, res) => {
    res.sendFile(path.resolve(__dirname, "index.html"));
});
app.listen(port);
console.log("Server started at " + port);
