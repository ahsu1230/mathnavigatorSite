var webpack = require("webpack");

module.exports = {
<<<<<<< HEAD
  entry: "./src/index.js",
  output: {
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
        test: /\.styl$/,
        loader: "style-loader!css-loader!stylus-loader",
      },
    ],
  },
=======
    entry: "./src/index.js",
    output: {
        filename: "./bundle.js"
    },
    module: {
        rules: [
            {
                test: /\.(js|jsx)$/,
                exclude: /node_modules/,
                use: {
                    loader: "babel-loader"
                }
            },
            {
                test: /\.styl$/,
                loader: "style-loader!css-loader!stylus-loader"
            }
        ]
    }
>>>>>>> a27fb3b5070f8e1928daed628fb9a9038d1e89b9
};
