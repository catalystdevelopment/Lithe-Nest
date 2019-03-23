// Copyright (c) 2018, The TurtleCoin Developers
// Copyright (c) 2019, The Lithe Project Development Team
//
// Please see the included LICENSE file for more information.
//

package walletdmanager

const (
	// DefaultTransferFee is the default fee. It is expressed in LITHE
	DefaultTransferFee float64 = 0.0001

	logWalletdCurrentSessionFilename     = "lithe-service-session.log"
	logWalletdAllSessionsFilename        = "lithe-service.log"
	logLithedCurrentSessionFilename      = "Lithed-session.log"
	logLithedAllSessionsFilename         = "Lithed.log"
	walletdLogLevel                      = "3" // should be at least 3 as I use some logs messages to confirm creation of wallet
	walletdCommandName                   = "lithe-service"
	lithedCommandName                    = "Lithed"
)
