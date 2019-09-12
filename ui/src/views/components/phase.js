/* eslint-disable no-unused-vars */
import {h} from 'hyperapp'

export const Phase = ({phase, render}) => (state, actions) => (
  phase === state.gameState.phase &&
  render({
    gameState: state.gameState,
    playerName: state.playerName,
    actions,
  })
)
