const webpack = require("webpack");
const path = require("path");

module.exports = {
    entry: [path.resolve(__dirname, "src/app/index.js")],
    output: {
        publicPath: "/dist/",
        filename: "./bundle.js",
    },
    module: {
        rules: [
            {
                test: /\.(js|jsx)$/,
                exclude: /node_modules/,
                use: {
                    loader: "babel-loader",
                },
            },
            {
                test: /\.(css|sass)$/,
                loader: "style-loader!css-loader!sass-loader",
            },
            {
                test: /\.(png|jpe?g|gif|svg|ttf)$/i,
                use: [{ loader: "file-loader" }],
            },
        ],
    },
};
