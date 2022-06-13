# Removes all empty keys and values in input.
def remove_empty:
  . | walk(
    if type == "object" then
      with_entries(
        select(
          .value != null and
          .value != "" and
          .value != [] and
          .key != null and
          .key != ""
        )
      )
    else .
    end
  )
;

# Converts decimal string to number.
def to_int:
  if . == null then . else .|tonumber end
;

# Converts "1" / "0" to boolean.
def to_bool:
  if . == null then . else
    if . == "1" then true else false end
  end
;

# convert comma separated strings to []string
def to_strings: 
  if . == null then [] else (.|split(",")) end
;

# Replace config in input.
. + {
  "config": {
    "posa": {
      "period": env.HIVE_KCC_POSA_BLOCK_INTERVAL|to_int,
      "epoch": env.HIVE_KCC_POSA_EPOCH|to_int,
      "ishikariInitialValidators": env.HIVE_KCC_POSA_ISHIKARI_INIT_VALIDATORS | to_strings,
      "ishikariInitialManagers": env.HIVE_KCC_POSA_ISHIKARI_INIT_MANAGERS | (if . == null then (env.HIVE_KCC_POSA_ISHIKARI_INIT_VALIDATORS | to_strings) else (.|split(",")) end ),
      "ishikariAdminAddress":  env.HIVE_KCC_POSA_ADMIN | (if . == null then "0x0000000000000000000000000000000000000000" else . end) 
    },
    "chainId": env.HIVE_CHAIN_ID|to_int,
    "homesteadBlock": 0,
    "daoForkSupport": true,
    "eip150Block": 0,
    "eip150Hash": "0x0000000000000000000000000000000000000000000000000000000000000000",
    "eip155Block": 0,
    "eip158Block": 0,
    "byzantiumBlock": 0,
    "constantinopleBlock": 0,
    "petersburgBlock": 0,
    "istanbulBlock": 0,
    "muirGlacierBlock": 0,
    "berlinBlock": 0,
    "cve_2021_39137Block":env.HIVE_CVE_2021_39137_BLOCK|to_int,
    "londonBlock": env.HIVE_FORK_LONDON|to_int,
    "ishikariBlock": env.HIVE_FORK_KCC_ISHIKARI| to_int,
  }|remove_empty
}
