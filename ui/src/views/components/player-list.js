/* eslint-disable no-unused-vars */
import {h} from 'hyperapp'

export const PlayerList = ({gameState, playerName}) => (
  <div class="card">
    <header>
      <h4>Players</h4>
    </header>

    <table>
      <tr>
        <th>Name</th>
        <th>Points</th>
      </tr>
      {
        Object.keys(gameState.players).map(name => {
          let player = gameState.players[name]
          return (
            <tr>
              { name === playerName && <td><strong>{name}</strong></td>}
              { name !== playerName && <td>{name}</td>}
              <td>{player.points}</td>
            </tr>
          )
        })
      }
    </table>
  </div>
)
