export {
  AVSResponse,
  AccessTokenResponse,
  AccountUpdaterResponse,
  AdditionalAmountRate,
  AdditionalDetails,
  AdditionalTransactionDetails,
  Address,
  Airline,
  AirlineAncillaryServiceCategory,
  AirlineTravelRoute,
  AliPay,
  AliPayPaymentMethod,
  AliPaySaleTransaction,
  Amount,
  AmountComponents,
  Authentication,
  AuthenticationRedirect,
  AuthenticationRedirectParams,
  AuthenticationVerificationRequest,
  BasicResponse,
  Billing,
  BillingAddress,
  BillingAddressPhone,
  CarRental,
  CarRentalExtraCharges,
  Card,
  CardFunction,
  CardInfo,
  CardInfoLookupRequest,
  CardInfoLookupResponse,
  CardVerificationRequest,
  ChinaDomestic,
  ChinaDomesticPaymentMethod,
  ChinaPnRSaleTransaction,
  ClientLocale,
  Contact,
  CreatePaymentToken,
  CurrencyConversion,
  Customer,
  CustomerAddress,
  CustomerAddressPhone,
  DCCExchangeRateRequest,
  Dcc,
  Device,
  DeviceNetworks,
  Document,
  DynamicPricing,
  DynamicPricingExchangeRateRequest,
  ErrorDetails,
  ErrorResponse,
  ExchangeRateRequest,
  ExchangeRateResponse,
  Expiration,
  FraudOrder,
  FraudOrderItems,
  FraudOrderShipToAddress,
  Frequency,
  IndustrySpecificExtensions,
  InstallmentOptions,
  Lodging,
  LodgingExtraCharges,
  Loyalty,
  Merchant,
  MerchantLocation,
  MerchantLocationMerchantAddress,
  ModelError,
  Order,
  OrderErrorResponse,
  OrderResponse,
  PayPal,
  PayPalPaymentMethod,
  Payment,
  PaymentCard,
  PaymentCardCreditTransaction,
  PaymentCardForcedTicketTransaction,
  PaymentCardPayerAuthTransaction,
  PaymentCardPaymentMethod,
  PaymentCardPaymentTokenizationRequest,
  PaymentCardPreAuthTransaction,
  PaymentCardSaleTransaction,
  PaymentFacilitator,
  PaymentIssuerResponse,
  PaymentMethod,
  PaymentMethodDetails,
  PaymentMethodPaymentSchedulesRequest,
  PaymentMethodType,
  PaymentPayMethod,
  PaymentSchedulesErrorResponse,
  PaymentSchedulesRequest,
  PaymentSchedulesResponse,
  PaymentTokenCreditTransaction,
  PaymentTokenDetails,
  PaymentTokenPaymentMethod,
  PaymentTokenPreAuthTransaction,
  PaymentTokenSaleTransaction,
  PaymentTokenizationErrorResponse,
  PaymentTokenizationRequest,
  PaymentTokenizationResponse,
  PaymentUrlErrorResponse,
  PaymentUrlRequest,
  PaymentUrlResponse,
  PaymentVerification3ds,
  PaymentVerificationAvs,
  PaymentVerificationCvv,
  PaypalCreditTransaction,
  PostAuthTransaction,
  PrimaryTransaction,
  ProcessorData,
  PurchaseCards,
  PurchaseCardsLevel2,
  PurchaseCardsLevel3,
  PurchaseCardsLevel3LineItems,
  RecurringPaymentDetails,
  RecurringPaymentDetailsResponse,
  ReferencedOrderPaymentSchedulesRequest,
  ReferencedOrderPaymentTokenizationRequest,
  RequestArgs,
  RequiredError,
  ResponseAmountComponents,
  ResponseType,
  ReturnTransaction,
  ScoreOnlyRequest,
  ScoreOnlyResponse,
  ScoreOnlyResponseFraudScore,
  ScoreOnlyResponseFraudScoreExplanations,
  SecondaryTransaction,
  Secure3dAuthenticationResult,
  Secure3dAuthenticationVerificationRequest,
  Secure3dResponse,
  Sepa,
  SepaMandate,
  SepaPaymentMethod,
  SepaSaleTransaction,
  Shipping,
  SoftDescriptor,
  SplitShipment,
  StoredCredential,
  SubMerchantData,
  SubMerchantSplit,
  Transaction,
  TransactionErrorResponse,
  TransactionOrigin,
  TransactionResponse,
  TransactionType,
  UnionPayAuthenticationRequest,
  UnionPayAuthenticationVerificationRequest,
  UsePaymentToken,
  VoidTransaction,
} from "../openapi/api";

// Wrapper specific models

export type Omit<T, K extends keyof T> = Pick<T, Exclude<keyof T, K>>;

export type ApiField<T, P = any> =
  T extends "region" ? {
    // The region where client wants to process the transaction
    region?: string;
  } :
  T extends "authorization" ? {
    // The access token previously generated with the access-tokens call. Use the format &#39;Bearer {access-token}&#39;.
    authorization?: string;
  } :
  T extends "storeId" ? {
    // An optional outlet ID for clients that support multiple stores in the same developer app
    storeId?: string;
  } :
  T extends "orderId" ? {
    // Gateway order identifier as returned in the parameter orderId
    orderId: string;
  } :
  T extends "transactionId" ? {
    // Gateway transaction identifier as returned in the parameter ipgTransactionId
    transactionId: string;
  } :
  T extends "tokenId" ? {
    // Identifies a payment token
    tokenId: string;
  } :
  T extends "payload" ? {
    payload: P;
  } : unknown;

