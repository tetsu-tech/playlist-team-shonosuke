
const path = require("path");
const webpack = require("webpack");

module.exports = {
  entry: "./src/index.js",
  mode: "development",
  module: {
    // ファイルをどのように変換すればよいのかのルールを設定。
    // testで入力するファイルの条件、excludeで除外する条件、
    //loaderで外部ライブラリのルールを参照する 
    rules: [
      {
        test: /\.(js|jsx)$/,
        exclude: /(node_modules|bower_components)/,
        loader: "babel-loader",
      },
      {
        test: /\.css$/,
        use: ["style-loader", "css-loader"]
      }
    ]
  },
  // ビルドの順番を設定
  resolve: { extensions: ["*", ".js", ".jsx"] },
  // ビルド後の設定
  // pathは、ビルド後のファイルを吐き出すフォルダ、
  // filenameはビルド後のファイル名を設定
  output: {
    path: path.resolve(__dirname, "dist/"),
    filename: "bundle.js"
  },
  // ローカルで起動するサーバの設定
  // contentBaseでブラウザからアクセスしたときのルート、
  // portはブラウザからアクセスするときのポート番号、
  // hotOnlyはファイルを更新したときに自動読み込みをする設定
  devServer: {
    contentBase: path.join(__dirname, "public/"),
    port: 8080,
    hotOnly: true
  },
  plugins: [new webpack.HotModuleReplacementPlugin()]
};