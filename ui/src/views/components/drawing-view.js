/* eslint-disable no-unused-vars */
import {h} from 'hyperapp'
import Painterro from 'painterro'

const drawCanvas = () => {
  if (!window.drawCanvas) {
    window.drawCanvas = Painterro({
      id: 'canvas',
      defaultTool: 'brush',
      hiddenTools: ['save', 'open', 'close'],
    })
  }

  window.drawCanvas.show()
}

export const DrawingView = ({gameState, playerName, actions}) => {
  const isCurrentPlayer = gameState.drawingPlayer === playerName

  return (
    <div class="card">
      <header>
        <h4>{gameState.drawingPlayer} is drawing</h4>
      </header>

      { isCurrentPlayer &&
        <div class="canvas-wrapper">
          <div id="canvas" style="width: 400px; height: 300px;" oncreate={drawCanvas}></div>
        </div>
      }

      { isCurrentPlayer &&
        <footer class="is-right">
          <a class="button primary" onclick={() => actions.submitImage(window.drawCanvas.imageSaver.asDataURL())}>Done</a>
        </footer>
      }
    </div>
  )
}
