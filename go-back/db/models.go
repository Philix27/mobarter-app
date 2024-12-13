// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type Adsstatus string

const (
	AdsstatusUP   Adsstatus = "UP"
	AdsstatusDOWN Adsstatus = "DOWN"
)

func (e *Adsstatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Adsstatus(s)
	case string:
		*e = Adsstatus(s)
	default:
		return fmt.Errorf("unsupported scan type for Adsstatus: %T", src)
	}
	return nil
}

type NullAdsstatus struct {
	Adsstatus Adsstatus
	Valid     bool // Valid is true if Adsstatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAdsstatus) Scan(value interface{}) error {
	if value == nil {
		ns.Adsstatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Adsstatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAdsstatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Adsstatus), nil
}

type Agentstatus string

const (
	AgentstatusTRADE     Agentstatus = "TRADE"
	AgentstatusBLOCKED   Agentstatus = "BLOCKED"
	AgentstatusINACTIVE  Agentstatus = "IN_ACTIVE"
	AgentstatusBANNED    Agentstatus = "BANNED"
	AgentstatusSUSPENDED Agentstatus = "SUSPENDED"
)

func (e *Agentstatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Agentstatus(s)
	case string:
		*e = Agentstatus(s)
	default:
		return fmt.Errorf("unsupported scan type for Agentstatus: %T", src)
	}
	return nil
}

type NullAgentstatus struct {
	Agentstatus Agentstatus
	Valid       bool // Valid is true if Agentstatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAgentstatus) Scan(value interface{}) error {
	if value == nil {
		ns.Agentstatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Agentstatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAgentstatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Agentstatus), nil
}

type Currencypair string

const (
	CurrencypairCUSDNGN Currencypair = "CUSD_NGN"
	CurrencypairUSDTNGN Currencypair = "USDT_NGN"
	CurrencypairCELONGN Currencypair = "CELO_NGN"
)

func (e *Currencypair) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Currencypair(s)
	case string:
		*e = Currencypair(s)
	default:
		return fmt.Errorf("unsupported scan type for Currencypair: %T", src)
	}
	return nil
}

type NullCurrencypair struct {
	Currencypair Currencypair
	Valid        bool // Valid is true if Currencypair is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCurrencypair) Scan(value interface{}) error {
	if value == nil {
		ns.Currencypair, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Currencypair.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCurrencypair) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Currencypair), nil
}

type Kycvalidationstatus string

const (
	KycvalidationstatusNONE     Kycvalidationstatus = "NONE"
	KycvalidationstatusAPPROVED Kycvalidationstatus = "APPROVED"
	KycvalidationstatusPENDING  Kycvalidationstatus = "PENDING"
	KycvalidationstatusDECLINED Kycvalidationstatus = "DECLINED"
)

func (e *Kycvalidationstatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Kycvalidationstatus(s)
	case string:
		*e = Kycvalidationstatus(s)
	default:
		return fmt.Errorf("unsupported scan type for Kycvalidationstatus: %T", src)
	}
	return nil
}

type NullKycvalidationstatus struct {
	Kycvalidationstatus Kycvalidationstatus
	Valid               bool // Valid is true if Kycvalidationstatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullKycvalidationstatus) Scan(value interface{}) error {
	if value == nil {
		ns.Kycvalidationstatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Kycvalidationstatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullKycvalidationstatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Kycvalidationstatus), nil
}

type Orderstatus string

const (
	OrderstatusCREATED   Orderstatus = "CREATED"
	OrderstatusCOMPLETED Orderstatus = "COMPLETED"
	OrderstatusAPPEAL    Orderstatus = "APPEAL"
)

func (e *Orderstatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Orderstatus(s)
	case string:
		*e = Orderstatus(s)
	default:
		return fmt.Errorf("unsupported scan type for Orderstatus: %T", src)
	}
	return nil
}

type NullOrderstatus struct {
	Orderstatus Orderstatus
	Valid       bool // Valid is true if Orderstatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullOrderstatus) Scan(value interface{}) error {
	if value == nil {
		ns.Orderstatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Orderstatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullOrderstatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Orderstatus), nil
}

type Ordertype string

const (
	OrdertypeBUY  Ordertype = "BUY"
	OrdertypeSELL Ordertype = "SELL"
)

func (e *Ordertype) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Ordertype(s)
	case string:
		*e = Ordertype(s)
	default:
		return fmt.Errorf("unsupported scan type for Ordertype: %T", src)
	}
	return nil
}

type NullOrdertype struct {
	Ordertype Ordertype
	Valid     bool // Valid is true if Ordertype is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullOrdertype) Scan(value interface{}) error {
	if value == nil {
		ns.Ordertype, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Ordertype.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullOrdertype) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Ordertype), nil
}

type Tradetype string

const (
	TradetypeBUY  Tradetype = "BUY"
	TradetypeSELL Tradetype = "SELL"
)

func (e *Tradetype) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Tradetype(s)
	case string:
		*e = Tradetype(s)
	default:
		return fmt.Errorf("unsupported scan type for Tradetype: %T", src)
	}
	return nil
}

type NullTradetype struct {
	Tradetype Tradetype
	Valid     bool // Valid is true if Tradetype is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullTradetype) Scan(value interface{}) error {
	if value == nil {
		ns.Tradetype, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Tradetype.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullTradetype) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Tradetype), nil
}

type Transactionpurpose string

const (
	TransactionpurposeAIRTIME  Transactionpurpose = "AIRTIME"
	TransactionpurposeDATA     Transactionpurpose = "DATA"
	TransactionpurposeTRANSFER Transactionpurpose = "TRANSFER"
	TransactionpurposeWITHDRAW Transactionpurpose = "WITHDRAW"
	TransactionpurposeREFUND   Transactionpurpose = "REFUND"
)

func (e *Transactionpurpose) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Transactionpurpose(s)
	case string:
		*e = Transactionpurpose(s)
	default:
		return fmt.Errorf("unsupported scan type for Transactionpurpose: %T", src)
	}
	return nil
}

type NullTransactionpurpose struct {
	Transactionpurpose Transactionpurpose
	Valid              bool // Valid is true if Transactionpurpose is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullTransactionpurpose) Scan(value interface{}) error {
	if value == nil {
		ns.Transactionpurpose, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Transactionpurpose.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullTransactionpurpose) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Transactionpurpose), nil
}

type Ad struct {
	ID              int64
	AgentID         pgtype.Int4
	PaymentMethodID pgtype.Int4
	AdStatus        NullAdsstatus
	CurrencyPair    NullCurrencypair
	LimitUpper      pgtype.Numeric
	LimitLower      pgtype.Numeric
	Rate            pgtype.Numeric
	Instructions    pgtype.Text
	CreatedAt       pgtype.Timestamp
	UpdatedAt       pgtype.Timestamp
}

type Agent struct {
	ID           int64
	UserID       pgtype.Int4
	AgentStatus  NullAgentstatus
	TradeCount   pgtype.Int4
	NickName     pgtype.Text
	Instructions []string
	CreatedAt    pgtype.Timestamp
	UpdatedAt    pgtype.Timestamp
}

type Airtime struct {
	ID        int64
	Network   string
	Country   string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type Author struct {
	ID   int64
	Name string
	Bio  pgtype.Text
}

type BankAccount struct {
	ID          int64
	BankName    string
	AccountName string
	AccountNo   int32
	UserID      pgtype.Int4
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

type Country struct {
	ID        int64
	Name      string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type CustomerCare struct {
	ID        int64
	Message   string
	UserID    pgtype.Text
	RoomID    pgtype.Text
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type DataPlan struct {
	ID          int64
	Network     string
	Country     string
	Amount      int32
	Duration    string
	Plan        string
	IsSupported string
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

type Kana struct {
	ID       int32
	Username string
}

type KycCredential struct {
	ID                           int32
	UserID                       int32
	Bvn                          pgtype.Text
	BvnStatus                    NullKycvalidationstatus
	BvnVerificationDate          pgtype.Timestamp
	BvnVerifiedBy                pgtype.Text
	Nin                          pgtype.Text
	NinStatus                    NullKycvalidationstatus
	NinVerificationDate          pgtype.Timestamp
	NinVerifiedBy                pgtype.Text
	HouseAddress                 pgtype.Text
	HouseAddressStatus           NullKycvalidationstatus
	HouseAddressVerificationDate pgtype.Timestamp
	HouseAddressVerifiedBy       pgtype.Text
	UtilityBill                  pgtype.Text
	UtilityBillStatus            NullKycvalidationstatus
	UtilityBillVerificationDate  pgtype.Timestamp
	UtilityBillVerifiedBy        pgtype.Text
	CreatedAt                    pgtype.Timestamp
	UpdatedAt                    pgtype.Timestamp
}

type Newsletter struct {
	ID        int64
	Email     string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type Order struct {
	ID           int64
	UserID       pgtype.Int4
	AgentID      pgtype.Int4
	OrderType    NullOrdertype
	OrderStatus  NullOrderstatus
	CurrencyPair NullCurrencypair
	Amount       pgtype.Numeric
	Rate         pgtype.Numeric
	CreatedAt    pgtype.Timestamp
	UpdatedAt    pgtype.Timestamp
}

type Transaction struct {
	ID        int64
	Name      string
	Purpose   Transactionpurpose
	Amount    int32
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type User struct {
	ID        int64
	Wallets   []string
	FirstName pgtype.Text
	LastName  pgtype.Text
	Dob       pgtype.Timestamp
	Email     string
	Phone     pgtype.Text
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type Waitlist struct {
	ID        int64
	Email     string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}
