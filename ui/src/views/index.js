/* eslint-disable no-unused-vars */
import {h} from 'hyperapp'
import { JoinView } from './components/join-view';
import { PlayerList } from './components/player-list';
import { WaitingView } from './components/waiting-view';
import { ResultView } from './components/result-view';
import { DrawingView } from './components/drawing-view';
import { Phase } from './components/phase';
import { GuessingView } from './components/guessing-view';
import { ChatView } from './components/chat-view';

export const view = (state, actions) => (
  <main class="container">

    { !state.connected &&
      <div class="row">
        <div class="col-4">
          <JoinView state={state} actions={actions} />
        </div>
      </div>
      }

      { state.connected &&
        <div class="row">
          <div class="col-4">
            <PlayerList gameState={state.gameState} playerName={state.playerName} />
            <ChatView chat={state.chat} sendAction={actions.sendChat} />
          </div>
          <div class="col">
            <Phase phase="waiting" render={WaitingView} />
            <Phase phase="drawing" render={DrawingView} />
            <Phase phase="guessing" render={GuessingView} />
            <Phase phase="result" render={ResultView} />
          </div>
        </div>
      }

  </main>
)
