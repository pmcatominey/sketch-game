import {location} from '@hyperapp/router'

export const state = {
  // routing
  location: location.state,

  connected: false,
  error: null,
  ws: null,

  playerName: '',
  gameState: null,
  chat: [],
}
