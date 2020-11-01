const webpack = require("webpack");
const { merge } = require("webpack-merge");
const common = require("./webpack.common.js");

module.exports = merge(common, {
    mode: "production",
    devServer: {
        host: "localhost",
        port: 9000,
        contentBase: "./",
        historyApiFallback: true,
    },
    plugins: [
        new webpack.DefinePlugin({
            "process.env.MATHNAV_ORION_HOST": JSON.stringify(
                "https://www.andymathnavigator.com:8001"
            ),
        }),
    ],
});
