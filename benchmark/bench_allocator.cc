/** Copyright 2021 Alibaba Group Holding Limited.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

#include <sys/mman.h>

#include <fstream>
#include <iostream>
#include <memory>
#include <string>
#include <thread>

#include "arrow/status.h"
#include "arrow/util/io_util.h"
#include "glog/logging.h"

#include "basic/ds/array.h"
#include "client/allocator.h"
#include "client/client.h"
#include "client/ds/object_meta.h"
#include "common/util/env.h"
#include "common/util/functions.h"

#define JEMALLOC_NO_DEMANGLE
#include "jemalloc/include/jemalloc/jemalloc.h"
#undef JEMALLOC_NO_DEMANGLE

#include "client/allocator.h"

using namespace vineyard;  // NOLINT(build/namespaces)

void alloc_with_malloc(size_t rounds, std::vector<size_t> const &sizes) {
  for (size_t round = 0; round < rounds; ++round) {
    for (size_t s: sizes) {
      free(malloc(s));
    }
  }
}

void alloc_with_jemalloc(size_t rounds, std::vector<size_t> const &sizes) {
  for (size_t round = 0; round < rounds; ++round) {
    for (size_t s: sizes) {
      vineyard_je_free(vineyard_je_malloc(s));
    }
  }
}

void alloc_with_vineyard_alloc(size_t rounds, std::vector<size_t> const &sizes,
                               VineyardAllocator<void> &allocator) {
  for (size_t round = 0; round < rounds; ++round) {
    for (size_t s: sizes) {
      allocator.Free(allocator.Allocate(s));
    }
  }
}

int main(int argc, char** argv) {
  if (argc < 2) {
    printf("usage ./bench_allocator <ipc_socket> <rounds> <allocations>");
    return 1;
  }
  std::string ipc_socket = std::string(argv[1]);
  auto rounds = std::stoll (argv[1]);

  Client client;
  VINEYARD_CHECK_OK(client.Connect(ipc_socket));
  LOG(INFO) << "Connected to IPCServer: " << ipc_socket;

  VineyardAllocator<void> allocator(client);

  std::vector<size_t> random_sizes(10)

  {
    LOG(INFO) << "Starting system malloc ...";
    auto beginTime = GetCurrentTime();

  }


  LOG(INFO) << "Finish allocator benchmarks...";
  client.Disconnect();

  return 0;
}
