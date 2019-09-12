/* eslint-disable no-unused-vars */
import {h} from 'hyperapp'

export const JoinView = ({state, actions}) => (
  <div class="card">
    <header>
      <h4>Join</h4>
    </header>

    { state.error && <p class="text-error">{state.error}</p> }

    <p>
      <input placeholder="Name" id="name" type="text" autocomplete="off" />
    </p>

    <footer class="is-right">
      <a class="button primary" onclick={() => {
        let playerName = document.querySelector('#name').value
        actions.join(playerName)
      }}>Join</a>
    </footer>
  </div>
)
