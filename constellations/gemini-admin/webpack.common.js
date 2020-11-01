const webpack = require("webpack");

module.exports = {
    entry: "./src/index.js",
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
                test: /\.(css)$/,
                loader: "style-loader!css-loader",
            },
            {
                test: /\.(sass)$/,
                loader: "style-loader!css-loader!sass-loader",
            },
            {
                test: /\.(png|jp(e*)g|svg|gif)$/,
                loader: "file-loader",
            },
        ],
    },
};
