const tmi = require("tmi.js")

const oauth = 'oauth:t36qo7sk3dboeo3eoyma4qy225q1wu'
const port = 8080;

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
    channels: ['guguch284']
}
);

client.on('message', (channel, tags, message, self) =>
{
    console.log(tags.username);
    console.log(message);
}
);

client.connect();
