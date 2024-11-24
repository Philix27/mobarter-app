import type { gql } from '@urql/svelte';
// import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = {
	[_ in K]?: never;
};
export type Incremental<T> =
	| T
	| { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
export type Omit<T, K extends keyof T> = Pick<T, Exclude<keyof T, K>>;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
	ID: { input: string; output: string };
	String: { input: string; output: string };
	Boolean: { input: boolean; output: boolean };
	Int: { input: number; output: number };
	Float: { input: number; output: number };
};

export type CountryResponse = {
	__typename?: 'CountryResponse';
	flag?: Maybe<Scalars['String']['output']>;
	name?: Maybe<Scalars['String']['output']>;
};

export type CreateBankAccountInput = {
	accountName: Scalars['String']['input'];
	accountNo: Scalars['String']['input'];
	bankName: Scalars['String']['input'];
	token?: InputMaybe<Scalars['String']['input']>;
	wallet: Scalars['String']['input'];
};

export type CreateBankAccountResponse = {
	__typename?: 'CreateBankAccountResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export type CreateUserInput = {
	firstName?: InputMaybe<Scalars['String']['input']>;
	lastName?: InputMaybe<Scalars['String']['input']>;
	walletAddress: Scalars['Int']['input'];
};

export type DataPlans = {
	__typename?: 'DataPlans';
	id?: Maybe<Scalars['String']['output']>;
	network?: Maybe<Scalars['String']['output']>;
	plan?: Maybe<Scalars['String']['output']>;
};

export type DeleteBankAccountInput = {
	accountId: Scalars['String']['input'];
	token?: InputMaybe<Scalars['String']['input']>;
	wallet: Scalars['String']['input'];
};

export type DeleteBankAccountResponse = {
	__typename?: 'DeleteBankAccountResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export type GetBankAccountResponse = {
	__typename?: 'GetBankAccountResponse';
	dexId?: Maybe<Scalars['Int']['output']>;
	downVotes?: Maybe<Scalars['Int']['output']>;
	id?: Maybe<Scalars['Int']['output']>;
	lossPercentage?: Maybe<Scalars['Float']['output']>;
	name?: Maybe<Scalars['String']['output']>;
	totalVotes?: Maybe<Scalars['Int']['output']>;
	upVotes?: Maybe<Scalars['Int']['output']>;
	winPercentage?: Maybe<Scalars['Float']['output']>;
};

export type GetBuyersInput = {
	vendor?: InputMaybe<VendorType>;
};

export type GetDataPlansInput = {
	category?: InputMaybe<Scalars['String']['input']>;
	network?: InputMaybe<NetworkProviders>;
};

export type GetDataPlansResponse = {
	[x: string]: any;
	__typename?: 'GetDataPlansResponse';
	dataPlans?: Maybe<Array<Maybe<DataPlans>>>;
};

export type GetP2PBuyersResponse = {
	__typename?: 'GetP2PBuyersResponse';
	accountName?: Maybe<Scalars['String']['output']>;
	accountNo?: Maybe<Scalars['String']['output']>;
	bankName?: Maybe<Scalars['String']['output']>;
	vendors?: Maybe<Array<Maybe<User>>>;
};

export type GetP2PSellersResponse = {
	__typename?: 'GetP2PSellersResponse';
	amount?: Maybe<Scalars['String']['output']>;
	name?: Maybe<Scalars['String']['output']>;
	rate?: Maybe<Scalars['String']['output']>;
};

export type GetSellersInput = {
	paymentMethod: Scalars['Int']['input'];
};

export type GetUserInput = {
	walletAddress: Scalars['Int']['input'];
};

export type Kyc_GetLevelResponse = {
	__typename?: 'Kyc_GetLevelResponse';
	accountName?: Maybe<Scalars['String']['output']>;
	accountNo?: Maybe<Scalars['String']['output']>;
	bankName?: Maybe<Scalars['String']['output']>;
};

export type Kyc_VerifyBvnInput = {
	bvn: Scalars['String']['input'];
};

export type Kyc_VerifyBvnResponse = {
	__typename?: 'Kyc_VerifyBvnResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export type Kyc_VerifyEmailInput = {
	email: Scalars['String']['input'];
	nin: Scalars['String']['input'];
	otp: Scalars['String']['input'];
};

export type Kyc_VerifyEmailResponse = {
	__typename?: 'Kyc_VerifyEmailResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export type Kyc_VerifyNinInput = {
	nin: Scalars['String']['input'];
};

export type Kyc_VerifyNinResponse = {
	__typename?: 'Kyc_VerifyNinResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export type Kyc_VerifyPhoneInput = {
	nin: Scalars['String']['input'];
	otp: Scalars['String']['input'];
	phone: Scalars['String']['input'];
};

export type Kyc_VerifyPhoneResponse = {
	__typename?: 'Kyc_VerifyPhoneResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export enum NetworkProviders {
	/** Regular user with limited access */
	Airtel = 'AIRTEL',
	/** Guest user with minimal access */
	Glo = 'GLO',
	/** Administrator with full access */
	Mtn = 'MTN'
}

export type PurchaseAirtimeInput = {
	amount: Scalars['Int']['input'];
	network?: InputMaybe<NetworkProviders>;
	phone?: InputMaybe<Scalars['String']['input']>;
	transactionHash: Scalars['String']['input'];
	walletAddress: Scalars['String']['input'];
};

export type PurchaseAirtimeResponse = {
	__typename?: 'PurchaseAirtimeResponse';
	message?: Maybe<Scalars['String']['output']>;
	network?: Maybe<Scalars['String']['output']>;
};

export type PurchaseDataInput = {
	amount: Scalars['Int']['input'];
	network?: InputMaybe<NetworkProviders>;
	phone?: InputMaybe<Scalars['String']['input']>;
	planId: Scalars['String']['input'];
	transactionHash: Scalars['String']['input'];
	walletAddress: Scalars['String']['input'];
};

export type PurchaseDataResponse = {
	__typename?: 'PurchaseDataResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export type RootMutations = {
	__typename?: 'RootMutations';
	BankAccount_Create?: Maybe<CreateBankAccountResponse>;
	BankAccount_Delete?: Maybe<DeleteBankAccountResponse>;
	JoinWaitList?: Maybe<WaitListResponse>;
	Kyc_VerifyBvn?: Maybe<Kyc_VerifyBvnResponse>;
	Kyc_VerifyEmail?: Maybe<Kyc_VerifyEmailResponse>;
	Kyc_VerifyNin?: Maybe<Kyc_VerifyNinResponse>;
	Kyc_VerifyPhone?: Maybe<Kyc_VerifyPhoneResponse>;
	PurchaseAirtime?: Maybe<PurchaseAirtimeResponse>;
	PurchaseData?: Maybe<PurchaseDataResponse>;
	User_Create?: Maybe<User_CreateResponse>;
};

export type RootMutationsBankAccount_CreateArgs = {
	input?: InputMaybe<CreateBankAccountInput>;
};

export type RootMutationsBankAccount_DeleteArgs = {
	input?: InputMaybe<DeleteBankAccountInput>;
};

export type RootMutationsJoinWaitListArgs = {
	country: Scalars['String']['input'];
	email: Scalars['String']['input'];
	firstName: Scalars['String']['input'];
	lastName: Scalars['String']['input'];
};

export type RootMutationsKyc_VerifyBvnArgs = {
	input?: InputMaybe<Kyc_VerifyBvnInput>;
};

export type RootMutationsKyc_VerifyEmailArgs = {
	input?: InputMaybe<Kyc_VerifyEmailInput>;
};

export type RootMutationsKyc_VerifyNinArgs = {
	input?: InputMaybe<Kyc_VerifyNinInput>;
};

export type RootMutationsKyc_VerifyPhoneArgs = {
	input?: InputMaybe<Kyc_VerifyPhoneInput>;
};

export type RootMutationsPurchaseAirtimeArgs = {
	input?: InputMaybe<PurchaseAirtimeInput>;
};

export type RootMutationsPurchaseDataArgs = {
	input?: InputMaybe<PurchaseDataInput>;
};

export type RootMutationsUser_CreateArgs = {
	input?: InputMaybe<CreateUserInput>;
};

export type RootQuery = {
	__typename?: 'RootQuery';
	Airtime_GetDataPlans?: Maybe<GetDataPlansResponse>;
	BankAccount_Get?: Maybe<GetBankAccountResponse>;
	BankAccount_GetAll?: Maybe<GetAllBankAccountResponse>;
	CountryList?: Maybe<CountryResponse>;
	GetBuyers?: Maybe<GetP2PBuyersResponse>;
	GetP2PSellers?: Maybe<GetP2PSellersResponse>;
	Kyc_GetLevel?: Maybe<Kyc_GetLevelResponse>;
	Transactions_GetAll?: Maybe<GetAllTransactionsResponse>;
	Transactions_GetOne?: Maybe<GetOneTransactionsResponse>;
	User_GetInfo?: Maybe<User_GetInfoResponse>;
};

export type RootQueryAirtime_GetDataPlansArgs = {
	input?: InputMaybe<GetDataPlansInput>;
};

export type RootQueryGetBuyersArgs = {
	input?: InputMaybe<GetBuyersInput>;
};

export type RootQueryGetP2PSellersArgs = {
	input?: InputMaybe<GetSellersInput>;
};

export type RootQueryUser_GetInfoArgs = {
	input?: InputMaybe<GetUserInput>;
};

export type User = {
	__typename?: 'User';
	id?: Maybe<Scalars['String']['output']>;
	name?: Maybe<Scalars['String']['output']>;
};

export type User_CreateResponse = {
	__typename?: 'User_CreateResponse';
	firstName?: Maybe<Scalars['String']['output']>;
	id?: Maybe<Scalars['String']['output']>;
	lastName?: Maybe<Scalars['String']['output']>;
	message?: Maybe<Scalars['String']['output']>;
	walletAddress?: Maybe<Scalars['String']['output']>;
};

export type User_GetInfoResponse = {
	__typename?: 'User_GetInfoResponse';
	firstName?: Maybe<Scalars['String']['output']>;
	id?: Maybe<Scalars['String']['output']>;
	lastName?: Maybe<Scalars['String']['output']>;
	message?: Maybe<Scalars['String']['output']>;
	walletAddress?: Maybe<Scalars['String']['output']>;
};

export enum VendorType {
	Buyers = 'BUYERS',
	Sellers = 'SELLERS'
}

export type GetAllBankAccountResponse = {
	__typename?: 'getAllBankAccountResponse';
	accountName?: Maybe<Scalars['String']['output']>;
	accountNo?: Maybe<Scalars['String']['output']>;
	bankName?: Maybe<Scalars['String']['output']>;
};

export type GetAllTransactionsResponse = {
	__typename?: 'getAllTransactionsResponse';
	country?: Maybe<Scalars['String']['output']>;
	firstName?: Maybe<Scalars['String']['output']>;
	lastName?: Maybe<Scalars['String']['output']>;
	walletAddress?: Maybe<Scalars['String']['output']>;
};

export type GetOneTransactionsResponse = {
	__typename?: 'getOneTransactionsResponse';
	country?: Maybe<Scalars['String']['output']>;
	firstName?: Maybe<Scalars['String']['output']>;
	lastName?: Maybe<Scalars['String']['output']>;
	walletAddress?: Maybe<Scalars['String']['output']>;
};

export type WaitListResponse = {
	__typename?: 'waitListResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export type GetCountryQueryVariables = Exact<{ [key: string]: never }>;

export type GetCountryQuery = {
	__typename?: 'RootQuery';
	CountryList?: { __typename?: 'CountryResponse'; name?: string | null } | null;
};

export type GetDataPlansQueryVariables = Exact<{
	input?: InputMaybe<GetDataPlansInput>;
}>;

export type GetDataPlansQuery = {
	__typename?: 'RootQuery';
	Airtime_GetDataPlans?: {
		__typename?: 'GetDataPlansResponse';
		dataPlans?: Array<{
			__typename?: 'DataPlans';
			plan?: string | null;
			network?: string | null;
			id?: string | null;
		} | null> | null;
	} | null;
};

export type KycVerifyBvnMutationVariables = Exact<{
	input?: InputMaybe<Kyc_VerifyBvnInput>;
}>;

export type KycVerifyBvnMutation = {
	__typename?: 'RootMutations';
	Kyc_VerifyBvn?: { __typename?: 'Kyc_VerifyBvnResponse'; message?: string | null } | null;
};

export type KycVerifyNinMutationVariables = Exact<{
	input?: InputMaybe<Kyc_VerifyNinInput>;
}>;

export type KycVerifyNinMutation = {
	__typename?: 'RootMutations';
	Kyc_VerifyNin?: { __typename?: 'Kyc_VerifyNinResponse'; message?: string | null } | null;
};

export type PurchaseDataMutationVariables = Exact<{
	input?: InputMaybe<PurchaseDataInput>;
}>;

export type PurchaseDataMutation = {
	__typename?: 'RootMutations';
	PurchaseData?: { __typename?: 'PurchaseDataResponse'; message?: string | null } | null;
};

export type BankAccountCreateMutationVariables = Exact<{
	input?: InputMaybe<CreateBankAccountInput>;
}>;

export type BankAccountCreateMutation = {
	__typename?: 'RootMutations';
	BankAccount_Create?: { __typename?: 'CreateBankAccountResponse'; message?: string | null } | null;
};

export type BankAccountDeleteMutationVariables = Exact<{
	input?: InputMaybe<DeleteBankAccountInput>;
}>;

export type BankAccountDeleteMutation = {
	__typename?: 'RootMutations';
	BankAccount_Delete?: { __typename?: 'DeleteBankAccountResponse'; message?: string | null } | null;
};

export type PurchaseAirtimeMutationVariables = Exact<{
	input?: InputMaybe<PurchaseAirtimeInput>;
}>;

export type PurchaseAirtimeMutation = {
	__typename?: 'RootMutations';
	PurchaseAirtime?: { __typename?: 'PurchaseAirtimeResponse'; message?: string | null } | null;
};

export type GetP2pSellersQueryVariables = Exact<{
	input?: InputMaybe<GetSellersInput>;
}>;

export type GetP2pSellersQuery = {
	__typename?: 'RootQuery';
	GetP2PSellers?: {
		__typename?: 'GetP2PSellersResponse';
		name?: string | null;
		amount?: string | null;
		rate?: string | null;
	} | null;
};

export type GetP2pBuyersQueryVariables = Exact<{
	input?: InputMaybe<GetBuyersInput>;
}>;

export type GetP2pBuyersQuery = {
	__typename?: 'RootQuery';
	GetBuyers?: {
		__typename?: 'GetP2PBuyersResponse';
		accountNo?: string | null;
		accountName?: string | null;
	} | null;
};

export type GetUserQueryVariables = Exact<{
	input?: InputMaybe<GetUserInput>;
}>;

export type GetUserQuery = {
	__typename?: 'RootQuery';
	User_GetInfo?: {
		__typename?: 'User_GetInfoResponse';
		message?: string | null;
		walletAddress?: string | null;
		firstName?: string | null;
		lastName?: string | null;
		id?: string | null;
	} | null;
};

export type CreateUserMutationVariables = Exact<{
	input?: InputMaybe<CreateUserInput>;
}>;

export type CreateUserMutation = {
	__typename?: 'RootMutations';
	User_Create?: {
		__typename?: 'User_CreateResponse';
		message?: string | null;
		lastName?: string | null;
		walletAddress?: string | null;
		firstName?: string | null;
		id?: string | null;
	} | null;
};

export const GetCountryDocument = gql`
	query GetCountry {
		CountryList {
			name
		}
	}
`;
export const GetDataPlansDocument = gql`
	query getDataPlans($input: GetDataPlansInput) {
		Airtime_GetDataPlans(input: $input) {
			dataPlans {
				plan
				network
				id
			}
		}
	}
`;
export const KycVerifyBvnDocument = gql`
	mutation kycVerifyBvn($input: Kyc_VerifyBvnInput) {
		Kyc_VerifyBvn(input: $input) {
			message
		}
	}
`;
export const KycVerifyNinDocument = gql`
	mutation kycVerifyNin($input: Kyc_VerifyNinInput) {
		Kyc_VerifyNin(input: $input) {
			message
		}
	}
`;
export const PurchaseDataDocument = gql`
	mutation purchaseData($input: PurchaseDataInput) {
		PurchaseData(input: $input) {
			message
		}
	}
`;
export const BankAccountCreateDocument = gql`
	mutation bankAccountCreate($input: CreateBankAccountInput) {
		BankAccount_Create(input: $input) {
			message
		}
	}
`;
export const BankAccountDeleteDocument = gql`
	mutation bankAccountDelete($input: DeleteBankAccountInput) {
		BankAccount_Delete(input: $input) {
			message
		}
	}
`;
export const PurchaseAirtimeDocument = gql`
	mutation purchaseAirtime($input: PurchaseAirtimeInput) {
		PurchaseAirtime(input: $input) {
			message
		}
	}
`;
export const GetP2pSellersDocument = gql`
	query getP2pSellers($input: GetSellersInput) {
		GetP2PSellers(input: $input) {
			name
			amount
			rate
		}
	}
`;
export const GetP2pBuyersDocument = gql`
	query getP2pBuyers($input: GetBuyersInput) {
		GetBuyers(input: $input) {
			accountNo
			accountName
		}
	}
`;
export const GetUserDocument = gql`
	query getUser($input: GetUserInput) {
		User_GetInfo(input: $input) {
			message
			walletAddress
			firstName
			lastName
			id
		}
	}
`;
export const CreateUserDocument = gql`
	mutation createUser($input: CreateUserInput) {
		User_Create(input: $input) {
			message
			lastName
			walletAddress
			firstName
			id
		}
	}
`;
export type GetCountryQueryStore = OperationStore<GetCountryQuery, GetCountryQueryVariables>;
export type GetDataPlansQueryStore = OperationStore<GetDataPlansQuery, GetDataPlansQueryVariables>;
export type KycVerifyBvnMutationStore = OperationStore<
	KycVerifyBvnMutation,
	KycVerifyBvnMutationVariables
>;
export type KycVerifyNinMutationStore = OperationStore<
	KycVerifyNinMutation,
	KycVerifyNinMutationVariables
>;
export type PurchaseDataMutationStore = OperationStore<
	PurchaseDataMutation,
	PurchaseDataMutationVariables
>;
export type BankAccountCreateMutationStore = OperationStore<
	BankAccountCreateMutation,
	BankAccountCreateMutationVariables
>;
export type BankAccountDeleteMutationStore = OperationStore<
	BankAccountDeleteMutation,
	BankAccountDeleteMutationVariables
>;
export type PurchaseAirtimeMutationStore = OperationStore<
	PurchaseAirtimeMutation,
	PurchaseAirtimeMutationVariables
>;
export type GetP2pSellersQueryStore = OperationStore<
	GetP2pSellersQuery,
	GetP2pSellersQueryVariables
>;
export type GetP2pBuyersQueryStore = OperationStore<GetP2pBuyersQuery, GetP2pBuyersQueryVariables>;
export type GetUserQueryStore = OperationStore<GetUserQuery, GetUserQueryVariables>;
export type CreateUserMutationStore = OperationStore<
	CreateUserMutation,
	CreateUserMutationVariables
>;

// export const GetCountryDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"GetCountry"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"CountryList"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]} as unknown as DocumentNode<GetCountryQuery, GetCountryQueryVariables>;
// export const GetDataPlansDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"getDataPlans"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"GetDataPlansInput"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"Airtime_GetDataPlans"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"dataPlans"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"plan"}},{"kind":"Field","name":{"kind":"Name","value":"network"}},{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]}}]} as unknown as DocumentNode<GetDataPlansQuery, GetDataPlansQueryVariables>;
// export const KycVerifyBvnDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"kycVerifyBvn"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"Kyc_VerifyBvnInput"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"Kyc_VerifyBvn"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"message"}}]}}]}}]} as unknown as DocumentNode<KycVerifyBvnMutation, KycVerifyBvnMutationVariables>;
// export const KycVerifyNinDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"kycVerifyNin"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"Kyc_VerifyNinInput"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"Kyc_VerifyNin"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"message"}}]}}]}}]} as unknown as DocumentNode<KycVerifyNinMutation, KycVerifyNinMutationVariables>;
// export const PurchaseDataDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"purchaseData"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"PurchaseDataInput"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"PurchaseData"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"message"}}]}}]}}]} as unknown as DocumentNode<PurchaseDataMutation, PurchaseDataMutationVariables>;
// export const BankAccountCreateDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"bankAccountCreate"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"CreateBankAccountInput"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"BankAccount_Create"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"message"}}]}}]}}]} as unknown as DocumentNode<BankAccountCreateMutation, BankAccountCreateMutationVariables>;
// export const BankAccountDeleteDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"bankAccountDelete"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"DeleteBankAccountInput"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"BankAccount_Delete"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"message"}}]}}]}}]} as unknown as DocumentNode<BankAccountDeleteMutation, BankAccountDeleteMutationVariables>;
// export const PurchaseAirtimeDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"purchaseAirtime"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"PurchaseAirtimeInput"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"PurchaseAirtime"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"message"}}]}}]}}]} as unknown as DocumentNode<PurchaseAirtimeMutation, PurchaseAirtimeMutationVariables>;
// export const GetP2pSellersDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"getP2pSellers"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"GetSellersInput"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"GetP2PSellers"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"amount"}},{"kind":"Field","name":{"kind":"Name","value":"rate"}}]}}]}}]} as unknown as DocumentNode<GetP2pSellersQuery, GetP2pSellersQueryVariables>;
// export const GetP2pBuyersDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"getP2pBuyers"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"GetBuyersInput"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"GetBuyers"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"accountNo"}},{"kind":"Field","name":{"kind":"Name","value":"accountName"}}]}}]}}]} as unknown as DocumentNode<GetP2pBuyersQuery, GetP2pBuyersQueryVariables>;
// export const GetUserDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"getUser"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"GetUserInput"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"User_GetInfo"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"message"}},{"kind":"Field","name":{"kind":"Name","value":"walletAddress"}},{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<GetUserQuery, GetUserQueryVariables>;
// export const CreateUserDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createUser"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"CreateUserInput"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"User_Create"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"message"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"walletAddress"}},{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<CreateUserMutation, CreateUserMutationVariables>;
