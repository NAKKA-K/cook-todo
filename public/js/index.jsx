import * as React from "react";
import { render } from "react-dom";
import { Provider } from "react-redux";
import { createStore, applyMiddleware } from "redux";
import thunkMiddleware from "redux-thunk";
import { composeWithDevTools } from "redux-devtools-extension";

import { loadTodos, fetchToken } from "./actions";
import { rootReducer } from "./reducers";
import { App } from "./components/App";

const store = createStore(
    rootReducer,
    composeWithDevTools(applyMiddleware(thunkMiddleware))
);

const view = (
    <Provider store={store}>
        <App />
    </Provider>
);
render(view, document.getElementById("root"));

store.dispatch(fetchToken());
store.dispatch(loadTodos());
