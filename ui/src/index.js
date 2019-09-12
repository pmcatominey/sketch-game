import {app} from 'hyperapp'
import {actions} from './actions/'
import {state} from './state/'
import {view} from './views/'
import {location} from '@hyperapp/router';

// window.main = withLogger(app)(state, actions, view, document.getElementById('app'))
window.main = app(state, actions, view, document.getElementById('app'))

location.subscribe(window.main.location);