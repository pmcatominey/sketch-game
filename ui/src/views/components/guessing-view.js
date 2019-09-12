/* eslint-disable no-unused-vars */
import {h} from 'hyperapp'

const submitGuess = actions => {
  let guess = document.querySelector('#guess').value
  actions.submitGuess(guess)
};

export const GuessingView = ({gameState, playerName, actions}) => {
  const isDrawing = gameState.drawingPlayer === playerName
  const player = gameState.players[playerName]
  const canGuess = player.guess === "" || player.guessIncorrect

  return (
    <div class="card">
      <header>
        <h4>{gameState.drawingPlayer} drew this</h4>
      </header>

      <p>
        <img src={gameState.image} />
      </p>

      { !isDrawing &&
        <footer class="is-right">
          { player.guessIncorrect && <p class="">Wrong answer, try again!</p> }
          <input id="guess" type="text" placeholder="Guess" value={player.guess} disabled={!canGuess} />
          <button class="button primary" onclick={() => submitGuess(actions)} disabled={!canGuess}>Send</button>
        </footer>
      }

      { isDrawing &&
        <p>
          <h4>Guesses</h4>
          <table>
            <tr><th>Guess</th><th>Actions</th></tr>
            { Object.keys(gameState.players)
                    .filter(name => gameState.players[name].guess !== "" && !gameState.players[name].guessIncorrect)
                    .map(name =>  (
                <tr>
                  <td>{gameState.players[name].guess}</td>
                  {/* <td><button class="button primary" onclick={() => console.log('nooo')} >Yes</button></td> */}
                  <td><button class="button primary" onclick={actions.markGuessCorrect.bind(this, name)} >Yes</button></td>
                  <td><button class="button danger" onclick={() => actions.markGuessIncorrect(name)} >No</button></td>
                </tr>
            ))}
          </table>
        </p>
      }
    </div>
  )
}
