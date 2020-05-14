var webpack = require("webpack");
var os = require("os");
var path = require("path");

module.exports = {
    entry: [
        'webpack-dev-server/client?http://' + os.hostname() + ':9000/',
        path.resolve(__dirname, 'src/index.js')
    ],
    output: {
        filename: "./dist/bundle.js",
    },
    devServer: {
        host: '0.0.0.0',
        port: 9000
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
        ],
    },
};
