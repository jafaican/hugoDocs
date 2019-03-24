// Copyright 2018 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package livereloadinject

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/gohugoio/hugo/transform"
)

func TestLiveReloadInject(t *testing.T) {
	doTestLiveReloadInject(t, "</body>")
	doTestLiveReloadInject(t, "</BODY>")
}

func doTestLiveReloadInject(t *testing.T, bodyEndTag string) {
	out := new(bytes.Buffer)
	in := strings.NewReader(bodyEndTag)

	tr := transform.New(New(1313))
	tr.Apply(out, in)

	expected := fmt.Sprintf(`<script data-no-instant>document.write('<script src="/livereload.js?port=1313&mindelay=10"></' + 'script>')</script>%s`, bodyEndTag)
	if out.String() != expected {
		t.Errorf("Expected %s got %s", expected, out.String())
	}
}
