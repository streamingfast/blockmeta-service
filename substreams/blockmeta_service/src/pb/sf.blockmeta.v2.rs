// @generated
/// Block Requests
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct NumToIdReq {
    #[prost(uint64, tag="1")]
    pub block_num: u64,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct IdToNumReq {
    #[prost(string, tag="1")]
    pub block_id: ::prost::alloc::string::String,
}
/// Block & BlockByTime Responses
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct BlockResp {
    #[prost(string, tag="1")]
    pub id: ::prost::alloc::string::String,
    #[prost(uint64, tag="2")]
    pub num: u64,
    #[prost(message, optional, tag="3")]
    pub time: ::core::option::Option<::prost_types::Timestamp>,
}
/// BlockByTime Requests
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct TimeReq {
    #[prost(message, optional, tag="1")]
    pub time: ::core::option::Option<::prost_types::Timestamp>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct RelativeTimeReq {
    #[prost(message, optional, tag="1")]
    pub time: ::core::option::Option<::prost_types::Timestamp>,
    #[prost(bool, tag="2")]
    pub inclusive: bool,
}
// @@protoc_insertion_point(module)
