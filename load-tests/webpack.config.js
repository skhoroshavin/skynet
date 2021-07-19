let path = require('path');

module.exports = {
    mode: 'production',
    entry: './index.js',
    output: {
        filename: 'index.js',
        path: path.resolve(__dirname, 'build'),
        libraryTarget: 'commonjs',
    },
    module: {
        rules: [
            {
                test: /\.js$/,
                use: {
                    loader: 'babel-loader',
                    options: {
                        presets: ['@babel/preset-env']
                    }
                }
            }
        ],
    },
    performance: {
        maxEntrypointSize: 500*1024,
        maxAssetSize: 500*1024
    },
    target: 'web',
    externals: /^(k6|https?\:\/\/)(\/.*)?/,
    devtool: 'source-map',
};
