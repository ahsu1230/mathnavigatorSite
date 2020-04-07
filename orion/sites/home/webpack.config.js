var webpack = require("webpack");

module.exports = {
<<<<<<< HEAD
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
=======
>>>>>>> c15f24dc4318ffae807d39aef3ef62f1b6948b26
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
<<<<<<< HEAD
>>>>>>> a27fb3b5070f8e1928daed628fb9a9038d1e89b9
=======
>>>>>>> c15f24dc4318ffae807d39aef3ef62f1b6948b26
};
