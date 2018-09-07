const path = require("path");

module.exports = {
    entry: "./public/js/index.jsx",
    output: {
        path: path.resolve(__dirname, "public/dist"),
        filename: "bundle.js"
    },
    resolve: {
        extensions: [".js", ".jsx"],
        modules: ["node_modules"]
    },
    module: {
        rules: [
            {
                test: /\.jsx?$/,
                exclude: /node_modules/,
                use: {
                    loader: "babel-loader",
                    options: {
                        presets: ["react"],
                        plugins: ["transform-object-rest-spread"]
                    }
                }
            }
        ]
    }
};
