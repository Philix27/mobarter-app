import type { CodegenConfig } from '@graphql-codegen/cli';

const config: CodegenConfig = {
	schema: 'http://localhost:8080/graphql',
	documents: ['./src/**/*.{ts,tsx}', './src/**/*.gql'],
	generates: {
		'./src/graphql/': {
			preset: 'client',
			presetConfig: {
				gqlTagName: 'gql'
			},
			plugins: ['typescript', 'typescript-operations', 'graphql-codegen-svelte-apollo']
		}
	},
	ignoreNoDocuments: true
};
export default config;
