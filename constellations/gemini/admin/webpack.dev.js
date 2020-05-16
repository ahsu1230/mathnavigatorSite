const merge = require("webpack-merge");
const common = require("./webpack.common.js");

module.exports = merge(common, {
    mode: "none",
    devServer: {
        host: "localhost",
        port: 9001,
        contentBase: "./",
    },
});
