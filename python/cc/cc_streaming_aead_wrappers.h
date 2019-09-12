// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
///////////////////////////////////////////////////////////////////////////////

#ifndef TINK_PYTHON_CC_CC_STREAMING_AEAD_WRAPPERS_H_
#define TINK_PYTHON_CC_CC_STREAMING_AEAD_WRAPPERS_H_

#include <memory>

#include "absl/strings/string_view.h"
#include "tink/streaming_aead.h"
#include "tink/util/statusor.h"
#include "tink/python/cc/output_stream_adapter.h"
#include "tink/python/cc/python_file_object_adapter.h"
#include "tink/python/cc/python_output_stream.h"

namespace crypto {
namespace tink {

// Wrapper function for StreamingAead.NewEncryptingStream
//
// It uses 'streaming_aead' to create an EncryptingStream that writes the
// ciphertext to ciphertext_destination through a PythonOutputStream, and
// returns an OutputStreamAdapter that wraps this EncryptingStream.
// Taking a raw pointer signals to CLIF that the object is borrowed - ownership
// is not taken, and the value is not copied.
util::StatusOr<std::unique_ptr<OutputStreamAdapter>> NewCcEncryptingStream(
    StreamingAead* streaming_aead, const absl::string_view aad,
    std::unique_ptr<PythonFileObjectAdapter> ciphertext_destination);

}  // namespace tink
}  // namespace crypto

#endif  // TINK_PYTHON_CC_CC_STREAMING_AEAD_WRAPPERS_H_
