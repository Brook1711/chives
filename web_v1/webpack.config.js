const {resolve} = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const OptimizeCssAssetsWebpackPlugin = require('optimize-css-assets-webpack-plugin');
module.exports = {
    mode:'development',
    /*mode:'development'or 'production',*/
    entry: {
        index: ['./src/index.js'],
    },

    output: {
        filename: "[name].js",
        path: resolve(__dirname, 'build'),
    },
    module: {
        rules: [
            {test: /\.css$/, use:[MiniCssExtractPlugin.loader, 'css-loader', 'postcss-loader']},
            {test: /\.less$/, use: [MiniCssExtractPlugin.loader, 'css-loader', 'less-loader', 'postcss-loader']},
            {test: /\.scss$/, use: [MiniCssExtractPlugin.loader, 'css-loader', 'sass-loader', 'postcss-loader']},
            /*{test: /\.(png|jpg|jpeg|gif)$/, use: ['url-loader', {loader:'file-loader', options: {} }]}*/
            {
                test: /\.html$/,

                loader: 'html-withimg-loader',
            },
            {
                test: /\.(png|jpg|jpeg|gif)$/,
                loader: 'file-loader',
                options: {
                    outputPath: 'images/',
                    publicPath: 'images/',
                    limit: 1024 * 8,
                    /*name: '[name][hash:10].[ext]',*/
                    name: '[name][hash:10].[ext]',
                    esModule: false
                }
            },


        ]
    },
    plugins: [
        new HtmlWebpackPlugin({
            template: "./src/index.html",
            filename: "index.html",
            chunks: ['index'],
            minify: {
                collapseInlineTagWhitespace:true,
                removeComments:true,
            }
        }),

        new MiniCssExtractPlugin(
            {
                filename: '[name].css',
            }
        ),
        /*
        new OptimizeCssAssetsWebpackPlugin()
        */
    ],

}