var webpack = require("webpack");
var path = require("path");

module.exports = (env) => {
    console.log("app environment: " + env.APP_ENV);
    return {
        entry: [path.resolve(__dirname, "src/index.js")],
        output: {
            filename: "./dist/bundle.js",
        },
        devServer: {
            host: "localhost",
            port: 9000,
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
};
