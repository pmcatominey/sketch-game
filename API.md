# API

## Notes

* All player names must match the regex `^\w+$`.

# Games

Games are rooms with players, they are created when the first person joins and
destroyed when the last person leaves.

### Joining a Game

Joining a game involves opening a WebSocket connection, providing a player name
and optionally a game key, without this a new game will be created.

Establish a WebSocket connection to the path:
`/api/join?name={playerName}&key={key}`

If there was an error joining the game, such as a player already joined with the
provided name, this will be sent through the socket and the connection closed.

# Protocol

## Game State


## Messages

### Error

Direction: Server => Client

Always followed by the connection being closed.

```json
{
    "type": "error",
    "data": {
        "error": "player already connected with provided name"
    }
}
```

### Send Chat Message

Direction: Client => Server

```json
{
    "type": "chat",
    "data": {
        "message": "hello"
    }
}
```

### Receive Chat Message

Direction: Server => Client

```json
{
    "type": "chat",
    "data": {
        "name": "tim",
        "message": "hello"
    }
}
```

### State

Direction: Server => Client

Contains an updated game state.

```json
{
    "type": "state",
    "data": {

    }
}
```

### Submit Image

Direction: Client => Server

```json
{
    "type": "submitImage",
    "data": {
        "image": "data:image/png;base64,aaaaaaa"
    }
}
```