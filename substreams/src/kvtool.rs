use std::time::SystemTime;

const TBL_PREFIX_BLOCK_NUMS: &str = "1";
const TBL_PREFIX_NUMS_BLOCK: &str = "2";
const IDX_PREFIX_TIMELINE_FWD: &str = "3";
const IDX_PREFIX_TIMELINE_BCK: &str = "4";

pub struct Keyer;
impl Keyer {
    pub fn new() -> Self {
        Keyer
    }
    pub fn pack_id_num_key(&self, block_num: u64, block_hash: &str) -> String {
        return format!(
            "{}:{}:{}",
            TBL_PREFIX_BLOCK_NUMS,
            block_hash,
            pack_u64_reverse(block_num)
        );
    }

    pub fn pack_num_id_key(&self, block_num: u64, block_hash: &str) -> String {
        return format!(
            "{}:{}:{}",
            TBL_PREFIX_NUMS_BLOCK,
            pack_u64_reverse(block_num),
            block_hash
        );
    }
    pub fn pack_sys_time_id_key(
        &self,
        fwd: bool,
        block_time: SystemTime,
        block_hash: &str,
    ) -> String {
        let unix_millis = sys_time_to_unix_millis(block_time);
        if fwd {
            return format!(
                "{}:{}:{}",
                IDX_PREFIX_TIMELINE_FWD,
                pack_u64(unix_millis),
                block_hash
            );
        }
        return format!(
            "{}:{}:{}",
            IDX_PREFIX_TIMELINE_BCK,
            pack_u64_reverse(unix_millis),
            block_hash
        );
    }
}
fn pack_u64_reverse(num: u64) -> String {
    pack_u64(u64::MAX - num)
}

fn pack_u64(num: u64) -> String {
    hex::encode(num.to_be_bytes())
}
pub fn sys_time_to_unix_millis(sys_time: SystemTime) -> u64 {
    sys_time
        .duration_since(SystemTime::UNIX_EPOCH)
        .unwrap()
        .as_millis() as u64
}
