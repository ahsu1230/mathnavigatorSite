const merge = require("webpack-merge");
const common = require("./webpack.common.js");

module.exports = merge(common, {
<<<<<<< HEAD
  mode: "development",
  devServer: {
    contentBase: "./",
  },
=======
    mode: "development",
    devServer: {
        contentBase: "./",
    },
>>>>>>> a27fb3b5070f8e1928daed628fb9a9038d1e89b9
});
