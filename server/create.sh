#!/bin/sh

mkdir -p php go 

thrift -r --gen go thrift/Service.thrift
thrift -r --gen php:server thrift/Service.thrift

