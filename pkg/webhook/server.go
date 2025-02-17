// Copyright © 2019 Banzai Cloud
//
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

package webhook

import (
	"net/http"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var (
	log = logf.Log.WithName("webhooks")
)

type webhookServer struct {
	client       client.Client
	scheme       *runtime.Scheme
	deserializer runtime.Decoder
}

func SetupServerHandlers(mgr ctrl.Manager, certDir string) {
	server := mgr.GetWebhookServer()
	server.CertDir = certDir
	mux := http.NewServeMux()
	whsrv := &webhookServer{
		client:       mgr.GetClient(),
		scheme:       mgr.GetScheme(),
		deserializer: serializer.NewCodecFactory(mgr.GetScheme()).UniversalDeserializer(),
	}
	mux.HandleFunc("/validate", whsrv.serve)
	server.Register("/validate", mux)
	return
}
