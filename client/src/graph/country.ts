import { gql } from '@urql/svelte';

export const GetCountry = gql`
	query GetCountry {
		CountryList {
			name
		}
	}
`;
