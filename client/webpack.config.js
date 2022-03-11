const path = require("path")

const HtmlWebpackPlugin = require("html-webpack-plugin")

const portArgIndex = process.argv.indexOf("--port")
const port = portArgIndex !== -1 ? parseInt(process.argv[portArgIndex + 1]) : 3000

if (!process.env.NODE_ENV) {
	process.env.NODE_ENV = "development"
}

const isProduction = process.env.NODE_ENV === "production"

module.exports = {
	entry: `./src/index.tsx`,
	mode: isProduction ? "production" : "development",
	target: "web",
	output: {
		filename: "[name].[contenthash].js",
		path: path.resolve(__dirname, "dist"),
		publicPath: "/",
	},
	cache: isProduction ? false : { type: "filesystem" },
	module: {
		rules: [
			{
				test: /\.(ts|tsx|js|jsx)$/,
				exclude: /node_modules/,
				use: {
					loader: "swc-loader",
					options: {
						jsc: {
							parser: {
								syntax: "typescript",
								tsx: true,
								decorators: true,
							},
							transform: {
								legacyDecorator: true,
								react: {
									runtime: "automatic",
								},
							},
						},
						env: {
							coreJs: 3,
						},
						minify: true,
					},
				},
			},
			{
				test: /\.css$/i,
				use: ["style-loader", "css-loader"],
			},
			{
				test: /\.(png|svg|jpg|jpeg|gif)$/i,
				type: "asset/resource",
				generator: {
					filename: "assets/[hash][ext][query]",
				},
			},
		],
	},
	resolve: {
		extensions: [".tsx", ".ts", ".js", "jsx"],
	},
	devtool: isProduction ? "source-map" : "inline-cheap-module-source-map",
	devServer: {
		historyApiFallback: true,
		static: path.join(__dirname, "public"),
		allowedHosts: ["*"],
		headers: {
			"Access-Control-Allow-Origin": "*",
		},
		port,
		open: true,
		client: {
			overlay: {
				warnings: true,
				errors: true,
			},
		},
	},
	optimization: {
		splitChunks: {
			chunks: "all",
		},
	},
	plugins: [
		new HtmlWebpackPlugin({
			template: "public/index.html",
			hash: true,
			filename: "index.html",
		}),
	],
}
