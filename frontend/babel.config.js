import webpack from "webpack";

module.exports = {
  presets: ["@vue/cli-plugin-babel/preset"],
  plugins: [
    new webpack.optimize.LimitChunkCountPlugin({
      maxChunks: 1
    })
  ]
};
