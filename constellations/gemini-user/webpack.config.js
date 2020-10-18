var webpack = require("webpack");
var path = require("path");

module.exports = (env) => {
    return {
        entry: [path.resolve(__dirname, "src/app/index.js")],
        output: {
            publicPath: "/dist",
            filename: "./bundle.js",
        },
        devServer: {
            host: "localhost",
            port: 9000,
        },
        plugins: [
            // Add environment variables
            new webpack.DefinePlugin({
                "process.env.MATHNAV_ORION_HOST": JSON.stringify(
                    env.MATHNAV_ORION_HOST
                ),
            }),
        ],
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
                {
                    test: /\.(png|jpe?g|gif|svg|ttf)$/i,
                    use: [{ loader: "file-loader" }],
                },
            ],
        },
    };
};
