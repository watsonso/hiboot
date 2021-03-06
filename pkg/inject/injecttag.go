// Copyright 2018 John Deng (hi.devops.io@gmail.com).
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

package inject

import (
	"github.com/hidevopsio/hiboot/pkg/factory"
	"github.com/hidevopsio/hiboot/pkg/utils/io"
	"github.com/hidevopsio/hiboot/pkg/utils/mapstruct"
	"github.com/hidevopsio/hiboot/pkg/utils/str"
	"reflect"
)

type injectTag struct {
	BaseTag
}

func init() {
	AddTag(new(injectTag))
}

func getInstance(cf factory.ConfigurableFactory, pkgName, name string) (retVal interface{}) {
	name = str.ToLowerCamel(name)
	retVal = cf.GetInstance(pkgName + "." + name)
	return
}

func (t *injectTag) Decode(object reflect.Value, field reflect.StructField, tag string) (retVal interface{}) {
	properties := t.ParseProperties(tag)

	// first, find if object is already instantiated
	if field.Type.Kind() == reflect.Ptr || field.Type.Kind() == reflect.Interface {
		// if object is not exist, then instantiate new object
		// parse tag and instantiate filed
		ft := field.Type
		if ft.Kind() == reflect.Ptr {
			ft = ft.Elem()
		}

		pkgName := io.DirName(ft.PkgPath())

		// get the user specific instance first
		retVal = getInstance(t.ConfigurableFactory, pkgName, tag)
		// else to find with the field name if above is not found
		if retVal == nil {
			retVal = getInstance(t.ConfigurableFactory, pkgName, field.Name)
		}
		// else to find with the type name if above is not found
		if retVal == nil {
			retVal = getInstance(t.ConfigurableFactory, pkgName, ft.Name())
		}
		// else create new instance at runtime
		if retVal == nil && field.Type.Kind() != reflect.Interface {
			o := reflect.New(ft)
			retVal = o.Interface()
		}
		// inject field value
		if properties.Count() != 0 {
			mapstruct.Decode(retVal, properties.Items())
		}
	} else {
		return
	}
	return retVal
}
