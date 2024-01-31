mod pb;
use std::vec;

use prost::Message;
use prost_types::value;
use substreams::pb::substreams::Clock;
use substreams_sink_kv::pb::sf::substreams::sink::kv::v1::KvOperations;

#[substreams::handlers::map]
fn map_clocks(clock: Clock) -> Result<KvOperations, substreams::errors::Error> {
    let mut kv_ops: KvOperations = Default::default();

    let mut keyIdNumber = String::new();
    let mut keyNumberId = String::new();

    let mut clockTimestampAsValue = Vec::new();
    clock.timestamp.unwrap().encode(&mut clockTimestampAsValue).expect("Buffer error on serialization");

    let mut blockNumberAsValue = Vec::new();
    clock.number.encode(&mut blockNumberAsValue).expect("Buffer error on serialization");

    kv_ops.push_new(keyIdNumber, &clockTimestampAsValue, 1);
    kv_ops.push_new(keyNumberId, &clockTimestampAsValue, 1);


    Ok(kv_ops)
}
