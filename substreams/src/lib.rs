use std::time::SystemTime;
use substreams::pb::substreams::Clock;
use substreams_sink_kv::pb::sf::substreams::sink::kv::v1::KvOperations;
const KEY_BLOCK_BY_ID_NUMBER_PREFIX: &str = "bi";
const KEY_BLOCK_BY_NUMBER_ID_PREFIX: &str = "bn";
const KEY_BLOCK_BY_TIMESTAMP_ID_PREFIX: &str = "bt";

#[substreams::handlers::map]
fn map_clocks(clock: Clock) -> Result<KvOperations, substreams::errors::Error> {
    let mut kv_ops: KvOperations = Default::default();

    let clock_timestamp = clock.timestamp.unwrap();
    let sys_time: SystemTime = clock_timestamp.try_into().unwrap();
    let unix_millis = sys_time.duration_since(SystemTime::UNIX_EPOCH).unwrap().as_millis();

    let key_id_number = format!("{}:{}:{}", KEY_BLOCK_BY_ID_NUMBER_PREFIX, clock.id, clock.number);
    let key_number_id = format!("{}:{}:{}", KEY_BLOCK_BY_NUMBER_ID_PREFIX, clock.number, clock.id);
    let key_timestamp_id = format!("{}:{}:{}", KEY_BLOCK_BY_TIMESTAMP_ID_PREFIX, unix_millis, clock.number);

    let unix_millis_bytes = unix_millis.to_be_bytes().to_vec();

    kv_ops.push_new(key_id_number, &unix_millis_bytes, 1);
    kv_ops.push_new(key_number_id, &unix_millis_bytes, 1);
    kv_ops.push_new(key_timestamp_id, clock.number.to_be_bytes(), 1);

    Ok(kv_ops)
}
