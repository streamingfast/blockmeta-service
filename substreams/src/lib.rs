use std::time::SystemTime;
use substreams::pb::substreams::Clock;
use substreams_sink_kv::pb::sf::substreams::sink::kv::v1::KvOperations;

mod kvtool;

#[substreams::handlers::map]
fn map_clocks(clock: Clock) -> Result<KvOperations, substreams::errors::Error> {
    let keyer = kvtool::Keyer::new();

    let mut kv_ops: KvOperations = Default::default();

    let clock_timestamp = clock.timestamp.unwrap();
    let sys_time: SystemTime = clock_timestamp.try_into().unwrap();

    let key_id_number = keyer.pack_id_num_key(clock.number, &clock.id);
    let key_number_id = keyer.pack_num_id_key(clock.number, &clock.id);
    let key_timestamp_id = keyer.pack_sys_time_id_key(true, sys_time, &clock.id);
    let key_reverse_timestamp_id = keyer.pack_sys_time_id_key(false, sys_time, &clock.id);

    let unix_millis_bytes = kvtool::sys_time_to_unix_millis(sys_time).to_be_bytes();

    kv_ops.push_new(key_id_number, &unix_millis_bytes, 1);
    kv_ops.push_new(key_number_id, &unix_millis_bytes, 1);
    kv_ops.push_new(key_timestamp_id, clock.number.to_be_bytes(), 1);
    kv_ops.push_new(key_reverse_timestamp_id, clock.number.to_be_bytes(), 1);

    Ok(kv_ops)
}
