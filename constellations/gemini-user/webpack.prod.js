const webpack = require("webpack");
const { merge } = require("webpack-merge");
const common = require("./webpack.common.js");

module.exports = merge(common, {
    mode: "production",
    plugins: [
        new webpack.DefinePlugin({
            "process.env.MATHNAV_ORION_HOST": JSON.stringify(
                "http://mathnav-orion-lb-1990674869.us-east-1.elb.amazonaws.com:8001"
            ),
        }),
    ],
});
