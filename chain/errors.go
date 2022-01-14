// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package chain

import (
	"errors"
)

var (
	// Block Correctness
	ErrTimestampTooEarly      = errors.New("block timestamp too early")
	ErrTimestampTooLate       = errors.New("block timestamp too late")
	ErrNoTxs                  = errors.New("no transactions")
	ErrInvalidCost            = errors.New("invalid block cost")
	ErrInvalidPrice           = errors.New("invalid price")
	ErrInsufficientSurplus    = errors.New("insufficient surplus difficulty")
	ErrParentBlockNotVerified = errors.New("parent block not verified or accepted")

	// Tx Correctness
	ErrInvalidMagic      = errors.New("invalid magic")
	ErrInvalidBlockID    = errors.New("invalid blockID")
	ErrInvalidSignature  = errors.New("invalid signature")
	ErrDuplicateTx       = errors.New("duplicate transaction")
	ErrInsufficientPrice = errors.New("insufficient price")

	// Execution Correctness
	ErrValueEmpty      = errors.New("value empty")
	ErrValueTooBig     = errors.New("value too big")
	ErrSpaceExpired    = errors.New("space expired")
	ErrKeyMissing      = errors.New("key missing")
	ErrInvalidKey      = errors.New("key is invalid")
	ErrAddressMismatch = errors.New("address does not match decoded space")
	ErrSpaceNotExpired = errors.New("space not expired")
	ErrSpaceMissing    = errors.New("space missing")
	ErrUnauthorized    = errors.New("sender is not authorized")
	ErrInvalidBalance  = errors.New("invalid balance")
	ErrNonActionable   = errors.New("transaction doesn't do anything")
)
