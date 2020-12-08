const webpack = require("webpack");
const { merge } = require("webpack-merge");
const common = require("./webpack.common.js");

module.exports = merge(common, {
    mode: "production",
    devServer: {
        host: "localhost",
        port: 9001,
        contentBase: "./",
    },
    plugins: [
        // Add environment variables
        new webpack.DefinePlugin({
            "process.env.MATHNAV_ORION_HOST": JSON.stringify(
                "https://www.andymathnavigator.com:8001"
            ),
        }),
    ],
});
