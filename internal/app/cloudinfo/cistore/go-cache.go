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

package cistore

import (
	"fmt"
	"io"
	"time"

	"github.com/goph/emperror"
	"github.com/goph/logur"
	"github.com/patrickmn/go-cache"

	"github.com/banzaicloud/cloudinfo/pkg/cloudinfo"
)

// cacheProductStore in memory cloud product information storer
type cacheProductStore struct {
	*cache.Cache
	// all items are cached with this expiry
	itemExpiry time.Duration
	log        logur.Logger
}

func (cis *cacheProductStore) DeleteRegions(provider, service string) {
	cis.Delete(cis.getKey(cloudinfo.RegionKeyTemplate, provider, service))
}

func (cis *cacheProductStore) DeleteZones(provider, service, region string) {
	cis.Delete(cis.getKey(cloudinfo.ZoneKeyTemplate, provider, service, region))
}

func (cis *cacheProductStore) DeleteImage(provider, service, regionId string) {
	cis.Delete(cis.getKey(cloudinfo.ImageKeyTemplate, provider, service, regionId))
}

func (cis *cacheProductStore) DeleteVersion(provider, service, region string) {
	cis.Delete(cis.getKey(cloudinfo.VersionKeyTemplate, provider, service, region))
}

func (cis *cacheProductStore) StoreRegions(provider, service string, val map[string][]cloudinfo.Region) {
	cis.Set(cis.getKey(cloudinfo.RegionKeyTemplate, provider, service), val, cis.itemExpiry)
}

func (cis *cacheProductStore) GetRegions(provider, service string) (map[string][]cloudinfo.Region, bool) {
	if res, ok := cis.get(cis.getKey(cloudinfo.RegionKeyTemplate, provider, service)); ok {
		return res.(map[string][]cloudinfo.Region), ok
	}
	return nil, false
}

func (cis *cacheProductStore) StoreZones(provider, service, region string, val []string) {
	cis.Set(cis.getKey(cloudinfo.ZoneKeyTemplate, provider, service, region), val, cis.itemExpiry)
}

func (cis *cacheProductStore) GetZones(provider, service, region string) ([]string, bool) {
	if res, ok := cis.get(cis.getKey(cloudinfo.ZoneKeyTemplate, provider, service, region)); ok {
		return res.([]string), ok
	}

	return nil, false
}

func (cis *cacheProductStore) StorePrice(provider, region, instanceType string, val cloudinfo.Price) {
	cis.Set(cis.getKey(cloudinfo.PriceKeyTemplate, provider, region, instanceType), val, cis.itemExpiry)
}

func (cis *cacheProductStore) GetPrice(provider, region, instanceType string) (cloudinfo.Price, bool) {
	if res, ok := cis.get(cis.getKey(cloudinfo.PriceKeyTemplate, provider, region, instanceType)); ok {
		return res.(cloudinfo.Price), ok
	}
	return cloudinfo.Price{}, false
}

func (cis *cacheProductStore) StoreVm(provider, service, region string, val []cloudinfo.VmInfo) {
	cis.Set(cis.getKey(cloudinfo.VmKeyTemplate, provider, service, region), val, cis.itemExpiry)
}

func (cis *cacheProductStore) GetVm(provider, service, region string) ([]cloudinfo.VmInfo, bool) {
	if res, ok := cis.get(cis.getKey(cloudinfo.VmKeyTemplate, provider, service, region)); ok {
		return res.([]cloudinfo.VmInfo), ok
	}

	return nil, false
}

func (cis *cacheProductStore) StoreImage(provider, service, regionId string, val []cloudinfo.Image) {
	cis.Set(cis.getKey(cloudinfo.ImageKeyTemplate, provider, service, regionId), val, cis.itemExpiry)
}

func (cis *cacheProductStore) GetImage(provider, service, regionId string) ([]cloudinfo.Image, bool) {
	if res, ok := cis.get(cis.getKey(cloudinfo.ImageKeyTemplate, provider, service, regionId)); ok {
		return res.([]cloudinfo.Image), ok
	}

	return nil, false
}

func (cis *cacheProductStore) StoreVersion(provider, service, region string, val []cloudinfo.LocationVersion) {
	cis.Set(cis.getKey(cloudinfo.VersionKeyTemplate, provider, service, region), val, cis.itemExpiry)
}

func (cis *cacheProductStore) GetVersion(provider, service, region string) ([]cloudinfo.LocationVersion, bool) {
	if res, ok := cis.get(cis.getKey(cloudinfo.VersionKeyTemplate, provider, service, region)); ok {
		return res.([]cloudinfo.LocationVersion), ok
	}

	return nil, false
}

func (cis *cacheProductStore) StoreStatus(provider string, val string) {
	cis.Set(cis.getKey(cloudinfo.StatusKeyTemplate, provider), val, cis.itemExpiry)
}

func (cis *cacheProductStore) GetStatus(provider string) (string, bool) {
	if res, ok := cis.Get(cis.getKey(cloudinfo.StatusKeyTemplate, provider)); ok {
		return res.(string), ok
	}

	return "", false
}

// Export writes the content of the store into the passed in writer
func (cis *cacheProductStore) Export(w io.Writer) error {
	if err := cis.Save(w); err != nil {
		cis.log.Error("failed to export the store", map[string]interface{}{"op": "export", "destination": "todo"})
		return emperror.WrapWith(err, "failed to export the store", "op", "export", "destination", "todo")
	}
	return nil
}

// Import loads the store data from the standard input
func (cis *cacheProductStore) Import(r io.Reader) error {
	if err := cis.Load(r); err != nil {
		cis.log.Error("failed to load store data", map[string]interface{}{"op": "import", "destination": "todo"})
		return emperror.WrapWith(err, "failed to load the store data", "op", "import", "destination", "todo")
	}
	return nil
}

func (cis *cacheProductStore) DeleteVm(provider, service, region string) {
	cis.Delete(cis.getKey(cloudinfo.VmKeyTemplate, provider, service, region))
}

func (cis *cacheProductStore) StoreServices(provider string, services []cloudinfo.Service) {
	cis.Set(cis.getKey(cloudinfo.ServicesKeyTemplate, provider), services, cis.itemExpiry)
}

func (cis *cacheProductStore) GetServices(provider string) ([]cloudinfo.Service, bool) {
	r, o := cis.Get(cis.getKey(cloudinfo.ServicesKeyTemplate, provider))
	return r.([]cloudinfo.Service), o
}

// NewCacheProductStore creates a new store instance.
// the backing cache is initialized with the defaultExpiration and cleanupInterval
func NewCacheProductStore(cloudInfoExpiration, cleanupInterval time.Duration, logger logur.Logger) cloudinfo.CloudInfoStore {
	return &cacheProductStore{
		cache.New(cloudInfoExpiration, cleanupInterval),
		cleanupInterval,
		logger,
	}
}

func (cis *cacheProductStore) getKey(keyTemplate string, args ...interface{}) string {
	return fmt.Sprintf(keyTemplate, args...)
}

func (cis *cacheProductStore) get(key string) (interface{}, bool) {
	if val, ok := cis.Get(key); ok && val != nil {
		return val, true
	}

	return nil, false
}

func (cis *cacheProductStore) Close() {
}
