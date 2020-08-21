const webpack = require("webpack");

module.exports = {
  filenameHashing: false,
  publicPath: "/goscope/",
  configureWebpack: {
    plugins: [
      new webpack.optimize.LimitChunkCountPlugin({
        maxChunks: 1
      })
    ]
  },
  chainWebpack: config => {
    config.optimization.delete("splitChunks");
  }
};
