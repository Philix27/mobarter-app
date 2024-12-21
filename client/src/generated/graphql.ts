import type { OperationStore } from '@urql/svelte';
import gql from 'graphql-tag';
// import { TypedDocumentNode as DocumentN
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

export type Admin_ApproveBvnInput = {
	bvn: Scalars['String']['input'];
	isApproved: Scalars['Boolean']['input'];
	userId: Scalars['String']['input'];
};

export type Admin_ApproveBvnResponse = {
	__typename?: 'Admin_ApproveBvnResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export type Admin_ApproveNinInput = {
	isApproved: Scalars['Boolean']['input'];
	nin: Scalars['String']['input'];
	userId: Scalars['String']['input'];
};

export type Admin_ApproveNinResponse = {
	__typename?: 'Admin_ApproveNinResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export type Admin_ApprovePassportInput = {
	isApproved: Scalars['Boolean']['input'];
	passport: Scalars['String']['input'];
	userId: Scalars['String']['input'];
};

export type Admin_ApprovePassportResponse = {
	__typename?: 'Admin_ApprovePassportResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export type Admin_ApproveResidentialAddressInput = {
	address: Scalars['String']['input'];
	isApproved: Scalars['Boolean']['input'];
	userId: Scalars['String']['input'];
};

export type Admin_ApproveResidentialAddressResponse = {
	__typename?: 'Admin_ApproveResidentialAddressResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export type Admin_BlockAccountInput = {
	isBlocked: Scalars['Boolean']['input'];
	reason: Scalars['String']['input'];
	userId: Scalars['String']['input'];
};

export type Admin_BlockAccountResponse = {
	__typename?: 'Admin_BlockAccountResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export type Admin_GetBlockedAccountsInput = {
	userId?: InputMaybe<Scalars['String']['input']>;
};

export type Admin_GetBlockedAccountsResponse = {
	__typename?: 'Admin_GetBlockedAccountsResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export type Admin_GetUserBankAccountsInput = {
	userId?: InputMaybe<Scalars['String']['input']>;
};

export type Admin_GetUserBankAccountsResponse = {
	__typename?: 'Admin_GetUserBankAccountsResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export type Admin_GetUserTransactionsInput = {
	userId?: InputMaybe<Scalars['String']['input']>;
};

export type Admin_GetUserTransactionsResponse = {
	__typename?: 'Admin_GetUserTransactionsResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export type Admin_GetUsersInput = {
	userId?: InputMaybe<Scalars['String']['input']>;
};

export type Admin_GetUsersResponse = {
	__typename?: 'Admin_GetUsersResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export type Admin_LoginInput = {
	email: Scalars['String']['input'];
	password: Scalars['String']['input'];
};

export type Admin_LoginResponse = {
	__typename?: 'Admin_LoginResponse';
	message?: Maybe<Scalars['String']['output']>;
	token?: Maybe<Scalars['String']['output']>;
};

export type Admin_SendEmail = {
	email: Scalars['String']['input'];
	message: Scalars['String']['input'];
};

export type Admin_SendEmailResponse = {
	__typename?: 'Admin_SendEmailResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export type Admin_SetRatesInput = {
	email: Scalars['String']['input'];
	message: Scalars['String']['input'];
};

export type Admin_SetRatesResponse = {
	__typename?: 'Admin_SetRatesResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export enum AdvertCountries {
	Ghana = 'GHANA',
	Kenya = 'KENYA',
	Nigeria = 'NIGERIA'
}

export enum AdvertPaymentMethods {
	BankAccount = 'BANK_ACCOUNT',
	MobileMoney = 'MOBILE_MONEY',
	Transfer = 'TRANSFER'
}

export enum AdvertStatus {
	Down = 'DOWN',
	Up = 'UP'
}

export type Advert_CreateInput = {
	password: Scalars['String']['input'];
	token: Scalars['String']['input'];
};

export type Advert_CreateResponse = {
	__typename?: 'Advert_CreateResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export type Advert_DeleteInput = {
	adId: Scalars['String']['input'];
	reason?: InputMaybe<Scalars['String']['input']>;
	token: Scalars['String']['input'];
};

export type Advert_DeleteResponse = {
	__typename?: 'Advert_DeleteResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export type Advert_GetAllInput = {
	country?: InputMaybe<AdvertCountries>;
	paymentMethod?: InputMaybe<AdvertPaymentMethods>;
	token: Scalars['String']['input'];
};

export type Advert_GetAllResponse = {
	__typename?: 'Advert_GetAllResponse';
	currency_pair?: Maybe<Scalars['String']['output']>;
	id?: Maybe<Scalars['String']['output']>;
	instructions?: Maybe<Scalars['String']['output']>;
	limit_lower?: Maybe<Scalars['String']['output']>;
	limit_upper?: Maybe<Scalars['String']['output']>;
	message?: Maybe<Scalars['String']['output']>;
	rate?: Maybe<Scalars['String']['output']>;
};

export type Advert_GetOneInput = {
	id: Scalars['String']['input'];
	token: Scalars['String']['input'];
};

export type Advert_GetOneResponse = {
	__typename?: 'Advert_GetOneResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export type Advert_UpdateStatusInput = {
	status?: InputMaybe<AdvertStatus>;
	token: Scalars['String']['input'];
};

export type Advert_UpdateStatusResponse = {
	__typename?: 'Advert_UpdateStatusResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export type Auth_SendEmailOtpResponse = {
	__typename?: 'Auth_SendEmailOtpResponse';
	message?: Maybe<Scalars['String']['output']>;
	token?: Maybe<Scalars['String']['output']>;
};

export type Auth_SendPhoneOtpResponse = {
	__typename?: 'Auth_SendPhoneOtpResponse';
	message?: Maybe<Scalars['String']['output']>;
	token?: Maybe<Scalars['String']['output']>;
};

export type Auth_SentOtpResponse = {
	__typename?: 'Auth_SentOtpResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export type Auth_VerifyOtpInput = {
	email: Scalars['String']['input'];
	otp: Scalars['String']['input'];
	token: Scalars['String']['input'];
};

export type Auth_VerifyOtpResponse = {
	__typename?: 'Auth_VerifyOtpResponse';
	message?: Maybe<Scalars['String']['output']>;
	token?: Maybe<Scalars['String']['output']>;
};

export type Country_Response = {
	__typename?: 'Country_Response';
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
	email: Scalars['String']['input'];
	password: Scalars['String']['input'];
	token: Scalars['String']['input'];
};

export type DataPlans = {
	__typename?: 'DataPlans';
	amount?: Maybe<Scalars['String']['output']>;
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

export type Issues = {
	__typename?: 'Issues';
	id?: Maybe<Scalars['String']['output']>;
	title?: Maybe<Scalars['String']['output']>;
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

export type Messages = {
	__typename?: 'Messages';
	from?: Maybe<Scalars['String']['output']>;
	msg?: Maybe<Scalars['String']['output']>;
};

export enum NetworkProviders {
	/** Regular user with limited access */
	Airtel = 'AIRTEL',
	/** Guest user with minimal access */
	Glo = 'GLO',
	/** Administrator with full access */
	Mtn = 'MTN'
}

export enum OrderStatus {
	Appeal = 'APPEAL',
	Completed = 'COMPLETED',
	Pending = 'PENDING'
}

export type Order_CloseInput = {
	orderId: Scalars['String']['input'];
	token: Scalars['String']['input'];
};

export type Order_CloseResponse = {
	__typename?: 'Order_CloseResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export type Order_CreateInput = {
	password: Scalars['String']['input'];
	token: Scalars['String']['input'];
};

export type Order_CreateResponse = {
	__typename?: 'Order_CreateResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export type Order_FeedbackInput = {
	comment: Scalars['String']['input'];
	rating: Scalars['String']['input'];
	token: Scalars['String']['input'];
};

export type Order_FeedbackResponse = {
	__typename?: 'Order_FeedbackResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export type Order_GetAllInput = {
	status?: InputMaybe<OrderStatus>;
};

export type Order_GetAllResponse = {
	__typename?: 'Order_GetAllResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export type Order_GetOneInput = {
	orderId: Scalars['String']['input'];
	token: Scalars['String']['input'];
};

export type Order_GetOneResponse = {
	__typename?: 'Order_GetOneResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export type Order_UpdateStatusInput = {
	status?: InputMaybe<OrderStatus>;
	token: Scalars['String']['input'];
};

export type Order_UpdateStatusResponse = {
	__typename?: 'Order_UpdateStatusResponse';
	message?: Maybe<Scalars['String']['output']>;
};

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
	Admin_ApproveBvn?: Maybe<Admin_ApproveBvnResponse>;
	Admin_ApproveNin?: Maybe<Admin_ApproveNinResponse>;
	Admin_ApprovePassport?: Maybe<Admin_ApprovePassportResponse>;
	Admin_ApproveResidentialAddress?: Maybe<Admin_ApproveResidentialAddressResponse>;
	Admin_BlockAccount?: Maybe<Admin_BlockAccountResponse>;
	Admin_Login?: Maybe<Admin_LoginResponse>;
	Admin_SendEmail?: Maybe<Admin_SendEmailResponse>;
	Admin_SetRates?: Maybe<Admin_SetRatesResponse>;
	Advert_Create?: Maybe<Advert_CreateResponse>;
	Advert_Delete?: Maybe<Advert_DeleteResponse>;
	Advert_UpdateStatus?: Maybe<Advert_UpdateStatusResponse>;
	Auth_SendEmailOtp?: Maybe<Auth_SendEmailOtpResponse>;
	Auth_SendPhoneOtp?: Maybe<Auth_SendPhoneOtpResponse>;
	Auth_SentOtp?: Maybe<Auth_SentOtpResponse>;
	Auth_VerifyOtp?: Maybe<Auth_VerifyOtpResponse>;
	BankAccount_Create?: Maybe<CreateBankAccountResponse>;
	BankAccount_Delete?: Maybe<DeleteBankAccountResponse>;
	JoinWaitList?: Maybe<WaitListResponse>;
	Kyc_VerifyBvn?: Maybe<Kyc_VerifyBvnResponse>;
	Kyc_VerifyEmail?: Maybe<Kyc_VerifyEmailResponse>;
	Kyc_VerifyNin?: Maybe<Kyc_VerifyNinResponse>;
	Kyc_VerifyPhone?: Maybe<Kyc_VerifyPhoneResponse>;
	Orders_Close?: Maybe<Order_CloseResponse>;
	Orders_CreateAdvert?: Maybe<Order_CreateResponse>;
	Orders_Feedback?: Maybe<Order_FeedbackResponse>;
	Orders_UpdateStatus?: Maybe<Order_UpdateStatusResponse>;
	PurchaseAirtime?: Maybe<PurchaseAirtimeResponse>;
	PurchaseData?: Maybe<PurchaseDataResponse>;
	Support_SendMessage?: Maybe<Support_SendMessageResponse>;
	User_AddBvn?: Maybe<User_AddBvnResponse>;
	User_AddNin?: Maybe<User_AddNinResponse>;
	User_Create?: Maybe<User_CreateResponse>;
	User_PhoneSendOtp?: Maybe<User_PhoneSendOtpResponse>;
	User_PhoneVerifyOtp?: Maybe<User_PhoneVerifyOtpResponse>;
	User_ResetPassword?: Maybe<User_ResetPasswordResponse>;
};

export type RootMutationsAdmin_ApproveBvnArgs = {
	input?: InputMaybe<Admin_ApproveBvnInput>;
};

export type RootMutationsAdmin_ApproveNinArgs = {
	input?: InputMaybe<Admin_ApproveNinInput>;
};

export type RootMutationsAdmin_ApprovePassportArgs = {
	input?: InputMaybe<Admin_ApprovePassportInput>;
};

export type RootMutationsAdmin_ApproveResidentialAddressArgs = {
	input?: InputMaybe<Admin_ApproveResidentialAddressInput>;
};

export type RootMutationsAdmin_BlockAccountArgs = {
	input?: InputMaybe<Admin_BlockAccountInput>;
};

export type RootMutationsAdmin_LoginArgs = {
	input?: InputMaybe<Admin_LoginInput>;
};

export type RootMutationsAdmin_SendEmailArgs = {
	input?: InputMaybe<Admin_SendEmail>;
};

export type RootMutationsAdmin_SetRatesArgs = {
	input?: InputMaybe<Admin_SetRatesInput>;
};

export type RootMutationsAdvert_CreateArgs = {
	input?: InputMaybe<Advert_CreateInput>;
};

export type RootMutationsAdvert_DeleteArgs = {
	input?: InputMaybe<Advert_DeleteInput>;
};

export type RootMutationsAdvert_UpdateStatusArgs = {
	input?: InputMaybe<Advert_UpdateStatusInput>;
};

export type RootMutationsAuth_SendEmailOtpArgs = {
	input?: InputMaybe<SendEmailOtpInput>;
};

export type RootMutationsAuth_SendPhoneOtpArgs = {
	input?: InputMaybe<SendPhoneOtpInput>;
};

export type RootMutationsAuth_SentOtpArgs = {
	email: Scalars['String']['input'];
	issuerAddress: Scalars['String']['input'];
	network: Scalars['String']['input'];
	phone: Scalars['String']['input'];
	transactionHash: Scalars['Int']['input'];
};

export type RootMutationsAuth_VerifyOtpArgs = {
	input?: InputMaybe<Auth_VerifyOtpInput>;
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

export type RootMutationsOrders_CloseArgs = {
	input?: InputMaybe<Order_CloseInput>;
};

export type RootMutationsOrders_CreateAdvertArgs = {
	input?: InputMaybe<Order_CreateInput>;
};

export type RootMutationsOrders_FeedbackArgs = {
	input?: InputMaybe<Order_FeedbackInput>;
};

export type RootMutationsOrders_UpdateStatusArgs = {
	input?: InputMaybe<Order_UpdateStatusInput>;
};

export type RootMutationsPurchaseAirtimeArgs = {
	input?: InputMaybe<PurchaseAirtimeInput>;
};

export type RootMutationsPurchaseDataArgs = {
	input?: InputMaybe<PurchaseDataInput>;
};

export type RootMutationsSupport_SendMessageArgs = {
	input?: InputMaybe<Support_SendMessageInput>;
};

export type RootMutationsUser_AddBvnArgs = {
	input?: InputMaybe<User_AddBvnInput>;
};

export type RootMutationsUser_AddNinArgs = {
	input?: InputMaybe<User_AddNinInput>;
};

export type RootMutationsUser_CreateArgs = {
	input?: InputMaybe<CreateUserInput>;
};

export type RootMutationsUser_PhoneSendOtpArgs = {
	input?: InputMaybe<User_PhoneSendOtpInput>;
};

export type RootMutationsUser_PhoneVerifyOtpArgs = {
	input?: InputMaybe<User_PhoneVerifyOtpInput>;
};

export type RootMutationsUser_ResetPasswordArgs = {
	input?: InputMaybe<User_ResetPasswordInput>;
};

export type RootQuery = {
	__typename?: 'RootQuery';
	Admin_GetBlockedAccounts?: Maybe<Admin_GetBlockedAccountsResponse>;
	Admin_GetUserBankAccounts?: Maybe<Admin_GetUserBankAccountsResponse>;
	Admin_GetUserTransactions?: Maybe<Admin_GetUserTransactionsResponse>;
	Admin_GetUsers?: Maybe<Admin_GetUsersResponse>;
	Advert_GetAll?: Maybe<Advert_GetAllResponse>;
	Advert_GetOne?: Maybe<Advert_GetOneResponse>;
	Airtime_GetDataPlans?: Maybe<GetDataPlansResponse>;
	BankAccount_Get?: Maybe<GetBankAccountResponse>;
	BankAccount_GetAll?: Maybe<GetAllBankAccountResponse>;
	CountryList?: Maybe<Country_Response>;
	Kyc_GetLevel?: Maybe<Kyc_GetLevelResponse>;
	Orders_GetAll?: Maybe<Order_GetAllResponse>;
	Orders_GetOne?: Maybe<Order_GetOneResponse>;
	P2P_GetBuyers?: Maybe<GetP2PBuyersResponse>;
	P2P_GetSellers?: Maybe<GetP2PSellersResponse>;
	Support_GetAllMessages?: Maybe<Support_GetAllMessagesResponse>;
	Support_GetIssueTopics?: Maybe<Support_GetIssueTopicsResponse>;
	Transactions_GetAll?: Maybe<Transactions_GetAllResponse>;
	Transactions_GetOne?: Maybe<Transactions_GetOneResponse>;
	User_GetInfo?: Maybe<User_GetInfoResponse>;
};

export type RootQueryAdmin_GetBlockedAccountsArgs = {
	input?: InputMaybe<Admin_GetBlockedAccountsInput>;
};

export type RootQueryAdmin_GetUserBankAccountsArgs = {
	input?: InputMaybe<Admin_GetUserBankAccountsInput>;
};

export type RootQueryAdmin_GetUserTransactionsArgs = {
	input?: InputMaybe<Admin_GetUserTransactionsInput>;
};

export type RootQueryAdmin_GetUsersArgs = {
	input?: InputMaybe<Admin_GetUsersInput>;
};

export type RootQueryAdvert_GetAllArgs = {
	input?: InputMaybe<Advert_GetAllInput>;
};

export type RootQueryAdvert_GetOneArgs = {
	input?: InputMaybe<Advert_GetOneInput>;
};

export type RootQueryAirtime_GetDataPlansArgs = {
	input?: InputMaybe<GetDataPlansInput>;
};

export type RootQueryOrders_GetAllArgs = {
	input?: InputMaybe<Order_GetAllInput>;
};

export type RootQueryOrders_GetOneArgs = {
	input?: InputMaybe<Order_GetOneInput>;
};

export type RootQueryP2P_GetBuyersArgs = {
	input?: InputMaybe<GetBuyersInput>;
};

export type RootQueryP2P_GetSellersArgs = {
	input?: InputMaybe<GetSellersInput>;
};

export type RootQuerySupport_GetAllMessagesArgs = {
	input?: InputMaybe<Support_GetAllMessagesInput>;
};

export type RootQuerySupport_GetIssueTopicsArgs = {
	input?: InputMaybe<Support_GetIssueTopicsInput>;
};

export type RootQueryTransactions_GetAllArgs = {
	input?: InputMaybe<Transactions_GetAllInput>;
};

export type RootQueryTransactions_GetOneArgs = {
	input?: InputMaybe<Transactions_GetOneInput>;
};

export type RootQueryUser_GetInfoArgs = {
	input?: InputMaybe<GetUserInput>;
};

export type SendEmailOtpInput = {
	email?: InputMaybe<Scalars['String']['input']>;
};

export type SendPhoneOtpInput = {
	country: Scalars['String']['input'];
	phone: Scalars['String']['input'];
};

export type Support_GetAllMessagesInput = {
	orderId: Scalars['String']['input'];
	token: Scalars['String']['input'];
};

export type Support_GetAllMessagesResponse = {
	__typename?: 'Support_GetAllMessagesResponse';
	message?: Maybe<Scalars['String']['output']>;
	messages?: Maybe<Array<Maybe<Messages>>>;
};

export type Support_GetIssueTopicsInput = {
	token: Scalars['String']['input'];
};

export type Support_GetIssueTopicsResponse = {
	__typename?: 'Support_GetIssueTopicsResponse';
	issues?: Maybe<Array<Maybe<Issues>>>;
	message?: Maybe<Scalars['String']['output']>;
};

export type Support_SendMessageInput = {
	message: Scalars['String']['input'];
	token: Scalars['String']['input'];
};

export type Support_SendMessageResponse = {
	__typename?: 'Support_SendMessageResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export type Transactions_GetAllInput = {
	orderId: Scalars['String']['input'];
	token: Scalars['String']['input'];
};

export type Transactions_GetAllResponse = {
	__typename?: 'Transactions_GetAllResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export type Transactions_GetOneInput = {
	orderId: Scalars['String']['input'];
	token: Scalars['String']['input'];
};

export type Transactions_GetOneResponse = {
	__typename?: 'Transactions_GetOneResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export type User = {
	__typename?: 'User';
	id?: Maybe<Scalars['String']['output']>;
	name?: Maybe<Scalars['String']['output']>;
};

export type User_AddBvnInput = {
	bvn: Scalars['String']['input'];
	token: Scalars['String']['input'];
};

export type User_AddBvnResponse = {
	__typename?: 'User_AddBvnResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export type User_AddNinInput = {
	nin: Scalars['String']['input'];
	token: Scalars['String']['input'];
};

export type User_AddNinResponse = {
	__typename?: 'User_AddNinResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export type User_CreateResponse = {
	__typename?: 'User_CreateResponse';
	email?: Maybe<Scalars['String']['output']>;
	message?: Maybe<Scalars['String']['output']>;
	userId?: Maybe<Scalars['String']['output']>;
};

export type User_GetInfoResponse = {
	__typename?: 'User_GetInfoResponse';
	firstName?: Maybe<Scalars['String']['output']>;
	id?: Maybe<Scalars['String']['output']>;
	lastName?: Maybe<Scalars['String']['output']>;
	message?: Maybe<Scalars['String']['output']>;
	walletAddress?: Maybe<Scalars['String']['output']>;
};

export type User_PhoneSendOtpInput = {
	bvn: Scalars['String']['input'];
	token: Scalars['String']['input'];
};

export type User_PhoneSendOtpResponse = {
	__typename?: 'User_PhoneSendOtpResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export type User_PhoneVerifyOtpInput = {
	bvn: Scalars['String']['input'];
	token: Scalars['String']['input'];
};

export type User_PhoneVerifyOtpResponse = {
	__typename?: 'User_PhoneVerifyOtpResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export type User_ResetPasswordInput = {
	email: Scalars['String']['input'];
	password: Scalars['String']['input'];
	token: Scalars['String']['input'];
};

export type User_ResetPasswordResponse = {
	__typename?: 'User_ResetPasswordResponse';
	message?: Maybe<Scalars['String']['output']>;
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

export type WaitListResponse = {
	__typename?: 'waitListResponse';
	message?: Maybe<Scalars['String']['output']>;
};

export type GetCountryQueryVariables = Exact<{ [key: string]: never }>;

export type GetCountryQuery = {
	__typename?: 'RootQuery';
	CountryList?: { __typename?: 'Country_Response'; name?: string | null } | null;
};

export const GetCountryDocument = gql`
	query GetCountry {
		CountryList {
			name
		}
	}
`;
export type GetCountryQueryStore = OperationStore<GetCountryQuery, GetCountryQueryVariables>;
