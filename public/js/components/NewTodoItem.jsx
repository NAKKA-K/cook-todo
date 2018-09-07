import React from "react";
import PropTypes from "prop-types";

import TextField from "@material-ui/core/TextField";
import Button from "@material-ui/core/Button";

export class NewTodoItem extends React.Component {
    constructor(props) {
        super(props);
        this.state = { title: "" };
    }
    setTitle(title) {
        this.setState({
            title: title
        });
    }
    postTask(title) {
        if (title === "") {
            return;
        }
        this.props.postTask(title);
    }
    render() {
        return (
            <div className="input-row">
                <TextField
                    id="box"
                    label="新たしい工程を追加"
                    value={this.state.title}
                    onChange={e => this.setTitle(e.target.value)}
                    margin="normal"
                />
                <Button
                    variant="contained"
                    className="btn recipe-button"
                    onClick={() => this.postTask(this.state.title)}
                >
                    追加
                </Button>
            </div>
        );
    }
}
NewTodoItem.propTypes = {
    postTask: PropTypes.func
};
