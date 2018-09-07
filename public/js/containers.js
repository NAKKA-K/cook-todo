import { connect } from "react-redux";

import { TodoItemList } from "./components/TodoItemList";
import { NewTodoItem } from "./components/NewTodoItem";
import { NewRecipeTodo } from "./components/NewRecipeTodo";

import {
    postTask,
    deleteTask,
    postToggle,
    deleteAllTask,
    postScrapingURL
} from "./actions";

export const ConnectedTodoItemList = connect(
    state => state,
    dispatch => ({
        postToggle: (task, tasks) => dispatch(postToggle(task, tasks)),
        deleteTask: (task, tasks) => dispatch(deleteTask(task, tasks)),
        deleteAllTask: () => dispatch(deleteAllTask())
    })
)(TodoItemList);

export const ConnectedNewTodoItem = connect(
    state => state,
    dispatch => ({
        postTask: title => dispatch(postTask(title))
    })
)(NewTodoItem);

export const ConnectedNewRecipeTodo = connect(
    state => state,
    dispatch => ({
        postScrapingURL: url => dispatch(postScrapingURL(url))
    })
)(NewRecipeTodo);
