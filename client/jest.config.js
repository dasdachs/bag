module.exports = {
	rootDir: "src",
	testEnvironment: "jsdom",
	transform: {
		"^.+\\.tsx?$": "@swc/jest",
	},
	moduleNameMapper: {
		"\\.(css)$": "identity-obj-proxy",
		"single-spa-react/parcel": "single-spa-react/lib/cjs/parcel.cjs",
	},
	setupFilesAfterEnv: ["./setupTests.ts"],
}
