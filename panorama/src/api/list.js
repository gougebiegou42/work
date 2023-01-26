import request from '@/utils/request';

// 获取数据
export function panoramaList(query) {
  return request({
    url: '/system/panorama1/List',
    method: 'get',
    params: query,
  });
}
