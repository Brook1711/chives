const {resolve} = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
module.exports = {
    mode:'production',
    /*mode:'development',*/
    entry: {
        one: ['./src/index.js'],
    },

    output: {
        filename: "[name].js",
        path: resolve(__dirname, 'build'),
    },
    module: {
        rules: [

        ]
    },
    plugins: [
        new HtmlWebpackPlugin({
            template: "./views/dplayer.html",
            filename: "build.html",
            minify: {
                collapseInlineTagWhitespace:true
            }
        })
    ],

}