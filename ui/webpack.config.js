const HtmlWebPackPlugin = require("html-webpack-plugin");

module.exports = {
        module: {
                rules: [
                        {
                                test: /\.js$/,
                                exclude: /node_modules/,
                                use: {
                                        loader: "babel-loader",
                                        options: {
                                                presets: [
                                                        '@babel/preset-env',
                                                        '@babel/preset-react'
                                                ],
                                                plugins: [
                                                        'styled-jsx/babel'
                                                ]
                                        }
                                }
                        },
                        {
                                test: /\.html$/,
                                use: {
                                        loader: 'html-loader'
                                }
                        }
                ]
        },
        plugins: [
                new HtmlWebPackPlugin({
                        template: "./src/index.html",
                        filename: "./index.html"
                })
        ]
};
