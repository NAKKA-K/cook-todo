import * as React from "react";
import PropTypes from "prop-types";

import Paper from "@material-ui/core/Paper";
import CircularProgress from "@material-ui/core/CircularProgress";

const paperStyle = { margin: "0.5em", padding: "5px" };

export class TodoItemList extends React.Component {
    toggle(task) {
        this.props.postToggle(task, this.props.tasks);
    }
    destroy(task) {
        this.props.deleteTask(task, this.props.tasks);
    }
    destroyAll() {
        this.props.deleteAllTask();
    }
    semiAutoToggle() {
        const { tasks } = this.props;
        for (let i = 0; i < tasks.length; i++) {
            if (tasks[i].completed == false) {
                return this.toggle(tasks[i]);
            }
        }
    }
    render() {
        if (this.props.isLoading === true) {
            return <CircularProgress />;
        }
        const tasks = this.props.tasks.map((v, i) => (
            <li key={i}>
                <input
                    className="toggle"
                    type="checkbox"
                    checked={v.completed}
                    onChange={() => {
                        this.toggle(v);
                    }}
                />
                <button className="btn destroy" onClick={() => this.destroy(v)}>
                    削除
                </button>
                <Paper style={paperStyle}>{v.title}</Paper>
            </li>
        ));
        return (
            <div>
                <div className="master-buttons-row">
                    <button
                        className="btn samiauto-button"
                        onClick={() => this.semiAutoToggle()}
                    >
                        次の工程に進む
                    </button>
                    <button
                        className="btn all-destroy-button"
                        onClick={() => this.destroyAll()}
                    >
                        全て削除
                    </button>
                </div>
                <ol id="todo-list">{tasks}</ol>
            </div>
        );
    }
}
TodoItemList.propTypes = {
    isLoading: PropTypes.bool,
    tasks: PropTypes.array,
    postToggle: PropTypes.func,
    deleteTask: PropTypes.func,
    deleteAllTask: PropTypes.func
};
