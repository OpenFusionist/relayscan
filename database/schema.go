package database

import (
	"github.com/flashbots/mev-boost-relay/common"
)

var (
	tableBase = common.GetEnv("DB_TABLE_PREFIX", "dev")

	TableBid = tableBase + "_bid"
	// TableExecutionPayload       = tableBase + "_execution_payload"
	// TableBuilderBlockSubmission = tableBase + "_builder_block_submission"
	// TableDeliveredPayload       = tableBase + "_payload_delivered"
	// TableBlockBuilder           = tableBase + "_blockbuilder"
)

var schema = `
CREATE TABLE IF NOT EXISTS ` + TableBid + ` (
	id          bigint GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
	inserted_at timestamp NOT NULL default current_timestamp,

	pubkey        varchar(98) NOT NULL,
	fee_recipient varchar(42) NOT NULL,
	timestamp     bigint NOT NULL,
	gas_limit     bigint NOT NULL,
	signature     text NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS ` + TableBid + `_pubkey_feerec_idx ON ` + TableBid + `(pubkey, fee_recipient);
`

// CREATE TABLE IF NOT EXISTS ` + TableExecutionPayload + ` (
// 	id          bigint GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
// 	inserted_at timestamp NOT NULL default current_timestamp,

// 	slot            bigint NOT NULL,
// 	proposer_pubkey varchar(98) NOT NULL,
// 	block_hash      varchar(66) NOT NULL,

// 	version     text NOT NULL, -- bellatrix
// 	payload 	json NOT NULL
// );

// CREATE UNIQUE INDEX IF NOT EXISTS ` + TableExecutionPayload + `_slot_pk_hash_idx ON ` + TableExecutionPayload + `(slot, proposer_pubkey, block_hash);

// CREATE TABLE IF NOT EXISTS ` + TableBuilderBlockSubmission + ` (
// 	id bigint GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
// 	inserted_at timestamp NOT NULL default current_timestamp,

// 	execution_payload_id bigint references ` + TableExecutionPayload + `(id) on delete set null,

// 	-- simulation & verification results
// 	sim_success boolean NOT NULL,
// 	sim_error   text    NOT NULL,

// 	-- bidtrace data
// 	signature            text NOT NULL,

// 	slot        bigint NOT NULL,
// 	parent_hash varchar(66) NOT NULL,
// 	block_hash  varchar(66) NOT NULL,

// 	builder_pubkey         varchar(98) NOT NULL,
// 	proposer_pubkey        varchar(98) NOT NULL,
// 	proposer_fee_recipient varchar(42) NOT NULL,

// 	gas_used   bigint NOT NULL,
// 	gas_limit  bigint NOT NULL,

// 	num_tx int NOT NULL,
// 	value  NUMERIC(48, 0),

// 	-- helpers
// 	epoch        bigint NOT NULL,
// 	block_number bigint NOT NULL,
// 	was_most_profitable boolean NOT NULL
// );

// CREATE INDEX IF NOT EXISTS ` + TableBuilderBlockSubmission + `_slot_idx ON ` + TableBuilderBlockSubmission + `("slot");
// CREATE INDEX IF NOT EXISTS ` + TableBuilderBlockSubmission + `_blockhash_idx ON ` + TableBuilderBlockSubmission + `("block_hash");
// CREATE INDEX IF NOT EXISTS ` + TableBuilderBlockSubmission + `_blocknumber_idx ON ` + TableBuilderBlockSubmission + `("block_number");
// CREATE INDEX IF NOT EXISTS ` + TableBuilderBlockSubmission + `_builderpubkey_idx ON ` + TableBuilderBlockSubmission + `("builder_pubkey");
// CREATE INDEX IF NOT EXISTS ` + TableBuilderBlockSubmission + `_simsuccess_idx ON ` + TableBuilderBlockSubmission + `("sim_success");
// CREATE INDEX IF NOT EXISTS ` + TableBuilderBlockSubmission + `_mostprofit_idx ON ` + TableBuilderBlockSubmission + `("was_most_profitable");
// CREATE INDEX IF NOT EXISTS ` + TableBuilderBlockSubmission + `_executionpayloadid_idx ON ` + TableBuilderBlockSubmission + `("execution_payload_id");

// CREATE TABLE IF NOT EXISTS ` + TableDeliveredPayload + ` (
// 	id bigint GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
// 	inserted_at timestamp NOT NULL default current_timestamp,

// 	execution_payload_id        bigint references ` + TableExecutionPayload + `(id) on delete set null,
// 	signed_blinded_beacon_block json,

// 	epoch bigint NOT NULL,
// 	slot  bigint NOT NULL,

// 	builder_pubkey         varchar(98) NOT NULL,
// 	proposer_pubkey        varchar(98) NOT NULL,
// 	proposer_fee_recipient varchar(42) NOT NULL,

// 	parent_hash  varchar(66) NOT NULL,
// 	block_hash   varchar(66) NOT NULL,
// 	block_number bigint NOT NULL,

// 	gas_used  bigint NOT NULL,
// 	gas_limit bigint NOT NULL,

// 	num_tx  int NOT NULL,
// 	value   NUMERIC(48, 0),

// 	UNIQUE (slot, proposer_pubkey, block_hash)
// );

// CREATE INDEX IF NOT EXISTS ` + TableDeliveredPayload + `_slot_idx ON ` + TableDeliveredPayload + `("slot");
// CREATE INDEX IF NOT EXISTS ` + TableDeliveredPayload + `_blockhash_idx ON ` + TableDeliveredPayload + `("block_hash");
// CREATE INDEX IF NOT EXISTS ` + TableDeliveredPayload + `_blocknumber_idx ON ` + TableDeliveredPayload + `("block_number");
// CREATE INDEX IF NOT EXISTS ` + TableDeliveredPayload + `_proposerpubkey_idx ON ` + TableDeliveredPayload + `("proposer_pubkey");
// CREATE INDEX IF NOT EXISTS ` + TableDeliveredPayload + `_builderpubkey_idx ON ` + TableDeliveredPayload + `("builder_pubkey");
// CREATE INDEX IF NOT EXISTS ` + TableDeliveredPayload + `_executionpayloadid_idx ON ` + TableDeliveredPayload + `("execution_payload_id");
// CREATE INDEX IF NOT EXISTS ` + TableDeliveredPayload + `_value_idx ON ` + TableDeliveredPayload + `("value");

// CREATE TABLE IF NOT EXISTS ` + TableBlockBuilder + ` (
// 	id bigint GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
// 	inserted_at timestamp NOT NULL default current_timestamp,

// 	builder_pubkey  varchar(98) NOT NULL,
// 	description    	text NOT NULL,

// 	is_high_prio    boolean NOT NULL,
// 	is_blacklisted  boolean NOT NULL,

// 	last_submission_id   bigint references ` + TableBuilderBlockSubmission + `(id) on delete set null,
// 	last_submission_slot bigint NOT NULL,

// 	num_submissions_total    bigint NOT NULL,
// 	num_submissions_simerror bigint NOT NULL,
// 	num_submissions_topbid   bigint NOT NULL,

// 	num_sent_getpayload bigint NOT NULL DEFAULT 0,

// 	UNIQUE (builder_pubkey)
// );
// `
