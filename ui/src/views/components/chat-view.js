/* eslint-disable no-unused-vars */
import {h} from 'hyperapp'

export const ChatView = ({chat, sendAction}) => (
  <div class="card chat-view">
    <header>
      <h4>Chat</h4>
    </header>

    <pre class="chat-box">
      { chat.map(msg => `${msg}
`) }
    </pre>

    <footer class="is-right">
      <input type="text" placeholder="Message" autocomplete="off" onkeyup={event => {
        if (event.code === 'Enter') {
          sendAction(event.target.value)
          event.target.value = ''
        }
      }} />
    </footer>
  </div>
)
