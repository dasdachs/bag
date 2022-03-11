module.exports = {
	extends: [
		"eslint:recommended",
		"plugin:import/recommended",
		"plugin:import/typescript",
		"plugin:react/recommended",
		"plugin:react/jsx-runtime",
		"plugin:react-hooks/recommended",
		"plugin:@typescript-eslint/recommended",
		"plugin:@typescript-eslint/recommended-requiring-type-checking",
		"plugin:prettier/recommended",
	],
	parser: "@typescript-eslint/parser",
	parserOptions: {
		project: "./tsconfig.json",
	},
	plugins: ["import", "react", "react-hooks", "@typescript-eslint", "prettier", "simple-import-sort"],
	rules: {
		"prettier/prettier": [
			2,
			{
				printWidth: 150,
			},
		],
		"react-hooks/rules-of-hooks": "error",
		"react-hooks/exhaustive-deps": "error",
		"simple-import-sort/imports": ["error", { groups: [["^react", "^\\u0000", "^@?\\w", "^[^.]", "^\\."]] }],
		"simple-import-sort/exports": "error",
	},
	settings: {
		react: {
			version: "detect",
		},
	},
}
