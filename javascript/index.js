const express = require('express');
const routes = require('./router');

const app = express();

const port = 3000;

app.use('/api/v1', routes);
app.listen(port, () => {
    console.log('Server started successfully and running at http://localhost:' + port)
})