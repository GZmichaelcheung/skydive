/*
 * Copyright (C) 2018 IBM, Inc.
 *
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 *
 */

package istio

import (
	"fmt"

	kiali "github.com/kiali/kiali/kubernetes"
	"github.com/skydive-project/skydive/topology/graph"
	"github.com/skydive-project/skydive/topology/probes/k8s"
)

type quotaSpecBindingHandler struct {
}

// Map graph node to k8s resource
func (h *quotaSpecBindingHandler) Map(obj interface{}) (graph.Identifier, graph.Metadata) {
	qsb := obj.(*kiali.QuotaSpecBinding)
	return graph.Identifier(qsb.GetUID()), k8s.NewMetadata(Manager, "quotaspecbinding", qsb, qsb.Name, qsb.Namespace)
}

// Dump k8s resource
func (h *quotaSpecBindingHandler) Dump(obj interface{}) string {
	qsb := obj.(*kiali.QuotaSpecBinding)
	return fmt.Sprintf("quotaspecbinding{Namespace: %s, Name: %s}", qsb.Namespace, qsb.Name)
}

func newQuotaSpecBindingProbe(client *kiali.IstioClient, g *graph.Graph) k8s.Subprobe {
	return k8s.NewResourceCache(client.GetIstioConfigApi(), &kiali.QuotaSpecBinding{}, "quotaspecbindings", g, &quotaSpecBindingHandler{})
}