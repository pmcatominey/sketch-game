## Game Protocol

The protocol is centered around the game state, rather than each request returning a response, the request instead may modify the state which is then broadcast to all clients. The state may omit fields depending on which player receives it, for example answers submitted will only be sent to the player whose turn it is, other players will receive an empty value.

### State

```json
{
    "phase": "drawing",
    "drawer": "alex", // will be empty if phase = waiting
    "image": "data:image/png;base64,", // will be empty unless phase = guessing
    "players": {
        "alex": {
            "points": 1,
            "guess": "", // will be empty unless phase = guessing and recipient is current drawer
            "judged": false, // indicates if last guess has been reviewed by drawer
            "correct": false // indicates if last guess was correct, ignore unless judged = true
        },
        "sam": {
            "points": 2,
            "guess": "whatamgonnado",
            "judged": true,
            "correct": true
        }
    }
}
```

- Games move through the following phases
    - `waiting` the game has not started yet
        - once there are at least 2 players and 50% have indicated they are ready, the game will start
        - once ready the order is determined at random
    - `drawing` the next player in the turn order is now able to draw
        - once finished the client submits the drawing
    - `guessing` the drawing is revealed
        - the state will be updated upon each answer submitted
        - the drawer can mark the answer either `correct` or `incorrect`
        - the turn ends when either:
            - the drawer chooses to end their turn
            - all players guess correctly
        - points are awarded:
            - 1 point per correct answer
            - 1 additional point for being first
        - the next player in the order becomes the drawer and phase changes to `guessing`

### Messages

Messages are JSON encoded objects sent between the server and client, all messages contain a `type` attribute.

#### State Update

`Direction: Server => Client`

Incidates the game state has been changed, the UI should be updated to reflect the new state.

```json
{
    "type": "state",
    "state": { ... see above }
}
```

#### Disconnect

`Direction: Server => Client`

Informs the client that the user has been disconnected, this may be due to game ending, being cleaned up due to inactivity.

```json
{
    "type": "disconnect"
}
```

#### Ready

`Direction: Client => Server`

Incidates the player is ready while the game is in the `waiting` phase.

```json
{
    "type": "waiting"
}
```

#### Submit Drawing

`Direction: Client => Server`

Indicates the player drawing has finished, the image data must be in the form of a data url representing a PNG image.

If this message is sent by a player not currently drawing it will be ignored.

```json
{
    "type": "submitDrawing",
    "image": "data:image/png;base64,..."
}
```

#### Submit Guess

`Direction: Client => Server`

Used by players to guess the answer.

```json
{
    "type": "guess",
    "guess": "bitconnect man"
}
```

#### Submit Guess Judgement

`Direction: Client => Server`

Used by the drawer to mark guesses.

```json
{
    "type": "guess",
    "nickname": "sam",
    "correct": true
}
```

#### End Turn

`Direction: Client => Server`

Used by the drawer to indicate their turn is over.

If sent by a player who is not currently the drawer, the message will be ignored.

```json
{
    "type": "endTurn"
}
```