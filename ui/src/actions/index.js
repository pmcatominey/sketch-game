import {location} from '@hyperapp/router'

let endpoint = window.location.host;
let wsProto = window.location.protocol.includes('https') ? 'wss:' : 'ws:'

if (window.location.host.includes('localhost')) {
  endpoint = 'localhost:8080'
}

// functions to handle received messages
// names match the message type
// arguments are message.data and state
const messageActions = {
  error: data => ({ error: data.error }),
  hello: data => {
    window.history.pushState(null, '', `?key=${data.key}`)

    return { connected: true }
  },
  state: gameState => ({ gameState }),
  chat: (data, state) => ({ chat: state.chat.concat(`${data.player}: ${data.message}`) }),
  event: (data, state) => ({ chat: state.chat.concat(data.event) }),
}

const send = (ws, type, data) => {
  ws.send(JSON.stringify({type, data}))
  return null;
}

export const actions = {
  location: location.actions,
 
  join: playerName => (state, actions) => {
    const key = new URLSearchParams(window.location.search).get("key");
    let ws = new WebSocket(`${wsProto}//${endpoint}/api/join?name=${playerName}&key=${key}`)

    ws.addEventListener('close', actions.onDisconnect)
    ws.addEventListener('message', actions.onMessage)

    return { ws, playerName, error: null }
  },

  onDisconnect: () => state => ({ connected: false, gameState: null, ws: null }),

  onMessage: event => state => {
    let msg = JSON.parse(event.data);
    let handler = messageActions[msg.type]
    
    return handler ? handler(msg.data, state) : null;
  },

  submitImage: image => (state, actions) => send(state.ws, 'submitImage', { image }),
  submitGuess: guess => (state, actions) => send(state.ws, 'submitGuess', { guess }),
  markGuessCorrect: player => (state, actions) => send(state.ws, 'markGuessCorrect', { player }),
  markGuessIncorrect: player => (state, actions) => send(state.ws, 'markGuessIncorrect', { player }),
  sendChat: message => (state, actions) => send(state.ws, 'chat', { message }),

}
