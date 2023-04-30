const tmi = require("tmi.js")
const fs = require("fs")
const toml = require('toml');
const express = require('express');
const bodyParser = require('body-parser');

const config = toml.parse(fs.readFileSync("config.toml"));
const oauth = config.oauth
const app = express();
let messages = [];

const client = new tmi.Client(
{
    options: {debug: true},
    connection: 
    {
        reconnect: true,
        secure: true
    },
    identity: 
    {
        username: 'LiveChatGetBot',
        password: oauth
    },
    channels: [config.channel]
}
);

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({
    extended: true,
}));

app.get('/chat', (request, response) => {
    response.send(messages);
    messages.length = 0;
    console.log(`info: get request in /chat `)
});

client.on('message', (channel, tags, message, self) =>
{
    messages.push(
        {
            username: tags.username,
            message: message
        }
    )
}
);

const server = app.listen(config.port, (error) => {
    if (error) return console.log(`Error: ${error}`);
    console.log(`info: Server listening on port ${server.address().port}`);
});

client.connect();
