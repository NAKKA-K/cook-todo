import * as React from "react";

import Typography from "@material-ui/core/Typography";

import {
    ConnectedTodoItemList,
    ConnectedNewTodoItem,
    ConnectedNewRecipeTodo
} from "../containers";

export const App = () => (
    <div>
        <Typography variant="display1" gutterBottom={true} id="cookdo-title">
            <img
                className="cookdo-title-logo"
                src="../../css/cookdo_logo.png"
            />
        </Typography>
        <ConnectedNewRecipeTodo />
        <ConnectedTodoItemList />
        <ConnectedNewTodoItem />
    </div>
);
