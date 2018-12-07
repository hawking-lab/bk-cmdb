/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.,
 * Copyright (C) 2017,-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the ",License",); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an ",AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package model

import (
	"configcenter/src/common/metadata"
	"configcenter/src/source_controller/coreservice/core"
	"configcenter/src/storage/dal"
)

var _ core.ModelOperation = (*modelManager)(nil)

type modelManager struct {
	*modelAttribute
	*modelClassification
	dbProxy dal.RDB
}

// New create a new model manager instance
func New(dbProxy dal.RDB) core.ModelOperation {

	coreMgr := &modelManager{dbProxy: dbProxy}

	coreMgr.modelAttribute = &modelAttribute{dbProxy: dbProxy, model: coreMgr}
	coreMgr.modelClassification = &modelClassification{dbProxy: dbProxy, model: coreMgr}

	return coreMgr
}

func (m *modelManager) CreateModel(ctx core.ContextParams, inputParam metadata.CreateModel) (*metadata.CreateOneDataResult, error) {
	return nil, nil
}
func (m *modelManager) SetModel(ctx core.ContextParams, inputParam metadata.SetModel) (*metadata.SetDataResult, error) {
	return nil, nil
}
func (m *modelManager) UpdateModel(ctx core.ContextParams, inputParam metadata.UpdateOption) (*metadata.UpdatedCount, error) {
	return nil, nil
}
func (m *modelManager) DeleteModel(ctx core.ContextParams, inputParam metadata.DeleteOption) (*metadata.DeletedCount, error) {
	return nil, nil
}
func (m *modelManager) CascadeDeleteModel(ctx core.ContextParams, inputParam metadata.DeleteOption) (*metadata.DeletedCount, error) {
	return nil, nil
}
func (m *modelManager) SearchModel(ctx core.ContextParams, inputParam metadata.QueryCondition) (*metadata.QueryResult, error) {
	return nil, nil
}
