# -*- coding: utf-8 -*-
import time
import logging

import pymongo.results
from bson.objectid import ObjectId
from sanic.request import Request

log = logging.getLogger(__name__)


def build_state(state=1, message=None, data=None):
    return dict(state=state, message=message, data=data)


def get_cond(cond: dict = None) -> dict:
    if cond is None:
        cond = {}
    if 'deleted' not in cond:
        cond['deleted'] = False
    return cond


# 不需要加入 deleted 条件
def get_cond_self(cond: dict = None) -> dict:
    if cond is None:
        cond = {}
    return cond


# key 不存在才需要加
def opt_data(data, key, value):
    if key not in data:
        data[key] = value
    return


class ModelMethod(object):

    def __init__(self, col, col_name: str, model, request: Request = None):
        self.col = col
        self.col_name = col_name
        self.model = model  # extends BaseModel
        # 根据 model 是否有 deleted 字段判断赋值 get_cond func
        self.get_cond = self.get_cond_func()
        self.request = request

    def get_cond_func(self):
        if 'deleted' in self.model.__fields__:
            return get_cond
        else:
            return get_cond_self

    def fill_user_field(self, data):
        if self.request:
            if 'user' in self.request.cookies and self.request.cookies['user']:
                opt_data(data, 'creator', self.request.cookies['user'].USER_NAME)
                opt_data(data, 'last_changer', self.request.cookies['user'].USER_NAME)
                return

        opt_data(data, 'creator', 'api')
        opt_data(data, 'last_changer', 'api')

        return

    async def to_dict(self, data):
        self.fill_user_field(data)
        meta = self.model(**data).dict()  # 校验 check
        log.info(f"校验模型 {self.col_name} 成功")
        return meta

    async def insert_one(self, data):
        try:
            opt_data(data, 'ctime', time.time())
            opt_data(data, 'mtime', time.time())
            insert_data = await self.to_dict(data)
            result = await self.col.insert_one(insert_data)
            log.info(f"新增数据到集合 {self.col_name} 成功,数据 ObjectID: {str(result.inserted_id)} ")
            return build_state(data={'_id': result.inserted_id})
        except Exception as e:
            log.error(f"新增数据到集合 {self.col_name} 失败, err: {str(e)}")
            return build_state(state=1, message="插入数据失败，检查提交数据")

    async def insert_many(self, data):
        try:
            for index, val in enumerate(data):  # 插入多条就不检测 ctime 和 mtime 了
                opt_data(data[index], 'ctime', time.time())
                opt_data(data[index], 'mtime', time.time())
            records = [await self.to_dict(single_data) for single_data in data]
            result = await self.col.insert_many(records)
            log.info(f"新增多条数据集合 {self.col_name} 成功,数据 ObjectID: {str(result.inserted_ids)}")
            return build_state(data={'_id': result.inserted_ids})
        except Exception as e:
            log.error(f"新增多条数据到集合 {self.col_name} 失败, err: {str(e)}")
            return build_state(state=1, message="插入数据失败，检查提交数据")

    async def update_one(self, cond, data, request=None):
        try:
            # 兼容 data
            if '$set' in data:
                data = data['$set']

            # 兼容 request
            request_instance = request if request is not None else self.request
            if request_instance:
                if 'user' in request_instance.cookies and request_instance.cookies['user']:
                    data['last_changer'] = request_instance.cookies['user'].USER_NAME

            if '_id' in data:
                data.pop('_id')
            # 兼容原先数据
            if isinstance(cond, str):  # id
                filter_cond = {'_id': ObjectId(cond)}
            elif isinstance(cond, ObjectId):
                filter_cond = {'_id': cond}
            else:  # dict
                filter_cond = cond

            opt_data(data, 'mtime', time.time())
            # pymongo.results.UpdateResult
            result = await self.col.update_one(self.get_cond(filter_cond), {'$set': data})
            log.info(f"更新集合 {self.col_name} 数据 {cond} 成功, "
                     f"match count: {result.matched_count}, modify count: {result.modified_count}")
            return build_state(message="数据更新成功")
        except Exception as e:
            log.error(f"更新集合 {self.col_name} 数据 {cond} 失败, err: {str(e)}")
            return build_state(state=2, message="更新数据失败")

    async def update_many(self, filter_cond, update, upsert=False):
        filter_cond = self.get_cond(filter_cond)
        update = {"$set": update} if '$set' not in update else update  # 数据兼容
        result = await self.col.update_many(filter_cond, update, upsert)
        log.info(f"更新集合 {self.col_name} 数据成功, "
                 f"match count: {result.matched_count}, modify count: {result.modified_count}")

    async def delete_one(self, delete_id, **kwargs):
        try:
            if kwargs.get('soft'):  # 软删除
                await self.col.update_one({'_id': ObjectId(delete_id)}, {'$set': {'deleted': True}})
            else:  # 默认兼容以前的硬删除
                await self.col.delete_one({'_id': ObjectId(delete_id)})

            log.info(f"删除集合 {self.col_name} 数据 {delete_id} 成功")
            return build_state()
        except Exception as e:
            log.error(f"删除集合 {self.col_name} 数据 {delete_id} 失败, err: {str(e)}")
            return build_state(state=2, message="删除数据失败")

    async def find_many(self, **kwargs):
        try:
            result = await self.col.find(kwargs).to_list(None)
            return build_state(data=result)
        except Exception as e:
            log.error(f"查询集合 {self.col_name} 多条数据失败, err: {str(e)}")
            return build_state(state=2, message=f"查询集合 {self.col_name} 多条数据失败")

    async def find(self, filter_cond=None, *args, **kwargs):
        to_list_opt = None
        if 'to_list_opt' in kwargs:
            to_list_opt = kwargs['to_list_opt']
            kwargs.pop('to_list_opt')

        return await self.col.find(self.get_cond(filter_cond), *args, **kwargs).to_list(to_list_opt)

    async def find_raw(self, filter_cond=None, *args, **kwargs):
        # object AsyncIOMotorCursor can't be used in 'await' expression -> original find func
        return self.col.find(self.get_cond(filter_cond), *args, **kwargs)

    async def find_one(self, filter_cond=None, *args, **kwargs):
        return await self.col.find_one(self.get_cond(filter_cond), *args, **kwargs)

    async def aggregate(self, matchStage=None, projectStage=None):
        pipeline = [{"$match": self.get_cond(matchStage)}]
        if projectStage is not None:
            pipeline.append({"$project": projectStage})
        return await self.col.aggregate(pipeline).to_list(None)
