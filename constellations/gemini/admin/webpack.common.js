var webpack = require("webpack");

module.exports = {
    entry: "./src/index.js",
    output: {
        publicPath: "/dist",
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
                test: /\.(css|styl|sass)$/,
                loader: "style-loader!css-loader!stylus-loader!sass-loader",
            },
        ],
    },
};
