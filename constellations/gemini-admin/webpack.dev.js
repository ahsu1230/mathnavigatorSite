const webpack = require("webpack");
const { merge } = require("webpack-merge");
const common = require("./webpack.common.js");

module.exports = merge(common, {
    mode: "none",
    devServer: {
        host: "localhost",
        port: 9001,
        contentBase: "./",
    },
    plugins: [
        // Add environment variables
        new webpack.DefinePlugin({
            "process.env.MATHNAV_ORION_HOST": JSON.stringify(
                "http://localhost:8001"
            ),
        }),
    ],
});
