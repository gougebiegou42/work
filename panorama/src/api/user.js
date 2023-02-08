import request from '@/utils/request';

// 获取数据
export function listSylas(query) {
  return request({
    url: '/system/sylas/List',
    method: 'get',
    params: query,
  });
}
