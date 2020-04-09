package models

import (
	"encoding/json"
	"github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"time"
)

type User struct {
	Id            uuid.UUID       // default: gen_random_uuid()
	CreatedAt     time.Time       // default: now()
	UpdatedAt     time.Time       // default: now()
	Name          string          // default:
	Mobile        *string         // default:
	Email         *string         // default:
	Qq            *string         // default:
	Wechat        *string         // default:
	Password      string          // default:
	Role          string          // default: 'user'::text
	Enabled       bool            // default: true
	Balance       decimal.Decimal // default: 0
	FrozenBalance decimal.Decimal // default: 0
	ParentId      *uuid.UUID      // default:
	Realname      *string         // default:
	IdentityNo    *string         // default:
	PromoCode     string          // default:
}

type Contact struct {
	Id        uuid.UUID // default: gen_random_uuid()
	CreatedAt time.Time // default: now()
	UpdatedAt time.Time // default: now()
	Platform  string    // default:
	Title     string    // default:
	Account   string    // default:
}

type FinanceSetting struct {
	Id                uuid.UUID       // default: gen_random_uuid()
	CreatedAt         time.Time       // default: now()
	UpdatedAt         time.Time       // default: now()
	MinBalanceReserve decimal.Decimal // default:
	MinWithdrawAmount decimal.Decimal // default:
	MaxWithdrawAmount decimal.Decimal // default:
}

type Notice struct {
	Id        uuid.UUID // default: gen_random_uuid()
	CreatedAt time.Time // default: now()
	UpdatedAt time.Time // default: now()
	Title     string    // default:
	Content   string    // default:
}

type SmsChannel struct {
	Id       uuid.UUID       // default: gen_random_uuid()
	Platform string          // default:
	Params   json.RawMessage // default: jsonb_build_object()
	Enabled  bool            // default: true
}

type GrabSetting struct {
	Id              uuid.UUID       // default: gen_random_uuid()
	Difficulty      decimal.Decimal // default:
	CreatedAt       time.Time       // default: now()
	UpdatedAt       time.Time       // default: now()
	CommissionRatio json.RawMessage // default: jsonb_build_object()
}

type GrabOrder struct {
	Id          uuid.UUID        // default: gen_random_uuid()
	CreatedAt   time.Time        // default: now()
	UpdatedAt   time.Time        // default: now()
	Amount      decimal.Decimal  // default:
	Status      string           // default: 'pending'::text
	FinishedAt  *time.Time       // default:
	UserId      *uuid.UUID       // default:
	PayPlatform string           // default:
	Commission  *decimal.Decimal // default:
	Account     *string          // default:
	Holder      *string          // default:
	Qrcode      *string          // default:
}

type Account struct {
	Id              uuid.UUID // default: gen_random_uuid()
	CreatedAt       time.Time // default: now()
	UpdatedAt       time.Time // default: now()
	Bank            string    // default:
	Account         string    // default:
	Holder          *string   // default:
	Enabled         bool      // default: true
	UserId          uuid.UUID // default:
	Qrcode          *string   // default:
	GrabDefault     bool      // default: false
	WithdrawDefault bool      // default: false
}
