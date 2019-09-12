/* eslint-disable no-unused-vars */
import {h} from 'hyperapp'

export const ResultView = ({gameState}) => (
  <div class="card">
    <header>
      <h4>{gameState.correctPlayer} wins!</h4>
    </header>

    <p>Winning answer: {gameState.correctAnswer}</p>

    <p>
      <img src={gameState.image} />
    </p>
  </div>
)
