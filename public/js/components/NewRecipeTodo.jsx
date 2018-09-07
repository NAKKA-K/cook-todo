import React from "react";
import PropTypes from "prop-types";

import TextField from "@material-ui/core/TextField";
import Button from "@material-ui/core/Button";

export class NewRecipeTodo extends React.Component {
    constructor(props) {
        super(props);
        this.state = { url: "" };
    }
    setURL(url) {
        this.setState({ url: url });
    }
    scrapingURL(url) {
        if (url === "") {
            return;
        }
        this.props.postScrapingURL(url);
    }

    render() {
        return (
            <div className="input-row">
                <TextField
                    id="box"
                    className="recipe-url-input"
                    label="クックパッドURLを貼ってください"
                    value={this.state.url}
                    onChange={e => this.setURL(e.target.value)}
                    margin="normal"
                />
                <Button
                    variant="contained"
                    className="btn recipe-button recipe-scraping-button"
                    onClick={() => this.scrapingURL(this.state.url)}
                >
                    調理手順の登録
                </Button>
            </div>
        );
    }
}
NewRecipeTodo.propTypes = {
    postScrapingURL: PropTypes.func
};
