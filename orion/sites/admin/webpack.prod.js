const merge = require("webpack-merge");
const common = require("./webpack.common.js");

module.exports = merge(common, {
<<<<<<< HEAD
  mode: "production",
=======
    mode: "production",
>>>>>>> a27fb3b5070f8e1928daed628fb9a9038d1e89b9
});
