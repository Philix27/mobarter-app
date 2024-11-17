import type { OperationStore } from '@urql/svelte';
import gql from 'graphql-tag';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
};

export type Auth_SentOtpResponse = {
  __typename?: 'Auth_SentOtpResponse';
  message?: Maybe<Scalars['String']['output']>;
};

export type CountryResponse = {
  __typename?: 'CountryResponse';
  flag?: Maybe<Scalars['String']['output']>;
  name?: Maybe<Scalars['String']['output']>;
};

export type CreateBankAccountResponse = {
  __typename?: 'CreateBankAccountResponse';
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

export type Kyc_GetLevelResponse = {
  __typename?: 'Kyc_GetLevelResponse';
  accountName?: Maybe<Scalars['String']['output']>;
  accountNo?: Maybe<Scalars['String']['output']>;
  bankName?: Maybe<Scalars['String']['output']>;
};

export type Kyc_VerifyBvnResponse = {
  __typename?: 'Kyc_VerifyBvnResponse';
  message?: Maybe<Scalars['String']['output']>;
};

export type Kyc_VerifyEmailResponse = {
  __typename?: 'Kyc_VerifyEmailResponse';
  message?: Maybe<Scalars['String']['output']>;
};

export type Kyc_VerifyNinResponse = {
  __typename?: 'Kyc_VerifyNinResponse';
  message?: Maybe<Scalars['String']['output']>;
};

export type Kyc_VerifyPhoneResponse = {
  __typename?: 'Kyc_VerifyPhoneResponse';
  message?: Maybe<Scalars['String']['output']>;
};

export type PurchaseAirtimeResponse = {
  __typename?: 'PurchaseAirtimeResponse';
  message?: Maybe<Scalars['String']['output']>;
};

export type PurchaseDataResponse = {
  __typename?: 'PurchaseDataResponse';
  message?: Maybe<Scalars['String']['output']>;
};

export type RootMutations = {
  __typename?: 'RootMutations';
  Auth_SentOtp?: Maybe<Auth_SentOtpResponse>;
  BankAccount_Create?: Maybe<CreateBankAccountResponse>;
  JoinWaitList?: Maybe<WaitListResponse>;
  Kyc_VerifyBvn?: Maybe<Kyc_VerifyBvnResponse>;
  Kyc_VerifyEmail?: Maybe<Kyc_VerifyEmailResponse>;
  Kyc_VerifyNin?: Maybe<Kyc_VerifyNinResponse>;
  Kyc_VerifyPhone?: Maybe<Kyc_VerifyPhoneResponse>;
  PurchaseAirtime?: Maybe<PurchaseAirtimeResponse>;
  PurchaseData?: Maybe<PurchaseDataResponse>;
};


export type RootMutationsAuth_SentOtpArgs = {
  email: Scalars['String']['input'];
  issuerAddress: Scalars['String']['input'];
  network: Scalars['String']['input'];
  phone: Scalars['String']['input'];
  transactionHash: Scalars['Int']['input'];
};


export type RootMutationsBankAccount_CreateArgs = {
  accountName: Scalars['String']['input'];
  accountNo: Scalars['String']['input'];
  bankName: Scalars['String']['input'];
};


export type RootMutationsJoinWaitListArgs = {
  country: Scalars['String']['input'];
  email: Scalars['String']['input'];
  firstName: Scalars['String']['input'];
  lastName: Scalars['String']['input'];
};


export type RootMutationsKyc_VerifyBvnArgs = {
  bvn: Scalars['String']['input'];
};


export type RootMutationsKyc_VerifyEmailArgs = {
  email: Scalars['String']['input'];
  otp: Scalars['String']['input'];
  token: Scalars['String']['input'];
};


export type RootMutationsKyc_VerifyNinArgs = {
  nin: Scalars['String']['input'];
};


export type RootMutationsKyc_VerifyPhoneArgs = {
  otp: Scalars['String']['input'];
  phone: Scalars['String']['input'];
  token: Scalars['String']['input'];
};


export type RootMutationsPurchaseAirtimeArgs = {
  amount: Scalars['String']['input'];
  issuerAddress: Scalars['String']['input'];
  network: Scalars['String']['input'];
  phone: Scalars['String']['input'];
  transactionHash: Scalars['Int']['input'];
};


export type RootMutationsPurchaseDataArgs = {
  amount: Scalars['String']['input'];
  issuerAddress: Scalars['String']['input'];
  network: Scalars['String']['input'];
  phone: Scalars['String']['input'];
  transactionHash: Scalars['Int']['input'];
};

export type RootQuery = {
  __typename?: 'RootQuery';
  BankAccount_Get?: Maybe<GetBankAccountResponse>;
  BankAccount_GetAll?: Maybe<GetAllBankAccountResponse>;
  CountryList?: Maybe<CountryResponse>;
  Kyc_GetLevel?: Maybe<Kyc_GetLevelResponse>;
  Transactions_GetAll?: Maybe<GetAllTransactionsResponse>;
  Transactions_GetOne?: Maybe<GetAllTransactionsResponse>;
  User_GetInfo?: Maybe<UserResponse>;
};

export type UserResponse = {
  __typename?: 'UserResponse';
  country?: Maybe<Scalars['String']['output']>;
  firstName?: Maybe<Scalars['String']['output']>;
  lastName?: Maybe<Scalars['String']['output']>;
  walletAddress?: Maybe<Scalars['String']['output']>;
};

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

export type WaitListResponse = {
  __typename?: 'waitListResponse';
  message?: Maybe<Scalars['String']['output']>;
};
