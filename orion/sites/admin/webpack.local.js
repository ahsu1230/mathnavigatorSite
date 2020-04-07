const merge = require("webpack-merge");
const common = require("./webpack.common.js");

module.exports = merge(common, {
<<<<<<< HEAD
<<<<<<< HEAD
  mode: "none",
  devServer: {
    contentBase: "./",
  },
=======
=======
>>>>>>> c15f24dc4318ffae807d39aef3ef62f1b6948b26
    mode: "none",
    devServer: {
        contentBase: "./",
    },
<<<<<<< HEAD
>>>>>>> a27fb3b5070f8e1928daed628fb9a9038d1e89b9
=======
>>>>>>> c15f24dc4318ffae807d39aef3ef62f1b6948b26
});
