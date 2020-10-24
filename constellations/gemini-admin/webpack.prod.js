const webpack = require("webpack");
const { merge } = require("webpack-merge");
const common = require("./webpack.common.js");

module.exports = merge(common, {
    mode: "production",
    plugins: [
        // Add environment variables
        new webpack.DefinePlugin({
            "process.env.MATHNAV_ORION_HOST": JSON.stringify(
                "http://ec2-3-80-123-134.compute-1.amazonaws.com:8001"
            ),
        }),
    ],
});
